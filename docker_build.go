package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func dockerBuild() {
	sharedServicesAccountID := os.Getenv("SHARED_SERVICES_ACCOUNT_ID")
	awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	dockerTag := os.Getenv("DOCKER_TAG")
	ecrName := os.Getenv("ECR_NAME")

	buildCmd := fmt.Sprintf("docker build -t %s .", ecrName)
	err := exec.Command("bash", "-c", buildCmd).Run()

	if err != nil {
		log.Fatalln("Error running docker build")
	}

	tagArgs := fmt.Sprintf(
		"%s:latest %s.dkr.ecr.%s.amazonaws.com/%s:%s",
		ecrName,
		sharedServicesAccountID,
		awsDefaultRegion,
		ecrName,
		dockerTag,
	)

	tagCmd := fmt.Sprintf("docker tag %s", tagArgs)
	err = exec.Command("bash", "-c", tagCmd).Run()

	if err != nil {
		log.Fatalln("Error running docker tag")
	}

	log.Println("Docker container built successfully!")
}
