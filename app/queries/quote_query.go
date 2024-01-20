package queries

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/models"
	"github.com/jmoiron/sqlx"
)

type QuoteQueries struct {
	*sqlx.DB
}

func (q *QuoteQueries) CreateQuote(qu *models.Quote) error {
	query := `INSERT INTO quotes (user_id, source, quote, num_views) VALUES (?, ?, ?, ?)`

	_, err := q.Exec(
		query,
		qu.UserID, qu.Source, qu.QuoteText, qu.NumViews,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *QuoteQueries) GetQuotes(userId string) ([]models.Quote, error) {
	var quotes []models.Quote

	query := `SELECT * FROM quotes WHERE user_id = ?`

	err := q.Select(&quotes, query, userId)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}
