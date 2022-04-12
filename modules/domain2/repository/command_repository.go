package repository

import (
	"codebase/go-codebase/modules/domain2/entity"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	// CQRS design pattern best practice. memisahkan DB untuk write dan read
	// bisa menambahkan config DB lain (Elasticsearch, redis, SQL, mongodb)
	DBWrite *sql.DB
	DBRead  *sql.DB
	logger  *logrus.Logger
}

func CreateRepository(dbWrite, dbRead *sql.DB, logger *logrus.Logger) *Repository {
	return &Repository{dbWrite, dbRead, logger}
}

func (r *Repository) InsertUser(ctx context.Context, user entity.Users) (err error) {
	query := "INSERT INTO sharing_session.users(id, fullname, no_hp, is_attend) values ($1, $2, $3, $4)"

	_, err = r.DBWrite.ExecContext(ctx, query, user.ID, user.Fullname, user.NoHP, user.IsAttend)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUserByID(ctx context.Context, id int, fullname string) (err error) {
	query := "UPDATE sharing_session.users SET fullname = $1 where id = $2"
	_, err = r.DBWrite.ExecContext(ctx, query, fullname, id)
	if err != nil {
		return
	}

	return
}
