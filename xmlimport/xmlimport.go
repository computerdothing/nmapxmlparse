package xmlimport

import (
    //"bytes"
    "encoding/xml"
    "io/ioutil"
    "nmapxmlparse/types"
)

// Attempt to read in the provided XML file and return a types.ScanResults
// Returns error verbatim on error
func ImportXmlFile(filename string) (types.NmapRun, error) {
  var result types.NmapRun

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return result, err
  }

  /*
  d := xml.NewDecoder(bytes.NewReader(data))

  for {
    token, _ := d.Token()
    if token == nil {
      break
    }
    // fixme: figure out wtf I'm doing here and do it more
    switch startElement := token.(type) {
      case xml.StartElement:
        if startElement.Name.Local == "nmaprun" {
          d.DecodeElement(&result, &startElement)
        }
    }
  }
  */

  err = xml.Unmarshal(data, &result)
  if err != nil {
    return result, err
  }

  return result, nil
}
