package types

type Country struct {
	Id          int    `json:"id" db:"id"`
	CountryName string `json:"country_name" db:"name"`
}

type TimeZone struct {
	Id        int    `json:"id" db:"id"`
	TimeZone  string `json:"time_zone" db:"time_zone"`
}

type GetTimezoneByCountryIdRequest struct {
	CountryId int `json:"country_id"`
}