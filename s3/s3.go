package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Config struct {
	EndpointURL string
	BucketName  string
}

func NewUploader(ctx context.Context, cfg Config) (*manager.Uploader, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           cfg.EndpointURL,
			SigningRegion: "ap-northeast-1",
		}, nil
	})
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(awsCfg)
	return manager.NewUploader(s3Client), nil
}
