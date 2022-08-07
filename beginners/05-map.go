package main

import "fmt"

func main() {
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)

	statePopulations := map[string]int{
		"California": 39_250_017,
		"Texas":      27_862_596,
		"Florida":    20_612_439,
	}
	sp := statePopulations
	statePopulations["Geogia"] = 10_310_371
	delete(statePopulations, "Florida")
	fmt.Println(statePopulations)
	fmt.Println(sp) // map is copied by reference
	fmt.Println(statePopulations["Texas"])
	fmt.Println(len(statePopulations))

	floridaPopulation, ok := statePopulations["Florida"]
	fmt.Printf("%v %v\n", floridaPopulation, ok)

}
