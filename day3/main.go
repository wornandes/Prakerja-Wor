package main

import (
	"fmt"
	"strconv"
)

// func ArrayMerge(arrayA, arrayB []string) []string {
// 	// your code here
// 	merged := make(map[string]bool)
// 	result := []string{}

// 	for _, item := range arrayA {
// 		merged[item] = true
// 	}

// 	for _, item := range arrayB {
// 		if _, ok := merged[item]; !ok {
// 			merged[item] = true
// 		}
// 	}

// 	for item := range merged {
// 		result = append(result, item)
// 	}

// 	return result

// }
// func main() {
// 	// Test cases
// 	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
// 	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
// 	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
// 	// ["sergei", "jin", "steve", "bryan"]
// 	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa",
// 		"law"}))
// 	// ["alisa", "yoshimitsu", "devil jin", "law"]
// 	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
// 	// ["devil jin", "sergei"]
// 	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
// 	// ["hwoarang"]
// 	fmt.Println(ArrayMerge([]string{}, []string{}))
// }

// func Mapping(slice []string) map[string]int {
// 	m := make(map[string]int)
// 	for _, item := range slice {
// 		m[item]++
// 	}
// 	return m
// }
// func main() {
// 	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
// 	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
// 	fmt.Println(Mapping([]string{}))                                         //map[]
// }

func munculSekali(angka string) []int {
	counter := make(map[int]int)

	for _, digit := range angka {
		num, _ := strconv.Atoi(string(digit))
		counter[num]++
	}

	angkaUnik := []int{}
	for num, count := range counter {
		if count == 1 {
			angkaUnik = append(angkaUnik, num)
		}
	}

	return angkaUnik

}
func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("8872504"))    // [8 7 2 5 4]
}
