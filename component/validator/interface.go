package validator

//go:generate mockery --name=BlacklistRepository
type BlacklistRepository interface {
	List() (items []string, err error)
}
