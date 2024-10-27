package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var _ = json.Marshal

func decodeBencode(bencodedString string) (interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		var firstColonIndex int

		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				firstColonIndex = i
				break
			}
		}

		lengthStr := bencodedString[:firstColonIndex]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", err
		}

		return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], nil
	} 


 } else if rune(bencodedString[0]) == 'l' {//+

		colonIndex := 2 // Start after "l" to find the length of the string
		for colonIndex < len(bencodedString) && bencodedString[colonIndex] != ':' {
			colonIndex++
		}

		// Extract the length string (characters between 'l' and ':')
		lengthStr := bencodedString[1:colonIndex]
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			fmt.Println("Error parsing length:", err)
			return "", err
		}

		// Extract the actual string based on the parsed length
		start := colonIndex + 1
		Real_string := bencodedString[start : start+length]
		return Real_string, nil
	} else {
		return "", fmt.Errorf("Only strings are supported at the moment")
	}


func main() {

	fmt.Println("Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {
		//
		//
		bencodedValue := os.Args[2]
		
		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
