# Template Go-Echo-HTMX-Tailwind

Foobar is a Python library for dealing with word pluralization.

## Installation

__Make__ is needed to simplify commands, you can install it from:

__Windows:__

[Make Download](https://gnuwin32.sourceforge.net/packages/make.htm)

or using choco:
```bash
choco install make
```

__Linux:__

```bash
sudo apt-get -y install make
```

You can download the template using:

```bash
git clone https://github.com/ArnulfoVargas/Template-Go-Echo-HTMX.git --branch v1.0
```
Then make a cd to the project and generate the go mod file:
```bash
go mod init <your repo url>
```

## Usage
There's some Makefile commands to launch the server:

```bash
# Download the dependencies (Echo, Templ, Godotenv)
make download

# Generate templ files
make templ

# Generate templ files, build an executable and run it
make build

# Run main without generate anything
make run
```
