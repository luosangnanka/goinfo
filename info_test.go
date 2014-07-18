package goinfo

import (
	"fmt"
	"testing"
)

func TestXminfo(t *testing.T) {
	info := NewXminfo()

	// cpu
	cpu, err := info.CPU()
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
	tcp, err := info.TCP()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tcp.String())

	// udp
	udp, err := info.UDP()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(udp.String())
}

func TestInfo(t *testing.T) {
	// cpu.
	cpuInfo, err := CPU()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("cpuinfo", cpuInfo)

	// disk.
	diskInfo, err := Disk()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("disk", diskInfo)

	// host.
	host, err := Host()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("host", host.String())

	// load.
	load, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("load", load)

	// mem.
	mem, err := Memory()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("mem", mem)

	// net.
	net, err := Net()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("net", net)

	// tcp.
	tcp, err := TCP()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tcp", tcp)

	// udp.
	udp, err := UDP()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("udp", udp)
}
