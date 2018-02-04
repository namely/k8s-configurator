# k8s-configurator

Generate environment-specific ConfigMaps from a single source Yaml.

[![Build Status](https://travis-ci.org/namely/k8s-configurator.svg?branch=master)](https://travis-ci.org/namely/k8s-configurator)
[![Coverage Status](https://coveralls.io/repos/github/namely/k8s-configurator/badge.svg?branch=mlh%2Ftravis)](https://coveralls.io/github/namely/k8s-configurator?branch=mlh%2Ftravis)

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

default:
  setting1: value1
  setting2: value2

overrides:
  int:
    setting1: an-override
```
