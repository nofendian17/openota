package city

import (
	"context"
	"github.com/nofendian17/openota/apigw/internal/infra/database"

	"github.com/nofendian17/openota/apigw/internal/entity"
)

type Repository interface {
	GetByID(ctx context.Context, ID string) (*entity.City, error)
	GetByStateID(ctx context.Context, StateID string) ([]*entity.City, error)
	GetAll(ctx context.Context) ([]*entity.City, error)

	Create(ctx context.Context, city *entity.City) error
	Update(ctx context.Context, ID string, city *entity.City) error
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

func (r *repository) GetByID(ctx context.Context, ID string) (*entity.City, error) {
	var city entity.City
	if err := r.db.GormDB.WithContext(ctx).First(&city, "id = ?", ID).Error; err != nil {

		return nil, err // Other errors
	}
	return &city, nil
}

func (r *repository) GetByStateID(ctx context.Context, StateID string) ([]*entity.City, error) {
	var cities []*entity.City
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where("state_id = ?", StateID).
		Where(condition).
		Order("precedence asc").
		Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.City, error) {
	var cities []*entity.City
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		Order("precedence asc").
		Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *repository) Create(ctx context.Context, city *entity.City) error {
	return r.db.GormDB.WithContext(ctx).Create(city).Error
}

func (r *repository) Update(ctx context.Context, ID string, city *entity.City) error {
	return r.db.GormDB.WithContext(ctx).Model(&entity.City{}).Where("id = ?", ID).Updates(city).Error
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	return r.db.GormDB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.City{}).Error
}
