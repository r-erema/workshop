package test

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func GoBuild(binDir, outputBinFile string, ldFlags ...string) {
	args := append([]string{
		"build", "-buildvcs=false", "-o", outputBinFile, "-ldflags",
	}, ldFlags...)
	args = append(args, binDir)

	//nolint:gosec //G204
	cmdOutput, err := exec.CommandContext(GinkgoT().Context(), "go", args...).
		Output()
	Expect(err).NotTo(HaveOccurred())
	_, err = fmt.Fprintf(GinkgoWriter, "Building binary: %d\n", cmdOutput)
	Expect(err).NotTo(HaveOccurred())

	GinkgoT().DeferCleanup(func() {
		err = os.Remove(outputBinFile)
		Expect(err).NotTo(HaveOccurred())
	})
}
