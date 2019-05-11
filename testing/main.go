package main

import (
	"fmt"

	"encoding/json"

	"github.com/dyerlab/dlabgo/genetic"
)

func main() {
	fmt.Println("genetic testing")

	alleles := [][]string{{"A", "B"}, {"B", "A"}, {"A", "A"}}

	fmt.Println(alleles)

	var loc1 genetic.Genotype
	var loc2 genetic.Genotype
	var loc3 genetic.Genotype

	loc1.Alleles = alleles[0][0:2]
	loc2.Alleles = alleles[1][0:2]
	loc3.Alleles = alleles[2][0:2]

	fmt.Println(loc1)
	fmt.Println(loc2)
	fmt.Println(loc3)

	var freqs genetic.Frequencies
	freqs.AddGenotype(loc1)
	freqs.AddGenotype(loc2)
	freqs.AddGenotype(loc3)

	fmt.Println(freqs)

	fmt.Println("Running some JSON tests")

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	locB, _ := json.Marshal(loc1)
	fmt.Println(string(locB))

	freB, _ := json.Marshal(freqs)
	fmt.Println(string(freB))

	fmt.Println("Done running json")

}
