/* examples from
https://qvault.io/golang/golang-logging-best-practices/#errors-not-strings
*/

package main

import (
	"errors"
	"fmt"
)

// use error  type
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

// wrap errors
func formatTime(hours, minutes int) (string, error) {
	return "", errors.New("there is an error.")
}

func formatTimeWithMessage(hours, minutes int) (string, error) {
	formatted, err := formatTime(hours, minutes) // formatTime does not exist
	if err != nil {
		return "", fmt.Errorf("formatTimeWithMessage: %v", err) // use "fmt.Errorf" instead errors.New
	}
	return "It is " + formatted + " o'clock", nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	_, err = formatTimeWithMessage(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
