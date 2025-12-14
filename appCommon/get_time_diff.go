package appCommon

import "time"

func CalcTimeDiff(a time.Time, b time.Time) time.Duration {
	diff := b.Sub(a)
	return diff
}
