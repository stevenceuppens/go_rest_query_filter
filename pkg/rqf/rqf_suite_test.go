package rqf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRqf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rqf Suite")
}
