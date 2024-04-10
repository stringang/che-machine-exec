//
// Copyright (c) 2019-2021 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package auth

import (
	"github.com/eclipse-che/che-machine-exec/cfg"
	"github.com/gin-gonic/gin"
)

func IsEnabled() bool {
	return cfg.UseBearerToken
}

func Authenticate(c *gin.Context) (string, error) {
	token, err := extractToken(c)
	if err != nil {
		return "", err
	}

	return token, nil
}
