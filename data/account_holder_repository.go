package data

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetAllAccountHolders(tx *sqlx.Tx) (accounts []*AccountHolder, err error) {
	err = tx.Select(&accounts, "SELECT id as ac_id, name as ac_name, email as ac_email,  alias as ac_alias FROM  account_holders  ORDER BY RANDOM()")
	return
}

func GetRandomAccount(tx *sqlx.Tx) (account AccountHolder, err error) {
	tx.Get(&account, "SELECT id as ac_id, name as ac_name, email as ac_email,  alias as ac_alias FROM  account_holders  ORDER BY RANDOM() LIMIT 1")
	return
}

func AddAccountHolder(tx *sqlx.Tx, accountHolder *AccountHolder) (err error) {
	log.Println("New accountHolder:", accountHolder)
	_, err = tx.NamedExec("INSERT INTO account_holders (name, email, alias) VALUES (:ac_name, :ac_email, :ac_alias)", accountHolder)
	return
}
