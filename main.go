package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func add(url string) {
	fmt.Println("add()")
	fmt.Println(url)
}

func list() {
	fmt.Println("list()")
}

func poll() {
	fmt.Println("poll()")
}

func remove(id int) {
	fmt.Println("remove()")
	fmt.Println(id)
}

func main() {
	usage := `Usage:
  rs add <url>
  rs list
  rs poll
  rs remove <id>`

	opts, _ := docopt.ParseDoc(usage)
	var conf struct {
		Add    bool
		List   bool
		Poll   bool
		Remove bool
		URL    string `docopt:"<url>"`
		ID     int    `docopt:"<id>"`
	}
	var error = opts.Bind(&conf)
	if error != nil {
		fmt.Println(error)
		return
	}

	if conf.Add {
		add(conf.URL)
	} else if conf.List {
		list()
	} else if conf.Poll {
		poll()
	} else if conf.Remove {
		remove(conf.ID)
	}
}
