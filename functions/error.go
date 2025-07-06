package functions

import "errors"

func Checknumber(n int) (string, error) {
	if n < 0 {
		return "this  is  nig number", errors.New("please provide the postive number")
	}
return  "this is a postive number",nil
}