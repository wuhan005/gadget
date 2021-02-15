package gadget

import (
	"fmt"
)

// MakeErrJSON makes the error response JSON for gin.
func MakeErrJSON(errCode int, msg string, a ...interface{}) (int, interface{}) {
	if errCode < 100 {
		errCode = 50000
	}
	// [100 - 999]
	if errCode < 1000 {
		errCode *= 100
	}
	// [1000 - 9999] or [10000 - ]
	if errCode < 10000 || errCode > 99999 {
		errCode = 50000
	}
	return errCode / 100, map[string]interface{}{"error": errCode, "msg": fmt.Sprintf(msg, a...)}
}

// MakeSuccessJSON makes the successful response JSON for gin.
func MakeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}
