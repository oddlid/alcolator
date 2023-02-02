package drink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_percentOf(t *testing.T) {
	assert.Equal(t, 50.0, percentOf(100, 50))
}

func Test_volumeRequiredToMatchAlcoholContent(t *testing.T) {
	v1 := 100.0
	p1 := 50.0
	p2 := 50.0
	assert.Equal(t, v1, volumeRequiredToMatchAlcoholContent(v1, p1, p2))
	assert.Zero(t, volumeRequiredToMatchAlcoholContent(0, p1, p2))
	assert.Zero(t, volumeRequiredToMatchAlcoholContent(v1, 0, p2))
	assert.Zero(t, volumeRequiredToMatchAlcoholContent(v1, p1, 0))

	assert.Equal(t, 1000.0, volumeRequiredToMatchAlcoholContent(500, 10, 5))
	assert.Equal(t, 500.0, volumeRequiredToMatchAlcoholContent(1000, 5, 10))

	// check that one 750ml bottle of 13% is about the same as 3x500ml of 6.4%
	t.Logf("750 ml of 13%% = %f ml of 6.4%%", volumeRequiredToMatchAlcoholContent(750, 13, 6.4))
}

func Test_Drink_AlcoholAmount(t *testing.T) {
	assert.Zero(t, (*Drink)(nil).AlcoholAmount(0))
	d := Drink{
		Percentage: 6.4,
	}
	assert.Equal(t, 32.0, d.AlcoholAmount(500))
}

func Test_Drink_WaterAmount(t *testing.T) {
	assert.Zero(t, (*Drink)(nil).WaterAmount(0))
	d := Drink{
		Percentage: 50.0,
	}
	assert.Equal(t, 50.0, d.WaterAmount(100))
}

func Test_DrinkContainer_AlcoholPrice(t *testing.T) {
	assert.Zero(t, (*DrinkContainer)(nil).AlcoholPrice())

	dc := DrinkContainer{}
	assert.Zero(t, dc.AlcoholPrice())

	dc.Volume = 100
	assert.Zero(t, dc.AlcoholPrice())

	dc.Price = 100
	assert.Zero(t, dc.AlcoholPrice())

	dc.Drink = Drink{Percentage: 50.0}
	assert.Equal(t, 2.0, dc.AlcoholPrice())
}

func Test_DrinkContainer_VolumeForPercentage(t *testing.T) {
	assert.Zero(t, (*DrinkContainer)(nil).VolumeForPercentage(0))

	dc := DrinkContainer{
		Drink: Drink{
			Percentage: 5,
		},
		Volume: 100,
	}
	assert.Equal(t, 50.0, dc.VolumeForPercentage(10))
}

func Test_DrinkContainer_adjustPriceFromVolume(t *testing.T) {
	assert.NotPanics(t, func() {
		(*DrinkContainer)(nil).adjustPriceFromVolume(0)
	})

	dc := DrinkContainer{}
	dc.adjustPriceFromVolume(100)
	assert.Zero(t, dc)

	dc.Price = 100
	dc.adjustPriceFromVolume(100)
	assert.Zero(t, dc.Volume)

	// verify that a new volume of 0 resets both price and volume to 0
	dc.Volume = 100
	dc.adjustPriceFromVolume(0)
	assert.Zero(t, dc)

	dc.Price = 100
	dc.Volume = 100
	dc.adjustPriceFromVolume(200)
	assert.Equal(t, 200.0, dc.Price)
	assert.Equal(t, 200.0, dc.Volume)
}

func Test_DrinkContainer_CompareTo(t *testing.T) {
	assert.Empty(t, (*DrinkContainer)(nil).CompareTo(nil))

	dc1 := DrinkContainer{
		Drink: Drink{
			Percentage: 10,
		},
		Volume: 100,
	}

	dc2 := DrinkContainer{
		Drink: Drink{
			Percentage: 5,
		},
		Volume: 100,
		Price:  100,
	}

	dc3 := dc1.CompareTo(&dc2)
	assert.Equal(t, dc2.Volume*2, dc3.Volume)
	assert.Equal(t, dc2.Price*2, dc3.Price)
}

func Benchmark_DrinkContainer_CompareTo(b *testing.B) {
	dc1 := DrinkContainer{
		Drink: Drink{
			Percentage: 10,
		},
		Volume: 100,
	}

	dc2 := DrinkContainer{
		Drink: Drink{
			Percentage: 5,
		},
		Volume: 100,
		Price:  100,
	}

	for i := 0; i < b.N; i++ {
		_ = dc1.CompareTo(&dc2)
	}
}
