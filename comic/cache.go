package comic

import (
	"fmt"
	"github.com/aimeelaplant/comiccruncher/internal/log"
	"go.uber.org/zap"
)

type Syncer interface {
	// Syncs appearances from postgres to redis. Returns the number of issues synced and an error if any.
	Sync(slug CharacterSlug) (int, error)
}

// A syncer to sync yearly appearances from Postgres to Redis.
type AppearancesSyncer struct {
	pgAppearanceRepository    *PGAppearancesByYearsRepository
	redisAppearanceRepository *RedisAppearancesByYearsRepository
}

// Gets all the character's appearances from the database and syncs them to Redis.
// returns the total number of issues synced and an error if any.
func (s *AppearancesSyncer) Sync(slug CharacterSlug) (int, error) {
	mainAppsPerYear, err := s.pgAppearanceRepository.Main(slug)
	if err != nil {
		return 0, err
	}
	if mainAppsPerYear.Aggregates != nil {
		log.COMIC().Info("main appearances to send to redis", zap.Int("total", mainAppsPerYear.Total()))
		err = s.redisAppearanceRepository.Set(mainAppsPerYear)
		if err != nil {
			return 0, err
		}
	}
	altAppsPerYear, err := s.pgAppearanceRepository.Alternate(slug)
	if err != nil {
		return 0, err
	}
	if altAppsPerYear.Aggregates != nil {
		log.COMIC().Info("alt appearances to send to redis", zap.Int("total", altAppsPerYear.Total()))
		s.redisAppearanceRepository.Set(altAppsPerYear)
	}
	all, err := s.pgAppearanceRepository.Both(slug)
	total := all.Total()
	log.COMIC().Info(
		"done syncing postgres appearances to redis!",
		zap.String("character", string(slug)),
		zap.String("appearances", fmt.Sprintf("%v", all.Aggregates)),
		zap.Int("total", total),
		zap.Error(err))
	return total, nil
}

// Returns a new appearances syncer
func NewAppearancesSyncer(container *PGRepositoryContainer, redisRepository *RedisAppearancesByYearsRepository) Syncer {
	return &AppearancesSyncer{
		pgAppearanceRepository:    container.AppearancesByYearsRepository(),
		redisAppearanceRepository: redisRepository,
	}
}

