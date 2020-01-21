package version

import (
	"flag"
	"fmt"

	"github.com/micro/cli"
	"time"
	"msp-git.connext.com.cn/module/pub/config"
	"github.com/hashicorp/consul/api"
	"path/filepath"
	"encoding/json"
)

var (
	// Version should be updated by hand at each release
	Version = "build_version"
	//will be overwritten automatically by the build system
	GitBranch string
	GitTag    string
	GitCommit string
	GitURL    string
	GoVersion string
	BuildTime string
)

const BuildInfoKeyPrefix = "services"

type BuildInfo struct {
	//// Version should be updated by hand at each release
	//Version string
	//will be overwritten automatically by the build system
	GitBranch string `json:"git_branch"`
	GitTag    string `json:"git_tag"`
	GitCommit string `json:"git_commit"`
	GitURL    string `json:"git_url"`
	GoVersion string `json:"go_version"`
	BuildTime string `json:"build_time"`
	WriteTime string `json:"write_time"`
}

type BuildItem struct {
	ServiceName string    `json:"service_name"`
	BuildInfo   BuildInfo `json:"build_info"`
}

var ShowVersion = flag.Bool("version", false, "current version")

// FullVersion formats the version to be printed
func FullVersion() string {
	return fmt.Sprintf("Version: %6s \nGit URL: %6s \nGit branch: %6s \nGit Tag: %6s \nGit commit: %6s \nGo version: %6s \nBuild time: %6s \n",
		Version, GitURL, GitBranch, GitTag, GitCommit, GoVersion, BuildTime)
}

var showBuildVersion bool

var buildVersionFlag = cli.BoolFlag{
	Name:        "build_version",
	Usage:       "show go build version info(git branch, git log...)",
	Destination: &showBuildVersion,
}

func ShowBuildVersion() bool {
	return showBuildVersion
}

func BuildVersionFlag() *cli.BoolFlag {
	return &buildVersionFlag
}

func GetBuildInfo() BuildInfo {
	return BuildInfo{
		GitURL:    GitURL,
		GitBranch: GitBranch,
		GitTag:    GitTag,
		GitCommit: GitCommit,
		GoVersion: GoVersion,
		BuildTime: BuildTime,
		WriteTime: time.Now().String()}
}

func WriteBuildInfo(serviceName string) error {
	value, err := json.Marshal(GetBuildInfo())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return writeBuildInfoToConsul(filepath.Join(BuildInfoKeyPrefix, serviceName), value)
}

func writeBuildInfoToConsul(key string, value []byte) error {
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: config.CConsulAddr()})
	if err != nil {
		return err
	}
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, nil)
	if err != nil {
		return err
	}
	return nil
}