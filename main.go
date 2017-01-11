package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	ansi "github.com/k0kubun/go-ansi"
)

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

type Character struct {
	charName, charNickname, charSurname, charGender string
}

func getRand(a string) string {

	time.Sleep(10 * time.Millisecond) //give rand.Seed some time to reset value
	rand.Seed(time.Now().UnixNano())

	file, _ := ioutil.ReadFile("")
	switch a {
	case "name":
		file, _ = ioutil.ReadFile("charNames.txt")
	case "surname":
		file, _ = ioutil.ReadFile("charSurNames.txt")
	case "nickname":
		file, _ = ioutil.ReadFile("charNickNames.txt")
	case "gender":
		gender := [2]string{"male", "female"}
		gRand := gender[rand.Intn(len(gender))]
		return gRand

	}

	strFile := string(file) //creates string from fileinput
	sliceFile := strings.Split(strFile, "\r\n")
	gRand := sliceFile[rand.Intn(len(sliceFile))] //randomzer from slice
	return gRand
}

func main() {
	anon := new(Character)
	for i := 1; i < 50; i++ {
		anon.charName = getRand("name")
		anon.charSurname = getRand("surname")
		anon.charNickname = getRand("nickname")
		anon.charGender = getRand("gender")
		ansi.Printf("\n%v \x1b[36m~%v~\x1b[0m %v, \x1b[33m%v\x1b[0m", anon.charName, anon.charNickname, anon.charSurname, anon.charGender)
	}

	fmt.Printf("\n\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // for building program interface
}
