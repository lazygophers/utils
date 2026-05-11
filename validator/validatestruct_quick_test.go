package validator

import (
	"fmt"
	"reflect"
	"testing"
)

// 快速测试数据结构
type QuickPerson struct {
	Name    string   `validate:"required"`
	Age     int      `validate:"gte=0,lte=150"`
	Email   string   `validate:"email"`
	Address Address  `validate:""`
	Tags    []string `validate:"dive,omitempty"`
}

func generateQuickPersonData(n int) []QuickPerson {
	people := make([]QuickPerson, n)
	for i := 0; i < n; i++ {
		people[i] = QuickPerson{
			Name:  fmt.Sprintf("Person%d", i),
			Age:   20 + i%50,
			Email: fmt.Sprintf("person%d@example.com", i),
			Address: Address{
				Street:  fmt.Sprintf("%d Street", i),
				City:    "City",
				ZipCode: "12345",
			},
			Tags: []string{"tag1", "tag2", "tag3"},
		}
	}
	return people
}

// Baseline: 当前实现
func (e *Engine) validateStruct_Baseline(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			if field.Kind() == reflect.Struct {
				e.validateStruct_Baseline(top, field, fieldName, errors)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Baseline(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						if elem.Kind() == reflect.Struct {
							e.validateStruct_Baseline(top, elem, elemFieldName, errors)
						} else if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Baseline(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		if field.Kind() == reflect.Struct {
			e.validateStruct_Baseline(top, field, fieldName, errors)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Baseline(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案1: 缓存 Kind
func (e *Engine) validateStruct_Opt1(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt1(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt1(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt1(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt1(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt1(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt1(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案2: 内联访问
func (e *Engine) validateStruct_Opt2(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt2(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt2(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt2(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt2(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt2(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt2(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案3: 对象池
func (e *Engine) validateStruct_Opt3(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt3(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt3(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt3(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt3(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt3(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt3(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案4: 组合优化 (Kind缓存 + 内联访问)
func (e *Engine) validateStruct_Opt4(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt4(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt4(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt4(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt4(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt4(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt4(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案5: 完整优化 (Kind缓存 + 内联访问 + 对象池)
func (e *Engine) validateStruct_Opt5(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt5(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt5(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt5(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt5(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt5(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt5(top, field.Elem(), fieldName, errors)
		}
	}
}

// 基准测试
func BenchmarkValidateStruct_Comparison_Simple(b *testing.B) {
	v, _ := New()
	person := QuickPerson{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Baseline(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt1-KindCache", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt1(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt2-InlineAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt2(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt3-ObjectPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt3(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt4-Combined", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt4(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt5-FullOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt5(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})
}

func BenchmarkValidateStruct_Comparison_Nested(b *testing.B) {
	v, _ := New()
	people := generateQuickPersonData(10)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Baseline(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt1-KindCache", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt1(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt2-InlineAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt2(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt3-ObjectPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt3(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt4-Combined", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt4(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt5-FullOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt5(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})
}
