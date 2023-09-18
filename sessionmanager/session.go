package sessionmanager

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmv1 "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/chrisdd3/session-manager-plugin/internal/session"
)

func StartSession(startSessionResp *ssm.StartSessionOutput, startSessionRequest *ssm.StartSessionInput, regionName string, profileName string, endpointUrl string, output io.Writer) error {
	// v1 parameters, the sessionmanager utility is using sdk v1
	parametersv1 := map[string][]*string{}
	for k, v := range startSessionRequest.Parameters {
		val := []*string{}
		for _, v := range v {
			val = append(val, &v)
		}
		parametersv1[k] = val
	}

	return session.StartSession(
		&ssmv1.StartSessionOutput{
			SessionId:  startSessionResp.SessionId,
			StreamUrl:  startSessionResp.StreamUrl,
			TokenValue: startSessionResp.TokenValue,
		},
		&ssmv1.StartSessionInput{
			DocumentName: startSessionRequest.DocumentName,
			Parameters:   parametersv1,
			Reason:       startSessionRequest.Reason,
			Target:       startSessionRequest.Target,
		},
		regionName,
		profileName,
		"", // meh
		os.Stdout)
}
