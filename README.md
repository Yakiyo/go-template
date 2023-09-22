# go-template
[![ci](https://github.com/Yakiyo/go-template/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Yakiyo/go-template/actions/workflows/ci.yml) ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Yakiyo/go-template) ![GitHub tag (with filter)](https://img.shields.io/github/v/tag/Yakiyo/go-template?label=version)

My own template repository for cli apps with Go

Some of the code structure and ideas were taken from [metafates/go-template](https://github.com/metafates/go-template/).

## Features
- use cobra for command line
- integrated viper with cobra to read config options first from cli, then use config files, else use defaults
- separate packages for each stuff for easier modification
- pretty logging with [charmbracelet/log](https://github.com/charmbracelet/log)
- allow setting logging levels from cli
- allow enabling/disabling colors from cli
- use [just](https://github.com/casey/just) for running tasks
