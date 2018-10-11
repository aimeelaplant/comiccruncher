package main

import (
	"fmt"
	"github.com/aimeelaplant/comiccruncher/comic"
	"github.com/aimeelaplant/comiccruncher/internal/log"
	"github.com/aimeelaplant/comiccruncher/internal/pgo"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"go.uber.org/zap"
	"os"
)

// Logs the error as fatal and exits.
func logIfError(err error) error {
	if err != nil {
		log.MIGRATIONS().Error("error", zap.Error(err))
	}
	return err
}

// Logs the error as fatal and exits.
func logWithResultIfError(_ orm.Result, err error) error {
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
			if err := logWithResultIfError(tx.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")); err != nil {
				return err
			}
		}
		opts := &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		}
		if err := logIfError(tx.CreateTable(&comic.Publisher{}, opts)); err != nil {
			return err
		}
		if err := logIfError(tx.CreateTable(&comic.Character{}, opts)); err != nil {
			return err
		}
		if err := logIfError(tx.CreateTable(&comic.CharacterSource{}, opts)); err != nil {
			return err
		}
		if err := logIfError(tx.CreateTable(&comic.CharacterSyncLog{}, opts)); err != nil {
			return err
		}
		if err := logIfError(tx.CreateTable(&comic.Issue{}, opts)); err != nil {
			return err
		}
		if err := logIfError(tx.CreateTable(&comic.CharacterIssue{}, opts)); err != nil {
			return err
		}
		updatedAtTriggers := []string{
			updatedAtTrigger("publishers"),
			updatedAtTrigger("characters"),
			updatedAtTrigger("character_sources"),
			updatedAtTrigger("character_sync_logs"),
			updatedAtTrigger("issues"),
			updatedAtTrigger("character_issues"),
		}
		for _, t := range updatedAtTriggers {
			if err := logWithResultIfError(tx.Exec(t)); err != nil {
				return err
			}
		}
		if err := logWithResultIfError(tx.Exec(`
			CREATE INDEX IF NOT EXISTS characters_publisher_id_idx ON characters(publisher_id) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS characters_name_odx ON characters(name) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS character_sources_character_id_idx ON character_sources(character_id) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS character_sync_logs_character_id_idx ON character_sync_logs(character_id);
			CREATE INDEX IF NOT EXISTS characters_name_idx_gin on characters USING GIN(name gin_trgm_ops) WHERE is_disabled = false;
			CREATE INDEX IF NOT EXISTS characters_other_name_idx_gin ON characters USING GIN(other_name gin_trgm_ops) WHERE is_disabled = false AND (other_name IS NOT NULL AND other_name != '');
		`)); err != nil {
			return err
		}
		if err := logWithResultIfError(tx.Exec(`
			ALTER TABLE IF EXISTS character_sources
  				ADD COLUMN IF NOT EXISTS vendor_other_name text NULL
		`)); err != nil {
			return err
		}
		// gonna have to add a default here. don't want to deal with null boolean values!
		if err := logWithResultIfError(tx.Exec(`
			ALTER TABLE IF EXISTS issues
			ADD COLUMN IF NOT EXISTS is_reprint bool NOT NULL DEFAULT FALSE
		`)); err != nil {
			return err
		}
		if os.Getenv("CC_ENVIRONMENT") != "test" {
			if err := logWithResultIfError(tx.Exec("INSERT INTO publishers (name, slug, created_at, updated_at) VALUES (?, ?, now(), now()) ON CONFLICT DO NOTHING;", "Marvel", "marvel")); err != nil {
				return err
			}
			if err := logWithResultIfError(tx.Exec("INSERT INTO publishers (name, slug, created_at, updated_at) VALUES (?, ?, now(), now()) ON CONFLICT DO NOTHING;", "DC Comics", "dc")); err != nil {
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
