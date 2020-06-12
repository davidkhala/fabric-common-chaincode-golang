package golang

import docker "github.com/fsouza/go-dockerclient"

func touchToForceVersion121() docker.HostConfig {
	return docker.HostConfig{}
}
