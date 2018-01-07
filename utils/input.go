package utils

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
/*
	Returns the input of the question
*/
func Question(question string) string {
	var input string
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}

func QuestionF(format string, question ...string) string {
	fmt.Printf(format, question)
	in := bufio.NewReader(os.Stdin)
	resp, _ := in.ReadString('\n')
	return strings.Replace(resp, "\n", "", -1)
}


