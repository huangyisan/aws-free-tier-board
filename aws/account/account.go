package account

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"sync"
)

type AWSAccount struct {
	sync     *sync.Mutex
	ak       string
	sk       string
	region   string
	CeClient *costexplorer.Client
}

func (c *AWSAccount) SetAccessKey(accessKey string) {
	c.ak = accessKey
}

func (c *AWSAccount) SetSecretKey(secretKey string) {
	c.sk = secretKey
}

func (c *AWSAccount) SetRegion(region string) {
	c.region = region
}

func (c *AWSAccount) SetStaticProvider() credentials.StaticCredentialsProvider {
	return credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     c.ak,
			SecretAccessKey: c.sk,
		},
	}
}

func (c *AWSAccount) SetClient() {
	staticProvider := c.SetStaticProvider()
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(staticProvider),
		config.WithRegion(c.region),
	)
	if err != nil {
		panic(err)
	}
	client := costexplorer.NewFromConfig(cfg)
	c.CeClient = client
}

func (c *AWSAccount) getAccessKey() string {
	return c.ak
}

func (c *AWSAccount) getSecretKey() string {
	return c.ak
}

func (c *AWSAccount) getRegion() string {
	return c.region
}

func BuildAWSAccount(accessKey, secretKey, region string) *AWSAccount {
	awsAccount := &AWSAccount{
		sync: &sync.Mutex{},
	}
	awsAccount.sync.Lock()
	defer awsAccount.sync.Unlock()
	awsAccount.SetAccessKey(accessKey)
	awsAccount.SetSecretKey(secretKey)
	awsAccount.SetRegion(region)
	awsAccount.SetClient()
	return awsAccount
}
