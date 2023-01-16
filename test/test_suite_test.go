package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

//nolint:gochecknoglobals // not sure how to avoid this
var (
	buildCmdPath   string
	detectCmdPath  string
	releaseCmdPath string
	runCmdPath     string
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
	var err error

	buildCmdPath, err = gexec.Build("github.com/hyperledger-labs/fabric-builder-k8s/cmd/build")
	Expect(err).NotTo(HaveOccurred())

	detectCmdPath, err = gexec.Build("github.com/hyperledger-labs/fabric-builder-k8s/cmd/detect")
	Expect(err).NotTo(HaveOccurred())

	releaseCmdPath, err = gexec.Build("github.com/hyperledger-labs/fabric-builder-k8s/cmd/release")
	Expect(err).NotTo(HaveOccurred())

	runCmdPath, err = gexec.Build("github.com/hyperledger-labs/fabric-builder-k8s/cmd/run")
	Expect(err).NotTo(HaveOccurred())

	// p := script.Exec("kind create cluster")
	// p.Stdout()
	// Expect(p.ExitStatus()).To(Equal(0))
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	// script.Exec("kind delete cluster").Stdout()
})
