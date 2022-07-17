package helper

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func CommitOrRollback(ctx context.Context, tx *sql.Tx, logger *logrus.Logger) {
	if err := recover(); err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
