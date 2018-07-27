package main

import (
	toml "github.com/sioncojp/tomlssm"
)

// Config ...config.toml
type Config struct {
	Port                 string `toml:"PORT"`
	BotID                string `toml:"BOT_ID"`
	BotToken             string `toml:"BOT_TOKEN"`
	VerificationToken    string `toml:"VERIFICATION_TOKEN"`
	ChannelID            string `toml:"CHANNEL_ID"`
	GitHubUsername       string `toml:"GITHUB_USERNAME"`
	GitHubToken          string `toml:"GITHUB_TOKEN"`
	GitHubRepositorySlug string `toml:"GITHUB_REPOSITORY_SLUG"`
	GitCloneLocalPath    string `toml:"GIT_CLONE_LOCAL_PATH"`
	GitCommitAuthorName  string `toml:"GIT_COMMIT_AUTHOR_NAME"`
	GitCommitAuthorEmail string `toml:"GIT_COMMIT_AUTHOR_EMAIL"`
	InfoPlistPath        string `toml:"INFOPLIST_PATH"`
	GitBranches          string `toml:"GIT_BRANCHES"`
}

// LoadToml ...tomlファイルを読み込む
func LoadToml(c, region string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(c, &config, region); err != nil {
		return nil, err
	}
	return &config, nil
}
