package nou

import "testing"

func TestNou(t *testing.T) {
	// DispUsage Test
	// no parameter
	DispUsage("")
	// set parameter
	DispUsage("nou test")

	// usageMessge Test
	usageMessage()

	// splitArgs Test
	splitArgs("/do/test/split/args/test")
}
