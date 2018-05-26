# repository-design-pattern

The repository design pattern that has following characteristics

* Easy to register new repositories with minimal repetitive code
* Single database connection (aws session) across all repositories
* Parent struct that contains shared field
* Each repository must implement standard Restful methods, put can contain any additional functions

Example:
---------
```go
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
```