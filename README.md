![templ](logo.svg) 

# Templ with icons provided by Radix

Simple package to use radix icons in GoLang a-h templ files.

## Installation

```bash
go get -u github.com/bbassie/go-templ-radix-icons
```

## Usage

Example usage inside of a `.templ` file:
```go
package pages

import "github.com/bbassie/go-templ-radix-icons/icons"

templ HomePage() {
    @icons.ActivityLogIcon()
    @icons.SectionIcon()
    @icons.SunIcon()
    @icons.MoonIcon()
    @icons.BellIcon()
}
```



### Thanks to the original authors of the `templ` and `radix-ui` projects.
- https://github.com/radix-ui/icons
- https://github.com/a-h/templ