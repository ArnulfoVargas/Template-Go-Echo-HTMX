# Template Go-Echo-HTMX-Tailwind

This template is made to start a project with a structure based on Echo, Templ, HTMX and TailwindCSS.

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
