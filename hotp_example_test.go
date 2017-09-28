package hotp_test

import (
	"fmt"

	"github.com/nasa9084/go-hotp"
)

func ExampleGenerate() {
	h := hotp.Generator{
		Secret: "12345678901234567890", // RFC4226 Appendix D Test Value
		Digit:  6,
	}
	fmt.Print(h.Generate())
	// Output:
	// 755224
}
