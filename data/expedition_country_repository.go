package data

import "github.com/jmoiron/sqlx"

func GetRandomExpeditionCountry(tx *sqlx.Tx) (country ExpeditionCountry, err error) {
	err = tx.Get(&country, "SELECT id as country_id, code as country_code, name as country_name FROM expedition_countries LIMIT 1")
	return
}

func FindExpeditionById(tx *sqlx.Tx, countryId int) (country ExpeditionCountry, err error) {
	err = tx.Get(&country, "SELECT id as country_id, code as country_code, name as country_name FROM expedition_countries WHERE id = $1", countryId)
	return
}
