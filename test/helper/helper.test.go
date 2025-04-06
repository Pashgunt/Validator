package test

const (
	DefaultErrorMessage       = "Error"
	DefaultIssetErrorCount    = 1
	DefaultNotIssetErrorCount = 0
	BlankString               = ""
)

type ResultValidator struct {
	count   int
	message string
}

func (r ResultValidator) Count() int {
	return r.count
}

func (r ResultValidator) Message() string {
	return r.message
}

func NewResultValidator(count int, message string) *ResultValidator {
	return &ResultValidator{count: count, message: message}
}
