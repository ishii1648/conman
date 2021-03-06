package config

import (
	"flag"
)

const (
	DefaultListen           = "/var/run/conmand.sock"
	DefaultLibRoot          = "/var/lib/conman"
	DefaultRunRoot          = "/var/run/conman"
	DefaultContainerLogRoot = "/var/log/conman/containers"
	DefaultStreamingAddr    = "127.0.0.1:8881"
	DefaultShimmyPath       = "/usr/local/bin/shimmy"
	DefaultRuntimePath      = "/usr/bin/runc"
	DefaultRuntimeRoot      = "/var/run/conman-runc"
)

type Config struct {
	Listen string

	// Root directory to store long-lived data (images, containers, etc).
	LibRoot string

	// Root directory to store state of the conman daemon.
	RunRoot string

	// Root directory to store container logs.
	ContainerLogRoot string

	// Streaming server host:port (for attach, exec, and port-forwarding).
	StreamingAddr string

	// Path to OCI runtime shim executable, aka shimmy.
	ShimmyPath string

	RuntimePath string

	RuntimeRoot string
}

func TestConfigFromFlags() *Config {
	cfg := &Config{
		LibRoot:     DefaultLibRoot,
		Listen:      DefaultListen,
		RunRoot:     DefaultRunRoot,
		RuntimeRoot: DefaultRuntimeRoot,
	}

	flag.StringVar(
		&cfg.ContainerLogRoot,
		"container-logs",
		DefaultContainerLogRoot,
		"Path to container logs root folder",
	)
	flag.StringVar(
		&cfg.RuntimePath,
		"runtime",
		DefaultRuntimePath,
		"Path to runc executable file",
	)
	flag.StringVar(
		&cfg.ShimmyPath,
		"shimmy",
		DefaultShimmyPath,
		"Path to shimmy executable file",
	)
	flag.Parse()

	return cfg
}
