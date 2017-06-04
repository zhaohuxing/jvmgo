package bubblesort

import "testing"
import "fmt"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 3, 4, 2, 1}
	BubbleSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("输出结果为:", values, "与期望结果{1,2,3,4,5}不符")
	}
	fmt.Println(values)
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 2, 4, 6}
	BubbleSort(values)

	if values[0] != 2 || values[1] != 4 || values[2] != 5 || values[3] != 5 || values[4] != 6 {
		t.Error("输出结果为:", values, "与期望结果{2,4,5,5,6}不符")
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)

	if values[0] != 5 {
		t.Error("输出结果:", values, "与期望结果{5}不符")
	}
}
