package rtinstances


import (
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/missioncontrol/utils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/io/httputils"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-cli/utils/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jfrogdev/jfrog-cli-go/jfrog-client/utils/errorutils"
)

func DetachLic(instanceName string, flags *DetachLicFlags) error {
	bucketId := flags.BucketId
	postContent := utils.LicenseRequestContent{
		Name: 	  	 instanceName,
		NodeID:	     flags.NodeId}
	requestContent, err := json.Marshal(postContent)
	if err != nil {
		return errorutils.CheckError(errors.New("Failed to marshal json. " + cliutils.GetDocumentationMessage()))
	}
	missionControlUrl := flags.MissionControlDetails.Url + "api/v1/buckets/" + bucketId + "/licenses";
	httpClientDetails := utils.GetMissionControlHttpClientDetails(flags.MissionControlDetails)
	resp, body, err := httputils.SendDelete(missionControlUrl, requestContent, httpClientDetails)
	if err != nil {
	    return err
	}
	if resp.StatusCode != 200 {
		return errorutils.CheckError(errors.New(resp.Status + ". " + utils.ReadMissionControlHttpMessage(body)))
	}
	fmt.Println("Mission Control response: " + resp.Status)
	return nil
}

type DetachLicFlags struct {
	MissionControlDetails *config.MissionControlDetails
	Interactive 	      bool
	NodeId 			      string
	BucketId 		      string
}
