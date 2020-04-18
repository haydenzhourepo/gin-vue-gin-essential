package repository

import (
	"github.com/jinzhu/gorm"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
)

type ICategoryRepository interface {
	Create(name string)	(*model.Category, error)
	Update(category model.Category, name string) (*model.Category, error)
	SelectById(id int) (*model.Category, error)
	DeleteById(id int) error
}

type CategoryRepository struct {
	DB *gorm.DB		
}

func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err !=nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name:      name,
	}

	if err := c.DB.Create(&category).Error; err !=nil {
		return nil, err
	}

	return &category, nil
}

func NewCategoryRepository() ICategoryRepository  {
	return CategoryRepository{DB:common.GetDB()}
}





