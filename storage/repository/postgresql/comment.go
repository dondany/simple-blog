package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/dondany/simple-blog/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresqlCommentRepo struct {
	dbpool *pgxpool.Pool
}

func NewPostresqlCommentRepository() domain.CommentRepository {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://root:root@localhost:5432/test_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	return &postgresqlCommentRepo{dbpool: dbpool}
}

func (repo *postgresqlCommentRepo) GetAll(ctx context.Context) ([]domain.Comment, error) {
	rows, err := repo.dbpool.Query(ctx, "select id, postid, author, content from comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Comment
	for rows.Next() {
		comment := domain.Comment{}
		err = rows.
			Scan(&comment.Id,
				&comment.PostId,
				&comment.Author,
				&comment.Content)
		if err != nil {
			return nil, err
		}
		result = append(result, comment)
	}
	return result, nil
}

func (repo *postgresqlCommentRepo) Get(ctx context.Context, id int64) (domain.Comment, error) {
	result := domain.Comment{}
	err := repo.dbpool.QueryRow(context.Background(), "select id, postid, author, content from comments where id=$1", id).
		Scan(&result.Id,
			&result.PostId,
			&result.Author,
			&result.Content)
	if err != nil {
		return result, err
	}
	return result, nil
}
