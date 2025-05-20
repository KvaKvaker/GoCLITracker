package main

import (
	//"flag"
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var helpNote = `-------------------------------------------------------------
To get help print <file.exe> --help
To add task print <file.exe> add "TaskName"
To pdate task print <file.exe> update <TaskId> "NewTaskName"
To delete task print <file.exe> delete <TaskId>
To mark task in progress print <file.exe> mark-in-progress <TaskId>
To mark task done print <file.exe> mark-done <TaskId>
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
		switch os.Args[1] {
		case "add":
			if len(os.Args) > 2 {
				var jsonFile *os.File
				if jsonCheck() {
					fmt.Println("Getting json file")
					jsonFile = jsonGet()
				} else {
					fmt.Println("Creating json file")
					jsonCreate()
					jsonFile = jsonGet()
				}
				defer jsonFile.Close()

				jsonInfo, err := io.ReadAll(jsonFile)
				if err != nil {
					fmt.Println(err)
				}

				var stru myStruct
				var date = time.Now().Format(time.DateTime)

				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				stru.Tasks = append(stru.Tasks, Task{
					ID:          len(stru.Tasks) + 1,
					Description: os.Args[2],
					Status:      "Not Done",
					CreatedAt:   date,
					UpdateAt:    date,
				})

				res, err := json.Marshal(stru)
				if err != nil {
					fmt.Println(err)
				}

				// Reset file pointer to beginning and truncate file
				jsonFile.Seek(0, 0)
				jsonFile.Truncate(0)

				_, err = jsonFile.Write(res)
				if err != nil {
					fmt.Println(err)
				}

				// fmt.Println(stru, ne, string(res))
			} else {
				fmt.Println("You should use name for task")
			}

		case "update":
			if jsonCheck() {

			} else {
				fmt.Println("You don't have tasks to update!")
			}
		case "delete":
			if jsonCheck() {

			} else {
				fmt.Println("You don't have tasks to delete!")
			}
		case "mark-in-progress":
			if jsonCheck() {

			} else {
				fmt.Println("You don't have tasks to edit!")
			}
		case "mark-done":
			if jsonCheck() {

			} else {
				fmt.Println("You don't have tasks to edit!")
			}
		case "list":
			if jsonCheck() {

			} else {
				fmt.Println("You don't have tasks to list!")
			}
		default:
			fmt.Println("Unknown command\nPrint --help to show all commands")
		}
	}
}

// Create json file
func jsonCreate() *os.File {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)

	s := myStruct{
		Tasks: []Task{},
	}

	res, err := json.Marshal(s)

	file.WriteString(string(res))

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return file
	}
}

func jsonGet() *os.File {
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0755)

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return file
	}
}

// Check json file
func jsonCheck() bool {
	file, err := os.Open("tasks.json")
	if err != nil {
		return false
	} else {
		file.Close()
		return true
	}

}
