package errortools

import "fmt"

func stringifyError(e error) string{
	return fmt.Sprintf("%+v", e)
}
