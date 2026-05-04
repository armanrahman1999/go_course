package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "12345"

	hash := sha256.Sum256([]byte(password))

	fmt.Println(hash)
	fmt.Printf("something %x\n", hash)
	salt, err := generateSalt()
	fmt.Println(salt, err) 
	hashedPassword := hashPassword(password, salt)
	fmt.Println(hashedPassword)
}

func generateSalt() ([]byte, error){ 
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string{

	saltedPassword := append(salt, []byte (password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])

}