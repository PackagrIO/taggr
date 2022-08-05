package main

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/packagrio/taggr/pkg"
	"github.com/packagrio/taggr/pkg/config"
	"github.com/packagrio/taggr/pkg/version"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var goos string
var goarch string

func main() {
	app := &cli.App{
		Name:     "taggr",
		Usage:    "Create tags in git repositories without cloning the repo",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []cli.Author{
			cli.Author{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},
		Before: func(c *cli.Context) error {

			packagrUrl := "github.com/packagrio/taggr"

			versionInfo := fmt.Sprintf("%s.%s-%s", goos, goarch, version.VERSION)

			subtitle := packagrUrl + utils.LeftPad2Len(versionInfo, " ", 53-len(packagrUrl))

			fmt.Fprintf(c.App.Writer, fmt.Sprintf(utils.StripIndent(
				`
			 ____   __    ___  __ _   __    ___  ____ 
			(  _ \ / _\  / __)(  / ) / _\  / __)(  _ \
			 ) __//    \( (__  )  ( /    \( (_ \ )   /
			(__)  \_/\_/ \___)(__\_)\_/\_/ \___/(__\_)
			%s

			`), subtitle))
			return nil
		},

		Commands: []cli.Command{
			{
				Name:  "create",
				Usage: "Create a tag",
				Action: func(c *cli.Context) error {

					configuration, _ := config.Create()
					if c.IsSet("scm") {
						configuration.Set(config.PACKAGR_SCM, c.String("scm"))
					}
					if c.IsSet("repo_full_name") {
						configuration.Set(config.PACKAGR_SCM_REPO_FULL_NAME, c.String("repo_full_name"))
					}
					if c.IsSet("repo_sha") {
						configuration.Set(config.PACKAGR_SCM_REPO_SHA, c.String("repo_sha"))
					}
					if c.IsSet("repo_tag") {
						configuration.Set(config.PACKAGR_SCM_REPO_TAG_NAME, c.String("repo_tag"))
					}

					fmt.Println("scm:", configuration.GetString(config.PACKAGR_SCM))
					fmt.Println("repo name:", configuration.GetString(config.PACKAGR_SCM_REPO_FULL_NAME))
					fmt.Println("sha:", configuration.GetString(config.PACKAGR_SCM_REPO_SHA))
					fmt.Println("tag:", configuration.GetString(config.PACKAGR_SCM_REPO_TAG_NAME))

					pipeline := pkg.Pipeline{}
					err := pipeline.Start(configuration)
					if err != nil {
						fmt.Printf("FATAL: %+v\n", err)
						os.Exit(1)
					}

					return nil
				},

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "scm",
						Value: "default",
						Usage: "The scm for the code, for setting additional SCM specific metadata",
					},

					&cli.StringFlag{
						Name:     "repo_full_name, name",
						Usage:    "The repository to create the tag within (eg. analogj/test)",
						Required: true,
					},

					&cli.StringFlag{
						Name:     "repo_sha, sha",
						Usage:    "The git commit to reference with this new tag",
						Required: true,
					},

					&cli.StringFlag{
						Name:     "repo_tag, tag",
						Usage:    "The tag to create in this repo (eg. v1.0.1)",
						Required: true,
					},

					&cli.BoolFlag{
						Name:  "dry_run",
						Usage: "When dry run is enabled, no data is written to file system",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
