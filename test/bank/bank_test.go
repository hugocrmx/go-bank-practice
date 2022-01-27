package bank

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hugocrmx/go-bank-practice/data"
	"github.com/jmoiron/sqlx"
)

type bankTesting struct {
	name string
	args struct{ id int }
	want string
}

var (
	bankData []bankTesting
	tx       *sqlx.Tx
)

func init() {
	bankData = []bankTesting{
		{
			name: "testing Banca Bunsan",
			args: struct{ id int }{id: 1},
			want: "1 11 Banca Bunsan",
		},
		{
			name: "testing Go Bank",
			args: struct{ id int }{id: 2},
			want: "2 22 Go Bank",
		},
		{
			name: "testing Aldo Institución Financiera",
			args: struct{ id int }{id: 3},
			want: "3 33 Aldo Institución Financiera",
		},
	}
	tx = data.DB.MustBegin()
}

func TestBanks(t *testing.T) {

	for _, bankTest := range bankData {
		bank, err := data.FindBankById(tx, bankTest.args.id)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		wantData := strings.Fields(bankTest.want)
		wantId, _ := strconv.Atoi(wantData[0])
		if bank.Id != wantId {
			t.Errorf("FindBankById:: %s Got: %v, Want:%v\n", bankTest.name, bank, bankTest.want)
		}
	}
}
