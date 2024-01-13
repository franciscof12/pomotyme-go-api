package models

import "gorm.io/gorm"

type Pomodoro struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	StudyTime uint   `json:"study_time"`
	BreakTime uint   `json:"break_time"`
	EndDate   string `json:"end_date"`
}

func GetAllPomodoros(db *gorm.DB) ([]Pomodoro, error) {
	var pomo []Pomodoro
	result := db.Find(&pomo)
	return pomo, result.Error
}

func GetPomodoroByID(db *gorm.DB, id string) (Pomodoro, error) {
	var pomo Pomodoro
	result := db.First(&pomo, id)
	return pomo, result.Error
}

func CreatePomodoro(db *gorm.DB, pomo *Pomodoro) error {
	result := db.Create(pomo)
	return result.Error
}

func UpdatePomodoro(db *gorm.DB, pomo *Pomodoro) error {
	result := db.Save(pomo)
	return result.Error
}

func DeletePomodoro(db *gorm.DB, pomo *Pomodoro) error {
	result := db.Delete(pomo)
	return result.Error
}
