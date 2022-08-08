package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	finalArray := []interface{}{}
	finalArray = CheckValue(nested, &finalArray)

	return finalArray

}

func CheckValue(value interface{}, currentArray *[]interface{}) []interface{} {
	valueType := reflect.ValueOf(value).Kind()
	if valueType.String() == "slice" {
		for _, v := range value.([]interface{}) {
			CheckValue(v, currentArray)
		}
	} else if value != nil {
		*currentArray = append(*currentArray, value)

	}
	return *currentArray

}
