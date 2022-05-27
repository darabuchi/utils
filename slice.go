package utils

import (
	"fmt"
	"reflect"
)

func PluckUint64(list interface{}, fieldName string) []uint64 {
	v := reflect.ValueOf(list)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		var result []uint64
		for i := 0; i < v.Len(); i++ {
			ev := v.Index(i)
			
			for ev.Kind() == reflect.Ptr {
				ev = ev.Elem()
			}
			
			if ev.Kind() != reflect.Struct {
				panic("element is not a struct")
			}
			
			if !ev.IsValid() {
				continue
			}
			
			et := ev.Type()
			_, ok := et.FieldByName(fieldName)
			if !ok {
				panic(fmt.Sprintf("field %s not found", fieldName))
			}
			
			field := ev.FieldByName(fieldName)
			if !field.IsValid() {
				continue
			}
			
			if field.Kind() != reflect.Uint64 {
				panic(fmt.Sprintf("field %s is not uint64", fieldName))
			}
			
			result = append(result, field.Uint())
		}
		
		return result
	
	default:
		panic("list must be an array or slice")
	}
}

func PluckString(list interface{}, fieldName string) []string {
	v := reflect.ValueOf(list)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		var result []string
		for i := 0; i < v.Len(); i++ {
			ev := v.Index(i)
			
			for ev.Kind() == reflect.Ptr {
				ev = ev.Elem()
			}
			
			if ev.Kind() != reflect.Struct {
				panic("element is not a struct")
			}
			
			if !ev.IsValid() {
				continue
			}
			
			et := ev.Type()
			_, ok := et.FieldByName(fieldName)
			if !ok {
				panic(fmt.Sprintf("field %s not found", fieldName))
			}
			
			field := ev.FieldByName(fieldName)
			if !field.IsValid() {
				continue
			}
			
			if field.Kind() != reflect.String {
				panic(fmt.Sprintf("field %s is not uint64", fieldName))
			}
			
			result = append(result, field.String())
		}
		
		return result
	
	default:
		panic("list must be an array or slice")
	}
}

// DiffSlice 传入两个slice
// 如果 a 或者 b 不为 slice 会 panic
// 如果 a 与 b 的元素类型不一致，也会 panic
// 返回的第一个参数为 a 比 b 多的，类型为 a 的类型
// 返回的第二个参数为 b 比 a 多的，类型为 b 的类型
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

// RemoveSlice 传入两个slice
// 如果 src 或者 rm 不为 slice 会 panic
// 如果 src 与 rm 的元素类型不一致，也会 panic
// 返回的第一个参数为 src 中不在 rm 中的元素，数据类型与 src 一致
func RemoveSlice(src interface{}, rm interface{}) interface{} {
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}
	
	bt := reflect.TypeOf(src)
	if bt.Kind() != reflect.Slice {
		panic("b is not slice")
	}
	
	atm := at.Elem()
	btm := bt.Elem()
	
	if atm.Kind() != btm.Kind() {
		panic("a and b are not same type")
	}
	
	m := map[interface{}]bool{}
	
	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}
	
	c := reflect.MakeSlice(at, 0, 0)
	av := reflect.ValueOf(src)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
			delete(m, av.Index(i).Interface())
		}
	}
	
	return c.Interface()
}

func KeyBy(list interface{}, fieldName string) interface{} {
	lv := reflect.ValueOf(list)
	
	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}
	
	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}
	
	if evs.Kind() != reflect.Struct {
		panic("list element is not struct")
	}
	
	if evs.Kind() != reflect.Struct {
		panic("element not struct")
	}
	
	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}
	
	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}
		
		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}
		
		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}
		
		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}
	
	return m.Interface()
}
