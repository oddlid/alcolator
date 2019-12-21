package alcolator


// Drink is the basic struct for calculation
type Drink struct {
	VolumeML int     // Amount in milliliter
	AlcPCT   float64 // Alcohol percentage
	Price    float64 // Price, in unspecified currency
	Name     string  // Name of the drink, for display purposes, not important really
	Currency string  // For display purposes
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

