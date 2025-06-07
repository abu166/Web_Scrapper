package domain

// Repository defines the interface for data storage
type Repository interface {
	SaveTablets(tablets []Tablet) error
	Initialize() error
}
