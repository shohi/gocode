package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println(url.PathEscape("lang:>50"))
}
