// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayVpcAttachments() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_transit_gateway_vpc_attachments",
		Resolver:  fetchEc2TransitGatewayVpcAttachments,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "transit_gateway_attachment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayAttachmentId"),
			},
			{
				Name:     "transit_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayId"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
			{
				Name:     "vpc_owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcOwnerId"),
			},
		},
	}
}
