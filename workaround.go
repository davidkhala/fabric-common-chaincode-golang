package golang

import docker "github.com/fsouza/go-dockerclient"

func touchToForceVersion130() {
	var MemorySwappiness int64
	var OOMKillDisable bool
	var config = docker.HostConfig{}
	config.MemorySwappiness = MemorySwappiness
	config.OOMKillDisable = OOMKillDisable
}
