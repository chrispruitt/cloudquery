// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_rds_subnet_groups",
		Resolver:  fetchRdsSubnetGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("rds"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "db_subnet_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupDescription"),
			},
			{
				Name:     "db_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupName"),
			},
			{
				Name:     "subnet_group_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetGroupStatus"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "supported_network_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedNetworkTypes"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
