package scm

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"github.com/packagrio/go-common/config"
	"github.com/packagrio/go-common/pipeline"
	taggrConfig "github.com/packagrio/taggr/pkg/config"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"strings"
)

type scmGithub struct {
	Client *github.Client
	Config config.BaseInterface

	isGithubActionEnv bool
}

func (g *scmGithub) Init(pipelineData *pipeline.Data, myConfig config.BaseInterface, httpClient *http.Client) error {
	g.Config = myConfig
	g.Config.SetDefault(config.PACKAGR_SCM_GITHUB_ACCESS_TOKEN_TYPE, "user")
	ctx := context.Background()

	//TODO: autopaginate turned on.
	if httpClient != nil {
		//primarily used for testing.
		g.Client = github.NewClient(httpClient)
	} else if githubToken, present := os.LookupEnv("GITHUB_TOKEN"); present && len(githubToken) > 0 {
		log.Printf("found GITHUB_TOKEN")
		g.Config.Set(config.PACKAGR_SCM_GITHUB_ACCESS_TOKEN, githubToken)
		if action, isAction := os.LookupEnv("GITHUB_ACTION"); isAction && len(action) > 0 {
			log.Printf("Running in a Github Action")
			//running as a github action.
			g.Config.Set(config.PACKAGR_SCM_GITHUB_ACCESS_TOKEN_TYPE, "app")
			g.isGithubActionEnv = true
		}
	} else if g.Config.IsSet(config.PACKAGR_SCM_GITHUB_ACCESS_TOKEN) {
		log.Printf("found PACKAGR_SCM_GITHUB_ACCESS_TOKEN")

		//already set, do nothing.
	} else {
		//no access token present
		return fmt.Errorf("github SCM requires an access token")
	}

	//create an authenticated client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Config.GetString(config.PACKAGR_SCM_GITHUB_ACCESS_TOKEN)},
	)
	tc := oauth2.NewClient(ctx, ts)

	if g.Config.IsSet(config.PACKAGR_SCM_GITHUB_API_ENDPOINT) {
		gheClient, err := github.NewEnterpriseClient(
			g.Config.GetString(config.PACKAGR_SCM_GITHUB_API_ENDPOINT),
			g.Config.GetString(config.PACKAGR_SCM_GITHUB_API_ENDPOINT),
			tc,
		)
		if err != nil {
			return err
		}
		g.Client = gheClient
	} else {
		g.Client = github.NewClient(tc)
	}

	return nil
}

func (g *scmGithub) GetClient() interface{} {
	return g.Client
}

func (g *scmGithub) CreateTagAtReference() error {
	parts := strings.Split(g.Config.GetString(config.PACKAGR_SCM_REPO_FULL_NAME), "/")

	ctx := context.Background()
	_, _, err := g.Client.Git.CreateRef(ctx, parts[0], parts[1], &github.Reference{
		Ref: github.String(fmt.Sprintf("refs/tags/%s", g.Config.GetString(taggrConfig.PACKAGR_SCM_REPO_TAG))),
		Object: &github.GitObject{
			SHA: github.String(g.Config.GetString(taggrConfig.PACKAGR_SCM_REPO_SHA)),
		},
	})
	return err
}
