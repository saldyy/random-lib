package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.
func Reverse (slice interface{}) {
  value := reflect.ValueOf(slice)
  fmt.Println(value.Elem().Len())
  i, j := 0, value.Elem().Len() - 1

  for ;i<=j; {
    left := value.Elem().Index(i)
    right := value.Elem().Index(j)

    tmp := reflect.New(left.Type()).Elem()
    tmp.Set(left)
    left.Set(right)
    right.Set(tmp)

    i++
    j--
  }
}

