/*
 * Frontend Service
 *
 * Frontend Microservice for the Giò Plants platform.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

import (
	"fmt"
)

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Mac  string `json:"mac"`
	Room string `json:"room"`
}

func (i *Device) Validate() (bool, error) {
	if i.Name == "" {
		return false, fmt.Errorf("invalid name")
	}

	if i.Mac == "" {
		return false, fmt.Errorf("invalid mac")
	}

	if i.Room == "" {
		return false, fmt.Errorf("invalid room")
	}

	return true, nil
}

type Reading struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Value             string `json:"value"`
	Unit              string `json:"unit"`
	CreationTimestamp string `json:"creation_timestamp"`
}
