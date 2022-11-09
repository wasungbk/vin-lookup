package main

import (
	"github.com/wasungbk/vin-lookup/service"
)

//   - `--raw` - output the raw JSON response of the API
//   - `--fields` - output the results of the fields that match the name or regular expression
//     provided to the argument
//   - `--sparse` - only output fields that have data in them and aren't empty or null
//   - `--meta` - output how many fields and results were returned to the response
//   - `--yaml` - output the results in the YAML format
//   - Note: if `--raw` and `--yaml` are both provided, whichever option comes later in the
//     arguments should be preferred
func main() {
	//parse flags

	// make call
	service.CallVin()

	// format output

}
