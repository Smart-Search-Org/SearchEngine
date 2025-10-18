package user_index_service

import "SmartSearch/internal/repository/user_index"

func IsUserHasIndex(userId string, indexName string) (bool, error) {
	_, err := user_index.FindByUserAndIndex(userId, indexName)
	if err != nil && err.Error() == "user index not found" {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
