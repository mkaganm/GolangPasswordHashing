package main

import (
	"GolangPasswordHashing/src/auth"
	"fmt"
)

func main() {

	password := "123456"

	hash, _ := auth.PasswordToHash(password)

	match := auth.CheckPassword(password, hash)

	fmt.Println("Password\t:", password)
	fmt.Println("Hash\t\t:", hash)
	fmt.Println("Match\t\t:", match)
}
