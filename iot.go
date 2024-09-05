package iot

import (
	"encoding/json"
	"io"
	"net/http"
)

type Device struct {
	Name     string
	Id       string
	Commands []string
	Url      string
}

func (d *Device) TriggerGetCommandsFromDevice() error {
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
	var data any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	fmt.println(data)
	return nil
}

type IOT struct {
	Devices []Device
}
