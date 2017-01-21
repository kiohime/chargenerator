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
	charNation, charGender, charMeta, charName, charNickname, charSurname string
	charID                                                                int
}

//variables of user input mode
var inputInt int
var inputNation int
var inputGender int
var inputMeta int

//variables of current index of current character
var genderIndex int
var nationIndex int
var metaIndex int

func main() {
	rand.Seed(time.Now().UnixNano())
	charArray := []*Character{}
	// SETTING MODES
	setCharNumber()
	// fmt.Println(inputInt)
	inputGender = setCharGender()
	inputNation = setCharNation()
	inputMeta = setCharMeta()

	//START GENERATING
	for i := 1; i <= inputInt; i++ {
		x := createChar(i)
		charArray = append(charArray, x)
		// fmt.Println(anon.charID, anon.charName, anon.charNickname, anon.charSurname, anon.charMeta, anon.charGender)
	}

	// fmt.Println(charArray[7])

	fmt.Printf("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // for building program interface
}
