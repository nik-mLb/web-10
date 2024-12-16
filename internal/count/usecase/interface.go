package usecase

type Provider interface {
	SelectCounter() (int, error)
	UpdateCounter(int) error
}