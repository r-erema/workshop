package cri_cni_flow_test

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/containerd/containerd/v2/client"
	_ "github.com/containerd/containerd/v2/cmd/containerd/builtins"
	"github.com/containerd/containerd/v2/integration/remote"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/r-erema/workshop/utils/test"
	"google.golang.org/grpc"
	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	mockBinFile         = "cni_mock_bin"
	mockLoopbackBinFile = "loopback"
	logFile             = "mock-cni.log"
)

var cli *client.Client

func TestCRI_CNI(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)

	RunSpecs(t, "CRI and CNI suite")
}

var _ = BeforeSuite(func() {
	cniBinDir, err := filepath.Abs("./cni_mock")
	Expect(err).NotTo(HaveOccurred())
	cniConfDir, err := filepath.Abs("./configs")
	Expect(err).NotTo(HaveOccurred())

	ldFlag := "-X=main.logFilePath=" + logFile
	test.GoBuild(cniBinDir, filepath.Join(cniBinDir, mockLoopbackBinFile), ldFlag)
	test.GoBuild(cniBinDir, filepath.Join(cniBinDir, mockBinFile), ldFlag)

	defaultNamespace := "test_ns_" + strconv.FormatInt(GinkgoT().RandomSeed(), 10)

	cli = test.RunContainerdEnv(GinkgoT().TempDir(), defaultNamespace, cniBinDir, cniConfDir)
})

var _ = Describe("CRI CNI Integration", func() {
	When("running a pod", func() {
		It("should trigger mock CNI plugin", func() {
			GinkgoT().DeferCleanup(func() {
				err := os.Remove(logFile)
				Expect(err).NotTo(HaveOccurred())
			})

			conn, ok := cli.Conn().(*grpc.ClientConn)
			Expect(ok).To(BeTrue())

			runtimeCli, err := remote.NewRuntimeService(conn.Target(), time.Second)
			Expect(err).NotTo(HaveOccurred())

			_, err = runtimeCli.RunPodSandbox(&v1.PodSandboxConfig{
				Metadata: &v1.PodSandboxMetadata{},
			}, "")
			Expect(err).To(MatchError(ContainSubstring("failed to find network info for sandbox")))

			Eventually(func() bool {
				logContent, err := os.ReadFile(logFile)
				if err != nil {
					return false
				}

				return len(logContent) > 0
			}, 5*time.Second, 500*time.Millisecond).Should(BeTrue())
		})
	})
})
