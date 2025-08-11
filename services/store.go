package services

import (
	"context"
	"errors"
	"time"

	"github.com/webtor-io/claims-provider/models"
	cs "github.com/webtor-io/common-services"
	"github.com/webtor-io/lazymap"
)

type Store struct {
	lazymap.LazyMap
	pg *cs.PG
}

func NewStore(pg *cs.PG) *Store {
	return &Store{
		pg: pg,
		LazyMap: lazymap.New(&lazymap.Config{
			Concurrency: 10,
			Expire:      60 * time.Second,
			ErrorExpire: 10 * time.Second,
			Capacity:    1000,
		}),
	}
}

func (s *Store) get(email string, patreonId string) (claims *models.Claims, err error) {
	claims = &models.Claims{}
	db := s.pg.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// If patreon_id is provided, use it for lookup (prioritize patreon_id over email)
	if patreonId != "" {
		_, err = db.QueryOneContext(ctx, claims, `select * from public.get_member_claims_by_patreon_id(?)`, patreonId)
	} else if email != "" {
		_, err = db.QueryOneContext(ctx, claims, `select * from public.get_member_claims(?)`, email)
	} else {
		return nil, errors.New("either email or patreon_id must be provided")
	}
	
	if err != nil {
		return nil, err
	}
	return
}

func (s *Store) Get(email string, patreonId string) (claims *models.Claims, err error) {
	// Create cache key based on which identifier is provided
	var cacheKey string
	if patreonId != "" {
		cacheKey = "patreon:" + patreonId
	} else {
		cacheKey = "email:" + email
	}
	
	v, err := s.LazyMap.Get(cacheKey, func() (interface{}, error) {
		return s.get(email, patreonId)
	})
	if err != nil {
		return nil, err
	}
	return v.(*models.Claims), nil
}
