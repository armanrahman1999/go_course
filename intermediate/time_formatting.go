package intermediate

import (
	"fmt"
	"time"
)

func main() {

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	fmt.Println(t, err)
}