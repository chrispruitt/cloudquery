// Code generated by codegen; DO NOT EDIT.

package athena

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkGroupQueryExecutions() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_work_group_query_executions",
		Resolver:  fetchAthenaWorkGroupQueryExecutions,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
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
				Name:     "work_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "execution_parameters",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExecutionParameters"),
			},
			{
				Name:     "query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Query"),
			},
			{
				Name:     "query_execution_context",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("QueryExecutionContext"),
			},
			{
				Name:     "query_execution_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryExecutionId"),
			},
			{
				Name:     "result_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultConfiguration"),
			},
			{
				Name:     "statement_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatementType"),
			},
			{
				Name:     "statistics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Statistics"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "work_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkGroup"),
			},
		},
	}
}
