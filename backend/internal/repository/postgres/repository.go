package postgres

import "github.com/sav1nbrave4code/APG3/pkg/db/postgres_db"

type repository struct {
	db *postgres_db.PostgresDb
}

func New(db *postgres_db.PostgresDb) *repository {
	return &repository{
		db: db,
	}
}
