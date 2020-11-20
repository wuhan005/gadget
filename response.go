package gadget

import "fmt"

// MakeErrJSON makes the error response JSON for gin.
func MakeErrJSON(errCode int, msg string, a ...interface{}) (int, interface{}) {
	return errCode / 100, map[string]interface{}{"error": errCode, "msg": fmt.Sprintf(msg, a...)}
}

// MakeSuccessJSON makes the successful response JSON for gin.
func MakeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}
