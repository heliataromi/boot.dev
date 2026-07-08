package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	// Don't edit above this line

	strIssues := string(issues)
	fmt.Println(strIssues)
}
