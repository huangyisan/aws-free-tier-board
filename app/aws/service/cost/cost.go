package cost

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

type Cost struct {
	start *string
	end   *string
}

func NewCost(start, end string) *Cost {
	return &Cost{
		&start,
		&end,
	}
}

func (c *Cost) DailyEC2DataTransferOutCostAndUsageInput() *costexplorer.GetCostAndUsageInput {
	Granularity := types.GranularityDaily
	Values := []string{"EC2: Data Transfer - Internet (Out)"}
	return &costexplorer.GetCostAndUsageInput{
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key:    types.DimensionUsageTypeGroup,
				Values: Values,
			},
		},
		Granularity: Granularity,
		TimePeriod: &types.DateInterval{
			Start: c.start,
			End:   c.end,
		},
		Metrics: []string{"UsageQuantity"},
	}
}

func (c *Cost) MonthlyEC2DataTransferOutCostAndUsageInput() *costexplorer.GetCostAndUsageInput {
	Granularity := types.GranularityMonthly
	Values := []string{"EC2: Data Transfer - Internet (Out)"}
	return &costexplorer.GetCostAndUsageInput{
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key:    types.DimensionUsageTypeGroup,
				Values: Values,
			},
		},
		Granularity: Granularity,
		TimePeriod: &types.DateInterval{
			Start: c.start,
			End:   c.end,
		},
		Metrics: []string{"UsageQuantity"},
	}
}

func GetCostAndUsage(CC costHandler, ctx context.Context, params *costexplorer.GetCostAndUsageInput) []UsageInfo {
	var usageInfos []UsageInfo
	res, err := CC.GetCostAndUsage(ctx, params)
	if err != nil {
		panic(err)
	}

	for _, v := range res.ResultsByTime {
		for _, v2 := range v.Total {
			usageInfos = append(usageInfos, UsageInfo{
				*v.TimePeriod.Start,
				*v2.Amount,
			})
		}
	}
	return usageInfos
}
