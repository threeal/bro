package internal

import (
	"fmt"
	"log"
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

func Quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}
