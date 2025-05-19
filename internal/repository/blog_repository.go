package repository

import (
	"api-blog-go/internal/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) FindAll(selects []string) ([]models.BlogList, error) {
	var blogs []models.BlogList
	err := r.db.Model(&models.Blog{}).Select(selects).Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) FindByID(id uint) (*models.Blog, error) {
	var blog models.Blog
	err := r.db.First(&blog, id).Error
	return &blog, err
}

func (r *BlogRepository) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) Delete(id uint) error {
	return r.db.Delete(&models.Blog{}, id).Error
}
