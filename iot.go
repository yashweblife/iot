package iot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Name string `json:"name"`
	Info string `json:"info"`
}
type DeviceInfo struct {
	Info     string    `json:"info"`
	Type     string    `json:"type"`
	Ip       string    `json:"ip"`
	Commands []Command `json:"commands"`
}
type Device struct {
	Name     string
	Id       string
	Commands []Command
	Url      string
}

func (d *Device) TriggerGetInfoFromDevice() (DeviceInfo, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", d.Url, nil)
	if err != nil {
		return DeviceInfo{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return DeviceInfo{}, err
	}
	if res.StatusCode != http.StatusOK {
		return DeviceInfo{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return DeviceInfo{}, err
	}
	var data DeviceInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return DeviceInfo{}, err
	}
	d.Commands = data.Commands
	return data, nil
}
func (d *Device) TriggerCommand(name string) error {
	fmt.Println("Triggering command: ", name)
	if d.Url == "" {
		return errors.New("url is empty")
	}
	client := http.DefaultClient
	req, err := http.NewRequest("POST", d.Url+name, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return err
	}
	type Outcome struct {
		Data string `json:"data"`
	}
	var jsonRes Outcome
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &jsonRes)
	fmt.Print("\n\n\n\n", jsonRes.Data, "\n\n\n\n")
	if err != nil {
		return err
	}
	return nil
}

type IOT struct {
	Devices []Device
}
