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

func GetTimeZoneService(countryType types.GetTimezoneByCountryIdRequest) (types.TimeZone, error) {
	timeZones, err := repositories.GestTimezoneByCountryId(countryType.CountryId)
	if err != nil {
		return types.TimeZone{}, err
	}

	return timeZones, nil
}