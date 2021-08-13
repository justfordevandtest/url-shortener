package blacklist

type StaticRepo struct {
	list []string
}

func New(list []string) (repo *StaticRepo, err error) {
	return &StaticRepo{
		list: list,
	}, nil
}

func (repo *StaticRepo) List() (items []string, err error) {
	return repo.list, nil
}
