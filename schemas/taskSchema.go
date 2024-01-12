package schemas

type TaskSchema struct {
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id" binding:"required"`
}
