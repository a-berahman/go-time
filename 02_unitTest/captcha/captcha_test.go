package captcha

import (
	"image"
	"testing"
)

func TestChallengeUserSuccess(t *testing.T) {
	got := ChallengeUser(stubChallenger("42"), stubPrompter("42"))
	if got != true {
		t.Fatal("exptected challengeuser to return true")
	}
}

func TestChallengeUserFail(t *testing.T) {
	got := ChallengeUser(stubChallenger("Hello"), stubPrompter("42"))
	if got != false {
		t.Fatal("exptected challengeuser to return false")
	}
}

type stubChallenger string
type stubPrompter string

func (c stubChallenger) Challenge() (image.Image, string) {
	return image.NewNRGBA(image.Rect(0, 0, 100, 100)), string(c)
}

func (p stubPrompter) Prompt(_ image.Image) string {
	return string(p)
}
