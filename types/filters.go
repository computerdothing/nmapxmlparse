package types

import (
	"nmapxmlparse/util"
)

type HostFilter interface {
	Pass(*Host) bool
}

func (nmr NmapRun) Filter(filters ...HostFilter) error {
	var finalForm []Host

	for _, h := range nmr.Hosts {
		good := true
		for _, f := range filters {
			if !f.Pass(&h) {
				good = false
				break
			}
		}
		if good {
			finalForm = append(finalForm, h)
		}
	}

	nmr.Hosts = finalForm

	return nil
}

//// ANY OPEN PORT ////
type anyOpenPort struct{}

func (hop anyOpenPort) Pass(h *Host) bool {
	for _, p := range h.Ports {
		if p.State.State == "open" {
			return true
		}
	}
	return false
}

func AnyOpenPort() HostFilter {
	return &anyOpenPort{}
}

///////////////////////

//// SELECT OPEN PORTS ////
type selectOpenPorts struct {
	portIds []int
}

func (sop selectOpenPorts) Pass(h *Host) bool {
	for _, p := range h.Ports {
		if util.ContainsInt(sop.portIds, p.Portid) {
			if p.State.State == "open" {
				return true
			}
		}
	}
	return false
}

func SelectOpenPorts(ports ...int) HostFilter {
	return &selectOpenPorts{ports}
}

///////////////////////////
