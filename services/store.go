package services

import (
	"context"
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

func (s *Store) get(email string) (claims *models.Claims, err error) {
	claims = &models.Claims{}
	db := s.pg.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.QueryOneContext(ctx, claims, `select * from public.get_member_claims(?)`, email)
	if err != nil {
		return nil, err
	}
	return
}

func (s *Store) Get(email string) (claims *models.Claims, err error) {
	v, err := s.LazyMap.Get(email, func() (interface{}, error) {
		return s.get(email)
	})
	if err != nil {
		return nil, err
	}
	return v.(*models.Claims), nil
}
