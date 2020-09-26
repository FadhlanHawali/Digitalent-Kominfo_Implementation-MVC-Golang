package utils

import "golang.org/x/crypto/bcrypt"

func HashGenerator(str string) (string, error) {
	hashedString,err := bcrypt.GenerateFromPassword([]byte(str),bcrypt.DefaultCost);if err!=nil{
		return "", err
	}

	return string(hashedString),nil
}

func HashComparator(hashedByte []byte,byte []byte) (error){
	err := bcrypt.CompareHashAndPassword(hashedByte,byte); if err != nil{
		return err
	}

	return nil
}
