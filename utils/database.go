package utils

import "gorm.io/gorm"

func HandleDbResponseError[dto any](result *gorm.DB, msg string, response *dto) (*dto, error) {
	if result.Error != nil {
		return nil, result.Error
	}

	return response, nil
}

func HandleDbArrayResponseError[dto any](result *gorm.DB, msg string, response []dto) ([]dto, error) {
	if result.Error != nil {
		return []dto{}, result.Error
	}

	return response, nil
}
