package services

import (
	"testing"
	"flag"
	"os"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/httpclient"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/services/utils/auth"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/artifactory/services/utils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/io/httputils"
	"errors"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/log"
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/utils/config"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/tests"
)

var RtUrl *string
var RtUser *string
var RtPassword *string
var RtApiKey *string
var RtSshKeyPath *string
var RtSshPassphrase *string
var LogLevel *string
var testsUploadService *UploadService
var testsSearchService *SearchService
var testsDeleteService *DeleteService
var testsDownloadService *DownloadService

const (
	RtTargetRepo              = "jfrog-cli-tests-repo1/"
	SpecsTestRepositoryConfig = "specs_test_repository_config.json"
	RepoDetailsUrl            = "api/repositories/"
	ClientIntegrationTests    = "jfrog-cli-go/jfrog-client/artifactory/services"
)

func init() {
	RtUrl = flag.String("rt.url", "http://localhost:8081/artifactory/", "Artifactory url")
	RtUser = flag.String("rt.user", "admin", "Artifactory username")
	RtPassword = flag.String("rt.password", "password", "Artifactory password")
	RtApiKey = flag.String("rt.apikey", "", "Artifactory user API key")
	RtSshKeyPath = flag.String("rt.sshKeyPath", "", "Ssh key file path")
	RtSshPassphrase = flag.String("rt.sshPassphrase", "", "Ssh key passphrase")
	LogLevel = flag.String("log-level", "INFO", "Sets the log level")
}

func TestMain(m *testing.M) {
	packages := tests.ExcludeTestsPackage(tests.GetPackages(), ClientIntegrationTests)
	tests.RunTests(packages)
	flag.Parse()
	log.Logger.SetLogLevel(log.GetCliLogLevel(*LogLevel))
	*RtUrl = cliutils.AddTrailingSlashIfNeeded(*RtUrl)
	InitArtifactoryServiceManager()
	createReposIfNeeded()
	result := m.Run()
	os.Exit(result)
}

func InitArtifactoryServiceManager() {
	createArtifactoryUploadManager()
	createArtifactorySearchManager()
	createArtifactoryDeleteManager()
	createArtifactoryDownloadManager()
}

func createArtifactorySearchManager() {
	testsSearchService = NewSearchService(httpclient.NewDefaultHttpClient())
	testsSearchService.ArtDetails = getArtDetails()
}

func createArtifactoryDeleteManager() {
	testsDeleteService = NewDeleteService(httpclient.NewDefaultHttpClient())
	testsDeleteService.ArtDetails = getArtDetails()
}

func createArtifactoryUploadManager() {
	testsUploadService = NewUploadService(httpclient.NewDefaultHttpClient())
	testsUploadService.ArtDetails = getArtDetails()
	testsUploadService.Threads = 3
}

func createArtifactoryDownloadManager() {
	testsDownloadService = NewDownloadService(httpclient.NewDefaultHttpClient())
	testsDownloadService.ArtDetails = getArtDetails()
	testsDownloadService.SetThreads(3)
}

func getArtDetails() *auth.ArtifactoryDetails {
	rtDetails := &config.ArtifactoryDetails{Url: *RtUrl, SshKeyPath: *RtSshKeyPath, SshPassphrase: *RtSshPassphrase}
	if !rtDetails.IsSsh() {
		if *RtApiKey != "" {
			rtDetails.ApiKey = *RtApiKey
		} else {
			rtDetails.User = *RtUser
			rtDetails.Password = *RtPassword
		}
	}
	artAuth, err := rtDetails.CreateArtAuthConfig()
	if err != nil {
		log.Error("Failed while attempting to authenticate with Artifactory: " + err.Error())
		os.Exit(1)
	}
	return artAuth
}

func artifactoryCleanUp(t *testing.T) {
	params := &utils.ArtifactoryCommonParams{Pattern: RtTargetRepo}
	toDelete, err := testsDeleteService.GetPathsToDelete(&DeleteParamsImpl{ArtifactoryCommonParams: params})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	var deleteItems []DeleteItem = make([]DeleteItem, len(toDelete))
	for i, item := range toDelete {
		deleteItems[i] = item
	}
	err = testsDeleteService.DeleteFiles(deleteItems, testsDeleteService)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func createReposIfNeeded() error {
	var err error
	var repoConfig string
	repo := RtTargetRepo
	if strings.HasSuffix(repo, "/") {
		repo = repo[0:strings.LastIndex(repo, "/")]
	}
	if !isRepoExist(repo) {
		repoConfig = filepath.Join(getTestDataPath(), "reposconfig", SpecsTestRepositoryConfig)
		err = execCreateRepoRest(repoConfig, repo)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func isRepoExist(repoName string) bool {
	artDetails := getArtDetails()
	artHttpDetails := artDetails.CreateArtifactoryHttpClientDetails()
	resp, _, _, err := httputils.SendGet(artDetails.Url+RepoDetailsUrl+repoName, true, artHttpDetails)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if resp.StatusCode != 400 {
		return true
	}
	return false
}

func execCreateRepoRest(repoConfig, repoName string) error {
	content, err := ioutil.ReadFile(repoConfig)
	if err != nil {
		return err
	}
	artHttpDetails := getArtDetails().CreateArtifactoryHttpClientDetails()

	artHttpDetails.Headers = map[string]string{"Content-Type": "application/json"}
	resp, _, err := httputils.SendPut(*RtUrl+"api/repositories/"+repoName, content, artHttpDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return errors.New("Fail to create repository. Reason local repository with key: " + repoName + " already exist\n")
	}
	log.Info("Repository", repoName, "was created.")
	return nil
}

func getTestDataPath() string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "testsdata")
}
