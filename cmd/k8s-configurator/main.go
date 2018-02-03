package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "github.com/ghodss/yaml"
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
	app.Description = "generate a set of environment-specific ConfigMap files from a single yaml"
	app.Flags = []cli.Flag{}
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:   "generate",
			Usage:  "generates a set of environment-specific ConfigMap files from a single yaml",
			Action: generateAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "input yaml file containing config information",
				},
				cli.StringFlag{
					Name:  "out, o",
					Value: "stdout",
					Usage: "output for generated files. defaults to stdout. if other than stdout a directory will be created.",
				},
				cli.StringFlag{
					Name:  "env, e",
					Value: "default",
					Usage: "the environment to generate. use `all` to generate all environments (only allowed with an output directory).",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func generateAction(ctx *cli.Context) error {
	configFile := ctx.String("file")
	if configFile == "" {
		return errors.New("missing parameter: file")
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	cfg := config.NewConfigFromYaml(file)

	env := ctx.String("env")
	out := ctx.String("out")
	all := cfg.OutputAll()

	if env == "all" {
		if out == "stdout" {
			return errors.New("cannot use stdout with all env output")
		}
		out := ctx.String("out")
		if _, err := os.Stat(out); os.IsNotExist(err) {
			os.Mkdir(out, 0744)
		}

		for k, v := range all {
			d, err := yaml.Marshal(v)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(fmt.Sprintf("%v/%v.yaml", out, k), d, 0644); err != nil {
				return err
			}
		}
	} else {
		e, exists := all[env]
		if !exists {
			return fmt.Errorf("requested env %v does not exist in input yaml", env)
		}
		d, err := yaml.Marshal(e)
		if err != nil {
			return err
		}
		fmt.Printf("%s", d)
	}
	return nil
}
