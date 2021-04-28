package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/dondany/simple-blog/domain"
)

type postgresqlPostRepo struct {
	dbpool *pgxpool.Pool
}

func PostgresqlPost() domain.PostRepository {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://root:root@localhost:5432/test_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	return &postgresqlPostRepo{dbpool: dbpool}
}

func (repo postgresqlPostRepo) FindAll(ctx context.Context) ([]domain.Post, error) {
	rows, err := repo.dbpool.Query(ctx, "select id, title, content from posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func (repo postgresqlPostRepo) Find(ctx context.Context, id int64) (domain.Post, error) {
	result := domain.Post{}
	err := repo.dbpool.QueryRow(ctx, "select id, title, content from posts where id=$1", id).Scan(&result.Id, &result.Title, &result.Content)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (repo postgresqlPostRepo) Create(ctx context.Context, post *domain.Post) error {
	_, err := repo.dbpool.Exec(ctx, `insert into posts (title, content) values ($1, $2)`, post.Title, post.Content)
	if err != nil {
		return err
	}
	return nil
}

func (repo postgresqlPostRepo) Delete(ctx context.Context, id int64) error {
	_, err := repo.dbpool.Exec(ctx, `delete from posts where id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
