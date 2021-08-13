package admin

import "shorturl/entity"

//go:generate mockery --name=Comp
type Comp interface {
	List(input *ListInput) (output *ListOutput, err error)
	Delete(input *DelInput) (err error)
}

//go:generate mockery --name=URLRepo
type URLRepo interface {
	List(page int, perPage int, filters map[string]interface{}) (total int, items []entity.URL, err error)
	Read(ID string) (ent *entity.URL, err error)
	Delete(ID string) (err error)
}

//go:generate mockery --name=Validator
type Validator interface {
	Validate(item interface{}) (err error)
}
