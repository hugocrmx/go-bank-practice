package creditcard

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hugocrmx/go-bank-practice/data"
	"github.com/jmoiron/sqlx"
)

type creditCardTesting struct {
	name string
	args struct{ id int }
	want string
}

var (
	creditCardData []creditCardTesting
	tx             *sqlx.Tx
)

func init() {
	creditCardData = []creditCardTesting{
		{
			name: "testing Banca Bunsan",
			args: struct{ id int }{id: 1},
			want: "1 11 Banca Bunsan",
		},
		{
			name: "testing Banca Bunsan",
			args: struct{ id int }{id: 2},
			want: "2 22 Go Bank",
		},
		{
			name: "testing Banca Bunsan",
			args: struct{ id int }{id: 3},
			want: "3 33 Aldo Institución Financiera",
		},
	}
	tx = data.DB.MustBegin()
}

func TestBanks(t *testing.T) {

	for _, creditCardTest := range creditCardData {
		creditcard, err := data.FindCreditCardById(tx, creditCardTest.args.id)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		wantData := strings.Fields(creditCardTest.want)
		wantId, _ := strconv.Atoi(wantData[0])
		if creditcard.Id != wantId {
			t.Errorf("FindBankById :: Got: %v, Want:%v\n", creditcard, creditCardTest.want)
		}
	}
}
