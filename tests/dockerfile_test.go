package test

import (
	"sort"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerFile(t *testing.T) {
	// Given
	tag := "gruntwork/nodejs-example:latest"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}
	files := []string{"app.js", "package-lock.json", "package.json", "node_modules"}
	sort.Strings(files)
	expectedOutput := strings.Join(files, "\n")

	// When
	docker.Build(t, "../", buildOptions)
	opts := &docker.RunOptions{
		Command: []string{"ls", "/usr/src/app"},
		Remove:  true,
	}
	output := docker.Run(t, tag, opts)

	// Then
	assert.Equal(t, expectedOutput, output)
}
