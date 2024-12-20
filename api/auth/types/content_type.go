package types

type Country struct {
	Id          int    `json:"id" db:"id"`
	CountryName string `json:"country_name" db:"name"`
}
