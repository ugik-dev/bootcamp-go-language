package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/models"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment models.Comment) (models.Comment, error) {
	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (models.Comment, error) {
	query := "SELECT * FROM comments where id = ? "
	result, err := repository.DB.QueryContext(ctx, query, id)
	comment := models.Comment{}
	if err != nil {
		return comment, err
	}
	defer result.Close()
	if result.Next() {
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("ID " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]models.Comment, error) {
	query := "SELECT * FROM comments"
	result, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var comments []models.Comment
	for result.Next() {
		comment := models.Comment{}
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
