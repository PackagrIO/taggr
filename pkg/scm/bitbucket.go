package scm

import (
	"github.com/packagrio/go-common/config"
	"github.com/packagrio/go-common/pipeline"
	"net/http"
)

type scmBitbucket struct{}

// configure method will generate an authenticated client that can be used to comunicate with Github
func (b *scmBitbucket) Init(pipelineData *pipeline.Data, myConfig config.BaseInterface, httpClient *http.Client) error {
	return nil
}

func (b *scmBitbucket) GetClient() interface{} {
	return nil
}

func (b *scmBitbucket) CreateTagAtReference() error {
	return nil
}
