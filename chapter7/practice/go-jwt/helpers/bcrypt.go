package helpers

import "golang.org/x/crypto/bcrypt"

// Melakukan hashing passworduser sebelum disimpan kedalam database
func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

/*
Digunakan untuk mengkomparasi password user yang sudah dihash dengan
password user yang di input ketika sedang melakukan login.
*/
func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
