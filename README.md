# BuildIt

BuildIt is a program to run any file that you defined.

# Install

## Without Golang

If you don't have [Golang](https://golang.org/) installed in your machine, do these commands:

```
git clone https://github.com/lissaferreira/buildit
cd buildit
sudo cp bin/buildit /usr/bin/buildit
```

## Golang

Install Golang with your package manager, and install BuildIt

To install BuildIt use:

`go get -u github.com/lissaferreira/buildit`

# Configuration

create a file `~/.buildit.yaml`, with the languages and compiler/runner commands, like this:

```yaml
js:
  - node
  - yarn serve
go:
  - go run
sh:
  - sh
php:
  - php
pdf:
  - zathura
rb:
  - ruby
```

# Usage

use `buildit [FILE]`, like `buildit main.go`.

Arguments can be passed only putting in the end, like `buildit index.js --argument=argument`
