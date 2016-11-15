package main

import (
	"fmt"
	"github.com/spf13/cobra"
  "nmapxmlparse/xmlimport"
	"os"
)

const (
	swVersion = "0.0.0"
  // xmloutputversion = "1.04"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(swVersion)
	},
}

var RootCmd = &cobra.Command{
	Use:   "nmapxmlparse",
	Short: "Parse NMap XML output files.",
	Long:  "Parse NMap XML output files and do useful things with it. Convert to other formats.",
}

var doCmd = &cobra.Command{
  Use: "do",
  Short: "Do default action",
  Run: func(cmd *cobra.Command, args []string) {
    for _, arg := range args {
      result, err := xmlimport.ImportXmlFile(arg)

      if err != nil {
        fmt.Printf("error: %v\n", err)
      }

			fmt.Printf("Scanner: %s\nArgs: %s\nStart: %s\nVersion: %s\n", result.Scanner, result.Args, result.StartStr, result.Version)
			fmt.Printf("Verbosity: %d\nDebugging Level: %d\n", result.Verbose.Level, result.Debugging.Level)

      for _, host := range result.Hosts {
        if len(host.Ports) > 0 {
          fmt.Printf("----------------------\n")
          fmt.Printf("%s\n", host.Address.Addr)

          for _, hostname := range host.Hostnames {
            fmt.Printf("\t%s  -  %s  -  %s\n", hostname.Name, hostname.HNType, host.Status.State)
          }

          for _, port := range host.Ports {
            fmt.Printf("\t\t%d  -  %s  -  %s  -  %s  -  %s\n", port.Portid, port.State.State, port.Service.Name, port.Service.Product, port.Service.Version)
          }
        }
      }
    }
  },
}

func init() {
  RootCmd.AddCommand(versionCmd)
  RootCmd.AddCommand(doCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
