package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
	"fmt"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	app := FlagSet()
	app.Action = func(c *cli.Context) error {
		if c.String("config") == "" {
			return errors.New("required -c option")
		}
		// load config
		config, err := LoadToml(c.String("config"), c.String("region"))
		if err != nil {
			return fmt.Errorf("failed to load toml file: %s", err)
		}

		repositoryOptions := RepositoryOptions{
			path: config.GitCloneLocalPath,
			credential: Credential{
				Username: config.GitHubUsername,
				Token:    config.GitHubToken,
			},
			slug: config.GitHubRepositorySlug,
			author: Author{
				Name:  config.GitCommitAuthorName,
				Email: config.GitCommitAuthorEmail,
			},
			infoPlistPath: config.InfoPlistPath,
			branches:      config.GitBranches,
		}

		log.Printf("[INFO] Start slack event listening")
		client := slack.New(config.BotToken)
		slackListener := &SlackListener{
			client:            client,
			botID:             config.BotID,
			channelID:         config.ChannelID,
			repositoryOptions: repositoryOptions,
		}
		go slackListener.ListenAndResponse()

		http.Handle("/interaction", interactionHandler{
			verificationToken: config.VerificationToken,
			channelID:         config.ChannelID,
			repositoryOptions: repositoryOptions,
		})

		log.Printf("[INFO] Server listening on :%s", config.Port)
		if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}

	return 0
}
