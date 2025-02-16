package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

func Marshal(v any) ([][]string, error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return nil, errors.New("expected a slice of struct")
	}
	structType := vv.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("expected a slice of struct")
	}

	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < vv.Len(); i++ {
		row, err := marshalOne(vv.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(t reflect.Type) []string {
	var row []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		curTag, ok := field.Tag.Lookup("csv")
		if !ok {
			continue
		}
		row = append(row, curTag)
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string
	vt := vv.Type()
	for i := 0; i < vt.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}
	return row, nil
}

func Unmarshal(data [][]string, v any) error {
	sliceValPointer := reflect.ValueOf(v)
	if sliceValPointer.Kind() != reflect.Pointer {
		return errors.New("expect a pointer to a slice of struct")
	}
	sliceVal := sliceValPointer.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("expect a pointer to a slice of struct")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("expect a pointer to a slice of struct")
	}

	header := data[0]
	namePos := make(map[string]int, len(header))
	for i, v := range header {
		namePos[v] = i
	}

	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(newVal, row, namePos)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func unmarshalOne(newVal reflect.Value, row []string,
	namePos map[string]int) error {
	vt := newVal.Type()
	for i := 0; i < vt.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := namePos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := newVal.Field(i)

		switch field.Kind() {
		case reflect.Int:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.String:
			field.SetString(val)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		default:
			return fmt.Errorf("cannot handle field of kind %v",
				field.Kind())
		}

	}
	return nil
}

func main() {
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith",42,false
Martha,37,"true"
`
	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []MyData
	Unmarshal(allData, &entries)
	fmt.Println(entries)

	//now to turn entries into output
	out, err := Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
