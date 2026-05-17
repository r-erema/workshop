package test

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/containerd/containerd/v2/client"
	_ "github.com/containerd/containerd/v2/cmd/containerd/builtins"
	"github.com/containerd/containerd/v2/cmd/containerd/command"
	"github.com/containerd/containerd/v2/cmd/containerd/server/config"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/pelletier/go-toml/v2"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	pollTimeout    = 5
	filePerm       = 0o600
	pollIntervalMs = 500
)

func RunContainerdEnv(testDir, defaultNamespace, cniBinDir, cniConfDir string) *client.Client {
	ginkgo.GinkgoT().Helper()

	ctx, cancel := context.WithCancel(ginkgo.GinkgoT().Context())
	defer cancel()

	sock := filepath.Join(testDir, "test_continerd.sock")
	configPath := filepath.Join(testDir, "test_continerd_conf.toml")

	cfg := config.Config{
		Root:    filepath.Join(testDir, "test_root"),
		State:   filepath.Join(testDir, "test_state"),
		TempDir: filepath.Join(testDir, "test_temp"),
		Plugins: map[string]any{
			"io.containerd.grpc.v1.cri": map[string]any{
				"cni": map[string]any{
					"bin_dir":  cniBinDir,
					"conf_dir": cniConfDir,
				},
			},
		},
	}

	cfgBytes, err := toml.Marshal(cfg)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	err = os.WriteFile(configPath, cfgBytes, filePerm)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	go func() {
		err = command.App().RunContext(ctx, []string{
			"",
			"--address", sock,
			"--config", configPath,
		})
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		ginkgo.DeferCleanup(cancel)
	}()

	cli, err := client.New(sock, client.WithDefaultNamespace(defaultNamespace))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	gomega.Eventually(func() grpc_health_v1.HealthCheckResponse_ServingStatus {
		resp, err := cli.HealthService().Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		return resp.GetStatus()
	},
		time.Second*pollTimeout,
		time.Millisecond*pollIntervalMs,
	).Should(gomega.Equal(grpc_health_v1.HealthCheckResponse_SERVING))

	return cli
}
