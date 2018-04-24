package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestView(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "View Suite")
}
