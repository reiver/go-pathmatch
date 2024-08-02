# go-pathmatch

A library that provides *pattern-matching* for paths, for the Go programming language.

For example, a path could be a file system path, or a path could be a path from a URL (such as an HTTP or HTTPS based URL).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-pathmatch

[![GoDoc](https://godoc.org/github.com/reiver/go-pathmatch?status.svg)](https://godoc.org/github.com/reiver/go-pathmatch)


## Example Usage
```go
import (
	"github.com/reiver/go-pathmatch"
)

// ...

var pattern pathmatch.Pattern

err := pathmatch.CompileTo(&pattern, "/users/{user_id}/vehicles/{vehicle_id}")
if nil != err {
	fmt.Fprintf(os.Stdout, "ERROR: %s\n", err)
	return
}

var userId    string
var vehicleId string

matched, err := pattern.Find("/users/bMM_kJFMEV/vehicles/o_bcU.RZGK", &userId, &vehicleId)
if nil != err {
	fmt.Fprintf(os.Stdout, "ERROR: %s\n", err)
	return
}

if !matched {
	fmt.Println("The patch did not match.")
	return
}

fmt.Println("The path matched!")

fmt.Printf("user_id     = %q \n", userId)     // user_id     = "bMM_kJFMEV"
fmt.Printf("vehicle_id  = %q \n", vehicleId)  // vehicle_id  = "o_bcU.RZGK"
```

Alternatively:
```go
import (
	"github.com/reiver/go-pathmatch"
)

// ...

var pattern patchmatch.Pattern

err := pathmatch.CompileTo(&pattern, "/users/{user_id}/vehicles/{vehicle_id}")
if nil != err {
	fmt.Fprintf(os.Stdout, "ERROR: %s\n", err)
	return
}

data := struct{
	UserId    string `match:"user_id"`
	VehicleId string `match:"vehicle_id"`
}{}

matched, err := pattern.FindAndLoad("/users/bMM_kJFMEV/vehicles/o_bcU.RZGK", &data)
if nil != err {
	fmt.Fprintf(os.Stdout, "ERROR: %s\n", err)
	return
}

if !matched {
	fmt.Println("The patch did not match.")
	return
}


fmt.Println("The path matched!")

fmt.Printf("user_id     = %q \n", data.UserId)     // user_id     = "bMM_kJFMEV"
fmt.Printf("vehicle_id  = %q \n", data.VehicleId)  // vehicle_id  = "o_bcU.RZGK"
```

## Import

To import package **pathmatch** use `import` code like the follownig:
```
import "github.com/reiver/go-pathmatch"
```

## Installation

To install package **pathmatch** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-pathmatch
```

## Author

Package **pathmatch** was written by [Charles Iliya Krempeaux](http://reiver.link)
