package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Category    string `json:"category"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}

func GetAllTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	result := db.Find(&tasks)
	return tasks, result.Error
}

func GetTaskByID(db *gorm.DB, id string) (Task, error) {
	var task Task
	result := db.First(&task, id)
	return task, result.Error
}

func CreateTask(db *gorm.DB, task *Task) error {
	result := db.Create(task)
	return result.Error
}

func UpdateTask(db *gorm.DB, task *Task) error {
	result := db.Save(task)
	return result.Error
}

func DeleteTask(db *gorm.DB, task *Task) error {
	result := db.Delete(task)
	return result.Error
}
