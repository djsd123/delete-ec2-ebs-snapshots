package snapshot

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

func GetSnapShots(connection *ec2.EC2, key string, val string) (output *ec2.DescribeSnapshotsOutput, err error) {

	snapShotInput := &ec2.DescribeSnapshotsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(fmt.Sprintf("tag:%s", key)),
				Values: []*string{
					aws.String(val),
				},
			},
		},
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
