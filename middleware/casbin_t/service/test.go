package service

import (
	"gorm.io/gorm"

	"gotest/middleware/casbin_t/models"
)

type TestInput struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TestService struct{ db *gorm.DB }

func NewTestService(db *gorm.DB) *TestService { return &TestService{db: db} }

func (s *TestService) Create(input TestInput, owner string) (*models.Test, error) {
	test := &models.Test{Name: input.Name, Content: input.Content, Owner: owner}
	return test, s.db.Create(test).Error
}

func (s *TestService) List() ([]models.Test, error) {
	var tests []models.Test
	return tests, s.db.Order("id DESC").Find(&tests).Error
}

func (s *TestService) Get(id uint) (*models.Test, error) {
	var test models.Test
	if err := s.db.First(&test, id).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (s *TestService) Update(id uint, input TestInput) (*models.Test, error) {
	test, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	test.Name, test.Content = input.Name, input.Content
	return test, s.db.Save(test).Error
}

func (s *TestService) Delete(id uint) error {
	result := s.db.Delete(&models.Test{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
