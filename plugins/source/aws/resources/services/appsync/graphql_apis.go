// Code generated by codegen; DO NOT EDIT.

package appsync

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GraphqlApis() *schema.Table {
	return &schema.Table{
		Name:      "aws_appsync_graphql_apis",
		Resolver:  fetchAppsyncGraphqlApis,
		Multiplex: client.ServiceAccountRegionMultiplexer("appsync"),
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "additional_authentication_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalAuthenticationProviders"),
			},
			{
				Name:     "api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
			{
				Name:     "authentication_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthenticationType"),
			},
			{
				Name:     "lambda_authorizer_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LambdaAuthorizerConfig"),
			},
			{
				Name:     "log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogConfig"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "open_id_connect_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OpenIDConnectConfig"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "uris",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Uris"),
			},
			{
				Name:     "user_pool_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolConfig"),
			},
			{
				Name:     "waf_web_acl_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WafWebAclArn"),
			},
			{
				Name:     "xray_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("XrayEnabled"),
			},
		},
	}
}
