package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Error biatch")
	}
	return 0, nil
}

func main() {
	x, err := sqrt(-10)
	fmt.Println(x, err)
}