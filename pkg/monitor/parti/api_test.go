package parti

import "testing"

func TestCheckPublicApi(t *testing.T) {
	userId := "465731"
	isLive, err := CheckPublicApi(userId)
	if err != nil {
		t.Errorf("Error checking public API for user %s: %v", userId, err)
		return
	}
	if !isLive {
		t.Errorf("Expected user %s to be live, but got false", userId)
	} else {
		t.Logf("User %s is live: %v", userId, isLive)
	}
}
