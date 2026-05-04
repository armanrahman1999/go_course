package intermediate

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixTime := now.Unix()
	t:= time.Unix(unixTime, 0)
	fmt.Println("unixTime", unixTime)
	fmt.Println("now", now)
	fmt.Println("t", t)
}