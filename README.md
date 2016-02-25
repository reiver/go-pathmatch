# go-pathmatch

A library that provides *pattern matching* for paths, for the Go programming language.

For example, a path could be a file system path, or a path could be a path from a URL (such as an HTTP or HTTPS based URL).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-pathmatch

[![GoDoc](https://godoc.org/github.com/reiver/go-pathmatch?status.svg)](https://godoc.org/github.com/reiver/go-pathmatch)


## Example Usage
```
import (
	"github.com/reiver/go-pathmatch"
)

// ...

pattern, err := pathmatch.Compile("/users/{user_id}/vehicles/{vehicle_id}}")
if nil != err {
    //@TODO
}

var userId    string
var vehicleId string

didMatch, err := pattern.Match("/users/bMM_kJFMEV/vehicles/o_bcU.RZGK", &userId, &vehicleId)

if nil != err {
    //@TODO
}

if didMatch {
    fmt.Println("The path matched!")

    fmt.Printf("user_id     = %q \n", userId)     // user_id     = "bMM_kJFMEV"
    fmt.Printf("vehicle_id  = %q \n", vehicleId)  // vehicle_id  = "o_bcU.RZGK"
} else {
    fmt.Println("The patch did not match.")
}
```

Alternatively:
```
import (
	"github.com/reiver/go-pathmatch"
)

// ...

pattern, err := pathmatch.Compile("/users/{user_id}/vehicles/{vehicle_id}}")
if nil != err {
    //@TODO
}

data := struct{
	UserId    string `match:"user_id"`
	VehicleId string `match:"vehicle_id"`
}{}

didMatch, err := pattern.MatchAndLoad("/users/bMM_kJFMEV/vehicles/o_bcU.RZGK", &data)

if nil != err {
    //@TODO
}

if didMatch {
    fmt.Println("The path matched!")

    fmt.Printf("user_id     = %q \n", data.UserId)     // user_id     = "bMM_kJFMEV"
    fmt.Printf("vehicle_id  = %q \n", data.VehicleId)  // vehicle_id  = "o_bcU.RZGK"
} else {
    fmt.Println("The patch did not match.")
}
```
