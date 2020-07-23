package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func GetEcsSession() *ecs.ECS {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		panic(err)
	}

	ecsSess := ecs.New(session.New(sess.Config))

	return ecsSess
}
