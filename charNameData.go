package main

//METATYPES LIST
// func metatypes() []string {
// 	metatype := []string{
// 		"human",
// 		"dwarf",
// 		"ork",
// 		"elf",
// 		"troll",
// 	}
// 	return metatype
// }

//NAMES LIST
func names(nationIndex int) []string {
	namesArr := []string{"NO ARRAY"}
	// fmt.Println("~~~~~~~~~~~~~~", nationIndex)
	switch nationIndex {
	case 0:
		namesArr = americanNames()
	case 1:
		namesArr = japaneseNames()
	case 2:
		namesArr = chineseNames()
	case 3:
		namesArr = indianNames()
	case 4:
		namesArr = russianNames()
	}

	return namesArr
}

//SURNAMES LIST
func surnames(nationIndex int) []string {
	surNamesArr := []string{"NO ARRAY"}
	switch nationIndex {
	case 0:
		surNamesArr = americanSurNames()
	case 1:
		surNamesArr = japaneseSurNames()
	case 2:
		surNamesArr = chineseSurNames()
	case 3:
		surNamesArr = indianSurNames()
	case 4:
		surNamesArr = russianSurNames()
	}
	return surNamesArr
}
