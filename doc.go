/*
Package pathmatch provides pattern matching for paths.

For example, a path could be a file system path, or a path could be a path from a URL (such as an HTTP or HTTPS based URL).

The matches can be loaded into variables (when using pathmatch.Find());
or can be loaded into a struct (when using pathmatch.Pattern.FindAndLoad()).

Example Usage:

	var pattern pathmatch.Pattern
	
	err := pathmatch.Compile(&pattern, "/users/{user_id}/vehicles/{vehicle_id}")
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

Alternate Example Usage:

	var pattern pathmatch.Pattern

	err := pathmatch.Compile(pattern *Pattern, "/users/{user_id}/vehicles/{vehicle_id}")
	if nil != err {
		//@TODO
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

*/
package pathmatch
