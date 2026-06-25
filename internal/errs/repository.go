package errs

func NewRepositoryError(message string) *RepositoryError {
	return &RepositoryError{
		Message: message,
	}
}

type RepositoryError struct {
	Message string
}

func (e *RepositoryError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Erro no repositório"
}
