package validator

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 测试用的复杂嵌套结构体
type Address struct {
	Street  string `validate:"required"`
	City    string `validate:"required"`
	ZipCode string `validate:"len=5"`
}

type Person struct {
	Name    string   `validate:"required"`
	Age     int      `validate:"gte=0,lte=150"`
	Email   string   `validate:"email"`
	Address Address  `validate:""`
	Tags    []string `validate:"dive,omitempty"`
}

type Company struct {
	Name    string    `validate:"required"`
	CEO     Person    `validate:"required"`
	Employees []Person `validate:"dive"`
}

// 生成测试数据
func generatePersonData(n int) []Person {
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = Person{
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

// 基准测试：当前实现
func BenchmarkValidateStruct_Current_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(person)
	}
}

func BenchmarkValidateStruct_Current_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(company)
	}
}

func BenchmarkValidateStruct_Current_Large(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(100),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(company)
	}
}

// 方案1: 缓存 field.Kind() 结果
func (e *Engine) validateStruct_Opt1_KindCache(top, current reflect.Value, namespace string, errors *ValidationErrors) {
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
				e.validateStruct_Opt1_KindCache(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt1_KindCache(top, elem, fieldName, errors)
				}
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
							e.validateStruct_Opt1_KindCache(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt1_KindCache(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt1_KindCache(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt1_KindCache(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt1_KindCache_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt1_KindCache(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt1_KindCache_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt1_KindCache(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案2: 减少字段名称字符串拼接（只在需要时拼接）
func (e *Engine) validateStruct_Opt2_LazyNamespace(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		var fieldName string
		if namespace == "" {
			fieldName = fieldType.Name
		} else {
			fieldName = namespace + "." + fieldType.Name
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			if field.Kind() == reflect.Struct {
				e.validateStruct_Opt2_LazyNamespace(top, field, fieldName, errors)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt2_LazyNamespace(top, field.Elem(), fieldName, errors)
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
						elemFieldName := fieldName + "[" + fmt.Sprintf("%d", j) + "]"

						if elem.Kind() == reflect.Struct {
							e.validateStruct_Opt2_LazyNamespace(top, elem, elemFieldName, errors)
						} else if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt2_LazyNamespace(top, elem.Elem(), elemFieldName, errors)
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt2_LazyNamespace(top, field, fieldName, errors)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt2_LazyNamespace(top, field.Elem(), fieldName, errors)
		}
	}
}

func BenchmarkValidateStruct_Opt2_LazyNamespace_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt2_LazyNamespace(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案3: 减少反射调用（内联字段访问）
func (e *Engine) validateStruct_Opt3_InlineAccess(top, current reflect.Value, namespace string, errors *ValidationErrors) {
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
				e.validateStruct_Opt3_InlineAccess(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt3_InlineAccess(top, elem, fieldName, errors)
				}
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
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt3_InlineAccess(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt3_InlineAccess(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt3_InlineAccess(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt3_InlineAccess(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt3_InlineAccess_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt3_InlineAccess(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案4: 使用 strings.Builder 替代字符串拼接
func (e *Engine) validateStruct_Opt4_StringBuilder(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		var fieldName string
		if namespace == "" {
			fieldName = fieldType.Name
		} else {
			var builder strings.Builder
			builder.Grow(len(namespace) + len(fieldType.Name) + 1)
			builder.WriteString(namespace)
			builder.WriteByte('.')
			builder.WriteString(fieldType.Name)
			fieldName = builder.String()
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt4_StringBuilder(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt4_StringBuilder(top, elem, fieldName, errors)
				}
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
						var builder strings.Builder
						builder.Grow(len(fieldName) + 12)
						builder.WriteString(fieldName)
						builder.WriteByte('[')
						builder.WriteString(fmt.Sprint(j))
						builder.WriteByte(']')
						elemFieldName := builder.String()

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt4_StringBuilder(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt4_StringBuilder(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt4_StringBuilder(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt4_StringBuilder(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt4_StringBuilder_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt4_StringBuilder(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案5: 使用 sync.Pool 对象池复用 fieldLevel
// 注意：对象池已在 engine.go 中定义，此处使用全局对象池

func (e *Engine) validateStruct_Opt5_ObjectPool(top, current reflect.Value, namespace string, errors *ValidationErrors) {
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
				e.validateStruct_Opt5_ObjectPool(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt5_ObjectPool(top, elem, fieldName, errors)
				}
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
						elemFieldName := fieldName + "[" + fmt.Sprint(j) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt5_ObjectPool(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt5_ObjectPool(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
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
			e.validateStruct_Opt5_ObjectPool(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt5_ObjectPool(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt5_ObjectPool_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt5_ObjectPool(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案6: 组合优化（Kind缓存 + 内联访问）
func (e *Engine) validateStruct_Opt6_Combined(top, current reflect.Value, namespace string, errors *ValidationErrors) {
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
				e.validateStruct_Opt6_Combined(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt6_Combined(top, elem, fieldName, errors)
				}
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
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt6_Combined(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt6_Combined(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt6_Combined(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt6_Combined(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt6_Combined_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt6_Combined(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt6_Combined_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt6_Combined(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案7: 组合优化 + 对象池
func (e *Engine) validateStruct_Opt7_FullCombined(top, current reflect.Value, namespace string, errors *ValidationErrors) {
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
				e.validateStruct_Opt7_FullCombined(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt7_FullCombined(top, elem, fieldName, errors)
				}
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
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt7_FullCombined(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt7_FullCombined(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
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
			e.validateStruct_Opt7_FullCombined(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt7_FullCombined(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt7_FullCombined_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt7_FullCombined(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt7_FullCombined_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt7_FullCombined(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案8: 预先提取常用值到局部变量
func (e *Engine) validateStruct_Opt8_LocalVars(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()
	tagName := e.tagName
	fieldNameFunc := e.fieldNameFunc

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

		tag := fieldType.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt8_LocalVars(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt8_LocalVars(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := fieldNameFunc(fieldType)

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
						elemFieldName := fieldName + "[" + fmt.Sprint(j) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt8_LocalVars(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt8_LocalVars(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt8_LocalVars(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt8_LocalVars(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt8_LocalVars_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt8_LocalVars(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案9: 减少函数调用（内联简单检查）
func (e *Engine) validateStruct_Opt9_InlinedChecks(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		// 内联 IsExported 检查
		if fieldType.PkgPath != "" {
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
				e.validateStruct_Opt9_InlinedChecks(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt9_InlinedChecks(top, elem, fieldName, errors)
				}
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
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt9_InlinedChecks(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt9_InlinedChecks(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
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
			e.validateStruct_Opt9_InlinedChecks(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt9_InlinedChecks(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt9_InlinedChecks_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt9_InlinedChecks(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案10: 完整优化组合
func (e *Engine) validateStruct_Opt10_AllOptimizations(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()
	tagName := e.tagName
	fieldNameFunc := e.fieldNameFunc

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if fieldType.PkgPath != "" {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt10_AllOptimizations(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt10_AllOptimizations(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := fieldNameFunc(fieldType)

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
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt10_AllOptimizations(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt10_AllOptimizations(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
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
			e.validateStruct_Opt10_AllOptimizations(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt10_AllOptimizations(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Large(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(100),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}
