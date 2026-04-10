package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"github/jdiaz7953/cli-todo-go/printErrors"
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

	//REPL for the cli
	for true{

		//Get the user input
		fmt.Print("$todo-cli ")
		scanner := bufio.NewScanner(os.Stdin)
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

			//Delete task based on ID
			taskId, _ := strconv.Atoi(input[1])
			for i := 0; i < len(allTask); i++ {
				if allTask[i].Id == taskId {
					allTask = slices.Delete(allTask, i, 1)
				}
			}

		case "mark-in-progress":
		}

		
		//Processes JSON Data
		allTaskJson, _ := json.Marshal(allTask)

		os.WriteFile("listOfTask.json", allTaskJson,0644)

	}
	//Display a help section to understand commands

}