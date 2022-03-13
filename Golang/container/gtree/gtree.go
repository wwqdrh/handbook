// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gtree provides concurrent-safe/unsafe tree containers.
//
// Some implements are from: https://github.com/emirpasic/gods
package gtree

import (
	"fmt"
	"strings"
)

func ComparatorString(a, b interface{}) int {
	return strings.Compare(fmt.Sprint(a), fmt.Sprint(b))
}

func Map(val interface{}) map[string]interface{} {
	return val.(map[string]interface{})
}
