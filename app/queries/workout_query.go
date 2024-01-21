package queries

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/models"
	"github.com/jmoiron/sqlx"
)

type WorkoutQueries struct {
	*sqlx.DB
}

func (q *WorkoutQueries) CreateWorkout(w *models.Workout) error {
	query := `INSERT INTO workouts (user_id, bench, overhead_press, squat, deadlift, current_day, note) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := q.Exec(
		query,
		w.UserID, w.Bench, w.OverheadPress, w.Squat, w.DeadLift, w.CurrentDay, w.Note,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *WorkoutQueries) UpdateWorkout(w *models.Workout) error {
	query := `
        UPDATE workouts 
        SET bench = ?, overhead_press = ?, squat = ?, deadlift = ?, current_day = ?, note = ?
        WHERE user_id = ?
    `

	_, err := q.Exec(
		query,
		w.Bench, w.OverheadPress, w.Squat, w.DeadLift, w.CurrentDay, w.Note, w.UserID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *WorkoutQueries) UpdateNote(n *models.Note) error {
	query := `UPDATE workouts SET note = ? WHERE user_id = ?`

	_, err := q.Exec(query, n.Note, n.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (q *WorkoutQueries) UpdateCurrentDay(c *models.CurrentDay) error {
	query := `UPDATE workouts SET current_day = ? WHERE user_id = ?`

	_, err := q.Exec(query, c.CurrentDay, c.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (q *WorkoutQueries) GetWorkout(userId string) (models.Workout, error) {
	workout := models.Workout{}

	query := `SELECT * FROM workouts WHERE user_id = ?`

	err := q.Get(&workout, query, userId)
	if err != nil {
		return workout, err
	}

	return workout, nil
}

func (q *WorkoutQueries) GetCurrentDay(userId string) (models.CurrentDay, error) {
	currentDay := models.CurrentDay{}

	query := `SELECT user_id, current_day FROM workouts WHERE user_id = ?`

	err := q.Get(&currentDay, query, userId)
	if err != nil {
		return currentDay, err
	}
	return currentDay, nil
}

func (q *WorkoutQueries) GetNote(userId string) (models.Note, error) {
	note := models.Note{}

	query := `SELECT user_id, note FROM workouts WHERE user_id = ?`

	err := q.Get(&note, query, userId)
	if err != nil {
		return note, err
	}
	return note, nil
}
