package integration_test

import (
	"github.com/bitfield/script"
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Integration tests", func() {
	It("Blends", func() {
		script.Echo("FAB").Stdout()
		script.Exec(detectCmdPath).Stdout()
	})
})
