package _map

import "reflect"

func GetFillValidateData(valueData reflect.Value) map[string]interface{} {
	data := make(map[string]interface{}, valueData.Len())

	for _, key := range valueData.MapKeys() {
		data[key.String()] = valueData.MapIndex(key).Interface()
	}

	return data
}
