# Template Go-Echo-HTMX-Tailwind

This template is made to start a project with a structure based on Echo, Templ, HTMX and TailwindCSS.

## Installation

__Make__ is needed to simplify commands, you can install it from:

### Windows:

[Make Download](https://gnuwin32.sourceforge.net/packages/make.htm)

or using choco:
```bash
choco install make
```

### Linux:

```bash
sudo apt-get -y install make
```

You can download the template using:

```bash
git clone https://github.com/ArnulfoVargas/Template-Go-Echo-HTMX.git --branch v1.0
```

or just using this template.

You should add a __.env__ file, use the following to make it work:
```
PORT=<Your Port>
```

Generate the go mod files using the next command: 
```
go mod init <name or repo url>
```

Once you create you will need to call the following commands:
```
make download
make generate templ
```

Now import the packages models and templates in ./handlers/handler.go and the package handlers in ./cmd/main.go, it should look like this:
__handler.go__
```
import (
	"<name or url used in your go mod file>/models"
	"<name or url used in your go mod file>/templates"
	"github.com/labstack/echo/v4"
)
```

__main.go__
```
import (
	"os"
	"strings"

	"<name or url used in your go mod file>/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)
```

Now run:
```
make run
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
