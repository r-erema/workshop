package test

import (
	"context"

	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
)

func cachedImageExists(ctx context.Context, imageName string) bool {
	dockerClient, err := testcontainers.NewDockerClientWithOpts(ctx)
	if err != nil {
		return false
	}
	defer func() {
		Expect(dockerClient.Close()).NotTo(HaveOccurred())
	}()

	_, err = dockerClient.Client.ImageInspect(ctx, imageName)

	return err == nil
}
