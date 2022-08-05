package awscost

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/app/aws/account"
	cost2 "trojan-dashboard/app/aws/service/cost"
)

var (
	ak     = "abc"
	sk     = "def"
	region = "ap-southeast-1"
)

func List(c *gin.Context) {
	start := "2022-07-01"
	end := "2022-08-01"

	awsAccount := account.BuildAWSAccount(ak, sk, region)
	client := cost2.NewClient(awsAccount)
	co := cost2.NewCost(start, end)
	//params := c.MonthlyEC2DataTransferOutCostAndUsageInput()
	params := co.DailyEC2DataTransferOutCostAndUsageInput()
	res := cost2.GetCostAndUsage(client.AWSAccount.CeClient, context.Background(), params)
	c.JSON(http.StatusOK, res)
}
