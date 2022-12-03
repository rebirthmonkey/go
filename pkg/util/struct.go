package util

import (
	"errors"
	"reflect"
)

func StructCopy(output interface{}, input interface{}, ignoreZero bool) error {
	srcv := reflect.ValueOf(input)
	dstv := reflect.ValueOf(output)
	srct := reflect.TypeOf(input)
	dstt := reflect.TypeOf(output)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		return errors.New("StructCopy error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		return errors.New("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(input).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			if !src.IsZero() || !ignoreZero {
				dst.Set(src)
				continue
			} else {
				continue
			}
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return nil
}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}
