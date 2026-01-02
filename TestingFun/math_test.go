package math

import "testing"

func TestTwoNumer(t *testing.T){
	get := Add(5,6)
	want := 11 
	if get!=want{
		t.Errorf("Add(2,3) = %d ; want %d ",get,want)
	}
}