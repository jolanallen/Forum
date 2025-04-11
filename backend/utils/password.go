package utils

import (

	"golang.org/x/crypto/bcrypt"
)


//vérification du hash = hash du mot de passe (à voir si j'ai fait la logique pour hashé le mdp au départ)
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}