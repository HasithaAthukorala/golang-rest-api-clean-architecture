package services

type ClientSet struct{}

func NewClients() (*ClientSet, error) {
	return &ClientSet{}, nil
}
