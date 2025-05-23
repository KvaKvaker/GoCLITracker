package main

import (
	"fmt"
	"github.com/KvaKvaker/GoCLITracker/task"
	"os"
)

var helpNote = `-------------------------------------------------------------
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

type myStruct struct {
	Tasks []Task
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdateAt    string `json:"updateAt"`
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		fmt.Println(helpNote)
	} else {
		var jsonFile *os.File
		switch os.Args[1] {
		case "add":
			task.Add(jsonFile)
		case "clear":
			task.Clear()
		case "update":
			task.Update(jsonFile)
		case "delete":
			task.Delete(jsonFile)
		case "mark-in-progress":
			task.MarkInProgress(jsonFile)
		case "mark-done":
			task.MarkDone(jsonFile)
		case "mark-todo":
			task.MarkToDo(jsonFile)
		case "list":
			task.List(jsonFile)
		default:
			fmt.Println("Unknown command\nPrint --help to show all commands")
		}
	}
}
