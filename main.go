package main

import (
	"bufio"
	"log"
	"os"
)

// This script helps you automate the process of comparing many ideas to pick one you like the most.
// It could be used to compare lists of anything really.

func main() {
	file, err := os.Open("ideas.txt")
	if err != nil {
		log.Fatalf("Failed to open ideas.txt file (one idea per line): %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ideas := []string{}
	for scanner.Scan() {
		ideas = append(ideas, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	top := 0
	for i := 1; i < len(ideas); i++ {
		log.Printf("a: %s\n", ideas[top])
		log.Printf("b: %s\n", ideas[i])
		log.Print("a or b?")

		choice, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		if string(choice) != "a" {
			top = i
		}
	}

	log.Printf("Top idea: %s\n", ideas[top])
}
