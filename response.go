package gadget

import "fmt"

// MakeErrJSON makes the error response JSON for gin.
func MakeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, map[string]interface{}{"error": errCode, "msg": fmt.Sprint(msg)}
}

// MakeSuccessJSON makes the successful response JSON for gin.
func MakeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}
