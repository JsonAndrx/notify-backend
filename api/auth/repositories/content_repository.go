package repositories

import (
	"notify-backend/api/auth/types"
	"notify-backend/api/utils/debug"
	"notify-backend/config/database"
)

func ListCountries() ([]types.Country, error) {
	client, err := database.GetDBConnection()
	if err != nil {
		return nil, err
	}

	rows, err := client.Query("SELECT id, name FROM countries")
	if err != nil {
		debug.LogError(err)
		return nil, err
	}

	var countries []types.Country
	for rows.Next() {
		var country types.Country
		err = rows.Scan(&country.Id, &country.CountryName)
		if err != nil {
			debug.LogError(err)
			return nil, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func GestTimezoneByCountryId(countryId int) (types.TimeZone, error) {
    client, err := database.GetDBConnection()
    if err != nil {
        return types.TimeZone{}, err
    }

    row := client.QueryRow("SELECT id, name FROM timezones WHERE country_id = $1", countryId)

    var timeZone types.TimeZone
    err = row.Scan(&timeZone.Id, &timeZone.TimeZone)
    if err != nil {
        debug.LogError(err)
        return types.TimeZone{}, err
    }

    return timeZone, nil
}
