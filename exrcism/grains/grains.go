package grains

import (
	"errors"
	"math"
)



func Total() uint64 {

	var grains uint64

	grains = (1 << 64) - 1

	return uint64(grains)
}


func Square(squares int) (uint64, error) {
	if squares <= 0 || squares > 64 {
		return 0, errors.New("wrong squares")
	}

	power := squares - 1

	grains := uint64(math.Pow(2, float64(power)))


	return grains, nil
}
