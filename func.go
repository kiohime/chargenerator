package main

import (
	"fmt"
	"math/rand"
	"strings"

	ansi "github.com/k0kubun/go-ansi"
)

/////////////////////////////////////////////////////////////////////
/////////////////////////////CHAR QUANTITY///////////////////////////
/////////////////////////////////////////////////////////////////////

func setCharNumber() int {
	var cond bool
	var input string
	for !cond {
		fmt.Print("Enter number of characters: ")
		fmt.Scanln(&input)
		inputInt = parseNumber(input)
		if inputInt < 0 {
			fmt.Println("Error in number input")
		} else {
			cond = true
		}
	}
	return inputInt
}

func parseNumber(s string) int {
	var value int
	if len(s) == 0 {
		value = -1
		// value = 10
		// fmt.Printf("Default option [%d] selected\n\n", value)
		return value
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

/////////////////////////////////////////////////////////////////////
/////////////////////////////CHAR GENDER/////////////////////////////
/////////////////////////////////////////////////////////////////////

func setCharGender() int {
	var cond bool
	var inputGen int
	for !cond {
		var input string
		fmt.Printf("\nEnter gender index (0 for \"male\", 1 for \"female\", 2 for \"unisex\"): ")
		fmt.Printf("Enter nationality index: ")
		fmt.Printf("\n0: Male\n1: Female\n2: Unisex")
		fmt.Printf("\nNOTHING: Random nationality\n")
		fmt.Scanln(&input)
		inputGen = parseGender(input)
		// fmt.Println("inputGen", inputGen)
		switch {
		case inputGen < 0:
			fmt.Println("Error in charGender input")
		case inputGen == 3:
			fmt.Printf("Random option selected\n\n")
			cond = true
		default:
			cond = true
		}
		// fmt.Println(inputGen)
	}
	return inputGen
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
		}
	}
	return value
}

func decideGend() int {
	switch {
	case inputGender == 3:
		genderIndex = rand.Intn(3) //change number of genders
	default:
		genderIndex = inputGender
	}
	return genderIndex
}

// int to sting for gender
func deferGend(dg int) string {
	var gender string
	switch dg { // add new genders here
	case 0:
		gender = "male"
	case 1:
		gender = "female"
	case 2:
		gender = "unisex"
	}
	return gender
}

/////////////////////////////////////////////////////////////////////
/////////////////////////////CHAR NATIONALITY////////////////////////
/////////////////////////////////////////////////////////////////////

func setCharNation() int {
	var cond bool
	var inputNat int
	for !cond {
		var input string
		fmt.Printf("Enter nationality index: ")
		fmt.Printf("\n0: American\n1: Japanese\n2: Chinese\n3: Indian\n4: Russian")
		fmt.Printf("\n5. Spanish\n6: Italian\n7: French\n8: German\nNOTHING: Random nationality\n")
		fmt.Scanln(&input)
		inputNat = parseNation(input)
		// fmt.Println(inputNat)
		// fmt.Println("inputGen", inputGen)
		switch {
		case inputNat < 0:
			fmt.Println("Error in nationality input")
		case inputNat == 9:
			fmt.Printf("Random option selected\n\n")
			cond = true
		default:
			cond = true
		}
		// fmt.Println(inputGen)
	}
	return inputNat
}

func parseNation(s string) int {
	var value int
	if len(s) == 0 {
		return 9
	}
	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '8':
			value = int(ch) - int('0') + value*10
		default:
			value = -1
		}
	}
	return value
}

func decideNat() int {
	switch {
	case inputNation == 9:
		nationIndex = rand.Intn(9) //change number of genders
	default:
		nationIndex = inputNation
	}
	// fmt.Println("nationindex", nationIndex)
	return nationIndex
}

func deferNat(dn int) string {
	var nation string
	switch dn { // add new genders here
	case 0:
		nation = "american"
	case 1:
		nation = "japanese"
	case 2:
		nation = "chinese"
	case 3:
		nation = "indian"
	case 4:
		nation = "russian"
	case 5:
		nation = "spanish"
	case 6:
		nation = "italian"
	case 7:
		nation = "french"
	case 8:
		nation = "german"
	}
	return nation
}

/////////////////////////////////////////////////////////////////////
/////////////////////////////CHAR METATYPE///////////////////////////
/////////////////////////////////////////////////////////////////////

func setCharMeta() int {
	var cond bool
	var inputMeta int
	for !cond {
		var input string
		fmt.Printf("Enter metatype index: ")
		fmt.Printf("Enter nationality index: ")
		fmt.Printf("\n0: Human\n1: Dwarf\n2: Elf\n3: Ork\n4: Troll")
		fmt.Printf("\nNOTHING: Random metatype\n")
		fmt.Scanln(&input)
		inputMeta = parseMeta(input)
		// fmt.Println(inputMeta)
		// fmt.Println("inputGen", inputGen)
		switch {
		case inputMeta < 0:
			fmt.Println("Error in metatype input")
		case inputMeta == 5:
			fmt.Printf("Random option selected\n\n")
			cond = true
		default:
			cond = true
		}
	}
	return inputMeta
}

func parseMeta(s string) int {
	var value int
	if len(s) == 0 {
		return 5
	}
	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '4':
			value = int(ch) - int('0') + value*10
		default:
			value = -1
		}
	}
	return value
}

func decideMeta() int {
	switch {
	case inputMeta == 5:
		metaIndex = rand.Intn(5) //change number of genders
	default:
		metaIndex = inputMeta
	}
	// fmt.Println(nationIndex)
	return metaIndex
}

func deferMeta(dm int) string {
	var meta string
	switch dm { // add new genders here
	case 0:
		meta = "human"
	case 1:
		meta = "dwarf"
	case 2:
		meta = "elf"
	case 3:
		meta = "ork"
	case 4:
		meta = "troll"
	}
	return meta
}

/////////////////////////////////////////////////////////////////////
/////////////////////////////ANOTHER/////////////////////////////////
/////////////////////////////////////////////////////////////////////

func getRandStr(a string) string {
	var file []string
	var gRand string
	switch a {
	case "name":
		file = names(nationIndex)
	case "surname":
		file = surnames(nationIndex)
	case "nickname":
		file = nicknames()
		// case "metatype":
		// 	file = metatypes()
	}
	gRand = file[rand.Intn(len(file))] //randomzer from slice
	return gRand
}

func createChar(i int) *Character {
	anon := new(Character)
	// 	// assert(i > 1, "yhmjfdyhjdjhdfyhj")
	anon.charID = i
	anon.charNation = deferNat(decideNat())
	anon.charMeta = deferMeta(decideMeta())
	anon.charGender = deferGend(decideGend())
	anon.charName = getRandStr("name")
	anon.charSurname = getRandStr("surname")
	anon.charNickname = getRandStr("nickname")
	ansi.Println(i, anon)
	// fmt.Println(anon.charID, anon.charName, "~"+anon.charNickname+"~", anon.charSurname, anon.charNation, anon.charMeta, anon.charGender)
	return anon
}

func colorDecider(cd string) string {
	var color string
	// black := "\x1b[30m"
	// red := "\x1b[31m"
	// green := "\x1b[32m"
	// yellow := "\x1b[33m"
	// blue := "\x1b[34m"
	// magenta := "\x1b[35m"
	cyan := "\x1b[36m"
	// white := "\x1b[37m"

	// blackB := "\x1b[30;1m"
	redB := "\x1b[31;1m"
	greenB := "\x1b[32;1m"
	yellowB := "\x1b[33;1m"
	blueB := "\x1b[34;1m"
	magentaB := "\x1b[35;1m"
	cyanB := "\x1b[36;1m"
	whiteB := "\x1b[37;1m"
	// fmt.Println("start", cyan)
	switch cd {
	case "nickname":
		color = cyan
	case "gender":
		switch genderIndex {
		case 0:
			color = cyanB
		case 1:
			color = magentaB
		case 2:
			color = whiteB
		}
	case "meta":
		switch metaIndex {
		case 0: //human
			color = whiteB
		case 1: //dwarf
			color = yellowB
		case 2: //elf
			color = greenB
		case 3: //ork
			color = redB
		case 4: //troll
			color = blueB

		}
	}

	return color
}

func (c *Character) String() string {
	return strings.Join([]string{c.charName, c.charSurname + ",", colorDecider("nickname") + c.charNickname + "\x1b[0m,", c.charNation, colorDecider("meta") + c.charMeta + "\x1b[0m", colorDecider("gender") + c.charGender + "\x1b[0m"}, " ")
}
