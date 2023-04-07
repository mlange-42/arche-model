package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextTime(t *testing.T) {
	start := time.Now()

	assert.Equal(t, start.Add(time.Second), nextTime(start, 1))
	assert.Equal(t, start.Add(time.Second/2), nextTime(start, 2))
	assert.Equal(t, start.Add(time.Second/60), nextTime(start, 60))
}
