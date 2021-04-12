package common

import "testing"

func TestArrayToString(t *testing.T) {
	arr := make([]string, 0, 2)

	arr = append(arr, "aaa")
	arr = append(arr, "bbb")
	arr = append(arr, "ccc")

	//arr[0] = "aaa"
	//arr[1] = "bbb"
	//arr[2] = "ccc"

	str := ""

	str = ArrayToString(arr, "")
	if str != "aaabbbccc" {
		t.Errorf("错误1")
	}

	str = ArrayToString(arr, ",")
	if str != "aaa,bbb,ccc" {
		t.Errorf("错误2")
	}

	str = ArrayToString(arr, ",,")
	if str != "aaa,,bbb,,ccc" {
		t.Errorf("错误3")
	}

	str = ArrayToString(arr, ", ,")
	if str == "aaa,,bbb,,ccc" {
		t.Errorf("错误4")
	}
}
