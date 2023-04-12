package model

import "time"

// nextTime calculates the next intended update.
func nextTime(last time.Time, fps float64) time.Time {
	if fps <= 0 {
		return last
	}
	dt := time.Second / time.Duration(fps)
	now := time.Now()
	if now.After(last.Add(200 * time.Millisecond)) {
		return now.Add(-10 * time.Millisecond)
	}
	return last.Add(dt)
}
