package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Character blablabla
type Character struct {
	charNation, charMeta, charGender, charName, charNickname, charSurname string
	charID                                                                int
}

func (c *Character) String() string {
	return strings.Join([]string{c.charNation, c.charMeta, c.charGender, c.charName, c.charNickname, c.charSurname}, " ")
}

var inputNation int
var inputGender int
var genderIndex int
var nationIndex int
var metaIndex int

func getRand(a string) string {
	var file []string
	var gRand string
	switch a {
	case "nation":
		// nationIndex = rand.Intn(5)
		nationIndex = inputNation
		return gRand

	case "gender":
		genderIndex = inputGender
		switch genderIndex {
		case 0:
			gRand = "male"
		case 1:
			gRand = "female"
		case 2:
			gRand = "unisex"
		}
		return gRand
	case "name":
		file = names(nationIndex)
	case "surname":
		file = surnames(nationIndex)
	case "nickname":

		file = nicknames()
	case "metatype":
		file = metatypes()
	}

	// strFile := string(file) //creates string from fileinput
	// sliceFile := strings.Split(strFile, "\r\n")
	gRand = file[rand.Intn(len(file))] //randomzer from slice
	return gRand
}

func createChar(i int) *Character {

	anon := new(Character)
	// 	// assert(i > 1, "yhmjfdyhjdjhdfyhj")
	anon.charID = i
	anon.charNation = getRand("nation")
	anon.charGender = getRand("gender")
	anon.charName = getRand("name")
	anon.charSurname = getRand("surname")
	anon.charNickname = getRand("nickname")
	anon.charMeta = getRand("metatype")

	fmt.Println(anon.charID, anon.charName, "~"+anon.charNickname+"~", anon.charSurname, anon.charMeta, anon.charGender)
	// fmt.Printf("genderindex = %v\n\n", genderIndex)
	return anon
}

func parseNumber(s string) int {
	var value int
	if len(s) == 0 {
		return -1
	}
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			value = int(ch) - int('0') + value*10
		} else {
			value = -1
			break
		}
	}
	return value
}

func parseGender(s string) int {
	var value int
	if len(s) == 0 {

		return 3
	}
	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '2':
			value = int(ch) - int('0') + value*10
		default:
			value = -1
			break
		}

	}
	return value
}

func setCharNumber() int {
	var cond bool
	var input string
	var inputInt int
	for !cond {
		fmt.Print("Enter number of characters: ")
		fmt.Scanln(&input)
		inputInt = parseNumber(input)
		// inputInt, err := strconv.ParseInt(input, 10, 64)
		if inputInt < 0 {
			fmt.Println("Error in charName input")
		} else {
			cond = true
		}
	}
	return inputInt
}

func setCharGender() int {
	var cond bool
	var input string
	var inputGen int
	for !cond {
		fmt.Printf("Enter gender index (0 for \"male\", 1 for \"female\", 2 for \"unisex\"): ")
		fmt.Scanln(&input)
		inputGen = parseGender(input)
		fmt.Println(inputGen)
		switch {
		case inputGen < 0:
			fmt.Println("Error in charGender input")
		case inputGen == 3:
			fmt.Println("Random option selected")
			cond = true
		default:
			cond = true
		}
	}
	return inputGen
}

func main() {
	rand.Seed(time.Now().UnixNano())
	charArray := []*Character{}
	inputInt := setCharNumber()

	inputGender = setCharGender()

	for i := 1; i <= inputInt; i++ {

		inputNation = 1

		// inputGender = rand.Intn(3)
		// inputGender = rand.Intn(3)

		x := createChar(i)
		charArray = append(charArray, x)
		// fmt.Println(anon.charID, anon.charName, anon.charNickname, anon.charSurname, anon.charMeta, anon.charGender)
	}

	// fmt.Println(charArray[7])

	fmt.Printf("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // for building program interface
}
