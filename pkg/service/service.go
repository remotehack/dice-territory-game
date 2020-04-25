// service package will hold all the things necessary for the service package
package service

type Service struct {
	websocket struct{}
	http      struct{}
}

func (s Service) Start() error {
	panic("implement me")
}

func (s Service) Stop() error {
	panic("implement me")
}
