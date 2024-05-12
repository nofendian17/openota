package repository

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/infra/database"

	"github.com/nofendian17/openota/coresvc/internal/entity"
)

type Repository interface {
	GetByID(ctx context.Context, ID string) (*entity.State, error)
	GetByCountryID(ctx context.Context, countryID string) ([]*entity.State, error)
	GetAll(ctx context.Context) ([]*entity.State, error)

	Create(ctx context.Context, state *entity.State) error
	Update(ctx context.Context, ID string, state *entity.State) error
	Delete(ctx context.Context, ID string) error
}

type repository struct {
	db *database.DB
}

func New(db *database.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, ID string) (*entity.State, error) {
	var state entity.State
	if err := r.db.GormDB.WithContext(ctx).First(&state, "id = ?", ID).Error; err != nil {

		return nil, err // Other errors
	}
	return &state, nil
}

func (r *repository) GetByCountryID(ctx context.Context, countryID string) ([]*entity.State, error) {
	var states []*entity.State
	if err := r.db.GormDB.WithContext(ctx).Where("country_id = ?", countryID).Find(&states).Error; err != nil {
		return nil, err
	}
	return states, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.State, error) {
	var states []*entity.State
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		Order("precedence asc").
		Find(&states).Error; err != nil {
		return nil, err
	}
	return states, nil
}

func (r *repository) Create(ctx context.Context, state *entity.State) error {
	return r.db.GormDB.WithContext(ctx).Create(state).Error
}

func (r *repository) Update(ctx context.Context, ID string, state *entity.State) error {
	return r.db.GormDB.WithContext(ctx).Model(&entity.State{}).Where("id = ?", ID).Updates(state).Error
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	return r.db.GormDB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.State{}).Error
}
