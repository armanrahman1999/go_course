package intermediates

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, base64 encoding")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
	decoded,err := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(decoded, err)
	fmt.Println(string(decoded))
}