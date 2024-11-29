// Copyright (C) 2024 The sql-gemini Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gemini

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// Docker is a struct that contains the docker client and the container id
type Docker struct {
	client *client.Client
}

// NewDocker creates a new Docker struct
func NewDocker() (*Docker, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Docker{client: client}, nil
}

// Run runs a Docker container with the specified image name
func (d *Docker) Run(imageName string) (string, error) {
	ctx := context.Background()

	// Pull the specified Docker image from DockerHub
	_, err := d.client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return "", err
	}

	// List all running containers
	containers, err := d.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", err
	}

	// Stop and remove all running containers
	for _, c := range containers {
		if err := d.client.ContainerStop(ctx, c.ID, container.StopOptions{}); err != nil {
			return "", err
		}
		if err := d.client.ContainerRemove(ctx, c.ID, container.RemoveOptions{}); err != nil {
			return "", err
		}
	}

	// Create a container from the pulled image
	resp, err := d.client.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, "")
	if err != nil {
		return "", err
	}

	// Start the created container
	if err := d.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}
