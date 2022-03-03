package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
)

func main() {
	f, err := os.Open("oscar_male.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	name_count := map[string]int{}

    
	for _, v:= range data {
		actor := v[3]
		if _, ok := name_count[actor]; ok {
			name_count[actor] += 1
		} else{
			name_count[actor] = 1
		}
	}
	// fmt.Println(name_count)

	for k, v := range name_count {
		if v > 1 {
			fmt.Println(k)
		}
	}
	// var get_values []int
	// for _, v := range name_count {
	// 	get_values = append(get_values, v)
	// }

	// fmt.Println(math.Max(get_values))
	
	

}
