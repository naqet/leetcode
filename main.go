package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Incorrect number of args")
	}
	problem := os.Args[1]
	problem = strings.ToLower(problem)
	problem = strings.Replace(problem, ". ", ".", 1)
	problem = strings.ReplaceAll(problem, " ", "-")

	err := os.Mkdir(problem, 0755)
	if err != nil {
		log.Fatalf("Problem dir already exists")
	}

	template, err := os.Open("template.txt")

	defer template.Close()

	if os.IsExist(err) {
		log.Fatalf("Template does not exist")
	}

	logic, err := os.Create(problem + "/logic_test.go")
	if err != nil {
		log.Fatalf("Error creating a logic file")
	}
	defer logic.Close()

	_, err = io.Copy(logic, template)

	if err != nil {
		log.Fatal("Error coping data to logic file")
	}

	// Copy the test command for given problem
	exec.Command("sh", "-c", "echo go test ./"+problem+"/logic_test.go "+"| pbcopy").Run()
	log.Println("Created new problem!")
}
