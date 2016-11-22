package main

import (
	"fmt"
	"nmapxmlparse/output"
	"nmapxmlparse/xmlimport"
	"os"

	"github.com/spf13/cobra"
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

		if err := output.CsvOut(result, args[1], altout); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	},
}

var altout bool

func init() {
	csvCmd.Flags().BoolVarP(&altout, "alternate-output", "a", false, "Do the alternate formatting of the CSV file")
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
