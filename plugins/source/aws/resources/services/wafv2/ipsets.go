// Code generated by codegen; DO NOT EDIT.

package wafv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ipsets() *schema.Table {
	return &schema.Table{
		Name:      "aws_wafv2_ipsets",
		Resolver:  fetchWafv2Ipsets,
		Multiplex: client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
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
				Name:     "addresses",
				Type:     schema.TypeInetArray,
				Resolver: resolveIpsetAddresses,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIpsetTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "ip_address_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPAddressVersion"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
		},
	}
}
