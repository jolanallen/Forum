package services

import (
	"Forum/backend/db"
	"errors"
)

func HasUserLikedPost(userID, postID uint64) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM post_likes WHERE userID = ? AND postID = ?`
	err := db.DB.QueryRow(query, userID, postID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddLikeToPost(userID, postID uint64) error {
	insert := `INSERT INTO post_likes (userID, postID) VALUES (?, ?)`
	_, err := db.DB.Exec(insert, userID, postID)
	if err != nil {
		return err
	}

	update := `UPDATE posts SET postLike = postLike + 1 WHERE postID = ?`
	_, err = db.DB.Exec(update, postID)
	return err
}

func RemoveLikeFromPost(userID, postID uint64) error {
	delete := `DELETE FROM post_likes WHERE userID = ? AND postID = ?`
	res, err := db.DB.Exec(delete, userID, postID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("like non trouvé")
	}

	update := `UPDATE posts SET postLike = postLike - 1 WHERE postID = ? AND postLike > 0`
	_, err = db.DB.Exec(update, postID)
	return err
}

func HasUserLikedComment(userID, commentID uint64) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM comment_likes WHERE userID = ? AND commentID = ?`
	err := db.DB.QueryRow(query, userID, commentID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddLikeToComment(userID, commentID uint64) error {
	insert := `INSERT INTO comment_likes (userID, commentID) VALUES (?, ?)`
	_, err := db.DB.Exec(insert, userID, commentID)
	if err != nil {
		return err
	}

	update := `UPDATE comments SET commentLike = commentLike + 1 WHERE commentID = ?`
	_, err = db.DB.Exec(update, commentID)
	return err
}

func RemoveLikeFromComment(userID, commentID uint64) error {
	delete := `DELETE FROM comment_likes WHERE userID = ? AND commentID = ?`
	res, err := db.DB.Exec(delete, userID, commentID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("like non trouvé")
	}

	update := `UPDATE comments SET commentLike = commentLike - 1 WHERE commentID = ? AND commentLike > 0`
	_, err = db.DB.Exec(update, commentID)
	return err
}
