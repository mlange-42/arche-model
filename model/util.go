package model

import "time"

// nextTime calculates the next intended update.
func nextTime(last time.Time, fps float64) time.Time {
	if fps <= 0 {
		return last
	}
	dt := time.Second / time.Duration(fps)
	now := time.Now()
	if now.After(last.Add(10 * dt)) {
		return now
	}
	return last.Add(dt)
}
