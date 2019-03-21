package synchronizer

type Sync interface {
	Start() error
	Stop()
}
