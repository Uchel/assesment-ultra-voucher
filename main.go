package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(strings ...string) [][]string {
	// buat map key string dan value slice string
	anagramMap := make(map[string][]string)

	for _, str := range strings {
		key := sortString(str)
		//key tidak bisa duplikat,jadi bisa pakai sort string sebagai key, valuenya bisa pakai string dari parameter
		anagramMap[key] = append(anagramMap[key], str)

	}

	var anagramSlice [][]string

	for _, v := range anagramMap {

		anagramSlice = append(anagramSlice, v)

	}

	// fmt.Println(anagramSlice)
	return anagramSlice

}

// Buat fungsi sort string untuk mapping di func anagram
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")

}

func main() {
	slice := []string{
		"cook", "save", "taste", "aves", "vase", "state", "map",
	}
	result := anagram(slice...)

	fmt.Println(result)
}
