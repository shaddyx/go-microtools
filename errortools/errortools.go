package errortools

import "fmt"

func StringifyError(e error) string{
	return fmt.Sprintf("%+v", e)
}
