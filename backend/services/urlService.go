package services

import (
	"strconv"
	"strings"
)

func ExtractPostIDFromURL(path string) (uint64, error) {
	postIDStr := strings.TrimPrefix(path, "/user/post/")
	postIDStr = strings.TrimSuffix(postIDStr, "/like")
	return strconv.ParseUint(postIDStr, 10, 64)
}

func ExtractCommentIDFromURL(path string) (uint64, error) {
	commentIDStr := strings.TrimPrefix(path, "/user/comment/")
	return strconv.ParseUint(commentIDStr, 10, 64)
}