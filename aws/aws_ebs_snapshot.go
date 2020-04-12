// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEbsSnapshot(client *Client) error {
	req := client.ec2conn.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})

	p := ec2.NewDescribeSnapshotsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Snapshots {
			fmt.Println(*r.SnapshotId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.StartTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
