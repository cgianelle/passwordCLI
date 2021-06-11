package main

import (
	"flag"
	"fmt"

	"github.com/cgianelle/password"
)

func main() {
	fmt.Println("Password Command Line Tool")
	number_of_characters := flag.Int("characters", 10, "number of password characters")
	flag.Parse()

	fmt.Println("numb:", *number_of_characters)

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

	passBldr := password.Builder(*number_of_characters)

	for i := 0; i < 10; i++ {
		myPassword := passBldr(lowercaseAlpha, uppercaseAlpha, numericCharacters, specialCharacters)
		fmt.Println(myPassword)
	}
}
