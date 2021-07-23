// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gadget

import (
	"encoding/json"
	"reflect"
)

// JSONEqual determines the two given JSON are equal or not.
// If the input is not JSON format, then make the input as string.
func JSONEqual(j1, j2 []byte) bool {
	var object1, object2 interface{}

	if err := json.Unmarshal(j1, &object1); err != nil {
		return string(j1) == string(j2)
	}
	if err := json.Unmarshal(j2, &object2); err != nil {
		return string(j1) == string(j2)
	}

	return reflect.DeepEqual(object1, object2)
}
