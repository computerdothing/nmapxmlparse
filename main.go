package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"nmapxmlparse/output"
	"nmapxmlparse/types"
	"nmapxmlparse/xmlimport"
	"os"
	"strconv"
	"strings"
)

const (
	swVersion = "1.0.0"
	// xmloutputversion = "1.04"
)

var rootCmd = &cobra.Command{
	Use:   "nmapxmlparse",
	Short: "Parse NMap XML output files.",
	Long:  "Parse NMap XML output files and do useful things with it. Convert to other formats.",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(swVersion)
	},
}

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Usage: nmapxmlparse csv <nmap-xml-input-filename> <csv-output-filename>",
	Long:  "Attempts to convert NMap XML data from input file to CSV and write that to outfile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: nmapxmlparse csv <nmap-xml-input-filename> <csv-output-filename>")
			return
		}

		result, err := xmlimport.ImportXmlFile(args[0])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		var filters []types.HostFilter

		if openPorts {
			filters = append(filters, types.AnyOpenPort())
		}

		if specificPortsStr != "" {
			filters = append(filters, types.SelectOpenPorts(processSpecficPortsStr(specificPortsStr)...))
		}

		if err := result.Filter(filters...); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		if err := output.CsvOut(result, args[1], altout); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	},
}

var (
	altout bool
	openPorts bool
	specificPortsStr string
)

func processSpecficPortsStr(pStr string) []int {
	var rv []int
	for _, s := range strings.Split(pStr, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
		rv = append(rv, i)
	}
	return rv
}

func init() {
	csvCmd.Flags().BoolVarP(&altout, "alternate-output", "a", false, "Do the alternate formatting of the CSV file.")
	rootCmd.PersistentFlags().BoolVarP(&openPorts, "open-hosts", "O", false, "Only show results for hosts with open ports.")
	rootCmd.PersistentFlags().StringVarP(&specificPortsStr, "specific-ports", "o", "", "Olny show results for hosts with specific open ports. (Comma separated list)")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(csvCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
