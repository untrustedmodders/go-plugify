package plugify

// TypeNotFoundException represents a type not found error
type TypeNotFoundException struct {
	message string
}

func (e *TypeNotFoundException) Error() string {
	return e.message
}

// NewTypeNotFoundException creates a new TypeNotFoundException
func NewTypeNotFoundException(message string) error {
	return &TypeNotFoundException{message: message}
}
