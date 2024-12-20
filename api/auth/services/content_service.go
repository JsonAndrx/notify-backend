package services

import (
	"notify-backend/api/auth/repositories"
	"notify-backend/api/auth/types"
)

func ListCountriesService() ([]types.Country, error) {
	countries, err := repositories.ListCountries()
	if err != nil {
		return nil, err
	}

	return countries, nil
}
