package db

// Err adds additional context to the error by storing the message and intended
// status code to be returned to the client.
type Err struct {
	Msg    string
	Status int
}

// Error implements Error interface that returns only the message.
func (e *Err) Error() string {
	return e.Msg
}
