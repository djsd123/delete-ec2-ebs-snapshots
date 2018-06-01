package snapshot

import (
	"../caller"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"time"
)

func GetSnapShots(connection *ec2.EC2, key string, val string) (output *ec2.DescribeSnapshotsOutput, err error) {

	var snapShotInput *ec2.DescribeSnapshotsInput

	// Get the owner ID a.k.a the aws account number
	callerResult, callerErr := caller.GetCaller()
	if callerErr != nil {
		log.Panic(callerErr)
	}
	ownerID := *callerResult.Account

	if len(key) != 0 && len(val) != 0 {
		snapShotInput = &ec2.DescribeSnapshotsInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String(fmt.Sprintf("tag:%s", key)),
					Values: []*string{
						aws.String(val),
					},
				},
				{
					Name: aws.String("owner-id"),
					Values: []*string{
						aws.String(ownerID),
					},
				},
			},
		}
	} else {
		snapShotInput = &ec2.DescribeSnapshotsInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("owner-id"),
					Values: []*string{
						aws.String(ownerID),
					},
				},
			},
		}
	}

	output, err = connection.DescribeSnapshots(snapShotInput)

	return output, err
}

func PruneSnapShots(connection *ec2.EC2, snapshotID string) (output *ec2.DeleteSnapshotOutput, err error) {

	deleteSnapShotShotInput := &ec2.DeleteSnapshotInput{
		SnapshotId: aws.String(snapshotID),
	}

	output, err = connection.DeleteSnapshot(deleteSnapShotShotInput)

	return output, err
}

func CheckSnapShotAge(SnapShotStartTimeStamp *time.Time, duration time.Duration) (isOlder bool, err error) {

	isOlder = SnapShotStartTimeStamp.Add(duration).Before(time.Now())

	return isOlder, err
}
