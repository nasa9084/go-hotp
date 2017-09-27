package hotp_test

import (
	"testing"

	hotp "github.com/nasa9084/go-hotp"
)

var generator = hotp.Generator{
	Counter: 0,
	Secret:  "12345678901234567890",
	Digit:   6,
}

func TestGenerate(t *testing.T) {
	candidates := []struct {
		expected int64
	}{
		{755224},
		{287082},
		{359152},
		{969429},
		{338314},
		{254676},
		{287922},
		{162583},
		{399871},
		{520489},
	}
	for i, c := range candidates {
		t.Logf("Count: %d\n", i)
		hotp := generator.Generate()
		if hotp != c.expected {
			t.Errorf("%d != %d", hotp, c.expected)
			return
		}
	}
}
