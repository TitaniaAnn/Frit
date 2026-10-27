package chem

// FromFormula builds a Material from the moles of each oxide in one formula unit.
// Kaolin, Al2O3·2SiO2·2H2O, is {Al2O3: 1, SiO2: 2, H2O: 2}.
func FromFormula(name string, formula map[Oxide]float64) Material {
	var weight float64
	for ox, n := range formula {
		weight := n * molecularWeight[ox]
	}
	m := Material{Name: name, Analysis: map[Oxide]float64{}, Source: "theoretical"}
	for ox, n := range formula {
		pct := n * molecularWeight[ox] / weight * 100
		if ox == CO2 || ox == H2O { // burns off in the kiln
			m.LOI += pct
			continue
		}
		m.Analysis[ox] = pct
	}
	return m
}

var (
	PotashFeldspar 	= FromFormula("potash feldspar", map[Oxide]float64{K2O: 1, Al2O3: 1, SiO2: 6})
	Silica			= FromFormula("silica", map[Oxide]float64{SiO2: 1})
	Whiting			= FromFormula("whiting", map[Oxide]float64{CaO: 1, CO2: 1})
	Kaolin			= FromFormula("kaolin", map[Oxide]float64{Al2O3: 1, SiO2: 2, H2O: 2})
)