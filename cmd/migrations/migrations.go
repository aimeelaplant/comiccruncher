package main

import (
	"fmt"
	"github.com/comiccruncher/comiccruncher/comic"
	"github.com/comiccruncher/comiccruncher/internal/log"
	"github.com/comiccruncher/comiccruncher/internal/pgo"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"go.uber.org/zap"
	"os"
)

var (
	materializedViews = map[string]map[comic.AppearanceType]uint{
		"mv_ranked_characters": {
			comic.Main | comic.Alternate: 0,
		},
		"mv_ranked_characters_main": {
			comic.Main: 0,
		},
		"mv_ranked_characters_alternate": {
			comic.Alternate: 0,
		},
		"mv_ranked_characters_marvel_main": {
			comic.Main: 1,
		},
		"mv_ranked_characters_dc_main": {
			comic.Main: 2,
		},
	}
	tables = []interface{}{
		&comic.Publisher{},
		&comic.Character{},
		&comic.CharacterSource{},
		&comic.CharacterSyncLog{},
		&comic.Issue{},
		&comic.CharacterIssue{},
	}
	updatedAtTriggers = []string{
		"publishers",
		"characters",
		"character_sources",
		"character_sync_logs",
		"issues",
		"character_issues",
	}
	opts = &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	execs = []string{
		trendingSQL("mv_trending_characters_marvel", 1),
		trendingSQL("mv_trending_characters_dc", 2),
		fmt.Sprintf(`CREATE UNIQUE INDEX IF NOT EXISTS %[1]s_id_idx ON %[1]s(id);`, "mv_trending_characters_marvel"),
		fmt.Sprintf(`CREATE UNIQUE INDEX IF NOT EXISTS %[1]s_id_idx ON %[1]s(id);`, "mv_trending_characters_dc"),
	}
)

// Logs the error as fatal and exits.
func logIfError(err error) error {
	if err != nil {
		log.MIGRATIONS().Error("error", zap.Error(err))
	}
	return err
}

// Logs the error as fatal and exits.
func logResultIfError(_ orm.Result, err error) error {
	if err != nil {
		log.MIGRATIONS().Error("error", zap.Error(err))
	}
	return err
}

// Logs an error if there's an error instantiating the db.
// Or logs when there's no error.
func logFatalIfError(err error, env string) {
	if err != nil {
		log.MIGRATIONS().Fatal("error instantiating database", zap.Error(err))
	}
	log.MIGRATIONS().Info("instantiated connection", zap.String("environment", env))
}

// Generates SQL for creating an `updated_at` column for the given `tableName`.
func updatedAtTrigger(tableName string) string {
	return fmt.Sprintf(`
		CREATE OR REPLACE FUNCTION update_updated_at_column() 
		RETURNS TRIGGER AS $$
		BEGIN
    		NEW.updated_at = now();
    	RETURN NEW; 
		END;
		$$ language 'plpgsql';
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'update_%[1]s_updated_at_column') THEN
				CREATE TRIGGER update_%[1]s_updated_at_column
				BEFORE UPDATE ON %[1]s FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
			END IF;
		END;
		$$`, tableName)
}

func mustInstance() *pg.DB {
	env := os.Getenv("CC_ENVIRONMENT")
	if env == "test" {
		db, err := pgo.InstanceTest()
		logFatalIfError(err, env)
		return db
	} else if env == "development" || env == "production" {
		db, err := pgo.Instance()
		logFatalIfError(err, env)
		return db
	}
	return nil
}

func main() {
	tx := mustInstance()
	err := logIfError(tx.RunInTransaction(func(tx *pg.Tx) error {
		if os.Getenv("CC_ENVIRONMENT") != "production" {
			if err := logResultIfError(tx.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")); err != nil {
				return err
			}
		}
		// create tables
		for _, table := range tables {
			if err := logIfError(tx.CreateTable(table, opts)); err != nil {
				return err
			}
		}
		// create triggers
		for _, t := range updatedAtTriggers {
			if err := logResultIfError(tx.Exec(updatedAtTrigger(t))); err != nil {
				return err
			}
		}
		// indexes
		if err := logResultIfError(tx.Exec(`
			CREATE INDEX IF NOT EXISTS characters_publisher_id_idx ON characters(publisher_id) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS characters_name_odx ON characters(name) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS character_sources_character_id_idx ON character_sources(character_id) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS character_sync_logs_character_id_idx ON character_sync_logs(character_id);
			CREATE INDEX IF NOT EXISTS characters_name_idx_gin on characters USING GIN(name gin_trgm_ops) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS characters_other_name_idx_gin ON characters USING GIN(other_name gin_trgm_ops) WHERE is_disabled = false AND (other_name IS NOT NULL AND other_name != '');
			CREATE INDEX IF NOT EXISTS issues_sale_date_idx ON issues(sale_date);
			CREATE INDEX IF NOT EXISTS issues_publication_date_idx ON issues(publication_date);
		`)); err != nil {
			return err
		}
		if err := logResultIfError(tx.Exec(`
			ALTER TABLE IF EXISTS character_sources
  				ADD COLUMN IF NOT EXISTS vendor_other_name text NULL
		`)); err != nil {
			return err
		}
		// gonna have to add a default here. don't want to deal with null boolean values!
		if err := logResultIfError(tx.Exec(`
			ALTER TABLE IF EXISTS issues
			ADD COLUMN IF NOT EXISTS is_reprint bool NOT NULL DEFAULT FALSE
		`)); err != nil {
			return err
		}
		if os.Getenv("CC_ENVIRONMENT") != "test" {
			if err := logResultIfError(tx.Exec("INSERT INTO publishers (name, slug, created_at, updated_at) VALUES (?, ?, now(), now()) ON CONFLICT DO NOTHING;", "Marvel", "marvel")); err != nil {
				return err
			}
			if err := logResultIfError(tx.Exec("INSERT INTO publishers (name, slug, created_at, updated_at) VALUES (?, ?, now(), now()) ON CONFLICT DO NOTHING;", "DC Comics", "dc")); err != nil {
				return err
			}
		}
		// views
		for view, t := range materializedViews {
			for ty, pubID := range t {
				if err := logResultIfError(tx.Exec(rankedCharactersSQL(view, ty, pubID))); err != nil {
					return err
				}
				if err := logResultIfError(tx.Exec(fmt.Sprintf(`CREATE UNIQUE INDEX IF NOT EXISTS %[1]s_id_idx ON %[1]s(id);`, view))); err != nil {
					return err
				}
			}
		}
		// trending views / other execs
		for _, exec := range execs {
			if err := logResultIfError(tx.Exec(exec)); err != nil {
				return err
			}
		}
		return nil
	}))

	if err != nil {
		log.MIGRATIONS().Fatal("error for transaction", zap.Error(err))
	}

	log.MIGRATIONS().Info("done")
}

func rankedCharactersSQL(name string, t comic.AppearanceType, publisherID uint) string {
	sql := fmt.Sprintf(`
		CREATE MATERIALIZED VIEW IF NOT EXISTS %s AS
		  SELECT
		  dense_rank() OVER (ORDER BY count(ci.id) DESC) AS issue_count_rank,
		  count(ci.id) as issue_count,
		  dense_rank() OVER (
		   ORDER BY
			 (
				count(ci.id)
				/
				(
				  CASE
					WHEN
					 (date_part('year', current_date)) -  min(date_part('year', i.sale_date)) = 0
					 THEN 1 -- avoid division by 0
					ELSE (date_part('year', current_date)) -  min(date_part('year', i.sale_date))
				  END
				)
			 ) DESC) AS average_per_year_rank,
		  round(count(ci.id)
		  /
		  (
			 CASE
			   WHEN
				 (date_part('year', current_date)) -  min(date_part('year', i.sale_date)) = 0
					   THEN 1 -- avoid division by 0
			   ELSE (date_part('year', current_date)) -  min(date_part('year', i.sale_date))
				 END
			 )::DECIMAL, 2) as average_per_year,
		  c.id,
		  c.publisher_id,
		  c.name,
		  c.other_name,
		  c.description,
		  c.image,
		  c.slug,
		  c.vendor_image,
		  c.vendor_url,
		  c.vendor_description,
          p.id as publisher__id,
          p.slug as publisher__slug,
          p.name as publisher__name
		  FROM characters c
			JOIN character_issues ci ON ci.character_id = c.id
			JOIN issues i ON i.id = ci.issue_id
            JOIN publishers p ON p.id = c.publisher_id
          WHERE 1=1
		`, name)
	if t == comic.Main {
		sql += " AND ci.appearance_type & 1::bit(8) > 0::bit(8)"
	}
	if t == comic.Alternate {
		sql += " AND ci.appearance_type & 2::bit(8) > 0::bit(8)"
	}
	if publisherID != 0 {
		sql += fmt.Sprintf(" AND c.publisher_id = %d", publisherID)
	}
	sql += ` AND c.is_disabled = false
				GROUP BY c.id, p.id
				ORDER BY issue_count_rank`
	return sql
}

func trendingSQL(name string, publisherID uint) string {
	sql := fmt.Sprintf(`
		CREATE MATERIALIZED VIEW IF NOT EXISTS %s AS
		SELECT
			   dense_rank() OVER (ORDER BY count(ci.id) DESC) AS issue_count_rank,
			   count(ci.id) as issue_count,
			   dense_rank() OVER (
					  ORDER BY
							 (
									count(ci.id)
									/
									(
										   CASE
										   WHEN
												  (date_part('year', current_date)) -  min(date_part('year', i.sale_date)) = 0
												  THEN 1 -- avoid division by 0
										   ELSE (date_part('year', current_date)) -  min(date_part('year', i.sale_date))
										   END
									)
							 ) DESC) AS average_per_year_rank,
			   round(count(ci.id)
							/
					 (
						 CASE
								WHEN
							 (date_part('year', current_date)) -  min(date_part('year', i.sale_date)) = 0
								   THEN 1 -- avoid division by 0
								ELSE (date_part('year', current_date)) -  min(date_part('year', i.sale_date))
							 END
						 )::DECIMAL, 2) as average_per_year,
			   c.id,
			   c.publisher_id,
			   c.name,
			   c.other_name,
			   c.description,
			   c.image,
			   c.slug,
			   c.vendor_image,
			   c.vendor_url,
			   c.vendor_description,
			   p.id as publisher__id,
			   p.slug as publisher__slug,
			   p.name as publisher__name
		FROM characters c
					JOIN character_issues ci ON ci.character_id = c.id
					JOIN issues i ON i.id = ci.issue_id
					JOIN publishers p ON p.id = c.publisher_id
		WHERE c.publisher_id = %d
		AND c.is_disabled = FALSE
		AND i.sale_date > date_trunc('month', CURRENT_DATE) - INTERVAL '1 year'
		AND ci.appearance_type & B'00000001' > 0::BIT(8)
		GROUP BY c.slug, c.id, c.name, c.other_name, p.id
		ORDER BY issue_count DESC
		LIMIT 50;`, name, publisherID)
	return sql
}
