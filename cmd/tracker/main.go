package main

import (
	"fmt"
	"github.com/KvaKvaker/GoCLITracker/task"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		fmt.Println(task.HelpNote)
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
