package queries

import (
	"fmt"
	"github.com/Jneville0815/fitness-app-backend-v2/app/models"
	"github.com/jmoiron/sqlx"
	"strconv"
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

func (q *QuoteQueries) DeleteQuote(userId string, quoteIdStr string) error {
	quoteId, err := strconv.Atoi(quoteIdStr)
	if err != nil {
		return fmt.Errorf("invalid quote_id: %v", err)
	}

	query := `DELETE FROM quotes WHERE id = ? AND user_id = ?`

	_, err = q.Exec(query, quoteId, userId)
	if err != nil {
		return err
	}

	return nil
}
