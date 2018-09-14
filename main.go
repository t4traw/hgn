package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "hgn"
	app.Usage = "hugo generate new post cli tool"
	app.Version = "0.1"

	app.Action = func(context *cli.Context) error {
		fmt.Println(context.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
