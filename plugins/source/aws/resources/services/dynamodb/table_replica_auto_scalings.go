// Code generated by codegen; DO NOT EDIT.

package dynamodb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TableReplicaAutoScalings() *schema.Table {
	return &schema.Table{
		Name:      "aws_dynamodb_table_replica_auto_scalings",
		Resolver:  fetchDynamodbTableReplicaAutoScalings,
		Multiplex: client.ServiceAccountRegionMultiplexer("dynamodb"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "table_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "global_secondary_indexes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalSecondaryIndexes"),
			},
			{
				Name:     "region_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionName"),
			},
			{
				Name:     "replica_provisioned_read_capacity_auto_scaling_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReplicaProvisionedReadCapacityAutoScalingSettings"),
			},
			{
				Name:     "replica_provisioned_write_capacity_auto_scaling_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReplicaProvisionedWriteCapacityAutoScalingSettings"),
			},
			{
				Name:     "replica_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaStatus"),
			},
		},
	}
}