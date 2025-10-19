package user_database_service

import "SmartSearch/internal/repository/user_index"

func DeleteIndex(userId string, indexName string) error {
	err := user_index.DeleteUserIndex(userId, indexName)
	if err != nil {
		return err
	}
	return nil
}
