package main

import (
	"fmt"
	"math"

	"github.com/Th3Whit3Wolf/go_intro/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(strutil.Reverse("hello"))
}
