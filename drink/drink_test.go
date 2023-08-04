package drink

import (
	"testing"
)

func TestVolumeAsDL(t *testing.T) {
	d := Drink{
		Name:     "Vodka",
		VolumeML: 750,
		AlcPCT:   37.5,
	}

	exp := 7.5
	dl := d.VolumeAsDL()

	if dl != exp {
		t.Errorf("Expected %f, got %f", exp, dl)
	}
}

func TestVolumeAsL(t *testing.T) {
	d := Drink{
		Name:     "Vodka",
		VolumeML: 750,
		AlcPCT:   37.5,
	}

	exp := 0.75
	l := d.VolumeAsL()

	if l != exp {
		t.Errorf("Expected %f, got %f", exp, l)
	}
}

func TestAlcML(t *testing.T) {
	d := Drink{
		Name:     "Pure Alc",
		VolumeML: 1000,
		AlcPCT:   96.0,
	}

	exp := 960
	ml := d.AlcML()

	if ml != exp {
		t.Errorf("Expected %d, got %d", exp, ml)
	}
}

func TestWaterML(t *testing.T) {
	d := Drink{
		Name:     "Beer",
		VolumeML: 500,
		AlcPCT:   5.0,
	}

	exp := 475
	ml := d.WaterML()

	if ml != exp {
		t.Errorf("Expected %d, got %d", exp, ml)
	}
}

func TestPricePerAlcML(t *testing.T) {
	d := Drink{
		Name:     "Pure Alc",
		VolumeML: 1000,
		AlcPCT:   96.0,
		Price:    960,
	}

	exp := 1.0
	ppml := d.PricePerAlcML()

	if ppml != exp {
		t.Errorf("Expected %f, got %f", exp, ppml)
	}
}
