// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mailx

import "net/mail"

// NewAddress returns a mail address with the given address and optional name.
func NewAddress(address string, a ...string) *mail.Address {
	addr := &mail.Address{
		Address: address,
	}
	if len(a) > 0 {
		addr.Name = a[0]
	}
	return addr
}
