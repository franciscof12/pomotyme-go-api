package schemas

type PomodoroSchema struct {
	UserID    uint   `json:"user_id" binding:"required"`
	StudyTime uint   `json:"study_time"`
	BreakTime uint   `json:"break_time"`
	EndDate   string `json:"end_date"`
}
