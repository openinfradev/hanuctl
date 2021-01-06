# hanuctl 
CLI tool for Declarative Lifecycle Management of Kubernetes Cluster

------

## What is the hanuctl

The hanuctl project is a CLI tool for declarative management of kubernetes and 
underlying infrastructure, mainly leveraging the Kubernetes Cluster API. It 
builds and manage a Kubernetes Cluster either on Baremetal or OpenStack VMs.

This project is the ongoing effort to produce next generation of existing 
[tacoplay](github.com/openinfradev/hanuctl). Currently, the hanuctl is at an very early development stage.


## Installation

Build after receiving the hanuctl source from the github repository. At this 
time, golang must be installed.

```sh
git clone https://github.com/openinfradev/hanuctl.git
cd hanuctl
git checkout v0.1
make build
mv bin/hanuctl /usr/local/bin
hanuctl -h
Find more information at: https://github.com/openinfradev/hanuctl
 
Usage:
  hanuctl [flags]
  hanuctl [command]
 
Available Commands
  completion  Generates bash completion scripts
  create      Create a resource from a file
  delete      Delete resources like a cluster, node
  get         Display one or many resources
  help        Help about any command
  version     Print this tool version information
 
Flags:
      --config string   config file (default is $HOME/.hanuctl.yaml)
  -h, --help            help for hanuctl
 
Use "hanuctl [command] --help" for more information about a command.

```
## Configuration

### Initialization
Use `tools/initialize.sh` for configuration. It copied yaml files to 
`/tmp/hanuctl/`, so you could fix it to your environment.

## Basic Usage

### Create and Delete Cluster 

```sh
hanuctl create cluster
hanuctl delete cluster
```

### Add Worker Node

```sh
hanuctl create node
```

## Features

- the lifecycle of a Kubernetes cluster

## Roadmap

- the lifecycle of a Kubernetes cluster on Baremetal
