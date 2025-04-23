package contract

import "reflect"

type ConstraintIsTypeInterface interface {
	DataType() reflect.Kind
	SetDataType(reflect.Kind)
	ConstraintInterface
}
