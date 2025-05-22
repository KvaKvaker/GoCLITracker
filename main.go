package main

import (
	//"flag"
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
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
			if len(os.Args) > 2 {
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

				var taskId int

				if len(stru.Tasks) != 0 {
					taskId = stru.Tasks[len(stru.Tasks)-1].ID + 1
				} else {
					taskId = 1
				}

				stru.Tasks = append(stru.Tasks, Task{
					ID:          taskId,
					Description: os.Args[2],
					Status:      "To do",
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
				} else {
					fmt.Println("Task added!")
				}

				// fmt.Println(stru, ne, string(res))
			} else {
				fmt.Println("You should use name for task")
			}
		case "clear":
			err := os.Remove("tasks.json")
			if err != nil {
				fmt.Println("You don't have tasks")
			} else {
				fmt.Println("All tasks deleted!")
			}
		case "update":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				var date = time.Now().Format(time.DateTime)
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					fmt.Println("You should text id and description to update task")
				} else if len(os.Args) > 3 {
					//fmt.Println(stru.Tasks)
					for idx, val := range stru.Tasks {
						if res, err := strconv.Atoi(os.Args[2]); err == nil {
							if val.ID == res {
								stru.Tasks[idx].Description = os.Args[3]
								stru.Tasks[idx].UpdateAt = date
							}
						} else {
							fmt.Println("Second argument should be integer")
							break
						}
					}
					//fmt.Println(stru.Tasks)

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
				}
			} else {
				fmt.Println("You don't have tasks to update!")
			}
		case "delete":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					fmt.Println("You should text id to delete task")
				} else if len(os.Args) > 2 {
					//fmt.Println(stru.Tasks)
					for idx, val := range stru.Tasks {
						if res, err := strconv.Atoi(os.Args[2]); err == nil {
							if val.ID == res {
								stru.Tasks = append(stru.Tasks[0:idx], stru.Tasks[idx+1:]...)
								fmt.Println("Task deleted!")
							}
						} else {
							fmt.Println("Second argument should be integer")
							break
						}
					}
					//fmt.Println(stru.Tasks)

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
				}
			} else {
				fmt.Println("You don't have tasks to delete!")
			}
		case "mark-in-progress":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				var date = time.Now().Format(time.DateTime)
				var flag bool
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					fmt.Println("You should text id to update task")
				} else if len(os.Args) > 2 {
					//fmt.Println(stru.Tasks)
					for idx, val := range stru.Tasks {
						if res, err := strconv.Atoi(os.Args[2]); err == nil {
							if val.ID == res {
								stru.Tasks[idx].Status = "In progress"
								stru.Tasks[idx].UpdateAt = date
								flag = true
								fmt.Println("Updated!")
							}
						} else {
							fmt.Println("Second argument should be integer")
							break
						}
					}
					if !flag {
						fmt.Println("There is no task with this id")
					}
					//fmt.Println(stru.Tasks)

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
				}
			} else {
				fmt.Println("You don't have tasks to edit!")
			}
		case "mark-done":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				var date = time.Now().Format(time.DateTime)
				var flag bool
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					fmt.Println("You should text id to update task")
				} else if len(os.Args) > 2 {
					//fmt.Println(stru.Tasks)
					for idx, val := range stru.Tasks {
						if res, err := strconv.Atoi(os.Args[2]); err == nil {
							if val.ID == res {
								stru.Tasks[idx].Status = "Done"
								stru.Tasks[idx].UpdateAt = date
								flag = true
								fmt.Println("Updated!")
							}
						} else {
							fmt.Println("Second argument should be integer")
							break
						}
					}
					if !flag {
						fmt.Println("There is no task with this id")
					}
					//fmt.Println(stru.Tasks)

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
				}
			} else {
				fmt.Println("You don't have tasks to edit!")
			}
		case "mark-todo":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				var date = time.Now().Format(time.DateTime)
				var flag bool
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					fmt.Println("You should text id to update task")
				} else if len(os.Args) > 2 {
					//fmt.Println(stru.Tasks)
					for idx, val := range stru.Tasks {
						if res, err := strconv.Atoi(os.Args[2]); err == nil {
							if val.ID == res {
								stru.Tasks[idx].Status = "To do"
								stru.Tasks[idx].UpdateAt = date
								flag = true
								fmt.Println("Updated!")
							}
						} else {
							fmt.Println("Second argument should be integer")
							break
						}
					}
					if !flag {
						fmt.Println("There is no task with this id")
					}
					//fmt.Println(stru.Tasks)

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
				}
			} else {
				fmt.Println("You don't have tasks to edit!")
			}
		case "list":
			if jsonCheck() {
				jsonFile = jsonGet()
				defer jsonFile.Close()
				jsonInfo, err := io.ReadAll(jsonFile)

				var stru myStruct
				err = json.Unmarshal(jsonInfo, &stru)
				if err != nil {
					fmt.Println(err)
				}

				if len(os.Args) == 2 {
					var flag bool
					for idx, val := range stru.Tasks {
						if idx == 0 {
							fmt.Printf("Task #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", idx+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
							flag = true
						} else {
							fmt.Printf("\n\nTask #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", idx+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
						}
					}
					if !flag {
						fmt.Println("No tasks!")
					}
				} else if len(os.Args) > 2 {
					switch os.Args[2] {
					case "done":
						var cnt int
						var flag bool
						for _, val := range stru.Tasks {
							if val.Status == "Done" {
								if cnt == 0 {
									fmt.Printf("Task #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								} else {
									fmt.Printf("\n\nTask #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								}
								cnt++
								flag = true
							}
						}
						if !flag {
							fmt.Println("No tasks!")
						}
					case "todo":
						var cnt int
						var flag bool
						for _, val := range stru.Tasks {
							if val.Status == "To do" {
								if cnt == 0 {
									fmt.Printf("Task #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								} else {
									fmt.Printf("\n\nTask #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								}
								cnt++
								flag = true
							}
						}
						if !flag {
							fmt.Println("No tasks!")
						}
					case "in-progress":
						var cnt int
						var flag bool
						for _, val := range stru.Tasks {
							if val.Status == "In progress" {
								if cnt == 0 {
									fmt.Printf("Task #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								} else {
									fmt.Printf("\n\nTask #%v\n  ID - %v\n  Description - %v\n  Status - %v\n  Date Create - %v\n  Date Update - %v", cnt+1, val.ID, val.Description, val.Status, val.CreatedAt, val.UpdateAt)
								}
								cnt++
								flag = true
							}
						}
						if !flag {
							fmt.Println("No tasks!")
						}
					default:
						fmt.Println("Unknown command")
					}
				}

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
