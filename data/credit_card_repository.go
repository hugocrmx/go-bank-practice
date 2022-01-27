package data

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	creditCardSelectBase = `
	SELECT
		card.ID,
		card.card_number,
		ac.ID AS ac_id,
		ac.NAME AS ac_name,
		ac.email AS ac_email,
		ac.ALIAS AS ac_alias,
		bank.ID AS bank_id,
		bank.code AS bank_code,
		bank.NAME AS bank_name,
		segment.ID AS segment_id,
		segment.code AS segment_code,
		segment.NAME AS segment_name,
		country.ID AS country_id,
		country.code AS country_code,
		country.NAME AS country_name 
	FROM
		credit_cards AS card
		JOIN account_holders AS ac ON ac.ID = card.id_account_holder
		JOIN bank_codes AS bank ON bank.ID = card.id_bank_code
		JOIN product_segments AS segment ON segment.ID = card.id_product_segment
		JOIN expedition_countries AS country ON country.ID = card.id_expedition_country 
	`
	orderById = `ORDER BY card.id`
	whereById = `WHERE card.id = `
)

func AddCreditCard(tx *sqlx.Tx, card *CreditCard) (id int64, err error) {
	card, err = GenerateCreditCardNumber(tx, card)
	row, err := tx.PrepareNamed(`
	INSERT
		INTO credit_cards
			(card_number, id_account_holder, id_bank_code, id_product_segment, id_expedition_country)
		VALUES
			(:card_number, :ac_id, :bank_id, :segment_id, :country_id)  RETURNING id`)

	err = row.Get(&id, card)

	return
}

func GetAllCreditCards(tx *sqlx.Tx) (creditCards []*CreditCard, err error) {
	err = tx.Select(&creditCards, `
	SELECT
		card.ID,
		card.card_number,
		ac.ID AS ac_id,
		ac.NAME AS ac_name,
		ac.email AS ac_email,
		ac.ALIAS AS ac_alias,
		bank.ID AS bank_id,
		bank.code AS bank_code,
		bank.NAME AS bank_name,
		segment.ID AS segment_id,
		segment.code AS segment_code,
		segment.NAME AS segment_name,
		country.ID AS country_id,
		country.code AS country_code,
		country.NAME AS country_name 
	FROM
		credit_cards AS card
		JOIN account_holders AS ac ON ac.ID = card.id_account_holder
		JOIN bank_codes AS bank ON bank.ID = card.id_bank_code
		JOIN product_segments AS segment ON segment.ID = card.id_product_segment
		JOIN expedition_countries AS country ON country.ID = card.id_expedition_country 
	`)
	return
}

func UpdateCreditCard(tx *sqlx.Tx, creditCard *CreditCard) (err error) {
	_, err = tx.NamedExec("UPDATE credit_cards SET card_number = :card_number WHERE id = :id", creditCard)
	return
}

func FindCreditCardById(tx *sqlx.Tx, id int) (creditCard CreditCard, err error) {
	tx.Get(&creditCard, creditCardSelectBase+whereById+"$1 ", id)
	return
}

func GenerateCreditCardNumber(tx *sqlx.Tx, cc *CreditCard) (creditCard *CreditCard, err error) {
	var bank Bank
	var segment ProductSegment
	var country ExpeditionCountry
	var accNumber string
	creditCard = cc

	//Get bank
	bank, err = FindBankById(tx, cc.Bank.Id)
	if err != nil {
		return
	}

	//Get segment
	if segment, err = FindSegmentById(tx, cc.ProductSegment.Id); err != nil {
		return
	}

	//Get country
	if country, err = FindExpeditionById(tx, cc.ExpeditionCountry.Id); err != nil {
		return
	}

	//Get valid card number
	accountNumber := initAccountAndCheckSum()
	for i := 0; i < len(accountNumber); i++ {
		accNumber += strconv.Itoa(accountNumber[i])
	}

	creditCard.CardNumber = strconv.Itoa(bank.Code) + strconv.Itoa(segment.Code) + strconv.Itoa(country.Code) + accNumber

	return
}

func initAccountAndCheckSum() []int {
	accountNumber := make([]int, 9)
	rand.Seed(time.Now().UnixNano())

	for {
		for i := 0; i < 9; i++ {
			accountNumber[i] = rand.Intn(9)
		}

		acc := 0
		for i, j := 0, 8; i < 9; i, j = i+1, j-1 {
			acc += accountNumber[i] * (j)
		}
		if acc%11 == 0 {
			return accountNumber
		}
	}
}
