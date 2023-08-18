package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError("", errRollback)
	} else {
		errCommit := tx.Commit()
		if err != nil {
			PanicIfError("Failed to Commit Transactions: ", errCommit)
		}
	}
}
