package phonenumber

import (
	"errors"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	reg := regexp.MustCompile("[^0-9]+")
	processedString := reg.ReplaceAllString(phoneNumber, "")
	if len(processedString) < 10 {
		return "", errors.New("incorrect number of digits")
	}
	if len(processedString) > 11 {
		return "", errors.New("more than 11 digits")
	}
	if len(processedString) == 11 {
		if processedString[0] != '1' {
			return "", errors.New("error, country code not one")
		}
		processedString = processedString[1:]
	}
	if processedString[0] == '1' || processedString[0] == '0' {
		return "", errors.New("error, first number is less than two")
	} else if processedString[3] == '0' || processedString[3] == '1' {
		return "", errors.New("error, first number is less than two")
	}
	return processedString, nil
}

func AreaCode(phoneNumber string) (string, error) {
	areaCode, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return areaCode[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	formattedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + formattedNumber[0:3] + ") " + formattedNumber[3:6] + "-" + formattedNumber[6:10], nil
}
