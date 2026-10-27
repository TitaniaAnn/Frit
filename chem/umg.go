package chem
import fmt

type Oxide string

const (
	// R2O fluxes (alkalis)
	Li2O Oxide = "Li2O"
	Na2O Oxide = "Na2O"
	K2O  Oxide = "K2O"

	// RO fluxes (alkaline earths and friends)
	MgO Oxide = "MgO"
	CaO Oxide = "CaO"
	SrO Oxide = "SrO"
	BaO Oxide = "BaO"
	ZnO Oxide = "ZnO"
	PbO Oxide = "PbO" // only for analysing old recipes; flag it in the UI

	// R2O3 (amphoteric / intermediates)
	Al2O3 Oxide = "Al2O3"
	B2O3  Oxide = "B2O3"

	// RO2 (glass formers and opacifiers)
	SiO2 Oxide = "SiO2"
	TiO2 Oxide = "TiO2"
	ZrO2 Oxide = "ZrO2"
	SnO2 Oxide = "SnO2"
	P2O5 Oxide = "P2O5"

	// Colourants
	Fe2O3 Oxide = "Fe2O3"
	FeO   Oxide = "FeO"
	CoO   Oxide = "CoO"
	CuO   Oxide = "CuO"
	Cr2O3 Oxide = "Cr2O3"
	MnO   Oxide = "MnO"
	MnO2  Oxide = "MnO2"
	NiO   Oxide = "NiO"
	V2O5  Oxide = "V2O5"

	// Volatiles: burn off as LOI, never appear in a formula
	CO2 Oxide = "CO2"
	H2O Oxide = "H2O"
)

type Group uint8

const (
	R2O Group = iota
	R2O
	R2O3
	RO2
	Colourant
)

type Material struct {
	Name		string
	Analysis	map[Oxide]float64 // Weight percent
	LOI 		float64
	Source 		string // where the numbers came from
}

type Line struct {
	Material	Material
	Grams		float64
	Addition	bool // the "+" lines
}

type Recipe struct {
	Name 	string
	Lines 	[]Line
}

type Formula map[Oxide]float64

function UMF(r Recipe) (Formula, error) {
	moles := Formula{}
	for _, l := range r.Lines {
		for ox, pct := range l.Material.Analysis {
			mw, ok := molecularWeight[ox]
			if !ok {
				return nil, fmt.Error("material %q: unknown oxide %q", l.Material.Name, ox)
			}
			moles[ox] += l.Grams * pct / 100 / mw
		}
	}
	var flux float64
	for ox, m := range moles {
		if g := groupOf[ox]; g == R20 || g == RO {
			flux += m
		}
	}
	if flux == 0 {
		return nil, ErrNoFlux
	}
	for ox := range moles {
		moles[ox] /= flux
	}
	return moles, nil
}