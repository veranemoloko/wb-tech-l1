package main

import "fmt"

func main() {

	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.0}

	groupData := make(map[int][]float64)

	var step int

	for _, t := range data {
		step = (int(t) / 10) * 10
		groupData[step] = append(groupData[step], t)

	}

	fmt.Println(groupData)

}
