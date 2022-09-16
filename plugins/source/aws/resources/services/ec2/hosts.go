// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_hosts",
		Resolver:  fetchEc2Hosts,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveHostArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "allocation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AllocationTime"),
			},
			{
				Name:     "allows_multiple_instance_types",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowsMultipleInstanceTypes"),
			},
			{
				Name:     "auto_placement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoPlacement"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "availability_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneId"),
			},
			{
				Name:     "available_capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AvailableCapacity"),
			},
			{
				Name:     "client_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientToken"),
			},
			{
				Name:     "host_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostId"),
			},
			{
				Name:     "host_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostProperties"),
			},
			{
				Name:     "host_recovery",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostRecovery"),
			},
			{
				Name:     "host_reservation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostReservationId"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:     "member_of_service_linked_resource_group",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MemberOfServiceLinkedResourceGroup"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "release_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReleaseTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
