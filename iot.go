package iot

import (
	"encoding/json"
	"io"
	"net/http"
)

type Command struct {
	Name string "json:name"
	Info string "json:info"
}
type DeviceInfo struct {
	Info     string    "json:info"
	Type     string    "json:type"
	ip       string    "json:ip"
	Commands []Command "json:commands"
}
type Device struct {
	Name     string
	Id       string
	Commands []string
	Url      string
}

func (d *Device) TriggerGetInfoFromDevice() error {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", d.Url, nil)
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
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var data DeviceInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	return nil
}

type IOT struct {
	Devices []Device
}
