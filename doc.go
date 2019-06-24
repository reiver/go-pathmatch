/*
Package pathmatch provides pattern matching for path templates.

A path template might look something like the following:

	/v1/users/{user_id}

Or:

	/account={account_name}/user={user_name}/message={message_hash}

Or:

	/backup/{folder_name}/
Or:

	/v2/select/{fields}/from/{table_name}/where/{filters}

This path template could be a file system path, or a path could be a path from a URL (such as an HTTP or HTTPS based URL).

To compile one of these pattern templates, you would do something such as:

	var template string = "/v1/users/{user_id}/messages/{message_id}"

	var pattern pathmatch.Pattern

	err := pathmatch.CompileTo(&pattern, template)
	if nil != err {
		fmt.Fprintf(os.Stdout, "ERROR: %s\n", err)
		return
	}

(In addition to the pathmatch.CompileTo() func, there is also the pathmatch.Compile(), and
pathmatch.MustCompile(). But pathmatch.CompileTo() is recommended over the other 2 options
for most cases.)

One you have the compiled pattern, you would either use pathmatch.Match(), pathmatch.Find(),
or pathmatch.FindAndLoad() depending on what you were trying to accomplish.

Example Usage

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

Alternate Example Usage

	var pattern pathmatch.Pattern

	err := pathmatch.CompileTo(pattern *Pattern, "/users/{user_id}/vehicles/{vehicle_id}")
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
