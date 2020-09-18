// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mailx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	mail := "foo@example.com"
	addr := NewAddress(mail)
	assert.Equal(t, mail, addr.Address)
	assert.Equal(t, "", addr.Name)

	name := "Foo"
	addr = NewAddress(mail, name)
	assert.Equal(t, mail, addr.Address)
	assert.Equal(t, name, addr.Name)

}
