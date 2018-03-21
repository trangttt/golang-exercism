package flatten

import (
    //"reflect"
    //"fmt"
)

func Flatten(input []interface{}) []interface{} {
    ret := []interface{}{}
    for _, v := range input {
        nv, ok := v.([]interface{})
        if ok {
            ret = append(ret, Flatten(nv)...)
        } else if v != nil {
            ret = append(ret, v)
        }
        //rt := reflect.TypeOf(v)
        //kind := rt.Kind()
        //fmt.Println(kind)
        //if kind == reflect.Array {
            //nv := v[0:len(v)]
        //} else
        //if kind == reflect.Slice {
            //nv, ok := v.([]interface{})
            //if ok {
                //ret = append(ret, Flatten(nv)...)
            //}
        //} else {
            //if v != nil {
                //ret = append(ret, v)
            //}
        //}
    }
    return ret
}

