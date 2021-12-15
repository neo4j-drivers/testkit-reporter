package entity

import (
	"fmt"
)

type SkippedTest struct {
	Method string
	Class  string
	Reason string

	FeatureFlags []string
}

func (s SkippedTest) String() string {
	return fmt.Sprintf("%s %s.%s: %s", s.FeatureFlags, s.Class, s.Method, s.Reason)
}
