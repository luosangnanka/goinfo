package goinfo

import (
	"fmt"
	"testing"
)

func TestXminfo(t *testing.T) {
	info := NewXminfo()

	// cpu
	cpu, err := info.Cpu()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cpu {
		fmt.Println(v.String())
	}

	// memory
	free, err := info.Memory()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(free.String())

	// traffic
	net, err := info.Net()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range net {
		fmt.Println(v.String())
	}

	// host
	host, err := info.Host()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(host)

	// disk
	disk, err := info.Disk()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range disk {
		fmt.Println(v.String())
	}

	// load
	load, err := info.Load()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(load.String())

	// tcp
	tcp, err := info.Tcp()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tcp.String())

	// udp
	udp, err := info.Udp()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(udp.String())
}
