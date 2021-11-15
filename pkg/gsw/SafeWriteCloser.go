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

type SafeWriteCloser struct {
	mutex  *sync.Mutex
	writer io.WriteCloser
}

func (w *SafeWriteCloser) Lock() {
	w.mutex.Lock()
}

func (w *SafeWriteCloser) Unlock() {
	w.mutex.Unlock()
}

func (w *SafeWriteCloser) Write(p []byte) (n int, err error) {

	if w.writer != nil {
		w.Lock()
		n, err := w.writer.Write(p)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

func (w *SafeWriteCloser) WriteUnsafe(p []byte) (n int, err error) {

	if w.writer != nil {
		return w.writer.Write(p)
	}

	return 0, nil
}

func (w *SafeWriteCloser) Close() error {

	if w.writer != nil {
		w.Lock()
		err := w.writer.Close()
		w.Unlock()
		return err
	}

	return nil
}

func (w *SafeWriteCloser) CloseUnsafe() error {

	if w.writer != nil {
		return w.writer.Close()
	}

	return nil
}

func NewSafeWriteCloser(w io.WriteCloser) *SafeWriteCloser {
	return &SafeWriteCloser{
		mutex:  &sync.Mutex{},
		writer: w,
	}
}
