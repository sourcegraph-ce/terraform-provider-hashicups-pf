// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build tools

package tools

import (
	// Ensure documentation generator is not removed from go.mod.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
