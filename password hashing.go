package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "Prathyusha123"
    hash, _ := HashPassword(password) 

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
output:
Password: Prathyusha123
Hash:     $2a$14$PMQLJjdbNLxRhk3EFalqVOoVFlHMwM7FGqvRZipG3VyQsKHqJTpuC
Match:    true

Program exited.