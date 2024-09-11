package iot

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Ip       string    "json:ip"
	Commands []Command "json:commands"
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
	if d.Url == "" {
		return errors.New("url is empty")
	}
	if len(d.Commands) == 0 {
		return errors.New("no commands")
	}
	check := false
	for _, v := range d.Commands {
		if v.Name == name {
			check = true
		}
	}
	if !check {
		return errors.New("command not found")
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
	var data string
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

type IOT struct {
	Devices []Device
}
