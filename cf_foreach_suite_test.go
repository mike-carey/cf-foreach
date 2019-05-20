package cfforeach_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfForeach(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfForeach Suite")
}
