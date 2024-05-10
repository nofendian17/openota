package airline

import (
	"context"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"github.com/nofendian17/openota/apigw/internal/infra/database"
)

type Repository interface {
	GetByID(ctx context.Context, ID string) (*entity.Airline, error)
	GetByCode(ctx context.Context, code string) (*entity.Airline, error)
	GetAll(ctx context.Context) ([]*entity.Airline, error)

	Create(ctx context.Context, airline *entity.Airline) error
	Update(ctx context.Context, ID string, airport *entity.Airline) error
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

func (r *repository) GetByID(ctx context.Context, ID string) (*entity.Airline, error) {
	var airline entity.Airline
	if err := r.db.GormDB.WithContext(ctx).First(&airline, "id = ?", ID).Error; err != nil {

		return nil, err // Other errors
	}
	return &airline, nil
}

func (r *repository) GetByCode(ctx context.Context, code string) (*entity.Airline, error) {
	var airline entity.Airline
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		First(&airline, "code = ?", code).Error; err != nil {

		return nil, err // Other errors
	}
	return &airline, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Airline, error) {
	var airlines []*entity.Airline
	condition := map[string]interface{}{
		"is_active": true,
	}
	if err := r.db.GormDB.WithContext(ctx).
		Where(condition).
		Order("precedence asc").
		Find(&airlines).Error; err != nil {
		return nil, err
	}
	return airlines, nil
}

func (r *repository) Create(ctx context.Context, airline *entity.Airline) error {
	return r.db.GormDB.WithContext(ctx).Create(airline).Error
}

func (r *repository) Update(ctx context.Context, ID string, airline *entity.Airline) error {
	return r.db.GormDB.WithContext(ctx).Model(&entity.Airline{}).Where("id = ?", ID).Updates(airline).Error
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	return r.db.GormDB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Airline{}).Error
}
