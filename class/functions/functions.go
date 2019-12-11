package functions

import (
	"fmt"
	"regexp"
)

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
func Strip_tags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content,"")
}
func Validate_data(content string) string {
	return Strip_tags(content)
}
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}