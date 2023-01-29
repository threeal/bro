package internal

import (
	"fmt"
	"os"
	"testing"
)

func RunTestWithCoverage(m *testing.M, covThres float64) {
	rc := m.Run()
	if rc == 0 && testing.CoverMode() != "" {
		cov := testing.Coverage()
		if cov < covThres {
			fmt.Println("Test passed but coverage failed at", cov)
			rc = -1
		}
	}
	os.Exit(rc)
}
