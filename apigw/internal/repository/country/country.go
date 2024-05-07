package country

import (
	"context"
	"github.com/nofendian17/openota/apigw/internal/infra/database"

	"github.com/nofendian17/openota/apigw/internal/entity"
)

type Repository interface {
	GetByID(ctx context.Context, ID string) (*entity.Country, error)
	GetByCode(ctx context.Context, code string) (*entity.Country, error)
	GetAll(ctx context.Context) ([]*entity.Country, error)

	Create(ctx context.Context, country entity.Country) error
	Update(ctx context.Context, ID string, country entity.Country) error
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

func (r *repository) GetByID(ctx context.Context, ID string) (*entity.Country, error) {
	var country entity.Country
	if err := r.db.GormDB.WithContext(ctx).First(&country, "id = ?", ID).Error; err != nil {
		return nil, err // Other errors
	}
	return &country, nil
}

func (r *repository) GetByCode(ctx context.Context, code string) (*entity.Country, error) {
	var country entity.Country
	if err := r.db.GormDB.WithContext(ctx).First(&country, "code = ?", code).Error; err != nil {

		return nil, err // Other errors
	}
	return &country, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Country, error) {
	var countries []*entity.Country
	if err := r.db.GormDB.WithContext(ctx).Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

func (r *repository) Create(ctx context.Context, country entity.Country) error {
	return r.db.GormDB.WithContext(ctx).Create(&country).Error
}

func (r *repository) Update(ctx context.Context, ID string, country entity.Country) error {
	return r.db.GormDB.WithContext(ctx).Model(&entity.Country{}).Where("id = ?", ID).Updates(&country).Error
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	return r.db.GormDB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Country{}).Error
}
