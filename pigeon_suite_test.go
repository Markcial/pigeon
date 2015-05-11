package Pigeon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPigeon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pigeon Suite")
}
