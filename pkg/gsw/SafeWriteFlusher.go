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

type SafeWriteFlusher struct {
	mutex  *sync.Mutex
	writer WriteFlusher
}

func (w *SafeWriteFlusher) Lock() {
	w.mutex.Lock()
}

func (w *SafeWriteFlusher) Unlock() {
	w.mutex.Unlock()
}

func (w *SafeWriteFlusher) Write(p []byte) (n int, err error) {

	if w.writer != nil {
		w.Lock()
		n, err := w.writer.Write(p)
		w.Unlock()
		return n, err
	}

	return 0, nil
}

func (w *SafeWriteFlusher) WriteUnsafe(p []byte) (n int, err error) {

	if w.writer != nil {
		return w.writer.Write(p)
	}

	return 0, nil
}

func (w *SafeWriteFlusher) Flush() error {

	if w.writer != nil {
		w.Lock()
		err := w.writer.Flush()
		w.Unlock()
		return err
	}

	return nil
}

func (w *SafeWriteFlusher) FlushUnsafe() error {

	if w.writer != nil {
		return w.writer.Flush()
	}

	return nil
}

func NewSafeWriteFlusher(w WriteFlusher) *SafeWriteFlusher {
	return &SafeWriteFlusher{
		mutex:  &sync.Mutex{},
		writer: w,
	}
}
