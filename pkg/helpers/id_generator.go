package helpers

import (
	"fmt"
)
func GenerateProductOrOrderID(prefix string,currentCount int) string {
	totalDigits := 8
	
	format := fmt.Sprintf("%s-%%0%dd", prefix, totalDigits)
	generatedID := fmt.Sprintf(format, currentCount)
	return generatedID
}



