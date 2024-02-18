package encrypt

import "golang.org/x/crypto/bcrypt"

func Hash(txt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	return string(bytes), err
}

func Check(requestedTxt, storedTxt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedTxt), []byte(requestedTxt))
	return err == nil
}
