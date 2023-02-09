package main

import (
  "github.com/theGOURL/encrypting"
)

func main(){
  password := "haggafhu1837uhud"
  generatePasswordEncrypted :=  encrypting.Encrypt(password,10);
  encrypting.ValidatePassword(password,generatePasswordEncrypted);
}
