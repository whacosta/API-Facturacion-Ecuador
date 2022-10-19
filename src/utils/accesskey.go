package utils

import "strconv"

type AccessKey struct {
	EmissionDate string
	VoucherType  string
	RUC          string
	Environment  string
	Serie        string
	Sequencial   string
	NumericCode  string
	EmissionType string
	CheckDigit   string
}

func (a *AccessKey) GenerateAccessKey() string {
	return a.EmissionDate + a.VoucherType + a.RUC + a.Environment + a.Serie + a.Sequencial + a.NumericCode + a.EmissionType + a.GetCheckDigit()
}

func (a *AccessKey) GetCheckDigit() string {
	digit1 := parseToInt(a.EmissionDate) * 3
	digit2 := parseToInt(a.VoucherType) * 2
	digit3 := parseToInt(a.RUC) * 7
	digit4 := parseToInt(a.Environment) * 6
	digit5 := parseToInt(a.EmissionDate) * 5
	digit6 := parseToInt(a.EmissionDate) * 4
	digit7 := parseToInt(a.EmissionDate) * 3
	digit8 := parseToInt(a.EmissionDate) * 2

	result := 11 - ((digit1 + digit2 + digit3 + digit4 + digit5 + digit6 + digit7 + digit8) % 11)
	if result == 11 {
		result = 0
	}

	if result == 10 {
		result = 1
	}

	return strconv.Itoa(result)
}

func parseToInt(key string) int {
	digit, err := strconv.Atoi(key)
	if err != nil {
		panic(err)
	}
	return digit
}
