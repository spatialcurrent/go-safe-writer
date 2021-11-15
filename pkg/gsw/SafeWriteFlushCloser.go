// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsw

import (
	"sync"
)

type SafeWriteFlushCloser struct {
	mutex  *sync.Mutex
	writer WriteFlushCloser
}

func (w *SafeWriteFlushCloser) Lock() {
	w.mutex.Lock()
}

func (w *SafeWriteFlushCloser) Unlock() {
	w.mutex.Unlock()
}

func (w *SafeWriteFlushCloser) Write(p []byte) (n int, err error) {

	if w.writer != nil {
		w.Lock()
		n, err := w.writer.Write(p)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

func (w *SafeWriteFlushCloser) WriteUnsafe(p []byte) (n int, err error) {

	if w.writer != nil {
		return w.writer.Write(p)
	}

	return 0, nil
}

func (w *SafeWriteFlushCloser) Flush() error {

	if w.writer != nil {
		w.Lock()
		err := w.writer.Flush()
		w.Unlock()
		return err
	}

	return nil
}

func (w *SafeWriteFlushCloser) FlushUnsafe() error {

	if w.writer != nil {
		return w.writer.Flush()
	}

	return nil
}

func (w *SafeWriteFlushCloser) Close() error {

	if w.writer != nil {
		w.Lock()
		err := w.writer.Close()
		w.Unlock()
		return err
	}

	return nil
}

func (w *SafeWriteFlushCloser) CloseUnsafe() error {

	if w.writer != nil {
		return w.writer.Close()
	}

	return nil
}

func NewSafeWriteFlushCloser(w WriteFlushCloser) *SafeWriteFlushCloser {
	return &SafeWriteFlushCloser{
		mutex:  &sync.Mutex{},
		writer: w,
	}
}
