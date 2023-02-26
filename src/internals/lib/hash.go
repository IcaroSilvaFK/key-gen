package lib

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(s string) (string, error) {

	byt, err := bcrypt.GenerateFromPassword([]byte(s), 12)

	if err != nil {

		log.Println(err)

		return "", err
	}

	return string(byt), nil
}
