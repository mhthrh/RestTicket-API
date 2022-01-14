package Utilitys

import (
	"net/mail"
	"regexp"
	"unicode"
)

func CheckPassword(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func CheckMail(address string) bool {
	if _, err := mail.ParseAddress(address); err != nil {
		return false
	}
	return true
}

func CheckPhoneNumber(CellNo string) bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(CellNo)
}

func CheckName(name string) bool {
	if name == "" {
		return false
	}
	if len(name) < 2 || len(name) > 40 {
		return false
	}
	return true
}
