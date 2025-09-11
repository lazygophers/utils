package candy

import (
	"reflect"
)

// DiffSlice 比较两个切片的差异
// 返回第一个切片中存在但第二个切片中不存在的元素，以及第二个切片中存在但第一个切片中不存在的元素
func DiffSlice(a interface{}, b interface{}) (interface{}, interface{}) {
	at := reflect.TypeOf(a)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}

	bt := reflect.TypeOf(b)
	if bt.Kind() != reflect.Slice {
		panic("b is not slice")
	}

	atm := at.Elem()
	btm := bt.Elem()

	if atm.Kind() != btm.Kind() {
		panic("a and b are not same type")
	}

	m := map[interface{}]reflect.Value{}

	bv := reflect.ValueOf(b)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = bv.Index(i)
	}

	c := reflect.MakeSlice(at, 0, 0)
	d := reflect.MakeSlice(bt, 0, 0)
	av := reflect.ValueOf(a)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()].IsValid() {
			c = reflect.Append(c, av.Index(i))
		} else {
			delete(m, av.Index(i).Interface())
		}
	}

	for _, value := range m {
		d = reflect.Append(d, value)
	}

	return c.Interface(), d.Interface()
}