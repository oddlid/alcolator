package drink

import (
	"sort"
)

// types and funcs to enable various sorting

// SortByVolume sorts a slice of Drinks by total volume in each drink
func (ds Drinks) SortByVolume() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].VolumeML < ds[j].VolumeML
	})
}

// SortByAlcPCT sorts a slice of Drinks by alcohol percent in each drink
func (ds Drinks) SortByAlcPCT() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].AlcPCT < ds[j].AlcPCT
	})
}

// SortByPrice sorts a slice of Drinks by price of each drink
func (ds Drinks) SortByPrice() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Price < ds[j].Price
	})
}

// SortByName sorts a slice of Drinks by each drink name
func (ds Drinks) SortByName() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Name < ds[j].Name
	})
}

// SortByAlcML sorts a slice of Drinks by amount of pure alcohol in each drink
func (ds Drinks) SortByAlcML() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].AlcML() < ds[j].AlcML()
	})
}

// SortByWaterML sorts a slice of Drinks by amount of water in each drink
func (ds Drinks) SortByWaterML() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].WaterML() < ds[j].WaterML()
	})
}

// SortByPricePerAlcML sorts a slice of Drinks by the price per ml of alcohol in each drink
func (ds Drinks) SortByPricePerAlcML() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].PricePerAlcML() < ds[j].PricePerAlcML()
	})
}

// SortByPricePerDrinkML sorts a slice of Drinks by the price per ml of the drink
func (ds Drinks) SortByPricePerDrinkML() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].PricePerDrinkML() < ds[j].PricePerDrinkML()
	})
}

