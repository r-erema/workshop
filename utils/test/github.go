package test

import (
	"os"

	"github.com/onsi/ginkgo/v2"
)

func SkipInGitHubActions() {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		ginkgo.Skip("Skip in Github Actions")
	}
}
