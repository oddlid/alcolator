package drink

type Drink struct {
	Name       string
	Comment    string
	Percentage float64
}

type DrinkContainer struct {
	VolumeUnit string
	PriceUnit  string
	Comment    string
	Drink      Drink
	Volume     float64
	Price      float64
}

func percentOf(volume, percent float64) float64 {
	return (volume / 100.0) * percent
}

func volumeRequiredToMatchAlcoholContent(volumeOfFirst, percentageOfFirst, percentageOfOther float64) float64 {
	if percentageOfFirst == percentageOfOther {
		return volumeOfFirst
	}
	if volumeOfFirst == 0.0 || percentageOfFirst == 0.0 || percentageOfOther == 0.0 {
		return 0.0
	}
	return (percentageOfFirst / percentageOfOther) * volumeOfFirst
}

func (d *Drink) AlcoholAmount(volume float64) float64 {
	if d == nil {
		return 0.0
	}
	return percentOf(volume, d.Percentage)
}

func (d *Drink) WaterAmount(volume float64) float64 {
	if d == nil {
		return 0.0
	}
	return volume - d.AlcoholAmount(volume)
}

func (dc *DrinkContainer) AlcoholAmount() float64 {
	if dc == nil {
		return 0.0
	}
	return dc.Drink.AlcoholAmount(dc.Volume)
}

func (dc *DrinkContainer) AlcoholPrice() float64 {
	if dc == nil {
		return 0.0
	}
	if dc.Volume == 0.0 || dc.Price == 0.0 {
		return 0.0
	}
	alcoholAmount := dc.Drink.AlcoholAmount(dc.Volume)
	if alcoholAmount == 0.0 {
		return 0.0
	}
	return dc.Price / alcoholAmount
}

// VolumeForPercentage returns the volume required of the given percentage to match the amount of alcohol
// in this DrinkContainer instance.
func (dc *DrinkContainer) VolumeForPercentage(percentage float64) float64 {
	if dc == nil {
		return 0.0
	}
	return volumeRequiredToMatchAlcoholContent(dc.Volume, dc.Drink.Percentage, percentage)
}

func (dc *DrinkContainer) adjustPriceFromVolume(newVolume float64) {
	if dc == nil {
		return
	}
	if newVolume == 0.0 {
		dc.Price = 0.0
		return
	}
	if dc.Price != 0 {
		dc.Price = (dc.Price / dc.Volume) * newVolume
	}
}

func (dc *DrinkContainer) CompareTo(other *DrinkContainer) DrinkContainer {
	if dc == nil || other == nil {
		return DrinkContainer{}
	}

	adjustedVolume := dc.VolumeForPercentage(other.Drink.Percentage)
	otherClone := *other
	otherClone.adjustPriceFromVolume(adjustedVolume)
	otherClone.Volume = adjustedVolume

	return otherClone
}
