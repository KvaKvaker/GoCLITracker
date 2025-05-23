package task

import (
	"encoding/json"
	"log"
	"os"
)

// Create json file
func JsonCreate() *os.File {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)

	s := TaskList{
		Tasks: []Task{},
	}

	res, err := json.Marshal(s)

	file.WriteString(string(res))

	if err != nil {
		log.Println(err)
		return nil
	} else {
		return file
	}
}

func JsonGet() *os.File {
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0755)

	if err != nil {
		log.Println(err)
		return nil
	} else {
		return file
	}
}

// Check json file
func JsonCheck() bool {
	file, err := os.Open("tasks.json")
	if err != nil {
		return false
	} else {
		file.Close()
		return true
	}

}
