package reader

import "github.com/hugocrmx/go-bank-practice/data"

type Reader interface {
	ReadFromFile(filePath string) (users []data.User, err error)
	ReadFromString(content string) (users []data.User, err error)
}
