package caller

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetCaller() (output *sts.GetCallerIdentityOutput, err error) {

	connection := sts.New(session.Must(session.NewSession()))
	callerInput := &sts.GetCallerIdentityInput{}

	return connection.GetCallerIdentity(callerInput)
}
