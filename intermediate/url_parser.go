package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	rawUrl := "https://www.youtube.com/watch?v=abc123#fragment"

	parsedUrl, err := url.Parse(rawUrl)
	fmt.Println(parsedUrl, err)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Port())
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Fragment)
	fmt.Println(parsedUrl.Path)
}