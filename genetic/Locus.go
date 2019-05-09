package genetic

// Locus structure collects genotypes
type Locus struct {
	Name              string
	genotypes         []Genotype
	AlleleFrequencies Frequencies
}

// AddGenotype is the entry for adding genetic information to this locus
func (l Locus) AddGenotype(g Genotype) {
	l.AlleleFrequencies.AddGenotype(g)
	l.genotypes = append(l.genotypes, g)
}

// Genotypes is the official getter (to enforce adding through AddGenotype())
func (l Locus) Genotypes() []Genotype {
	return l.genotypes
}
