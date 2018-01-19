package TestGoUnitTestCases

import (
	"testing"
)

func TestStringInSlice(t *testing.T)  {
	if ok := StringInSlice("a",[]string{"a","b"});ok {
		t.Log("pass 1'")
	}else {
		t.Errorf("failed")
	}
}

func TestFailedStringInSlice(t *testing.T) {
	if ok := StringInSlice("c", []string{"a", "b"}); ok {
		t.Errorf("failed")
	} else {
		t.Log("pass 2")
	}
}

func TestSuccessStringInSlice(t *testing.T)  {
	type args struct {
		a string
		list []string
	}

/*	tests := []struct{
		name string
		args args
		want bool
	}{

		
	}*/


}