package defer_statement

import (
	"fmt"
)
 
func Demo2(number int32) string {
	defer DeferFunc()
	
	if number%2 == 0 {
}
