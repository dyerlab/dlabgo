package genetic

import (
	"testing"
)

func TestFrequencies(t *testing.T) {

	var freqs Frequencies

	var vals = freqs.Alleles()

	if len(vals) != 0 {
		t.Error("Empty frequency object not zero alleles")
	}

	alleles := [][]string{{"A", "B"}, {"B", "A"}, {"A", "A"}}
	var loc1 Genotype
	var loc2 Genotype
	var loc3 Genotype
	loc1.Alleles = alleles[0][0:2]
	loc2.Alleles = alleles[1][0:2]
	loc3.Alleles = alleles[2][0:2]

	freqs.AddGenotype(loc1)
	freqs.AddGenotype(loc2)
	freqs.AddGenotype(loc3)

	vals = freqs.Alleles()
	if len(vals) != 2 {
		t.Errorf("Freq should have two alleles but had %d instead", len(vals))
	}

	if freqs.GetFrequency("Z") != 0.0 {
		t.Error("Bad frequency for missing allele.")
	}
	if freqs.GetFrequency("A") != 4.0/6.0 {
		t.Error("Bad frequency for present allele.")
	}
}
