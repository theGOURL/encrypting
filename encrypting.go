package encrypting

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "github.com/theGOURL/warning"
)

//This function generates encrypted passwords through the password received by parameter
func Encrypt(password string, hashcode int) []byte{
  encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashcode);
	  warning.PRINT_DEFAULT_ERRORS(err,"unable to generate encryption for this password");

	fmt.Println("Encrypted Password:",string(encryptedPassword));
  return encryptedPassword;
}

//This function receives the encrypted password and the password validating if it is the same password
func ValidatePassword(password string, encryptedPassword []byte){
  validatePassword := bcrypt.CompareHashAndPassword(encryptedPassword, []byte(password));
    if validatePassword != nil{
    fmt.Println("Incorrect Password");
  }
  fmt.Println("Aproved");
}
