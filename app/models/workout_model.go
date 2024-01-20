package models

import "github.com/google/uuid"

type Workout struct {
	UserID        uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Bench         float64   `db:"bench" json:"bench" validate:"required,gte=0"`
	OverheadPress float64   `db:"overhead_press" json:"overhead_press" validate:"required,gte=0"`
	Squat         float64   `db:"squat" json:"squat" validate:"required,gte=0"`
	DeadLift      float64   `db:"deadlift" json:"deadlift" validate:"required,gte=0"`
	CurrentDay    int       `db:"current_day" json:"current_day" validate:"min=0,max=15"`
	Note          string    `db:"note" json:"note"`
}

type Note struct {
	UserID uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Note   string    `db:"note" json:"note" validate:"required"`
}

type CurrentDay struct {
	UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	CurrentDay int       `db:"current_day" json:"current_day" validate:"required,min=0,max=15"`
}
