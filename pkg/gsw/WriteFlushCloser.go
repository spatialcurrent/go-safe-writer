// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsw

import (
	"io"
)

type WriteFlushCloser interface {
	io.Writer
	io.Closer
	Flusher
}
