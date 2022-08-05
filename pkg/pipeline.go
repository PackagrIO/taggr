package pkg

import (
	"fmt"
	"github.com/packagrio/go-common/pipeline"
	"github.com/packagrio/taggr/pkg/config"
	"github.com/packagrio/taggr/pkg/scm"
	"os"
)

type Pipeline struct {
	Data   *pipeline.Data
	Config config.Interface
	Scm    scm.Interface
}

func (p *Pipeline) Start(configData config.Interface) error {
	// Initialize Pipeline.
	p.Config = configData
	p.Data = new(pipeline.Data)

	sourceScm, err := scm.Create(p.Config.GetString(config.PACKAGR_SCM), p.Data, p.Config, nil)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}
	p.Scm = sourceScm

	//p.Scm.Client

	return p.Scm.CreateTagAtReference()
}
