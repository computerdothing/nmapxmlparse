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
  ScanInfo NmapScanInfo `xml:"scaninfo"`
  Verbose VerbosityLevel `xml:"verbose"`
  Debugging DebuggingLevel `xml:"debugging"`
  Hosts []Host `xml:"host"`
}
/////////////////////

//// NMAPSCANINFO ////
type NmapScanInfo struct {
  XMLName xml.Name `xml:"scaninfo"`
  ScanType string `xml:"type,attr"`
  Protocol string `xml:"protocol,attr"`
  NumServices int `xml:"numservices,attr"`
  Services string `xml:"services,attr"`
}
//////////////////

//// VERBOSITYLEVEL ////
type VerbosityLevel struct {
  XMLName xml.Name `xml:"verbose"`
  Level int `xml:"level,attr"`
}
////////////////////////

//// DEBUGGINGLEVEL ////
type DebuggingLevel struct {
  XMLName xml.Name `xml:"debugging"`
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
type Status struct {
  State string `xml:"state,attr"`
  Reason string `xml:"reason,attr"`
  ReasonTtl int `xml:"reason_ttl,attr"`
}
////////////////

//// ADDRESS ////
type Address struct {
  Addr string `xml:"addr,attr"`
  Addrtype string `xml:"addrtype,attr"`
}
/////////////////

//// HOSTNAMES ////
type Hostname struct {
  XMLName xml.Name `xml:"hostname"`
  Name string `xml:"name,attr"`
  HNType string `xml:"type,attr"`
}
///////////////////

//// PORTS ////
type PortState struct {
  XMLName xml.Name `xml:"state"`
  State string `xml:"state,attr"`
  Reason string `xml:"reason,attr"`
  ReasonTtl int `xml:"reason_ttl,attr"`
}

type Service struct {
  XMLName xml.Name `xml:"service"`
  Name string `xml:"name,attr"`
  ServiceFP string `xml:"servicefp,attr"`
  Product string `xml:"product,attr"`
  Version string `xml:"version,attr"`
  Extrainfo string `xml:extrainfo,attr`
  Hostname string `xml:hostname,attr`
  OsType string `xml:ostype,attr`
  Tunnel string `xml:tunnel,attr`
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
///////////////

//// TIMES ////
type Times struct {
  XMLName xml.Name `xml:"times"`
  Srtt int `xml:"srtt,attr"`
  Rttvar int `xml:"rttvar,attr"`
  To int `xml:"to,attr"`
}
///////////////
