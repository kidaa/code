/**
 * Created by Michael on 2015/7/31.
 */
package redis
import (
	"errors"
	"reflect"
	"strconv"
)

//转码到缓存
func Marshal(pro interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(pro)
	var ret = map[string]interface{}{}
	val = val.Elem()
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, errors.New("Is not a Struct!")
	}

	for i := 0; i < typ.NumField(); i++ {

		kind := val.Field(i).Kind()
		name := typ.Field(i).Name
		field := val.Field(i)
		if kind == reflect.String {
			ret[name] = field.String()
		} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
			ret[name] = field.Int()
		} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
			ret[name] = field.Uint()
		} else if kind == reflect.Float32 || kind == reflect.Float64 {
			ret[name] = field.Float()
		} else if kind == reflect.Bool {
			ret[name] = field.Bool()
		}
	}
	return ret, nil
}
/*
//解码到对象
func Unmarshal(client Client, key string, pro interface{}) error {
	val := reflect.ValueOf(pro)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return errors.New("Is not a Struct!")
	}

	for i := 0; i < typ.NumField(); i++ {

		kind := val.Field(i).Kind()
		log.Infoln(typ.Field(i).Name, kind)
		f := val.FieldByName(typ.Field(i).Name)

		if kind == reflect.String {
			str, _ := client.Hget(key, typ.Field(i).Name)
			f.SetString(string(str))
		} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
			str, _ := client.Hget(key, typ.Field(i).Name)
			s, _ := StringToInt(str)
			f.SetInt(s)
		} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
			str, _ := client.Hget(key, typ.Field(i).Name)
			s, _ := StringToUint(str)
			f.SetUint(s)

		} else if kind == reflect.Float32 || kind == reflect.Float64 {
			str, _ := client.Hget(key, typ.Field(i).Name)
			s, _ := StringToFloat(str)
			f.SetFloat(s)
		} else if kind == reflect.Bool {
			str, _ := client.Hget(key, typ.Field(i).Name)
			s, _ := StringToBool(str)
			f.SetBool(s)
		} else {
			log.Infoln("field [%s] not set!%v %v n", typ.Field(i).Name, kind, typ.Field(i).Type)
		}
	}
	return nil
}*/



//字符串到Uint32
func StringToUint(b []byte) (uint64, error) {
	s := string(b)
	i, err := strconv.Atoi(s)
	if err == nil {
		return uint64(i), err
	} else {
		return 0, err
	}
}

//字符串到Int
func StringToInt(b []byte) (int64, error) {
	s := string(b)
	i, err := strconv.Atoi(s)
	if err == nil {
		return int64(i), err
	} else {
		return 0, err
	}
}

//字符串到float
func StringToFloat(b []byte) (float64, error) {
	s := string(b)
	i, err := strconv.Atoi(s)
	if err == nil {
		return float64(i), err
	} else {
		return 0, err
	}
}

//字符串到float
func StringToBool(b []byte) (bool, error) {
	s := string(b)
	i, err := strconv.Atoi(s)
	if err != nil {
		return false, nil
	}
	if i == 0 {
		return false, nil
	} else {
		return true, nil
	}
}