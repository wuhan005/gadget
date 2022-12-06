// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gadget

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRetry(t *testing.T) {
	attempts := 3

	// Test success
	t.Run("normal", func(t *testing.T) {
		err := Retry(attempts, func() error {
			return nil
		})
		require.Nil(t, err)
	})

	t.Run("retry fail", func(t *testing.T) {
		err := Retry(attempts, func() error {
			// Always return error
			return errors.New("failed")
		})
		require.NotNil(t, err)
	})

	t.Run("success after 2 attempts", func(t *testing.T) {
		var i int
		err := Retry(attempts, func() error {
			i++
			if i == 2 {
				return nil
			}
			return errors.New("failed")
		})
		require.Nil(t, err)
	})
}
