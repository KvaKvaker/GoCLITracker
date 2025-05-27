# GO CLI Task Tracker

This is a cli tracker written in golang, it can add, delete, change tasks, and also change task statuses

## How to run

Clone the repository in you folder and run the following command:

```bash
git clone https://github.com/KvaKvaker/GoCLITracker.git
cd cmd/tracker
```

Run the following command to build and run the project:

```bash
go build
./tracker or ./tracker.exe # To see the list of available commands

# To add a task
./tracker.exe add "Some task"

# To update a task
./tracker.exe update 1 "Updated description"

# To delete a task
./tracker.exe delete 1

# To mark a task as in progress/done/todo
./tracker.exe mark-in-progress 1
./tracker.exe mark-done 1
./tracker.exe mark-todo 1

# To list all tasks
./tracker.exe list
./tracker.exe list done
./tracker.exe list todo
./tracker.exe list in-progress

# To clear all tasks
./tracker.exe clear
```
