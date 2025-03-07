package utils

import (

	"golang.org/x/crypto/bcrypt"
)



// Hasher un pass
func HashPassword(password string) string {
	pw := []byte(password) // Bcrypt a besoin de tableaux de bytes pour fonctionner
	result, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	return string(result) 
// Votre résultat ressemblera à genre : $2y$10$umu1eHzL2bfMZf41QBOL.OamYiIyDc6LGZNXRuyu1c41x2vAZpSQm
}

// Pour comparer un hash de mot de passe de votre BDD au string d'un MDP user
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
// Si le code retourne une erreur, c'est que les mots de pase ne correspondent pas
}