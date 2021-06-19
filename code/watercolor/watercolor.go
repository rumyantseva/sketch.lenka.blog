package watercolor

import (
	"fmt"
	"log"
	"strings"
)

const (
	WNLF1 string = "*"
	WNLF2 string = "**"
	WNLF3 string = "***"
)

type Staining string

const (
	NonStaining  Staining = "non-staining"
	SemiStaining Staining = "semi-staining"
	FullStaining Staining = "full-staining"
)

type Transparency string

const (
	TrTransp     Transparency = "transparent"
	TrSemiTransp Transparency = "semi-transparent"
	TrSemiOpaque Transparency = "semi-opaque"
	PrOpaque     Transparency = "opaque"
)

type Paint struct {
	Brand         string
	ID            string
	Title         string
	Pigments      []string
	Lightfastness string
	Staining      Staining
	Transparency  Transparency
	Granulating   bool
}

func (p *Paint) ParseWhiteNights(raw string) error {
	// Parse Lightfastness
	if strings.Contains(raw, "✱✱✱") {
		p.Lightfastness = WNLF3
	} else if strings.Contains(raw, "✱✱") {
		p.Lightfastness = WNLF2
	} else if strings.Contains(raw, "✱") {
		p.Lightfastness = WNLF1
	} else {
		return fmt.Errorf("can't parse lightfastness")
	}

	// Parse Staining
	if strings.Contains(raw, "◮") {
		p.Staining = SemiStaining
	} else if strings.Contains(raw, "△") {
		p.Staining = NonStaining
	} else if strings.Contains(raw, "▲") {
		p.Staining = FullStaining
	} else {
		log.Printf("can't parse staining")
	}

	// Parse Transparency
	if strings.Contains(raw, "□") {
		p.Transparency = TrSemiTransp
	} else if strings.Contains(raw, "■") {
		p.Transparency = TrSemiOpaque
	} else if strings.Contains(raw, "◨") {
		p.Transparency = TrSemiTransp
	} else {
		return fmt.Errorf("can't parse transparency")
	}

	// Parse Granulating
	if strings.Contains(raw, `"G"`) {
		p.Granulating = true
	}

	return nil
}
