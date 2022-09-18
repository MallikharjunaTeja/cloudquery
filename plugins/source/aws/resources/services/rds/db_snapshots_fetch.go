package rds

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsDbSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().RDS
	var input rds.DescribeDBSnapshotsInput
	for {
		output, err := svc.DescribeDBSnapshots(ctx, &input)
		if err != nil {
			return nil
		}
		res <- output.DBSnapshots
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRDSDBSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	tags := map[string]*string{}
	for _, t := range s.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}

func resolveRDSDBSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().RDS
	out, err := svc.DescribeDBSnapshotAttributes(
		ctx,
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = c.Region
		},
	)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if out.DBSnapshotAttributesResult == nil {
		return nil
	}

	b, err := json.Marshal(out.DBSnapshotAttributesResult.DBSnapshotAttributes)
	if err != nil {
		return err
	}
	return resource.Set(column.Name, b)
}