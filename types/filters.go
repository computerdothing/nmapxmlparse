package types

type HostFilter struct {}

type HostFilterChain []HostFilter

func (nmr NmapRun) Filter(fc HostFilterChain) NmapRun {
  var rv NmapRun

  // copy meta stuff over since we're just filtering hosts
  rv.Scanner = nmr.Scanner
  rv.Args = nmr.Args
  rv.Start = nmr.Start
  rv.StartStr = nmr.StartStr
  rv.Version = nmr.Version
  rv.XmlOutputVersion = nmr.XmlOutputVersion
  rv.ScanInfo = nmr.ScanInfo
  rv.Verbose = nmr.Verbose
  rv.Debugging = nmr.Debugging


}
