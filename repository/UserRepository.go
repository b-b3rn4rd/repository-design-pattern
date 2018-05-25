package repository

import "github.com/b-b3rn4rd/repository-design-pattern/model"

type UserRepository struct {
	DynamoRepository
}

func (r *UserRepository) List() (interface{}, error) {
	var users []model.User

	r.db.Table("User").Scan().All(&users)
	return users, nil
}

func (r *UserRepository) Get(id interface{}) (interface{}, error) {
	var user []model.User
	r.db.Table("User").Get("Id", id).One(&user)
	return user, nil
}

func (r *UserRepository) Create(entity interface{}) (interface{}, error) {
	err := r.db.Table("User").Put(entity.(model.User)).Run()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) Update(id interface{}, entity interface{}) (bool, error) {
	err := r.db.Table("User").Update("Id", id).Set("Email", entity.(model.User).Email).Run()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) Delete(id interface{}) (bool, error) {
	err := r.db.Table("User").Delete("Id", id).Run()
	if err != nil {
		return false, err
	}

	return false, nil
}
