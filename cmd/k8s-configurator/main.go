package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	config "github.com/namely/k8s-configurator"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

const (
	// Version defines the current version of k8s-pipeliner
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "k8s-configurator"
	app.Description = "create a set of environment-specific ConfigMap files from a single yaml"
	app.Flags = []cli.Flag{}
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "creates a set of environment-specific ConfigMap files from a single yaml",
			Action: createAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func createAction(ctx *cli.Context) error {
	configFile := ctx.Args().First()
	if configFile == "" {
		return errors.New("missing parameter: file")
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	cfg := config.NewConfigFromYaml(file)

	for k, v := range cfg.OutputAll() {
		d, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(fmt.Sprintf("%v.yaml", k), d, 0644); err != nil {
			return err
		}
	}
	return nil
}
