package pagerduty

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/heimweh/go-pagerduty/pagerduty"
)

func dataSourcePagerDutyEventOrchestrationServiceCacheVariable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePagerDutyEventOrchestrationServiceCacheVariableRead,
		Schema: map[string]*schema.Schema{
			"service": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"condition": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: dataSourceEventOrchestrationCacheVariableConditionSchema,
				},
			},
			"configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: dataSourceEventOrchestrationCacheVariableConfigurationSchema,
				},
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourcePagerDutyEventOrchestrationServiceCacheVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceEventOrchestrationCacheVariableRead(ctx, d, meta, pagerduty.CacheVariableTypeService)
}
