package math 

import "testing"

func TestAddElemnt(t *testing.T){
	cases := []struct {
		a , b , want int
	}{
		{1,2,3},
		{4,5,9},
		{0,0,0},
	}

	for _ , c := range cases{
		got:= Add(c.a,c.b)
		want := c.want 
		if got!= want{
			t.Errorf("Add(%d,%d) = %d; want %d", c.a, c.b, got, c.want)
		}
	}

}

