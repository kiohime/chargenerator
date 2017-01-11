package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// JUST ECHO, FOR SOME REASON
// func echo(names []string) {
// 	s, sep := "", ""
// 	for _, arg := range names[0:] {
// 		s += sep + arg
// 		sep = "\n"
// 	}
// 	fmt.Println(s)
// 	return
// }

// readLines reads a whole file into memory
// // and returns a slice of its lines.
// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

func main() {
	readName, _ := ioutil.ReadFile("charNames.txt")
	// if err != nil {
	// 	return
	// }
	strName := string(readName)             //creates string from fileinput
	names := strings.Split(strName, "\r\n") //creates slice from ctreated string, using beginning of stroke and newline as separator

	for i := 1; i < 50; i++ {
		time.Sleep(10 * time.Millisecond) //give rand.Seed some time to reset value
		rand.Seed(time.Now().UnixNano())
		randName := names[rand.Intn(len(names))] //randomzer from slice

		fmt.Printf("Random name is: %v\n", randName)
	}

	fmt.Printf("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // for building program interface

}
