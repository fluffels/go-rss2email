package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
)

func add(url string) {
    fmt.Println("add()")
}

func main() {
    usage := `rss2email

Usage:
  rs add <url>
  rs list
  rs poll
  rs remove <id>`

    args, _ := docopt.ParseDoc(usage)
    var opts struct {
        Add bool
        List bool
        Poll bool
        Remove bool
        Url string `docopt:"url"`
        Id int
    };
    args.Bind(&opts);

    fmt.Println(opts);
    fmt.Println(args);

    if (opts.Add) {
        add(opts.Url);
    }
}
