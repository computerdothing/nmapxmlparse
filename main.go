package main

import (
	"fmt"
	"nmapxmlparse/output"
	"nmapxmlparse/xmlimport"
	"os"

	"github.com/spf13/cobra"
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

var rootCmd = &cobra.Command{
	Use:   "nmapxmlparse",
	Short: "Parse NMap XML output files.",
	Long:  "Parse NMap XML output files and do useful things with it. Convert to other formats.",
}

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Usage: nmapxmlparse csv <nmap-xml-input-filename> <csv-output-filename>",
	Long:  "Attempts to convert NMap XML data from input file to CSV and write that to outfile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: nmapxmlparse csv <nmap-xml-input-filename> <csv-output-filename>")
		}

		result, err := xmlimport.ImportXmlFile(args[0])
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		if err := output.CsvOut(result, args[1], altout); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Usage: nmapxmlparse do <nmap-xml-input-filename>",
	Long:  "Currently, just converts the Nmap XML data in the into a really ugly and possibly less useful output format and prints it",
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

var altout bool

func init() {
	csvCmd.Flags().BoolVarP(&altout, "alternate-output", "a", false, "Do the alternate formatting of the CSV file")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(csvCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
