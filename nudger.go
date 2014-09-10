package main

import (
	"bytes"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/sgrimee/nudger/lib"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"strings"
)

var config *nudger.ConfigType

func main() {

	// initialize config object
	var err error
	config, err = nudger.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config: %#v", config)

	// launch web service
	m := martini.Classic()
	m.Get("/", ShowHelp)
	m.Group("/items", func(r martini.Router) {
		r.Get("/", ListItems)
		r.Get("/:item", NudgeItem)
		r.Post("/:item", NudgeItem)
	})
	m.Run()
}

func ShowHelp() string {
	text := "This trivial service tiggers an action on pre-defined items via "
	text += "HTTP requests.\n"
	text += "To action an item, send a GET or POST to its url /items/:item"
	text += "List available items with GET /items"
	return text
}

func ListItems() string {
	text := fmt.Sprintf("%s", listDir(config.ItemsDir))
	return text
}

// return a slice with entries in the given directory
func listDir(itemsDir string) []string {
	f, err := os.Open(itemsDir)
	if err != nil {
		log.Fatal("Cannot open dir ", itemsDir)
	}
	items, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal("Cannot get list of items from ", itemsDir)
	}
	return items
}

// Execute the action on the item
// only if it is in the authorized list
func NudgeItem(params martini.Params, r *http.Request) string {
	//requestBody, err := ioutil.ReadAll(r.Body)
	for _, validItem := range listDir(config.ItemsDir) {
		if validItem == params["item"] {
			// execute the action
			//out, err := exec.Command(config.NudgeCmd, config.NudgeArgs, params["item"]).Output()
			cmd := exec.Command(config.NudgeCmd, config.NudgeArgs, params["item"])
			//cmd.Stdin = strings.NewReader(string(r.Body))
			cmd.Stdin = r.Body
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				return err.Error()
			}
			result := fmt.Sprintf("Result: %s", out)
			return result
		}
	}
	return "Item not in authorized items list."
}
