package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/b-b3rn4rd/repository-design-pattern/repository"
	"github.com/guregu/dynamo"
	"github.com/sirupsen/logrus"
)

func main() {
	rr := repository.NewRepositoryRegistry(
		dynamo.New(session.New(), &aws.Config{Region: aws.String("us-west-2")}),
		logrus.New(),
		&repository.UserRepository{},
	)

	rr.MustRepository("UserRepository").List()
}
