package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type countryRepository struct {
	BaseRepository
}

func NewCountryRepository(db *gorm.DB) repository.CountryRepository {
	return &countryRepository{
		BaseRepository: NewBaseRepository(db, "country"),
	}
}

func (r *countryRepository) Create(country *entity.Country) (*entity.Country, error) {
	err := r.db.Create(country).Error
	if err != nil {
		return nil, err
	}

	// Cache the newly created country
	_ = r.cacheDetail("byid", country, country.ID)
	_ = r.cacheDetail("byname", country, country.Name)
	_ = r.cacheDetail("byshortcode", country, country.ShortCode)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return country, nil
}

func (r *countryRepository) GetByID(id uint) (*entity.Country, error) {
	var country entity.Country
	err := r.getCachedDetail("byid", &country, func() error {
		return r.db.First(&country, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) GetByName(name string) (*entity.Country, error) {
	var country entity.Country
	err := r.getCachedDetail("byname", &country, func() error {
		return r.db.Where("name = ?", name).First(&country).Error
	}, name)
	
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) GetByShortCode(shortCode string) (*entity.Country, error) {
	var country entity.Country
	err := r.getCachedDetail("byshortcode", &country, func() error {
		return r.db.Where("short_code = ?", shortCode).First(&country).Error
	}, shortCode)
	
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) List() ([]*entity.Country, error) {
	var countries []*entity.Country
	err := r.getCachedList("list", &countries, func() error {
		return r.db.Find(&countries).Error
	})
	
	return countries, err
}

func (r *countryRepository) Update(country *entity.Country) (*entity.Country, error) {
	err := r.db.Save(country).Error
	if err != nil {
		return nil, err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", country, country.ID)
	_ = r.cacheDetail("byname", country, country.Name)
	_ = r.cacheDetail("byshortcode", country, country.ShortCode)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return country, nil
}

func (r *countryRepository) Delete(id uint) error {
	var country entity.Country
	if err := r.db.First(&country, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&country).Error; err != nil {
		return err
	}

	// Invalidate all caches for this country
	r.invalidateAllCache()
	
	return nil
}
