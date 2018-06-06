package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/djsd123/delete-ec2-ebs-snapshots/snapshot"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	clientSession       = session.Must(session.NewSession())
	connection          = ec2.New(clientSession)
	snapShotKeyFilter   = os.Getenv("SNAPSHOT_TAG_KEY")
	snapShotValueFilter = os.Getenv("SNAPSHOT_TAG_VALUE")
	getDays             = os.Getenv("DAYS_OLD")
)

func main() {

	// Convert `DAYS_OLD` value to an integer
	daysOldValue, err := strconv.Atoi(getDays)
	if err != nil {
		fmt.Println(err)
		log.Panicf("Failed to parse DAYS_OLD, must be an integer: %s", err)
	}

	snapShots, err := snapshot.GetSnapShots(connection, snapShotKeyFilter, snapShotValueFilter)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, snapShot := range snapShots.Snapshots {

		snapShotID := *snapShot.SnapshotId

		// Cast integer to type Duration
		numberOfDays := time.Duration(daysOldValue)
		days := time.Hour * 24 * numberOfDays

		isTwoWeeksOrOlder, err := snapshot.CheckSnapShotAge(snapShot.StartTime, days)
		if err != nil {
			log.Printf("Error while evaluating the snapshot's age %s: %s", snapShotID, err.Error())
		}

		if isTwoWeeksOrOlder != false {
			_, err := snapshot.PruneSnapShots(connection, snapShotID)
			if err != nil {
				log.Printf("Error while pruning snapshot %s: %s", snapShotID, err.Error())
			}
		}
	}
}
