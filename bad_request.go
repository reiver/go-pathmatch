package pathmatch

// BadRequest is used to represent one of the types of errors that could be returned when
// calling the pathmatch.Compile func, the pathmatch.Pattern.Match method, or the pathmatch.Pattern.MatchAndLoad
// method. The meaning of this type of error is that the problem was due to something whomever called the func or method did.
//
// For example, maybe the uncompiled pattern passed to the pathmatch.Compile() func had
// a syntax error in it. Or, also for example, maybe the type of parameter passed to
// pathmatch.Pattern.Find() was of the wrong type. Etc.
//
// Example usage of BadRequest with pathmatch.Compile():
//
//	pattern, err := pathmatch.Compile("/fruits/{there_is_an_error_in_here/") // ← The uncompiled pattern there has an error in it.
//	if nil != err {
//		switch err.(type) { // ← Note that we are using a Go type-switch here.
//	
//		case pathmatch.BadRequest: // ← Here we are detecting if the error returned is a pathmatch.BadRequest.
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalError:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pathmatch.Compile() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
//
// Somewhat continuing this example (although without the error in the uncompiled pattern), we might then
// use the pathmatch.Pattern.Find() method, which could also generate an error that fits the BadRequest
// interface.
//
// Example usage of BadRequest with pathmatch.Pattern.Find():
//
//	pattern, err := pathmatch.Compile("/users/{user_id}/cards/{fruit_id}")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.BadRequest:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalError:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pathmatch.Compile() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
//	
//	var userId string
//	var cardId string
//	
//	didMatch, err := pattern.Find("/users/8sN.oP/cards/X3j_T4", userId, cardId)
//	if nil != err {
//		switch err.(type) { // ← Note that we are using a Go type-switch here.
//	
//		case pathmatch.BadRequest: // ← Here we are detecting if the error returned is a pathmatch.BadRequest.
//	
//			fmt.Printf("Something you did when you called pattern.Find() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalError:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pattern.Find() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
//
// Note that one can get more specific than just a BadRequest. For example:
// NotEnoughArguments, PatternSyntaxError, UnsupportedArgumentType,
// and StructFieldWrongType.
//
// To be able to detect those more specific error types, put them BEFORE the "case pathmatch.BadRequest:"
// in the type switch. For example:
//
//	pattern, err := pathmatch.Compile("/users/{user_id}/cards/{fruit_id}")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.PatternSyntaxError: // ← Here we are detecting if the error returned was due to a syntax error, in the uncompiled pattern. Also note that it comes BEFORE the 'pathmatch.BadRequest' case; THAT IS IMPORTANT!
//	
//			fmt.Printf("The uncompiled pattern passed to pathmatch.Compile() had a syntax error in it. The error message describing the syntax error is....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.BadRequest:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalError:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pathmatch.Compile() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
type BadRequest interface {
	error
	BadRequest()
}
