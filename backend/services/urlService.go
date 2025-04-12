package services

import (
	"errors"
	"strconv"
	"strings"
)
// /backend/services/likeService.go						/backend/services/commentService.go
// func UserAddComment; UserEditComment; UserDeleteComment ;ToggleLikePost ;ToggleLikeComment 
// recuper une info de type ID à partir de l'url
func ExtractIDFromURL(path string) (uint64, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, errors.New("URL invalide")
	}
	return strconv.ParseUint(parts[2], 10, 64)
}

