package qsort

import "testing"
import "fmt"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 5, 2, 4, 6}
	QuickSort(values)
	if values[0] != 2 || values[1] != 4 || values[2] != 5 || values[3] != 5 || values[4] != 6 {
		t.Error("输出结果:", values, "与期望结果:{2, 4, 5, 5, 6}不符")
	}
	fmt.Println("TestQuickSort1测试输出数据:", values)
}

func TestQuickSort2(t *testing.T) {
	values := []int{6, 5, 4, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 ||
		values[2] != 3 || values[3] != 4 ||
		values[4] != 5 || values[5] != 6 {
		t.Error("输出结果:", values, "与期望结果:{1, 2, 3, 4, 5, 6}不符")
	}
	fmt.Println("TestQuickSort2测试输出数据:", values)
}

func TestQuickSort3(t *testing.T) {
	values := []int{6, 8, 9, 1, 34, 56, 78, -5}
	QuickSort(values)

	if values[0] != -5 || values[1] != 1 || values[2] != 6 ||
		values[3] != 8 || values[4] != 9 || values[5] != 34 ||
		values[6] != 56 || values[7] != 78 {
		t.Error("输出结果:", values, "与期望结果:{-5, 1, 6, 8, 9, 34, 56, 78}不符")
	}
	fmt.Println("TestQuickSort3测试输出数据:", values)
}
