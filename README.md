[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/WildEgor/container-data-structures)](https://goreportcard.com/report/github.com/WildEgor/container-data-structures)

Contains data structures additionally for [container](https://pkg.go.dev/container)

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Installation
```shell
go get github.com/WildEgor/container-data-structures
```

## Usage

Project structure:
- .github
    - workflows
        - release.yml - run semantic release
        - testing.yml - run checks and tests
    - internal
      - list.go - generic list impl
    - pkg
        - orderedmap - ordered map impl

Structures:
- [x] [OrderedMap](./examples/orderedmap/main.go)
- [x] [Stack](./examples/stack/main.go)

## Contributing

Please, use git cz for commit messages!

```shell
git clone https://github.com/WildEgor/container-data-structures
cd container-data-structures
git checkout -b feature-or-fix-branch
git add .
git cz
git push --set-upstream-to origin/feature-or-fix-branch
```