package scm

import (
	"github.com/packagrio/go-common/config"
	"github.com/packagrio/go-common/pipeline"
	"net/http"
)

// Create mock using:
// mockgen -source=scm/interface.go -destination=scm/mock/mock_scm.go
type Interface interface {

	// init method will generate an authenticated client that can be used to comunicate with Scm
	Init(pipelineData *pipeline.Data, config config.BaseInterface, client *http.Client) error

	GetClient() interface{}
	CreateTagAtReference() error
}
