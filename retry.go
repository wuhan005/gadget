// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gadget

// Retry will retry a function a specified number of times before giving up.
// The function retries on any error.
func Retry(attempts int, f func() error) error {
	for i := 0; ; i++ {
		err := f()
		if err == nil {
			return nil
		}
		if i >= (attempts - 1) {
			return err
		}
	}
}
