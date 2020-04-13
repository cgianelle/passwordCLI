package main

import (
	"fmt"

	"github.com/cgianelle/password"
)

func main() {
	fmt.Println("Password Command Line Tool")

	lowercaseAlpha := password.Password{
		Characters: "abcdefghijklmnopqrstuvwxyz",
	}

	uppercaseAlpha := password.Password{
		Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	numericCharacters := password.Password{
		Characters: "0123456789",
	}

	specialCharacters := password.Password{
		Characters: "!@#$%^&*(){}[]\\/?,.<>~`",
	}

	passBldr := password.PasswordBuilder(10)

	for i := 0; i < 10; i++ {
		myPassword := passBldr(lowercaseAlpha, uppercaseAlpha, numericCharacters, specialCharacters)
		fmt.Println(myPassword)
	}
}
