package test

import (
	"path/filepath"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func RunCluster() *kubernetes.Clientset {
	env := &envtest.Environment{
		BinaryAssetsDirectory: filepath.Join("..", "..", ".bin", "k8s", "envtest-v1.36.0-linux-amd64"),
	}

	k8sConfig, err := env.Start()
	Expect(err).NotTo(HaveOccurred())

	ginkgo.DeferCleanup(func() {
		err = env.Stop()
		Expect(err).NotTo(HaveOccurred())
	})

	k8sClientset, err := kubernetes.NewForConfig(k8sConfig)
	Expect(err).NotTo(HaveOccurred())

	return k8sClientset
}
