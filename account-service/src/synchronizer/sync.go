package synchronizer

type Sync interface {
	Start() error
	Stop()
}
type sync struct {
	stop chan bool
}

func NewSync() Sync {
	return &sync{
		stop: make(chan bool),
	}
}

func (s *sync) Start() error {
	select {
	case <-s.stop:
		s.quit
	}
	return nil
}

func (s *sync) Stop() {

}
