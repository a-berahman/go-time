package drone

import "image"

//Drone ...
type Drone struct{}

//Vector ...
type Vector struct{}

//Target ...
type Target struct{}

// NavigateTo applies any required changes to the drone's speed
// vector so that its eventual position matches dst.
func (d *Drone) NavigateTo(dst Vector) error {
	//...
	return nil
}

// Position returns the current drone position vector.
func (d *Drone) Position() Vector {
	//...
	return Vector{}
}

// Speed returns the current drone speed vector.
func (d *Drone) Speed() Vector {
	//...
	return Vector{}
}

// CaptureImage records and returns an image of the drone's field of
// view using the on-board drone camera.
func (d *Drone) CaptureImage() (*image.RGBA, error) {
	//...
	return nil, nil
}
