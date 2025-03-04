package _struct

import (
	strhelper "github.com/Pashgunt/Validator/internal/helper/string"
	"reflect"
)

const (
	prefixGetterMethod = "Get"
)

func GetFillValidateData(valueData reflect.Value) map[string]interface{} {
	data := make(map[string]interface{}, valueData.NumField())

	for i := 0; i < valueData.NumField(); i++ {
		if valueData.Field(i).IsValid() && !valueData.Field(i).CanInterface() {
			data[valueData.Type().Field(i).Name] = getValueFieldByMethod(
				valueData,
				getGetterMethodByField(valueData.Type().Field(i)),
			)

			continue
		}

		data[valueData.Type().Field(i).Name] = valueData.Field(i).Interface()
	}

	return data
}

func getValueFieldByMethod(valueData reflect.Value, methodName string) interface{} {
	method := valueData.Addr().MethodByName(methodName)

	if !method.IsValid() || method.Type().NumIn() != 0 {
		return nil
	}

	return method.Call(nil)[0].Interface()
}

func getGetterMethodByField(valueData reflect.StructField) string {
	return prefixGetterMethod + string(strhelper.CapitalizeStringLetterByIndex(valueData.Name, 0))
}
