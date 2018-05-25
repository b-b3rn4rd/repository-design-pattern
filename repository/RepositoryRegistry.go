package repository

import (
	"reflect"

	"fmt"

	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Configure(*dynamo.DB, *logrus.Logger)
	List() (interface{}, error)
	Get(id interface{}) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Update(id interface{}, entity interface{}) (bool, error)
	Delete(id interface{}) (bool, error)
}

type DynamoRepository struct {
	db     *dynamo.DB
	logger *logrus.Logger
}

func (r *DynamoRepository) List() (interface{}, error) {
	r.logger.Panic("the following method must be implemented in the child struct")
	return nil, nil
}

func (r *DynamoRepository) Get(id interface{}) (interface{}, error) {
	r.logger.Panic("the following method must be implemented in the child struct")
	return nil, nil
}

func (r *DynamoRepository) Create(entity interface{}) (interface{}, error) {
	r.logger.Panic("the following method must be implemented in the child struct")
	return false, nil
}

func (r *DynamoRepository) Update(id interface{}, entity interface{}) (bool, error) {
	r.logger.Panic("the following method must be implemented in the child struct")
	return false, nil
}

func (r *DynamoRepository) Delete(id interface{}) (bool, error) {
	r.logger.Panic("the following method must be implemented in the child struct")
	return false, nil
}

func (r *DynamoRepository) Configure(db *dynamo.DB, logger *logrus.Logger) {
	r.db = db
	r.logger = logger
}

type RepositoryRegistry struct {
	registry map[string]Repository

	db     *dynamo.DB
	logger *logrus.Logger
}

func NewRepositoryRegistry(db *dynamo.DB, logger *logrus.Logger, repository ...Repository) *RepositoryRegistry {
	r := &RepositoryRegistry{
		db:     db,
		logger: logger,
	}

	r.registerRepositories(repository)
	return r
}

func (r *RepositoryRegistry) registerRepositories(repositories []Repository) {
	for _, repository := range repositories {
		repository.Configure(r.db, r.logger)
		r.registry[reflect.TypeOf(repository).String()] = repository
	}
}

func (r *RepositoryRegistry) Repository(repositoryName string) (Repository, error) {
	if repository, ok := r.registry[repositoryName]; ok {
		return repository, nil
	}

	return nil, errors.New(fmt.Sprintf("Repository %s does not exist", repositoryName))
}

func (r *RepositoryRegistry) MustRepository(repositoryName string) (repository Repository) {
	repository, err := r.Repository(repositoryName)
	if err != nil {
		r.logger.Panic(err.Error())
	}

	return repository
}
