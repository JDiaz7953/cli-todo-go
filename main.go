package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github/jdiaz7953/cli-todo-go/printErrors"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	Id int           `json:"id"`
	Description string  `json:"description"`
	Status string	    `json:"status"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main(){
	var allTask []Todo
	idCounter := 1
	session := true
	scanner := bufio.NewScanner(os.Stdin)

	//Get persistant data
	data, err := os.ReadFile("listOfTask.json")
	if err == nil {
		json.Unmarshal(data, &allTask)
		idCounter = len(allTask) + 1
	}


	//REPL for the cli
	for session {

		//Get the user input
		fmt.Print("$todo-cli ")
		scanner.Scan()
		userInput := scanner.Text()
		input := strings.Fields(userInput)
		userInput = input[0]
		

		switch userInput {
		case "help":

			//Sends the help dictionary to the user 
			if len(input) >= 2 {
				fmt.Println("")
				fmt.Println(input[1], "is not a valid argument")
				fmt.Println("")
			}else{
				printHelp()
			}
		case "add":

			if len(input) == 1 {
				fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			}

			//Adds a task to the user json file 
			textAfterCommand := strings.Join(input[1:], " ")
			if printErrors.QuoteError(textAfterCommand) {break}

			start := strings.Index(textAfterCommand, "\"")
			end := strings.LastIndex(textAfterCommand, "\"")

			if printErrors.ClosingQuoteError(start, end) {break}

			description := textAfterCommand[start+1:end]

			invalidCommand := strings.TrimSpace(textAfterCommand[end+1:])
			if  printErrors.BadCommand(invalidCommand) {break}


			task := Todo{
				Id: idCounter,
				Description: description,
				Status: "todo", 
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				}
				allTask = append(allTask, task)
				idCounter++

		case "update":

			if len(input) == 1 {
			   fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			   break
			}

			taskId, _:= strconv.Atoi(input[1])
			textAfterCommand := strings.Join(input[2:], " ")

			//This checks if there is quotes at all
			if printErrors.QuoteError(textAfterCommand) {break}

			start := strings.Index(textAfterCommand, "\"")
			end := strings.LastIndex(textAfterCommand, "\"")

			//This checks if the quotes are closed
			if printErrors.ClosingQuoteError(start, end){break}

			description := textAfterCommand[start+1:end]

			invalidCommand := strings.TrimSpace(textAfterCommand[end+1:])
			if printErrors.BadCommand(invalidCommand){break}

			for i := 0; i < len(allTask); i++ {
				if allTask[i].Id == taskId {
					allTask[i].Description = description
				}
			}

		case "delete":
			if len(input) == 1 {
			   fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			   break
			}

			//Delete task based on ID
			taskId, _ := strconv.Atoi(input[1])
			for i := 0; i < len(allTask); i++ {
				if allTask[i].Id == taskId {
					allTask = slices.Delete(allTask, i, i+1)
				}
			}

		case "mark-in-progress":
			if len(input) >= 3 {
				fmt.Println("")
				fmt.Println(input[2], "is not a valid argument")
				fmt.Println("")
				break
			}else if len(input) == 1 {
			   fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			   break
			}
			//Mark a task based on the ID as in-progress
			taskId, _ := strconv.Atoi(input[1])
			for i := 0; i < len(allTask); i++ {
				if allTask[i].Id == taskId {
					allTask[i].Status = "in-progress"
				}
			}
		case "mark-done":
			if len(input) >= 3 {
				fmt.Println("")
				fmt.Println(input[2], "is not a valid argument")
				fmt.Println("")
				break
			}else if len(input) == 1 {
			   fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			   break
			}
			taskId, _ := strconv.Atoi(input[1])
			for i := 0; i < len(allTask); i++ {
				if allTask[i].Id == taskId {
					allTask[i].Status = "done"
				}
			}

		case "list":
			if len(input) >= 3 {
				fmt.Println("")
				fmt.Println(input[2], "is not a valid argument")
				fmt.Println("")
				break
			}else if len(input) == 1 {
			   fmt.Println(`you must add a valid argument (type "help" to see commands)`)
			   break
			}

			if input[1] == "done" {
				for i := 0; i < len(allTask); i++ {
					if allTask[i].Status == "done" {
						fmt.Println(allTask[i].Description)
						fmt.Println(" ")
					}
				}
			}
			if input[1] == "todo" {
				for i := 0; i < len(allTask); i++ {
					if allTask[i].Status == "todo" {
						fmt.Println(allTask[i].Description)
						fmt.Println(" ")
					}
				}
			}
			if input[1] == "in-progress" {
				for i := 0; i < len(allTask); i++ {
					if allTask[i].Status == "in-progress" {
						fmt.Println(allTask[i].Description)
						fmt.Println(" ")
					}
				}
			}
		case "save-exit":
			session = false

		default:
			fmt.Println(" ")
			fmt.Println(`please enter a valid command (type "help" to see commands)`)
		}

	}
	//Display a help section to understand commands

		//Processes JSON Data
		allTaskJson, _ := json.Marshal(allTask)
		os.WriteFile("listOfTask.json", allTaskJson,0644)

}