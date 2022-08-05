package cost

import (
	"trojan-dashboard/app/aws/account"
)

type Client struct {
	AWSAccount *account.AWSAccount
}

func NewClient(awsAccount *account.AWSAccount) *Client {
	return &Client{
		AWSAccount: awsAccount,
	}
}

type UsageInfo struct {
	Date  string `json:"date"`
	Usage string `json:"usage"`
}
