package datasource

type Datasource interface {
	IsAvailable() bool
	AvailabilityChanges() bool
	ConfigRoot() string
	FetchMetadata() ([]byte, error)
	FetchUserdata() ([]byte, error)
	Type() string
}
