package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main() {
	f, _ := os.Open("oscar_male.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()

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
}
