package data

import (
	"github.com/jmoiron/sqlx"
)

func GetRandomBank(tx *sqlx.Tx) (bank Bank, err error) {
	err = tx.Get(&bank, "SELECT id as bank_id, code as bank_code, name as bank_name FROM bank_codes ORDER BY RANDOM() LIMIT 1")
	return
}

func GetAllBanks(tx *sqlx.Tx) (banks []*Bank, err error) {
	err = tx.Select(&banks, "SELECT id as bank_id, code as bank_code, name as bank_name FROM bank_codes ORDER BY RANDOM()")
	return
}

func AddBank(tx *sqlx.Tx, bank *Bank) (err error) {
	_, err = tx.NamedExec("INSERT INTO bank_codes (code, name) VALUES (:code, :name)", bank)
	return
}

func UpdateBank(tx *sqlx.Tx, bank *Bank) (err error) {
	_, err = tx.NamedExec("UPDATE bank_codes SET code = :code, name = :name WHERE id = :id", bank)
	return
}

func FindBankById(tx *sqlx.Tx, bankId int) (bank Bank, err error) {
	err = tx.Get(&bank, "SELECT id as bank_id, code as bank_code, name as bank_name FROM bank_codes WHERE id = $1", bankId)
	return
}
