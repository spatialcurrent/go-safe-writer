// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsw

import (
	"io"
	"sync"
)

type SafeWriter struct {
	mutex  *sync.Mutex
	writer io.Writer
}

func (w *SafeWriter) Lock() {
	w.mutex.Lock()
}

func (w *SafeWriter) Unlock() {
	w.mutex.Unlock()
}

func (w *SafeWriter) Write(p []byte) (n int, err error) {

	if w.writer != nil {
		w.Lock()
		n, err := w.writer.Write(p)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

func (w *SafeWriter) WriteUnsafe(p []byte) (n int, err error) {

	if w.writer != nil {
		return w.writer.Write(p)
	}

	return 0, nil
}

func NewSafeWriter(w io.Writer) *SafeWriter {
	return &SafeWriter{
		mutex:  &sync.Mutex{},
		writer: w,
	}
}
