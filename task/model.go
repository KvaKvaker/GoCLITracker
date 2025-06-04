package task

const (
	StatusTodo       = "To do"
	StatusInProgress = "In progress"
	StatusDone       = "Done"
)

var HelpNote = `-------------------------------------------------------------
To get help print <file.exe> --help
To clear all task list <file.exe> clear
To add task print <file.exe> add "TaskDescription"
To update task print <file.exe> update <TaskId> "NewTaskDescription"
To delete task print <file.exe> delete <TaskId>
To mark task in progress print <file.exe> mark-in-progress <TaskId>
To mark task done print <file.exe> mark-done <TaskId>
To mark task todo print <file.exe> mark-todo <TaskId>
To get tasks list print <file.exe> list <optional mod>
-With list command you can use done, todo, in-progress
-------------------------------------------------------------
`

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
