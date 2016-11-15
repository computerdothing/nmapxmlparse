package types

import (
  //"net"
  "encoding/xml"
)
//// SCANRESULTS ////
type NmapRun struct {
  Scanner string `xml:"scanner,attr"`
  Args string `xml:"args,attr"`
  Start int `xml:"start,attr"`
  StartStr string `xml:"startstr,attr"`
  Version string `xml:"version,attr"`
  XmlOutputVersion string `xml:"outputversion,attr"`
  ScanInfo NmapScanInfo `xml:"nmaprun>scaninfo"`
  Verbose VerbosityLevel `xml:"nmaprun>verbose"`
  Debugging DebuggingLevel `xml:"nmaprun>debugging"`
  Hosts []Host `xml:"nmaprun>host"`
}
/////////////////////

//// NMAPSCANINFO ////
type NmapScanInfo struct {
  ScanType string `xml:"type,attr"`
  Protocol string `xml:"protocol,attr"`
  NumServices int `xml:"numservices,attr"`
  Services string `xml:"services,attr"`
}
//////////////////

//// VERBOSITYLEVEL ////
type VerbosityLevel struct {
  Level int `xml:"level,attr"`
}
////////////////////////

//// DEBUGGINGLEVEL ////
type DebuggingLevel struct {
  Level int `xml:"level,attr"`
}
////////////////////////

//// HOST ////
type Host struct {
  Starttime int `xml:"starttime,attr"`
  Endtime int `xml:"endtime,attr"`
  Status Status `xml:"status"`
  Address Address `xml:"address"`
  Hostnames []Hostname `xml:"hostnames>hostname"`
  Ports []Port `xml:"ports>port"`
  Times Times `xml:"times"`
}
//////////////

//// STATUS ////
/*
type HostState int

const (
  HostDown HostState = 0
  HostUp HostState = 1
)

type HostStateReason int

const (
  EchoReply HostStateReason = 0
)
*/

type Status struct {
  //State HostState `xml:"state,attr"`
  State string `xml:"state,attr"`
  //Reason HostStateReason `xml:"reason,attr"`
  Reason string `xml:"reason,attr"`
  ReasonTtl int `xml:"reason_ttl,attr"`
}
////////////////

//// ADDRESS ////
/*
type AddressType int

const (
  TypeIPV4 AddressType = 0
  TypeIPV6 AddressType = 1
)
*/

type Address struct {
  //Addr net.IP `xml:"addr,attr"`
  Addr string `xml:"addr,attr"`
  //Addrtype AddressType `xml:"addrtype,attr"`
  Addrtype string `xml:"addrtype,attr"`
}
/////////////////

//// HOSTNAMES ////
/*
type HostnameType int

const (
  PTR HostnameType = 0
)
*/

type Hostname struct {
  XMLName xml.Name `xml:"hostname"`
  Name string `xml:"name,attr"`
  //HNType HostnameType `xml:"type,attr"`
  HNType string `xml:"type,attr"`
}
///////////////////

//// PORTS ////
/*
type Protocol int

const (
    Tcp Protocol = 0
    Udp Protocol = 1
)

type PortStateState int

const (
  PortOpen PortStateState = 0
  PortClosed PortStateState = 1
  PortFiltered PortStateState = 2
)

type PortStateReason int

const (
  SynAck PortStateReason = 0
  NoResponse PortStateReason = 1
)
*/

type PortState struct {
  XMLName xml.Name `xml:"state"`
  //State PortStateState  `xml:"state,attr"`
  State string `xml:"state,attr"`
  //Reason PortStateReason `xml:"reason,attr"`
  Reason string `xml:"reason,attr"`
  ReasonTtl int `xml:"reason_ttl,attr"`
}

/*
type ServiceMethod int

const (
  Probed ServiceMethod = 0
)
*/

type Service struct {
  XMLName xml.Name `xml:"service"`
  Name string `xml:"name,attr"`
  ServiceFP string `xml:"servicefp,attr"`
  Product string `xml:"product,attr"`
  Version string `xml:"version,attr"`
  //Method ServiceMethod `xml:"method,attr"`
  Method string `xml:"method,attr"`
  Conf int `xml:"conf,attr"`
}

type Port struct {
  XMLName xml.Name `xml:"port"`
  Protocol string `xml:"protocol,attr"`
  Portid int `xml:"portid,attr"`
  State PortState `xml:"state"`
  Service Service `xml:"service"`
  Cpe string `xml:"cpe"`
}

/*
type XtraPorts struct {
  XMLName xml.Name `xml:"extraports"`
  //State PortStateState `xml:"state,attr"`
  State string `xml:"state,attr"`
  Count int `xml:"count,attr"`
}

type Ports struct {
  XMLName xml.Name `xml:"ports"`
  ExtraPorts XtraPorts `xml:"extraports"`
  Ports []Port `xml:"ports>port"`
}
*/

///////////////

//// TIMES ////
type Times struct {
  XMLName xml.Name `xml:"times"`
  Srtt int `xml:"srtt,attr"`
  Rttvar int `xml:"rttvar,attr"`
  To int `xml:"to,attr"`
}
///////////////
