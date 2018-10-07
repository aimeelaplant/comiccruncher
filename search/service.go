package search

import (
	"github.com/aimeelaplant/comiccruncher/comic"
	"github.com/go-pg/pg"
)

// The interface for searching.
type Searcher interface {
	// Searches for characters by their names.
	Characters(name string, limit, offset int) ([]*comic.Character, error)
}

type Service struct {
	db *pg.DB
}

// Searches for characters by their names using postgres' trigram similarities.
func (s *Service) Characters(name string, limit, offset int) ([]*comic.Character, error) {
	characters := make([]*comic.Character, 0)
	err := s.db.
		Model(&characters).
		Column("character.*", "Publisher").
		Where("character.is_disabled = ?", false).
		Where("(character.name % ?0 OR character.other_name % ?0)", name).
		OrderExpr(`
			CASE 
				WHEN SIMILARITY(character.name, ?0) > SIMILARITY(character.other_name, ?0)
				THEN SIMILARITY(character.name, ?0)
				ELSE SIMILARITY(character.other_name, ?0)
			END DESC`, name).
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func NewSearchService(db *pg.DB) Searcher {
	return &Service{
		db: db,
	}
}
