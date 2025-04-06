package test

import (
	validator "github.com/Pashgunt/Validator"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"testing"
)

const (
	validatorNameWordCount = "WordCount"
)

type wordCountArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newWordCountArgs(value interface{}, min int, max int, minMessage string, maxMessage string) *wordCountArgs {
	return &wordCountArgs{
		value:      value,
		constraint: validator.NewWordCount(min, max, minMessage, maxMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestWordCountValidator(t *testing.T) {
	tests := []struct {
		name string
		wordCountArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name: "test isset error message " + validatorNameWordCount + " validator",
			wordCountArgs: *newWordCountArgs(
				"str",
				5,
				10,
				testhelper.DefaultErrorMessage,
				testhelper.DefaultErrorMessage,
			),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name: "test blank error message " + validatorNameWordCount + " validator",
			wordCountArgs: *newWordCountArgs(
				"string str2 str3",
				0,
				10,
				testhelper.DefaultErrorMessage,
				testhelper.DefaultErrorMessage,
			),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.WordCountValidator{}).Process(
				test.wordCountArgs.constraint,
				test.wordCountArgs.value,
				test.wordCountArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameWordCount+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameWordCount+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
