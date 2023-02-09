package encrypting

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "github.com/theGOURL/warning"
)

func Encrypt(password string, hashcode int) []byte{
  encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashcode);
	  warning.PRINT_DEFAULT_ERRORS(err,"unable to generate encryption for this password");

	fmt.Println("Encrypted Password:",string(encryptedPassword));
  return encryptedPassword;
}

func ValidatePassword(password string, encryptedPassword []byte){
  validatePassword := bcrypt.CompareHashAndPassword(encryptedPassword, []byte(password));
    if validatePassword != nil{
    fmt.Println("Incorrect Password");
  }
  fmt.Println("Aproved");
}
