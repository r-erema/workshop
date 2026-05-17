package test

import (
	"context"

	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go/modules/ollama"
)

const (
	ollamaImage       = "ollama/ollama:latest"
	cachedOllamaImage = "ollama-with-model:latest"
	ollamaModel       = "tinyllama:latest" // Smallest available model (~637MB)
)

func StartOllamaContainer(ctx context.Context) *ollama.OllamaContainer {
	if cachedImageExists(ctx, cachedOllamaImage) {
		container, err := ollama.Run(ctx, cachedOllamaImage)
		Expect(err).NotTo(HaveOccurred())

		return container
	}

	container, err := ollama.Run(ctx, ollamaImage)
	Expect(err).NotTo(HaveOccurred())

	_, _, err = container.Exec(ctx, []string{"ollama", "pull", ollamaModel})
	Expect(err).NotTo(HaveOccurred())

	err = container.Commit(ctx, cachedOllamaImage)
	Expect(err).NotTo(HaveOccurred())

	return container
}
