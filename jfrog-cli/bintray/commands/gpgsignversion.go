package commands

import (
	"errors"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/bintray/utils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/utils/config"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/io/httputils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/log"
	"fmt"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/errorutils"
	clientutils "github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils"
)

func GpgSignVersion(versionDetails *utils.VersionDetails, passphrase string, bintrayDetails *config.BintrayDetails) error {
	if bintrayDetails.User == "" {
		bintrayDetails.User = versionDetails.Subject
	}
	url := bintrayDetails.ApiUrl + "gpg/" + versionDetails.Subject + "/" +
			versionDetails.Repo + "/" + versionDetails.Package +
			"/versions/" + versionDetails.Version

	var data string
	if passphrase != "" {
		data = "{ \"passphrase\": \"" + passphrase + "\" }"
	}

	log.Info("GPG signing version...")
	httpClientsDetails := utils.GetBintrayHttpClientDetails(bintrayDetails)
	resp, body, err := httputils.SendPost(url, []byte(data), httpClientsDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errorutils.CheckError(errors.New("Bintray response: " + resp.Status + "\n" + clientutils.IndentJson(body)))
	}

	log.Debug("Bintray response:", resp.Status)
	log.Info("GPG signed version", versionDetails.Version, ", details:")
	fmt.Println(clientutils.IndentJson(body))
	return nil
}
