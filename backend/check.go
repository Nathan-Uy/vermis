package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Paste your stored hash here
	hashedPassword := "$2a$10$VknV9eFdLndWPCBl7Vm13eEwQ2yBYNETuc.S90tAZW6ET5xmYxBYe" // complete the hash
	plainPassword := "admin"

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		fmt.Println("❌ Password does NOT match!")
	} else {
		fmt.Println("✅ Password matches!")
	}
}
