package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//Character blablabla
type Character struct {
	charNation, charMeta, charGender, charName, charNickname, charSurname string
	charID                                                                int
}

var inputGender int
var genderIndex int
var nationIndex int

func getRand(a string) string {
	var file []string
	var gRand string
	switch a {
	case "nation":
		nationIndex = rand.Intn(5)
		// nationIndex = 1
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

	fmt.Println(anon.charID, anon.charName, anon.charNickname, anon.charSurname, anon.charMeta, anon.charGender)
	// fmt.Printf("genderindex = %v\n\n", genderIndex)

	return anon
}

func main() {
	rand.Seed(time.Now().UnixNano())
	charArray := []*Character{}

	for i := 1; i < 50; i++ {
		inputGender = 2
		// inputGender = rand.Intn(3)
		x := createChar(i)
		charArray = append(charArray, x)
		// fmt.Println(anon.charID, anon.charName, anon.charNickname, anon.charSurname, anon.charMeta, anon.charGender)
	}

	// fmt.Println(charArray[7])

	fmt.Printf("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // for building program interface
}
