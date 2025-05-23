package task

import (
	"os"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	var jsonFile *os.File
	res, err := Add(jsonFile)
	if err != nil {
		t.Errorf("We got some error! %v", err)
	} else {
		t.Log(res)
	}

	time.Sleep(10000 * time.Millisecond)
	os.Remove("tasks.json")
	//some
}
