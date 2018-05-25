package main

import (
	"./snapshot"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
		log.Panicf("The type of the value you provided is %T. %v needs to be an integer", daysOldValue, daysOldValue)
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

		isSnapShotTwoWeeksOld, err := snapshot.CheckSnapShotAge(snapShot.StartTime, days)
		if err != nil {
			log.Printf("Error while evaluating the snapshot's age %s: %s", snapShotID, err.Error())
		}

		if isSnapShotTwoWeeksOld != false {
			_, err = snapshot.PruneSnapShots(connection, snapShotID)
			if err != nil {
				log.Printf("Error while pruning snapshot %s: %s", snapShotID, err.Error())
			}
		}

	}
}
