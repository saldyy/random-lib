package main

import (
  "reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fnV.Set(fnI)
}

// TODO:completes implMap function.
func implMap(inpt []reflect.Value) []reflect.Value {
	switch inpt[1].Kind() {
	case reflect.Slice:
		{
			length := inpt[1].Len()

			for i := 0; i < length; i++ {
				elem := inpt[1].Index(i)
				newElem := inpt[0].Call([]reflect.Value{elem})[0]
				elem.Set(newElem)
			}
      break
		}
	case reflect.Map:
		{
      for _, key := range inpt[1].MapKeys(){
				elem := inpt[1].MapIndex(key)
				newElem := inpt[0].Call([]reflect.Value{elem})[0]
				inpt[1].SetMapIndex(key, newElem)
			}

      break
		}
	}

	return []reflect.Value{inpt[1]}
}

