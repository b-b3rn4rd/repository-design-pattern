# repository-design-pattern

The repository design pattern that has following characteristics

* Easy to register new repositories with minimal repetitive code
* Single database connection (aws session) across all repositories
* Parent struct that contains shared fields
* Each repository must implement standard Restful methods, however can contain any additional functions
* All repositories are stored in a registry and can be passed down the application as a single struct

Example:
---------
main.go

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

UserRepository.go
```go
package repository

import (
	"github.com/b-b3rn4rd/repository-design-pattern/model"
	"github.com/pkg/errors"
)

type UserRepository struct {
	DynamoRepository
}

func (r *UserRepository) List() (interface{}, error) {
	r.logger.Debug("listing all users")

	var users []model.User

	err := r.db.Table("User").Scan().All(&users)
	if err != nil {
		return nil, errors.Wrap(err, "error while listing users")
	}

	return users, nil
}

func (r *UserRepository) Get(id interface{}) (interface{}, error) {
	r.logger.WithField("id", id).Debug("retrieving single user")

	var user []model.User

	err := r.db.Table("User").Get("Id", id).One(&user)
	if err != nil {
		return nil, errors.Wrap(err, "error while retrieving single user")
	}

	return user, nil
}

func (r *UserRepository) Create(entity interface{}) (interface{}, error) {
	r.logger.WithField("user", entity).Debug("create user")

	err := r.db.Table("User").Put(entity.(model.User)).Run()
	if err != nil {
		return false, errors.Wrap(err, "error while creating user")
	}

	return true, nil
}

func (r *UserRepository) Update(id interface{}, entity interface{}) (bool, error) {
	r.logger.WithField("user", entity).Debug("update user")

	err := r.db.Table("User").Update("Id", id).Set("Email", entity.(model.User).Email).Run()
	if err != nil {
		return false, errors.Wrap(err, "error while updating user")
	}

	return true, nil
}

func (r *UserRepository) Delete(id interface{}) (bool, error) {
	r.logger.WithField("id", id).Debug("delete user")

	err := r.db.Table("User").Delete("Id", id).Run()
	if err != nil {
		return false, errors.Wrap(err, "error while deleting user")
	}

	return false, nil
}

func (r *UserRepository) GetGroups(id interface{}) (interface{}, error) {
	r.logger.WithField("id", id).Debug("get user groups")

	return nil, errors.New("error retrieving user's groups")
}

```