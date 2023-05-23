//go project where i can add, update and delete tasks from a list. A task manager
// that can be used from the command line.
// The tasks are saved in a file called tasks.json
// This file is located in the same directory as the executable file of the program.
// The tasks are saved in the file in JSON format.
// The program has the following commands:
// - add: add a new task to the list
// - list: show the list of tasks
// - do: mark a task as completed
// - rm: remove a task from the list
// - update: update the description of a task
// - help: show the help menu
// - exit: exit the program
//
// The program can be run from the command line as follows:
// task_manager.exe <command> <parameters>
// The program can also be run as follows:
// task_manager.exe
// In this case, the program will display the help menu.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// The Task structure is used for saving the tasks in the file
// The structure has the following fields:
// - Id: the id of the task
// - Description: the description of the task
// - Completed: true if the task is completed, false otherwise
// - Date: the date when the task was created
type Task struct {
	Id          int
	Description string
	Completed   bool
	Date        time.Time
}

// The TaskList structure is used for saving the tasks in memory
// The structure has the following fields:
// - Tasks: the list of tasks
// - NextId: the id of the next task
type TaskList struct {
	Tasks  []Task
	NextId int
}

// The TaskListManager structure is used for managing the tasks
// The structure has the following fields:
// - TaskList: the list of tasks
// - FileName: the file name where the tasks are saved
// - Scanner: the scanner used for reading the user input
// - Reader: the reader used for reading the tasks from the file
// - Writer: the writer used for saving the tasks to the file
type TaskListManager struct {
	TaskList TaskList
	FileName string
	Scanner  *bufio.Scanner
	Reader   *bufio.Reader
	Writer   *bufio.Writer
}

// The TaskManager interface is used for managing the tasks
// The interface has the following methods:
// - AddTask: add a new task to the list
// - ListTasks: show the list of tasks
// - DoTask: mark a task as completed
// - RemoveTask: remove a task from the list
// - UpdateTask: update the description of a task
// - SaveTasks: save the tasks to the file
// - ReadTasks: read the tasks from the file
// - Exit: exit the program
type TaskManager interface {
	AddTask()
	ListTasks()
	DoTask()
	RemoveTask()
	UpdateTask()
	SaveTasks()
	ReadTasks()
	Exit()
}

// The NewTaskListManager function creates a new instance of the TaskListManager structure
// The function has the following parameters:
// - fileName: the file name where the tasks are saved
// The function returns an instance of the TaskListManager structure
func NewTaskListManager(fileName string) *TaskListManager {
	return &TaskListManager{
		TaskList: TaskList{
			Tasks:  []Task{},
			NextId: 1,
		},
		FileName: fileName,
		Scanner:  bufio.NewScanner(os.Stdin),
		Reader:   bufio.NewReader(os.Stdin),
		Writer:   bufio.NewWriter(os.Stdout),
	}
}

// The AddTask method adds a new task to the list
// The method has no parameters
// The method has no return values

// The ListTasks method shows the list of tasks
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) ListTasks() {
	if len(taskListManager.TaskList.Tasks) == 0 {
		fmt.Fprintln(taskListManager.Writer, "The task list is empty")
		taskListManager.Writer.Flush()
		return
	}

	fmt.Fprintln(taskListManager.Writer, "Id\tDescription\tCompleted\tDate")
	taskListManager.Writer.Flush()
	for _, task := range taskListManager.TaskList.Tasks {
		fmt.Fprintf(taskListManager.Writer, "%d\t%s\t%t\t%s\n", task.Id, task.Description, task.Completed, task.Date.Format("2006-01-02 15:04:05"))
		taskListManager.Writer.Flush()
	}
}

// The DoTask method marks a task as completed
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) DoTask() {
	fmt.Fprintln(taskListManager.Writer, "Enter the id of the task:")
	taskListManager.Writer.Flush()
	id, _ := taskListManager.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for index, task := range taskListManager.TaskList.Tasks {
		if fmt.Sprintf("%d", task.Id) == id {
			taskListManager.TaskList.Tasks[index].Completed = true
			return
		}
	}

	fmt.Fprintln(taskListManager.Writer, "The task with the specified id was not found")
	taskListManager.Writer.Flush()
}

// The RemoveTask method removes a task from the list
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) RemoveTask() {
	fmt.Fprintln(taskListManager.Writer, "Enter the id of the task:")
	taskListManager.Writer.Flush()
	id, _ := taskListManager.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for index, task := range taskListManager.TaskList.Tasks {
		if fmt.Sprintf("%d", task.Id) == id {
			taskListManager.TaskList.Tasks = append(taskListManager.TaskList.Tasks[:index], taskListManager.TaskList.Tasks[index+1:]...)
			return
		}
	}

	fmt.Fprintln(taskListManager.Writer, "The task with the specified id was not found")
	taskListManager.Writer.Flush()
}

// The UpdateTask method
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) UpdateTask() {
	fmt.Fprintln(taskListManager.Writer, "Enter the id of the task:")
	taskListManager.Writer.Flush()
	id, _ := taskListManager.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for index, task := range taskListManager.TaskList.Tasks {
		if fmt.Sprintf("%d", task.Id) == id {
			fmt.Fprintln(taskListManager.Writer, "Enter the new description of the task:")
			taskListManager.Writer.Flush()
			description, _ := taskListManager.Reader.ReadString('\n')
			description = strings.TrimSpace(description)

			taskListManager.TaskList.Tasks[index].Description = description
			return
		}
	}

	fmt.Fprintln(taskListManager.Writer, "The task with the specified id was not found")
	taskListManager.Writer.Flush()
}

// The SaveTasks method saves the tasks to the file
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) SaveTasks() {
	file, err := os.Create(taskListManager.FileName)
	if err != nil {
		fmt.Fprintln(taskListManager.Writer, err)
		taskListManager.Writer.Flush()
		return
	}
	defer file.Close()

	for _, task := range taskListManager.TaskList.Tasks {
		_, err = fmt.Fprintf(file, "%d|%s|%t|%s\n", task.Id, task.Description, task.Completed, task.Date.Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Fprintln(taskListManager.Writer, err)
			taskListManager.Writer.Flush()
			return
		}
	}
}

// The ReadTasks method reads the tasks from the file
// The method has no parameters
// The method has no return values
func (taskListManager *TaskListManager) ReadTasks() {
	file, err := os.Open(taskListManager.FileName)
	if err != nil {
		fmt.Fprintln(taskListManager.Writer, err)
		taskListManager.Writer.Flush()
		return
	}
	defer file.Close()

	taskListManager.TaskList.Tasks = []Task{}
	taskListManager.TaskList.NextId = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		id := fields[0]
		description := fields[1]
		completed := fields[2]
		date := fields[3]

		task := Task{}
		task.Id, err = strconv.Atoi(id)
		if err != nil {
			fmt.Fprintln(taskListManager.Writer, err)
			taskListManager.Writer.Flush()
			return
		}
		task.Description = description
		task.Completed, err = strconv.ParseBool(completed)
		if err != nil {
			fmt.Fprintln(taskListManager.Writer, err)
			taskListManager.Writer.Flush()
			return
		}
		task.Date, err = time.Parse("", date)
		if err != nil {
			fmt.Fprintln(taskListManager.Writer, err)
			taskListManager.Writer.Flush()
			return
		}

		taskListManager.TaskList.Tasks = append(taskListManager.TaskList.Tasks, task)
		taskListManager.TaskList.NextId++
	}
}

// exit the program
func (taskListManager *TaskListManager) Exit() {
	fmt.Fprintln(taskListManager.Writer, "Goodbye!")
	taskListManager.Writer.Flush()
	os.Exit(0)
}
func (taskListManager *TaskListManager) Help() {
	fmt.Fprintln(taskListManager.Writer, "The following commands are available:")
	fmt.Fprintln(taskListManager.Writer, "add: add a new task to the list")
	fmt.Fprintln(taskListManager.Writer, "list: show the list of tasks")
	fmt.Fprintln(taskListManager.Writer, "do: mark a task as completed")
	fmt.Fprintln(taskListManager.Writer, "rm: remove a task from the list")
	fmt.Fprintln(taskListManager.Writer, "update: update the description of a task")
	fmt.Fprintln(taskListManager.Writer, "help: show the help menu")
	fmt.Fprintln(taskListManager.Writer, "exit: exit the program")
	taskListManager.Writer.Flush()
}

// The main function is the entry point of the program
func main() {
	taskListManager := NewTaskListManager("tasks.json")
	//loop until the user enters the exit command
	for {
		fmt.Fprintln(taskListManager.Writer, "Enter a command:")
		taskListManager.Writer.Flush()
		taskListManager.Scanner.Scan()
		command := taskListManager.Scanner.Text()

		switch command {
		case "list":
			taskListManager.ListTasks()
		case "do":
			taskListManager.DoTask()
		case "rm":
			taskListManager.RemoveTask()
		case "update":
			taskListManager.UpdateTask()
		case "help":
			taskListManager.Help()
		case "exit":
			taskListManager.Exit()
		default:
			fmt.Fprintln(taskListManager.Writer, "Invalid command")
			taskListManager.Writer.Flush()
		}

	}
}
