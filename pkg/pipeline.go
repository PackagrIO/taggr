package pkg

import (
	"fmt"
	"github.com/packagrio/go-common/pipeline"
	"github.com/packagrio/go-common/scm"
	"github.com/packagrio/taggr/pkg/config"
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

	p.Data.GitHeadInfo = &pipeline.ScmCommitInfo{
		Sha: p.Config.GetString(config.PACKAGR_SCM_REPO_SHA),
		Repo: &pipeline.ScmRepoInfo{
			FullName: p.Config.GetString(config.PACKAGR_SCM_REPO_FULL_NAME),
		},
	}

	return p.Scm.CreateTagAtReference(p.Config.GetString(config.PACKAGR_SCM_REPO_TAG_NAME))
}
