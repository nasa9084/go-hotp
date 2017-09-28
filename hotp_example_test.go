package hotp_test

import (
	"fmt"

	"github.com/nasa9084/go-hotp"
)

func ExampleGenerate() {
	h := hotp.Generator{
		Secret: "some shared secret",
		Digit:  6,
	}
	fmt.Print(h.Generate())
}
