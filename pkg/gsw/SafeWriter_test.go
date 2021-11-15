// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsw

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeWriter(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	sw := NewSafeWriter(buf)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			_, _ = sw.Write([]byte(fmt.Sprintf("%d\n", i)))
			wg.Done()
		}(i)
	}
	wg.Wait()
	lines := strings.Split(buf.String(), "\n")
	counts := map[string]int{}
	for _, line := range lines {
		if _, ok := counts[line]; !ok {
			counts[line] = 0
		}
		counts[line] += 1
	}
	for i := 1; i <= 10; i++ {
		assert.Equalf(
			t,
			1,
			counts[strconv.Itoa(i)],
			"expected %d for %d, but found %d",
			1,
			i,
			counts[strconv.Itoa(i)],
		)
	}
}
