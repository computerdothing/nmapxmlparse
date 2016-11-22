package output

import (
	"encoding/csv"
	"fmt"
	"nmapxmlparse/types"
	"nmapxmlparse/util"
	"os"
	"strconv"
)

// Convert 'data' to CSV format and write it to 'outfile'
// Return error verbatim on error and nil on success
func CsvOut(data types.NmapRun, outfile string, altout bool) error {
	if util.Exists(outfile) {
		return fmt.Errorf("File %s already exists.", outfile)
	}

	f, err := util.SafeOpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(PrepareCsv(data, altout))
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

// Take the NmapRun structure, add field headings and convert it to our CSV format
// Returns data as csv ready [][]string
func PrepareCsv(data types.NmapRun, altout bool) [][]string {
	var result [][]string

	fieldHeadings := []string{"Host IP", "Host Scan Start", "Host Scan End", "Host State", "Host State Reason",
		"Hostnames", "OS", "OS-match Accuracy", "OS Type", "OS Vendor", "OS Family", "OS Gen", "Last Boot",
		"Port ID", "Port State", "Port State Reason", "Service Name", "Service Product", "Service Version",
		"Service Hostname", "Service OSType", "Service Extrainfo", "Service Tunnel", "Service Method",
		"Service Confidence", "Service FP String"}

	result = append(result, fieldHeadings)

	for _, host := range data.Hosts {
		var hostnames string

		for i, hostname := range host.Hostnames {
			var singleHN string
			if i+1 == len(host.Hostnames) {
				singleHN = fmt.Sprintf("%s : %s", hostname.Name, hostname.HNType)
			} else {
				singleHN = fmt.Sprintf("%s : %s --", hostname.Name, hostname.HNType)
			}
			hostnames += singleHN
		}

		if len(host.Ports) == 0 {
			var tmp []string
			tmp = append(tmp, host.Address.Addr, strconv.Itoa(host.Starttime), strconv.Itoa(host.Endtime),
				host.Status.State, host.Status.Reason, hostnames, host.Os.OsMatch.Name, strconv.Itoa(host.Os.OsMatch.Accuracy),
				host.Os.OsMatch.OsClass.Type, host.Os.OsMatch.OsClass.Vendor, host.Os.OsMatch.OsClass.OsFamily,
				host.Os.OsMatch.OsClass.OsGen, host.Uptime.Lastboot, "", "", "", "", "", "", "", "", "", "", "", "", "",)
				result = append(result, tmp)
		}

		for i, port := range host.Ports {
			var thisRow []string
			if altout && (i != 0) {
				thisRow = append(thisRow, "", "", "", "", "", "", "", "", "", "", "", "", "", strconv.Itoa(port.Portid),
					port.State.State, port.State.Reason, port.Service.Name, port.Service.Product, port.Service.Version,
					port.Service.Hostname, port.Service.OsType, port.Service.Extrainfo, port.Service.Tunnel,
					port.Service.Method, strconv.Itoa(port.Service.Conf), port.Service.ServiceFP)
			} else {
				thisRow = append(thisRow, host.Address.Addr, strconv.Itoa(host.Starttime), strconv.Itoa(host.Endtime),
					host.Status.State, host.Status.Reason, hostnames, host.Os.OsMatch.Name, strconv.Itoa(host.Os.OsMatch.Accuracy),
					host.Os.OsMatch.OsClass.Type, host.Os.OsMatch.OsClass.Vendor, host.Os.OsMatch.OsClass.OsFamily,
					host.Os.OsMatch.OsClass.OsGen, host.Uptime.Lastboot, strconv.Itoa(port.Portid), port.State.State,
					port.State.Reason, port.Service.Name, port.Service.Product, port.Service.Version, port.Service.Hostname,
					port.Service.OsType, port.Service.Extrainfo, port.Service.Tunnel, port.Service.Method,
					strconv.Itoa(port.Service.Conf), port.Service.ServiceFP)
			}
			result = append(result, thisRow)
		}
	}
	return result
}
