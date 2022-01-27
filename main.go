package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/api"
	"github.com/hugocrmx/go-bank-practice/data"
	"github.com/hugocrmx/go-bank-practice/reader"
)

const (
	csvFilePath = "./resources/users.csv"
)

var (
	g *gin.Engine
)

func init() {
	g = gin.Default()
	populateAccountHolders()

}

func main() {
	api.CreateCreditCardRoutes(g)
	api.CreateBankRoutes(g)
	api.CreateProductSegmentRoutes(g)
	api.CreateAccountHolderRoutes(g)
	api.CreateExpeditionCountryRoutes(g)
	g.Run()
}

func populateAccountHolders() {
	tx := data.DB.MustBegin()
	accounts, err := data.GetAllAccountHolders(tx)
	if err != nil {
		log.Printf("populateAccountHolders() error:%v\n", err.Error())
	}

	if len(accounts) == 0 {
		var reader reader.Reader = &reader.CSVReader{}
		users, err := reader.ReadFromFile(csvFilePath)
		if err != nil {
			log.Printf("ReadFromFile(csvFilePath) error: %v\n", err)
		}

		for _, v := range users {
			data.AddAccountHolder(tx, &data.AccountHolder{Name: v.Name, Email: v.Email, Alias: v.Alias})
		}
		tx.Commit()
	}
}
