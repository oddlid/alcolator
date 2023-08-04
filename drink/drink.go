package drink

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

// Drink is the basic struct for calculation
type Drink struct {
	VolumeML int     `json:"millilitres"`        // Amount in milliliter
	AlcPCT   float64 `json:"percentage"`         // Alcohol percentage
	Price    float64 `json:"price"`              // Price, in unspecified currency
	Name     string  `json:"name"`               // Name of the drink, for display purposes, not important really
	Currency string  `json:"currency,omitempty"` // For display purposes
}

type Drinks []Drink

// Wrapper for strconv.Atoi
func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

// Wrapper for strconv.ParseFloat
func pfloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0
	}
	return val
}

func DrinkFromStringFields(name, volume, percent, price, currency string) Drink {
	return Drink{
		Name:     name,
		VolumeML: atoi(volume),
		AlcPCT:   pfloat(percent),
		Price:    pfloat(price),
		Currency: currency,
	}
}

// VolumeAsDL returns the drinks volume in DL (deciliter)
func (d Drink) VolumeAsDL() float64 {
	return float64(d.VolumeML) / 100.0
}

// VolumeAsL returns the drinks volume in L (liter)
func (d Drink) VolumeAsL() float64 {
	return float64(d.VolumeML) / 1000.0
}

// AlcML returns the number of millilitres of pure alcohol in the drink
func (d Drink) AlcML() int {
	return int((float64(d.VolumeML) / 100) * d.AlcPCT)
}

// WaterML returns the number of millilitres of water in the drink
func (d Drink) WaterML() int {
	return d.VolumeML - d.AlcML()
}

// PricePerAlcML returns the price per milliliter of alcohol
func (d Drink) PricePerAlcML() float64 {
	return d.Price / float64(d.AlcML())
}

// PricePerDrinkML returns the price per milliliter of the drink itself
func (d Drink) PricePerDrinkML() float64 {
	return d.Price / float64(d.VolumeML)
}

func (ds *Drinks) Load(r io.Reader) error {
	jb, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(jb, ds)
}

func (ds *Drinks) LoadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = ds.Load(file); err != nil {
		return err
	}
	return nil
}

func (ds Drinks) Save(w io.Writer) (int, error) {
	jb, err := json.MarshalIndent(ds, "", "\t")
	if err != nil {
		return 0, err
	}
	jb = append(jb, '\n')
	return w.Write(jb)
}

func (ds Drinks) SaveFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = ds.Save(file); err != nil {
		return err
	}
	return nil
}
