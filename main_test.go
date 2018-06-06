package main_test

import (
	"testing"
	"strconv"
	"reflect"
)


func TestTypeOfValues(t *testing.T) {

	getDays := "21"
	daysOldValue, _ := strconv.Atoi(getDays)
	kindInt := reflect.Int

	if reflect.TypeOf(daysOldValue).Kind() != kindInt {
		t.Logf("%T needs to be of type %T", daysOldValue, kindInt)
		t.Fail()
	}

}
