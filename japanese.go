package main

import (
	"fmt"
)

func japaneseNames() []string {
	names := []string{"NO JAPANESE NAMES ARRAY"}
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	switch genderIndex {
	case 0:
		names = []string{
			//MALE
			"Saito",
			"Kouta",
			"Sousuke",
			"Hiro",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	case 1:
		names = []string{
			//FEMALE
			"Himiko",
			"Chidori",
			"Nana",
			"Sakura",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	case 2:
		names = []string{
			//UNISEX
			"Hajime",
			"Jun",
			"Yuki",
			"Masa",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	default:
		fmt.Println("Error in Japanese Names Array")
		return nil
	}
	return names
}

//############################################################################################################################################################################################
func japaneseSurNames() []string {
	surNames := []string{"NO JAPANESE SURNAMES ARRAY"}
	surNames = []string{
		"Tanaka",
		"Yamato",
		"Okazaki",
		"Kawaguchi",
	}
	return surNames
}
