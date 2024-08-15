package main

import (
	"fmt"
	"math"
	"time"

	"github.com/notaud/gwintils/mouse"
)

const (
	Version = "v0.0.5"
)

func GetVersion() string {
	return Version
}

func main() {
	for {
		currentPosition, err := mouse.GetPosition()
		if err != nil {
			panic(err)
		}

		mouse.Move(currentPosition.X, currentPosition.Y)

		time.Sleep(1 * time.Second)
	}
}

func MoveMouse(targetX, targetY int32, speed uint) error {
	if speed == 0 {
		mouse.Move(targetX, targetY)
		return nil
	}

	const updateInterval = 10 * time.Millisecond
	minDistance := float64(speed) * updateInterval.Seconds()

	for {
		currentPosition, err := mouse.GetPosition()
		if err != nil {
			return err
		}

		dx := float64(targetX) - float64(currentPosition.X)
		dy := float64(targetY) - float64(currentPosition.Y)
		distance := math.Sqrt(dx*dx + dy*dy)

		if distance < 1 {
			// We've reached the target (or close enough)
			mouse.Move(targetX, targetY)

			return nil
		}

		// Calculate the movement for this step
		moveDistance := math.Min(minDistance, distance)
		ratio := moveDistance / distance

		newXFloat := float64(currentPosition.X) + dx*ratio

		var newX int32
		if currentPosition.X < targetX {
			newX = int32(math.Max(math.Ceil(newXFloat), float64(currentPosition.X+1)))
		} else {
			newX = int32(float64(currentPosition.X) + dx*ratio)
		}

		// newY := int32(float64(currentPosition.Y) + dy*ratio)

		fmt.Println(currentPosition.X, newX)
		err = mouse.Move(currentPosition.X, 1000)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
}
