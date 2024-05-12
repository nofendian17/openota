package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/infra/database"
)

type Repository interface {
	GetByID(ctx context.Context, ID string) (*entity.Airport, error)
	GetByCode(ctx context.Context, code string) (*entity.Airport, error)
	GetAll(ctx context.Context) ([]*entity.Airport, error)

	Create(ctx context.Context, airport *entity.Airport) error
	Update(ctx context.Context, ID string, airport *entity.Airport) error
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

func (r *repository) GetByID(ctx context.Context, ID string) (*entity.Airport, error) {
	var airport entity.Airport
	if err := r.db.GormDB.WithContext(ctx).First(&airport, "id = ?", ID).Error; err != nil {

		return nil, err // Other errors
	}
	return &airport, nil
}

func (r *repository) GetByCode(ctx context.Context, code string) (*entity.Airport, error) {
	var airport entity.Airport
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		First(&airport, "code = ?", code).Error; err != nil {

		return nil, err // Other errors
	}
	return &airport, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Airport, error) {
	var airports []*entity.Airport
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		Order("precedence asc").
		Find(&airports).Error; err != nil {
		return nil, err
	}
	return airports, nil
}

func (r *repository) Create(ctx context.Context, airport *entity.Airport) error {
	return r.db.GormDB.WithContext(ctx).Create(airport).Error
}

func (r *repository) Update(ctx context.Context, ID string, airport *entity.Airport) error {
	return r.db.GormDB.WithContext(ctx).Model(&entity.Airport{}).Where("id = ?", ID).Updates(airport).Error
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	return r.db.GormDB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Airport{}).Error
}
