package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/b-b3rn4rd/repository-design-pattern/repository"
	"github.com/guregu/dynamo"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	rr := repository.NewRepositoryRegistry(
		dynamo.New(sess),
		logger,
		&repository.UserRepository{},
	)

	users, err := rr.MustRepository("UserRepository").List()
	if err != nil {
		logger.WithError(err).Fatal("an error has occurred")
	}

	fmt.Println(users)
}
