package alcolator


type Drink struct {
	Name     string
	VolumeML int
	AlcPCT   float64
	Price    float64
}

func (d *Drink) VolumeAsDL() float64 {
	return float64(d.VolumeML) / 100.0
}

func (d *Drink) VolumeAsL() float64 {
	return float64(d.VolumeML) / 1000.0
}

func (d *Drink) AlcML() int {
	return int((float64(d.VolumeML) / 100) * d.AlcPCT)
}

func (d *Drink) WaterML() int {
	return d.VolumeML - d.AlcML()
}

func (d *Drink) PricePerAlcML() float64 {
	return d.Price / float64(d.AlcML())
}

func (d *Drink) PricePerDrinkML() float64 {
	return d.Price / float64(d.VolumeML)
}

