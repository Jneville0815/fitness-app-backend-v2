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

func (q *WorkoutQueries) GetWorkout(userId string) (models.Workout, error) {
	workout := models.Workout{}

	query := `SELECT * FROM workouts WHERE user_id = ?`

	err := q.Get(&workout, query, userId)
	if err != nil {
		return workout, err
	}

	return workout, nil
}
