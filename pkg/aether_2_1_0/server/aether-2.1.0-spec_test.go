// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"gotest.tools/assert"
	"testing"
)

func Test_Aether210Spec(t *testing.T) {
	swagger, err := GetSwagger()
	assert.NilError(t, err)
	assert.Assert(t, swagger != nil)
}
