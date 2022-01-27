package data

import "fmt"

type Bank struct {
	Id   int    `db:"bank_id" json:"id" binding:"required"`
	Code int    `db:"bank_code"`
	Name string `db:"bank_name"`
}

type ProductSegment struct {
	Id   int    `db:"segment_id" json:"id" binding:"required"`
	Code int    `db:"segment_code"`
	Name string `db:"segment_name"`
}

type ExpeditionCountry struct {
	Id   int    `db:"country_id" json:"id" binding:"required"`
	Code int    `db:"country_code"`
	Name string `db:"country_name"`
}

type AccountHolder struct {
	Id    int    `db:"ac_id" json:"id" binding:"required"`
	Name  string `db:"ac_name"`
	Email string `db:"ac_email"`
	Alias string `db:"ac_alias"`
}

type CreditCard struct {
	Id                int    `db:"id" json:"id"`
	CardNumber        string `db:"card_number" json:"card-number"`
	AccountHolder     `json:"account-holder"`
	Bank              `json:"bank"`
	ProductSegment    `json:"product-segment"`
	ExpeditionCountry `json:"expedition-country"`
}

func (b *Bank) String() string {
	return fmt.Sprintf("Bank#String{Id: %d, Code: %d, Name: %s}", b.Id, b.Code, b.Name)
}

func (c *CreditCard) String() string {
	return fmt.Sprintf("CreditCard#String{Id: %d, CardNumber: %s, Bank: %v}", c.Id, c.CardNumber, c.CardNumber)
}

func (ps *ProductSegment) String() string {
	return fmt.Sprintf("ProductSegment#String{Id: %d, Code: %d, Name: %s}", ps.Id, ps.Code, ps.Name)
}
