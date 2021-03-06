package main

import (
	"time"

	"barista.run"
	"barista.run/modules/battery"
	"barista.run/modules/bluetooth"
	"barista.run/modules/clock"
	//"barista.run/modules/funcs"
	"barista.run/modules/netinfo"
	"barista.run/modules/volume"
	"barista.run/modules/volume/alsa"
	"barista.run/modules/wlan"

	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/gbanegas/i3-bar/blocks"
	"github.com/gbanegas/i3-bar/xbacklight"
	//"github.com/gbanegas/i3-bar/yubikey"
)

func main() {
	//yu := yubikey.New().Output(blocks.Yubi)
	ly := kbdlayout.New().Output(blocks.Layout)
	br := xbacklight.New().Output(blocks.Brightness)
	snd := volume.New(alsa.DefaultMixer()).Output(blocks.Snd)
	bat := battery.All().Output(blocks.Bat)
	wi := wlan.Named("wifi").Output(blocks.WLAN)
	nt := netinfo.New().Output(blocks.Net)
	ti := clock.Local().Output(time.Second, blocks.Clock)

	adapter, mac := "hci0", "08:71:90:14:99:D8"//, "bluez_sink.09_A5_C1_A6_5C_77.headset_head_unit"
	//snd2 := volume.Sink(_).Output(blocks.Snd2)
	blD := bluetooth.Device(adapter, mac).Output(blocks.Blue)
	//bl := funcs.Every(time.Second, blocks.Bluetooth)

	panic(barista.Run(
		ly, blD, br, snd, bat, wi, nt, ti,
	))

}
