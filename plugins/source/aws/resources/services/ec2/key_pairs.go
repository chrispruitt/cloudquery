// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func KeyPairs() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_key_pairs",
		Resolver:  fetchEc2KeyPairs,
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
				Resolver: resolveKeyPairArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "key_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyFingerprint"),
			},
			{
				Name:     "key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyName"),
			},
			{
				Name:     "key_pair_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyPairId"),
			},
			{
				Name:     "key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyType"),
			},
			{
				Name:     "public_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicKey"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
