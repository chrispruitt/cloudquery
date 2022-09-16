// Code generated by codegen; DO NOT EDIT.

package cloudfront

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Distributions() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudfront_distributions",
		Resolver:  fetchCloudfrontDistributions,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudfrontDistributionTags,
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
				Name:     "distribution_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DistributionConfig"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "in_progress_invalidation_batches",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InProgressInvalidationBatches"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "active_trusted_key_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedKeyGroups"),
			},
			{
				Name:     "active_trusted_signers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedSigners"),
			},
			{
				Name:     "alias_icp_recordals",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AliasICPRecordals"),
			},
		},
	}
}
