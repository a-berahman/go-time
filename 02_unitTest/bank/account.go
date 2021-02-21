package bank

import (
	"fmt"
	"strings"
)

type Banker interface {
	GetAccountNumber(string) string
}

type Saman struct{}

func NewSaman() *Saman {
	return &Saman{}
}

// GetAccountNumber gets account number from IBAN
func (s *Saman) GetAccountNumber(iban string) string {

	return fmt.Sprintf("%s-%s-%s-%s", firstPart(iban), secondPart(iban), thirdPart(iban), fourthPart(iban))
}

func firstPart(iban string) (res string) {
	if iban[8] == '0' {
		res = iban[9:12]
	} else {
		res = iban[8:12]
	}
	return res
}
func secondPart(iban string) (res string) {
	if iban[12] == '0' {
		res = iban[13:15]

	} else {
		res = iban[12:15]
	}
	return res
}
func thirdPart(iban string) string {
	return strings.Replace(iban[15:23], "0", "", -1)
}
func fourthPart(iban string) string {
	return iban[25:26]
}
