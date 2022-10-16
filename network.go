package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"syscall"
)

const (
	DefaultMTU = 1500
)

func createBridge(name string) (*netlink.Bridge, error) {
	br := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{
			Name:   name,
			MTU:    DefaultMTU,
			TxQLen: -1,
		},
	}

	err := netlink.LinkAdd(br)
	if err != nil && err != syscall.EEXIST {
		return nil, err
	}

	l, err := netlink.LinkByName(name)
	if err != nil {
		return nil, fmt.Errorf("could not lookup %q: %v", name, err)
	}
	newBr, ok := l.(*netlink.Bridge)
	if !ok {
		return nil, fmt.Errorf("%q already exists but is not a bridge", name)
	}

	if err := netlink.LinkSetUp(br); err != nil {
		return nil, err
	}

	return newBr, nil
}
