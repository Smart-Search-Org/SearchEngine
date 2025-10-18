package user_index

import (
	"SmartSearch/internal/models/database"
	"SmartSearch/internal/utils"
	"errors"

	"gorm.io/gorm"
)

func FindByUserAndIndex(userID string, indexName string) (*database.UserIndex, error) {
	var userIndex database.UserIndex

	err := utils.DB.
		Where("user_id = ? AND index_name = ?", userID, indexName).
		First(&userIndex).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user index not found")
	}
	if err != nil {
		return nil, err
	}

	return &userIndex, nil
}

func CreateUserIndex(userID string, indexName string) (*database.UserIndex, error) {
	userIndex := &database.UserIndex{
		UserId:    userID,
		IndexName: indexName,
	}

	if err := utils.DB.Create(userIndex).Error; err != nil {
		return nil, err
	}

	return userIndex, nil
}
