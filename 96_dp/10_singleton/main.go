package main

import (
	"fmt"

	"github.com/rebirthmonkey/go/96_dp/10_singleton/hungry"
	"github.com/rebirthmonkey/go/96_dp/10_singleton/lazy"
)

func main() {
	fmt.Println("the mode of hungry is", hungry.GetInsOr())
	fmt.Println("the mode of layz is", lazy.GetInsOr())
}
