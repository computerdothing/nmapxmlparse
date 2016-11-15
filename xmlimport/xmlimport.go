package xmlimport

import (
    "bytes"
    "encoding/xml"
    "io/ioutil"
    "nmapxmlparse/types"
)

// Attempt to read in the provided XML file and return a types.ScanResults
// Returns error verbatim on error
func ImportXmlFile(filename string) ([]types.Host, error) {
  //var result types.NmapRun
  var result []types.Host

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return result, err
  }

  d := xml.NewDecoder(bytes.NewReader(data))

  for {
    token, _ := d.Token()
    if token == nil {
      break
    }
    // fixme: figure out wtf I'm doing here and do it more
    switch startElement := token.(type) {
      case xml.StartElement:
        if startElement.Name.Local == "host" {
          var ahost types.Host
          d.DecodeElement(&ahost, &startElement)
          result = append(result, ahost)
        }
    }
  }

  return result, nil
}
