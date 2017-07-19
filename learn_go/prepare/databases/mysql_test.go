package databases

import "testing"
import "fmt"

func TestInitDB(t *testing.T) {
	db := Open()
	fmt.Println(db)
}
