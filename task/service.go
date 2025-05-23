package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func Add(jsonFile *os.File) (int, error) {
	if len(os.Args) > 2 {
		if JsonCheck() {
			fmt.Println("Getting json file")
			jsonFile = JsonGet()
		} else {
			fmt.Println("Creating json file")
			JsonCreate()
			jsonFile = JsonGet()
		}
		defer jsonFile.Close()

		jsonInfo, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		var stru TaskList
		var date = time.Now().Format(time.DateTime)

		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
			return 0, err
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
			Status:      StatusTodo,
			CreatedAt:   date,
			UpdateAt:    date,
		})

		res, err := json.Marshal(stru)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		// Reset file pointer to beginning and truncate file
		jsonFile.Seek(0, 0)
		jsonFile.Truncate(0)

		_, err = jsonFile.Write(res)
		if err != nil {
			log.Println(err)
			return 0, err
		} else {
			fmt.Println("Task added!")
			return taskId, nil
		}
	} else {
		fmt.Println("You should use name for task")
		return 0, errors.New("You should use name for task")
	}
}

func Clear() {
	err := os.Remove("tasks.json")
	if err != nil {
		fmt.Println("You don't have tasks")
	} else {
		fmt.Println("All tasks deleted!")
	}
}

func Update(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		var date = time.Now().Format(time.DateTime)
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
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
						fmt.Println("Updated!")
					}
				} else {
					fmt.Println("Second argument should be integer")
					break
				}
			}
			//fmt.Println(stru.Tasks)

			res, err := json.Marshal(stru)
			if err != nil {
				log.Println(err)
			}

			// Reset file pointer to beginning and truncate file
			jsonFile.Seek(0, 0)
			jsonFile.Truncate(0)

			_, err = jsonFile.Write(res)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("You don't have tasks to update!")
	}
}

func Delete(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
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
				log.Println(err)
			}

			// Reset file pointer to beginning and truncate file
			jsonFile.Seek(0, 0)
			jsonFile.Truncate(0)

			_, err = jsonFile.Write(res)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("You don't have tasks to delete!")
	}
}

func MarkInProgress(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		var date = time.Now().Format(time.DateTime)
		var flag bool
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
		}

		if len(os.Args) == 2 {
			fmt.Println("You should text id to update task")
		} else if len(os.Args) > 2 {
			//fmt.Println(stru.Tasks)
			for idx, val := range stru.Tasks {
				if res, err := strconv.Atoi(os.Args[2]); err == nil {
					if val.ID == res {
						stru.Tasks[idx].Status = StatusInProgress
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
				log.Println(err)
			}

			// Reset file pointer to beginning and truncate file
			jsonFile.Seek(0, 0)
			jsonFile.Truncate(0)

			_, err = jsonFile.Write(res)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("You don't have tasks to edit!")
	}
}

func MarkDone(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		var date = time.Now().Format(time.DateTime)
		var flag bool
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
		}

		if len(os.Args) == 2 {
			log.Println("You should text id to update task")
		} else if len(os.Args) > 2 {
			//fmt.Println(stru.Tasks)
			for idx, val := range stru.Tasks {
				if res, err := strconv.Atoi(os.Args[2]); err == nil {
					if val.ID == res {
						stru.Tasks[idx].Status = StatusDone
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
				log.Println(err)
			}

			// Reset file pointer to beginning and truncate file
			jsonFile.Seek(0, 0)
			jsonFile.Truncate(0)

			_, err = jsonFile.Write(res)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("You don't have tasks to edit!")
	}
}

func MarkToDo(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		var date = time.Now().Format(time.DateTime)
		var flag bool
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
		}

		if len(os.Args) == 2 {
			fmt.Println("You should text id to update task")
		} else if len(os.Args) > 2 {
			//fmt.Println(stru.Tasks)
			for idx, val := range stru.Tasks {
				if res, err := strconv.Atoi(os.Args[2]); err == nil {
					if val.ID == res {
						stru.Tasks[idx].Status = StatusTodo
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
				log.Println(err)
			}

			// Reset file pointer to beginning and truncate file
			jsonFile.Seek(0, 0)
			jsonFile.Truncate(0)

			_, err = jsonFile.Write(res)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("You don't have tasks to edit!")
	}
}

func List(jsonFile *os.File) {
	if JsonCheck() {
		jsonFile = JsonGet()
		defer jsonFile.Close()
		jsonInfo, err := io.ReadAll(jsonFile)

		var stru TaskList
		err = json.Unmarshal(jsonInfo, &stru)
		if err != nil {
			log.Println(err)
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
					if val.Status == StatusDone {
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
					if val.Status == StatusTodo {
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
					if val.Status == StatusInProgress {
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
}
