package encrypting

func ExampleEncrypt(){
   password := "haggafhu1837uhud";
   generatePasswordEncrypted :=  []byte("$2a$10$r4zkLfYMLY5meIBlRLTT2eEeeA0fVwZUTqZ3HibVSgSxFSS2UImGK");
     ValidatePassword(password,generatePasswordEncrypted);
  //Output:
  //Aproved
}

func ExampleValidatePassword(){
   password := "haggafhu1837uhud";
   generatePasswordEncrypted :=  []byte("$2a$10$r4zkLfYMLY5meIBlRLTT2eEeeA0fVwZUTqZ3HibVSgSxFSS2UImGK");
     ValidatePassword(password,generatePasswordEncrypted);
  //Output:
  //Aproved
}
