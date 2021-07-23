// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gadget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONEqual(t *testing.T) {
	for _, tc := range []struct {
		name         string
		json1, json2 []byte
		want         bool
	}{
		{
			name:  "empty json",
			json1: []byte(""),
			json2: []byte(""),
			want:  true,
		},
		{
			name:  "empty object",
			json1: []byte("{}"),
			json2: []byte("{}"),
			want:  true,
		},
		{
			name:  "normal json",
			json1: []byte(`{"Hello": "E99p1ant", "Waifu": "Elaina"}`),
			json2: []byte(`{"Waifu":"Elaina","Hello":"E99p1ant"}`),
			want:  true,
		},
		{
			name:  "not equal",
			json1: []byte(`{"Hello": "E99p1ant", "Waifu": "Elaina"}`),
			json2: []byte(`{"Hello": "E99p1ant", "Waifu": "Mashiro"}`),
			want:  false,
		},
		{
			name:  "bad format",
			json1: []byte(`{"Hello": ""`),
			json2: []byte(`{Hello":"E99p1ant"}`),
			want:  false,
		},
		{
			name:  "bad format equal",
			json1: []byte(`{"Hello": "", 123`),
			json2: []byte(`{"Hello": "", 123`),
			want:  true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := JSONEqual(tc.json1, tc.json2)
			assert.Equal(t, tc.want, got)
		})
	}
}
