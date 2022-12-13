package main

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}

	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
