// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func ListEfsFileSystem(client *Client) ([]Resource, error) {
	req := client.efsconn.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{})

	var result []Resource

	p := efs.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FileSystems {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_efs_file_system",
				ID:        *r.FileSystemId,
				Region:    client.efsconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
