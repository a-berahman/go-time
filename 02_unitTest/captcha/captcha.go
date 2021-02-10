package captcha

import (
	"image"
)

//Challenger is implemented by objects that can generate CAPTCHA image challenges
type Challenger interface {
	Challenge() (img image.Image, imgText string)
}

//Prompter is implemented by objects that display a CAPTCHA image to the user,
//ask them to type their contents and return back their response
type Prompter interface {
	Prompt(img image.Image) string
}

// ChallengeUser requests a challenge from c and prompts the user for an answer
// using p. If the user's answer matches the challenge then ChallengeUser
// returns true.
func ChallengeUser(c Challenger, p Prompter) bool {
	img, expAnswer := c.Challenge()
	userAnswer := p.Prompt(img)

	if expAnswer == userAnswer {
		return true
	}
	return false
}
