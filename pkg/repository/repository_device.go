/*
 * Frontend Service
 *
 * Frontend Microservice for the Giò Plants platform.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gio-frontend-ms/pkg/model"
	"gio-frontend-ms/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type DeviceRepository struct {
	devicesServiceUrl string
}

// Returns a device with a specific id
func (r *DeviceRepository) Get(roomId string, id string) (*model.Device, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices/%s", r.devicesServiceUrl, roomId, id)

	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting data for device %s", id)
	}

	var d model.Device
	err = json.NewDecoder(resp.Body).Decode(&d)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

// Returns all readings of a device.
// limit: limits the result set
// name: filters the data by name
func (r *DeviceRepository) GetReadings(roomId string, id string, limit int, name string) ([]model.Reading, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices/%s/readings", r.devicesServiceUrl, roomId, id)

	var queryParams []string
	if limit > 0 {
		queryParams = append(queryParams, fmt.Sprintf("limit=%d", limit))
	}

	if name != "" {
		queryParams = append(queryParams, fmt.Sprintf("name=%s", name))
	}

	// Add query params if needed
	if len(queryParams) > 0 {
		u = u + "?"

		u = u + strings.Join(queryParams, "&")
	}

	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting data for device %s", id)
	}

	var rs []model.Reading
	err = json.NewDecoder(resp.Body).Decode(&rs)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

// Returns all devices in a room
func (r *DeviceRepository) GetAll(roomId string) ([]model.Device, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices", r.devicesServiceUrl, roomId)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting data for devices")
	}

	var d []model.Device
	err = json.NewDecoder(resp.Body).Decode(&d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

// Transform an ActionData object to a buffer of bytes
func getActionBodyData(data *model.ActionData) *bytes.Buffer {
	if data != nil {
		body, err := json.Marshal(data)
		if err == nil {
			return bytes.NewBuffer(body)
		}
	}

	return nil
}

// Triggers an action on a device
func (r *DeviceRepository) TriggerAction(roomId string, deviceId string, actionName string, actionData *model.ActionData) error {
	u := fmt.Sprintf("%s/rooms/%s/devices/%s/actions/%s", r.devicesServiceUrl, roomId, deviceId, actionName)

	bodyData := getActionBodyData(actionData)
	resp, err := utils.DoPost(u, bodyData)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error while performing the operation: %d - %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	return nil
}

var devicesRepository *DeviceRepository

func NewDeviceRepository() (*DeviceRepository, error) {
	serviceHost := os.Getenv("API_GATEWAY_HOST")
	servicePort := os.Getenv("API_GATEWAY_PORT")
	if devicesRepository == nil {
		u := fmt.Sprintf("http://%s:%s", serviceHost, servicePort)
		log.Printf("ApiGateway URL: %s\n", u)

		serviceUrl, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		devicesRepository = &DeviceRepository{serviceUrl.String()}
	}

	return devicesRepository, nil
}
