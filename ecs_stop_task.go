package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/rennanbadaro/ci-packages/utils"
)

func stopECSTasks() {
	ecsSession := utils.GetEcsSession()

	listTaskInput := &ecs.ListTasksInput{
		Cluster: aws.String("development"),
		Family:  aws.String("onb-dev-compliance-service"),
	}

	result, err := ecsSession.ListTasks(listTaskInput)
	if err != nil {
		panic(err)
	}

	taskArns := result.TaskArns

	for i := 0; i < len(taskArns); i++ {
		fmt.Println("Task ARN:", *taskArns[i])
		splittedArn := strings.Split(*taskArns[i], "/")

		taskID := splittedArn[len(splittedArn)-1]

		stopTaskInput := &ecs.StopTaskInput{
			Task:    aws.String(taskID),
			Cluster: aws.String("development"),
		}

		if _, err := ecsSession.StopTask(stopTaskInput); err != nil {
			panic(err)
		}
	}
}
