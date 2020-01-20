# bump

[![Go status](https://github.com/haya14busa/bump/workflows/Go/badge.svg)](https://github.com/haya14busa/bump/actions)
[![release](https://github.com/haya14busa/bump/workflows/release/badge.svg)](https://github.com/haya14busa/bump/actions?query=workflow%3Arelease)
[![releases](https://img.shields.io/github/release/haya14busa/bump.svg)](https://github.com/haya14busa/bump/releases)

**bump** returns next semantic version tag.

## Installation

```shell
# Install latest version. (Install it into ./bin/ by default).
$ curl -sfL https://raw.githubusercontent.com/haya14busa/bump/master/install.sh| sh -s

# Specify installation directory ($(go env GOPATH)/bin/) and version.
$ curl -sfL https://raw.githubusercontent.com/haya14busa/bump/master/install.sh| sh -s -- -b $(go env GOPATH)/bin [vX.Y.Z]

# In alpine linux (as it does not come with curl by default)
$ wget -O - -q https://raw.githubusercontent.com/haya14busa/bump/master/install.sh| sh -s [vX.Y.Z]

# homebrew / linuxbrew
$ brew install haya14busa/tap/bump
$ brew upgrade haya14busa/tap/bump

# Go
$ go get github.com/haya14busa/bump
```

## Usage

```
# Usage: bump [major,minor,patch (default=patch)]
$ git tag
v0.0.1
v0.0.2
v0.0.3
v0.0.4
$ bump patch
v0.0.5
$ bump minor
v0.1.0
$ bump major
v1.0.0

# Annotate next patch version tag.
$ git tag -a $(bump patch) -m $(bump patch)
```

