package customflags

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCustomflags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Customflags Suite")
}
