package iot

import (
	"fmt"
	"testing"
)

func TestTriggerGetCommandsFromDevice(t *testing.T) {
	var d Device = Device{
		Name: "test",
		Id:   "test",
		Url:  "http://192.168.0.29:81/",
	}
	vaL, err := d.TriggerGetInfoFromDevice()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("19:::: ", d.Url)
	fmt.Println("20:::: ", vaL.ip)
	if d.Url != vaL.ip {
		t.Fatal("url error")
	}
}

func TestTriggerCommand(t *testing.T) {
	var d Device = Device{
		Name: "test",
		Id:   "test",
		Url:  "http://192.168.0.29:81/",
	}
	_, err := d.TriggerGetInfoFromDevice()
	fmt.Println(d.Url)
	if err != nil {
		t.Fatal(err)
	}
	err = d.TriggerCommand("test")
	if err != nil {
		t.Fatal(err)
	}
}
