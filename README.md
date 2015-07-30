# ![Routify](http://praegress.us/routify-logo.png)
A usable, efficient and simple router implementation for Go.

## Overview
Routify is a route generation tool and routing package. The primary goal is to provide a simple and usable way of handling routing in Go. The router's performance and garbage creation overhead is comparable to other Go routing packages, but at present this isn't a primary goal - it's void of exotic data structures.

## Installation
    go get -u github.com/martingallagher/routify

The following assumes your Go $GOPATH/bin is on your $PATH environmental variable (`export PATH=$PATH:$GOPATH/bin`).

## Routing Example
Routes are defined in a simple YAML file. For example:

```yaml
# Method -> Route format
GET:
  /:                        indexHandler
  blog:                     blogHandler
  blog/$year:               blogArchiveHandler
  blog/$year/$month:        blogArchiveHandler
  blog/$year/$month/$day:   blogArchiveHandler

# Route -> Method format
blog/post:
  POST: 	newPost
  DELETE:	deletePost

# Params defines URL paramters to be captured and validated.
# URL components prefixed with "$" with no matching validation function
# will be captured but not validated.
params:
  $year:   IsYear
  $month:  IsMonth
  $day:    IsDay
```

# Traditional Routing
```go
r := &Router{}
r.AddValidator(":year", router.IsYear)
r.AddValidator(":month", router.IsMonth)
r.AddValidator(":day", router.IsDay)

if err := r.Add("GET", "blog/archives/:year/:month/:day", blogArchivesHandler); err != nil {
	// Handle error
}
```

# Accessing Parameters
```go
_, params, err := routes.Get(r) // Handle error

// String value, if not found returns empty string
v := params.Get("year")

// Integer helpers
v, err := params.GetInt("year") // int64
v, err := params.GetUint("id") // uint64

// Scanner interface
type month time.Month

func (m *month) Scan(i interface{}) error {
	v, ok := i.(string)

	if !ok {
		return errors.New("unsupported type")
	}

	c, err := strconv.Atoi(v)

	if err != nil {
		return err
	} else if c < 0 || c > 12 {
		return errors.New("month out of bounds")
	}

	*m = month(c)

	return nil
}

// Scan example
var m month

err := params.Scan("month", &m)

// Assuming month=02
fmt.Println(time.Month(m).String() == "February")
```

# `routify` Command Line Tool
The routify tool generates the Go routes file. Run `routify -h` for full options.

**Example:**

`routify -i routes.yaml -p blog -v routes`

## Using `go generate`
Routify works great in tandem with `go generate`, making route generation easy with the standard Go tools.

**Example `generate.go` file:**

```go
// Generate routes.go for the blog package
//go:generate routify -i routes.yaml -p blog
// gofmt the routify Go output
//go:generate gofmt -w -s routes.go

package blog
```

# Server Example
A simple reference example is available in the `example` directory.

# Contributions
Bug fixes and feature requests welcome.

# Contributors
- [Martin Gallagher](http://martingallagher.com/)