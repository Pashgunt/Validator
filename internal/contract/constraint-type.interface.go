package contract

import "reflect"

type ConstraintIsTypeInterface interface {
	DataType() reflect.Kind
	ConstraintInterface
}
