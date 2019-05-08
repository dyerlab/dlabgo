package main

import (
	"fmt"

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

}
