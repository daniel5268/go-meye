package domain

type DomainError struct {
	Section string
	Code    string
	Err     error
}

const (
	CodeBindError               = "bind_error"
	CodeValidationError         = "validation_error"
	CodeInvalidCredentialsError = "invalid_credentials_error"
	CodeSignTokenError          = "sign_token_error"
	CodeUserNotFoundError       = "user_not_found_error"
	CodeRepositoryError         = "repository_error"
	CodeHashError               = "hash_error"
	CodeForbiddenError          = "forbidden_error"
	CodeUserAlreadyCreatedError = "user_already_created_error"
)

func (e *DomainError) Error() string {
	return e.Err.Error()
}

func NewDomainError(section, code string, err error) *DomainError {
	return &DomainError{
		Section: section,
		Code:    code,
		Err:     err,
	}
}
