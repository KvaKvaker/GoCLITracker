package task

const (
	StatusTodo       = "To do"
	StatusInProgress = "In progress"
	StatusDone       = "Done"
)

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdateAt    string `json:"updateAt"`
}
