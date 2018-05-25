# delete-ec2-ebs-snapshots
AWS Lambda Package to delete ebs snapshots

Usage
-----

__Environment Variables__
- `DAYS_OLD` - (Required) The age of snapshots in __days__ that will be targeted (__Integer__)
- `SNAPSHOT_TAG_KEY`   - (Optional) The key of the tag associated with the ec2 snapshot
- `SNAPSHOT_TAG_VALUE` - (Optional) The value of the tag associated with the ec2 snapshot

Optional Filters on ec2 tags by setting the following environment variables

Example:
  
Only target snapshots that are two weeks or more old and have a tag with key:`etcd` and value:`1`

```
DAYS_OLD=14
SNAPSHOT_TAG_KEY=etcd
SNAPSHOT_TAG_VALUE=1
```


