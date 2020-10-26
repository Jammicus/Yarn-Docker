package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerfiles(t *testing.T) {
	// Configure the tag to use on the Docker image.

	type image struct {
		path        string
		tag         string
		nodeVersion string
		yarnVersion string
	}

	images := []image{
		{
			path:        "Centos/",
			tag:         "local/centos-latest",
			nodeVersion: "12.18.4",
			yarnVersion: "1.22.5",
		},
		{
			path:        "Debian/",
			tag:         "local/debian-latest",
			nodeVersion: "12.18.4",
			yarnVersion: "1.22.5",
		},
		{
			path:        "Ubuntu/",
			tag:         "local/ubuntu-latest",
			nodeVersion: "12.18.4",
			yarnVersion: "1.22.5",
		},
	}

	for _, image := range images {

		buildOptions := &docker.BuildOptions{
			Tags: []string{image.tag},
		}

		// Build the Docker image.
		docker.Build(t, image.path, buildOptions)

		opts := &docker.RunOptions{Command: []string{"node", "--version"}}
		output := docker.Run(t, image.tag, opts)
		assert.Equal(t, "v"+image.nodeVersion, output)

		opts = &docker.RunOptions{Command: []string{"yarn", "--version"}}
		output = docker.Run(t, image.tag, opts)
		assert.Equal(t, image.yarnVersion, output)
	}
}
