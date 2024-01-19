package database

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/queries"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
	*queries.WorkoutQueries
	*queries.QuoteQueries
}

func OpenDBConnection() (*Queries, error) {
	var (
		db  *sqlx.DB
		err error
	)

	db, err = MysqlConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries:    &queries.UserQueries{DB: db},
		WorkoutQueries: &queries.WorkoutQueries{DB: db},
		QuoteQueries:   &queries.QuoteQueries{DB: db},
	}, nil
}
