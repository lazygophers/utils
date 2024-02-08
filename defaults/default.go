package defaults

import (
	"reflect"
	"strconv"
)

func SetDefaults(value interface{}) error {
	vv := reflect.ValueOf(value)
	switch vv.Type().Kind() {
	case reflect.Ptr:
		// do nothing
	case reflect.Slice:
		return nil
	case reflect.Map:
		return nil
	default:
		panic("invalid out type, not ptr")
	}

	for vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	switch vv.Type().Kind() {
	case reflect.Struct:
		// do nothing
	case reflect.Slice:
		return nil
	case reflect.Map:
		return nil
	default:
		panic("invalid out elem type, not struct")
	}

	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		fieldType := vv.Type().Field(i)

		defaultStr, ok := fieldType.Tag.Lookup("default")
		if !ok {
			continue
		}

		if defaultStr == "-" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if defaultStr != "" && field.String() == "" {
				field.SetString(defaultStr)
			}

		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			if defaultStr != "" && field.Uint() != 0 {
				val, err := strconv.ParseUint(defaultStr, 10, 64)
				if err == nil {
					field.SetUint(val)
				}
			}

		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			if defaultStr != "" && field.Int() != 0 {
				val, err := strconv.ParseInt(defaultStr, 10, 64)
				if err == nil {
					field.SetInt(val)
				}
			}

		case reflect.Map:
		// do nothing
		case reflect.Ptr:
		// do nothing
		case reflect.Slice:
			// do nothing
		default:
			panic("unknown kind " + field.Kind().String())
		}
	}

	return nil
}
