package server

import "gocarch/config"

type (
	container struct {
		Config config.MainConfig
	}

	opts struct {
		Config *config.MainConfig
	}
)

func newContainer(o *opts) *container {
	return &container{
		Config: *o.Config,
	}
}
