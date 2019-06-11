# k8s-configurator

Generate environment-specific ConfigMaps from a single source Yaml.

[![Build Status](https://travis-ci.org/namely/k8s-configurator.svg?branch=master)](https://travis-ci.org/namely/k8s-configurator)
[![Coverage Status](https://coveralls.io/repos/github/namely/k8s-configurator/badge.svg?branch=master)](https://coveralls.io/github/namely/k8s-configurator?branch=master)

## How to install

Pre-compiled binaries:

1. On macOS via Homebrew:

   - Setup [Namely's Homebrew tap](https://github.com/namely/homebrew-tap)
   - Install the pre-built binary `brew install k8s-configurator`

2. On Linux or Windows Subsystem for Linux:

   - Run `curl -sL https://github.com/namely/k8s-configurator/releases/download/v0.0.3/k8s-configurator_0.0.3_linux_amd64.tar.gz | sudo tar -xzf - -C /usr/local/bin/`

Compile from source:

 - Ensure that the `bin/` folder under the configured `$GOPATH` is on `$PATH`
 - Run `go get -u github.com/namely/k8s-configurator/cmd/k8s-configurator`

## Usage

```
k8s-configurator generate <source file> <env>
```

This will read the input yaml and generate a Kubernets V1 ConfigMap for a given env
to stdout. This allows you to pass the output directly to kubectl:

```
k8s-configurator generate <source file> <env> | kubectl create -f -
```

## Input File and Envs

The input yaml file is a simplified version of a ConfigMap to
which makes specifying settings and overrides simple. You include
the name and namespace of the ConfigMap, a default set of values,
and any applicable overrides. k8s-configurator will merge the overrides
with the default based on the environment specified.

Example yaml:

```
name: foo
namespace: system
annotations:
    strategy.spinnaker.io/versioned: false

default:
  setting1: value1
  setting2: value2

overrides:
  int:
    setting1: an-override
```
