package mobileNet

import "github.com/a-berahman/go-time/01_solid/01_single-responsibility/drone"

// MobileNet performs target detection for drones using the
// SSD MobileNet V1 NN.
// For more info on this model see:
// https://github.com/tensorflow/models/tree/master/research/object_detection
type MobileNet struct {
	//...
}

// DetectTargets captures an image of the drone's field of view and feeds
// it to a neural network to detect and classify interesting nearby
// targets.
func (m *MobileNet) DetectTargets(d *drone.Drone) ([]*drone.Target, error) {
	//...
	return nil, nil
}
