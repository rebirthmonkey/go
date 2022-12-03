package util

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/thoas/go-funk"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IsNotEmpty checks if the specified object is empty
func IsNotEmpty(obj interface{}) bool {
	return !funk.IsEmpty(obj)
}

func Equal(obj1, obj2 interface{}) bool {
	return funk.Equal(obj1, obj2) || (funk.IsEmpty(obj1) && funk.IsEmpty(obj2))
}

func MapToTyped(objType reflect.Type, data map[string]interface{}) (interface{}, error) {
	objval := reflect.New(objType)
	obj := objval.Interface()
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: func(f reflect.Type, t reflect.Type, data interface{}) (ret interface{}, err error) {
			if f.Kind() != reflect.String {
				return data, nil
			}
			if t != reflect.TypeOf(v1.NewTime(time.Now())) {
				return data, nil
			}

			var ti time.Time
			ti, err = time.Parse(time.RFC3339, data.(string))
			if err != nil {
				return
			}

			ret = v1.NewTime(ti)
			return
		},
		WeaklyTypedInput: true,
		Result:           obj,
		ZeroFields:       true,
		ErrorUnused:      false,
	})
	return obj, decoder.Decode(data)
}
