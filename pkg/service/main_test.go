package service

import (
	"testing"

	"github.com/threeal/bro/pkg/internal"
)

func TestMain(m *testing.M) {
	internal.RunTestWithCoverage(m, 1)
}
