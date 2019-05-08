package genetic

import (
	"fmt"
	"sort"
	"strconv"
)

// Frequencies a struct for allele frequences at a locus
type Frequencies struct {
	counts map[string]float64
	n      float64
	ndip   float64
	nhet   float64
}

// AddGenotype is method for adding data to the struct
func (f *Frequencies) AddGenotype(g Genotype) {

	fmt.Printf("Adding genotype %s\n", g)

	if f.counts == nil {
		f.counts = make(map[string]float64)
	}

	for _, v := range g.Alleles {
		fmt.Printf("adding %s\n", v)
		f.counts[v] += 1.0
	}

	f.n += float64(g.Ploidy())
	if g.Ploidy() == 2 {
		f.ndip += 1.0
		if g.IsHeterozygote() {
			f.nhet += 1.0
		}
	}

}

// GetFrequency returns specific allele frequencies
func (f Frequencies) GetFrequency(allele string) float64 {
	if v, found := f.counts[allele]; found {
		return v / f.n
	}
	return 0.0
}

// Alleles returns all the alleles entered
func (f Frequencies) Alleles() []string {
	var ret []string
	for k := range f.counts {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// String overload
func (f Frequencies) String() string {
	ret := "Allele Frequencies:\n"
	for k := range f.counts {
		freq := f.GetFrequency(k)
		ret += k + " : " + strconv.FormatFloat(freq, 'f', -1, 64) + "\n"
	}
	return ret
}
