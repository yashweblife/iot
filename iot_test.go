package iot

import (
	"testing"
)

func TestTriggerGetCommandsFromDevice(t *testing.T) {
	var d Device = Device{
		Name: "test",
		Id:   "test",
		Url:  "http://192.168.0.29:81/",
	}
	val, err := d.TriggerGetInfoFromDevice()
	if err != nil {
		t.Fatal(err)
	}
	if d.Url != val.Ip {
		t.Fatal("ip not equal")
	}
}

func TestTriggerCommand(t *testing.T) {

	var d Device = Device{
		Name: "test",
		Id:   "test",
		Url:  "http://192.168.0.29:81/",
	}

	val, err := d.TriggerGetInfoFromDevice()

	if err != nil {
		t.Fatal(err)
	}

	var command = val.Commands[0].Name
	err = d.TriggerCommand(command)

	if err != nil {
		t.Fatal(err)
	}
}
