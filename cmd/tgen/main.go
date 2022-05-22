package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/clefever/truchetgen"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Truchet Tile Pattern Generator.

	Usage:
	  tgen <input-file> [<output-file>] 
	  tgen -h | --help
	
	Options:
	  -h --help     Show this screen.`

	opts, _ := docopt.ParseDoc(usage)
	input, _ := opts.String("<input-file>")
	output := parseString(opts, "<output-file>", "output.svg")

	jsonFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var layout truchetgen.TileLayout
	json.Unmarshal(jsonBytes, &layout)

	if err := layout.Draw(output); err != nil {
		log.Fatal(err)
	}
}

func parseString(opts docopt.Opts, name string, defaultValue string) string {
	opt := opts[name]
	if opt == nil {
		return defaultValue
	} else {
		return opt.(string)
	}
}
