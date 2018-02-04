package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	config "github.com/namely/k8s-configurator"
	"github.com/urfave/cli"
)

const (
	// Version defines the current version of k8s-pipeliner
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "k8s-configurator"
	app.Description = "generate an environment-specific ConfigMap from a single source yaml"
	app.Flags = []cli.Flag{}
	app.Version = Version
	app.UsageText = "generate input-file env"

	app.Commands = []cli.Command{
		{
			Name:      "generate",
			Usage:     "generate an environment-specific ConfigMap from a single source yaml",
			Action:    generateAction,
			ArgsUsage: "input-file env",
			UsageText: "generate input-file env",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}

func generateAction(ctx *cli.Context) error {
	configFile := ctx.Args().Get(0)
	if configFile == "" {
		return errors.New("missing file parameter")
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	env := ctx.Args().Get(1)
	if len(env) == 0 {
		env = "default"
	}

	return config.Generate(file, env, os.Stdout)

}
