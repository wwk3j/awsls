// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamInstanceProfile(client *Client) error {
	req := client.iamconn.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{})

	p := iam.NewListInstanceProfilesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.InstanceProfiles {
			fmt.Println(*r.InstanceProfileName)

			fmt.Printf("CreatedAt: %s\n", *r.CreateDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
