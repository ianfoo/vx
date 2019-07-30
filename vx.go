package vx

import "regexp"

var (
	regexProceed = regexp.MustCompile(`(?i:^\s*(kk?|yeh|ya+s)\s*$)`)
	regexCancel  = regexp.MustCompile(`(?i:^\s*(na[hw](?:\s*dawg)?|w/e|ur\s+face|gtfo|no\s+u)\s*$)`)
)

// Proceed indicates whether an operation should proceed given a response.
// The first return value indicates whether a response indicates whether to
// proceed. The second return value is false if the response is not included
// in the set of the acceptable responses, in which case the first return
// value should be ignored, unless you roll like that.
func Proceed(resp string) (proceed, valid bool) {
	if regexProceed.MatchString(resp) {
		return true, true
	}
	if regexCancel.MatchString(resp) {
		return false, true
	}
	return false, false
}
