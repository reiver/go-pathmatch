package pathmatch


// BadRequestComplainer is used to represent one of the types of errors that could be returned when
// calling the pathmatch.Compile func, the pathmatch.Pattern.Match method, or the pathmatch.Pattern.MatchAndLoad
// method. The meaning of this type of error is that the problem was due to something whomever called the func or method did.
//
// For example, maybe the uncompiled pattern passed to the pathmatch.Compile() func had
// a syntax error in it. Or, also for example, maybe the type of parameter passed to
// pathmatch.Pattern.Match() was of the wrong type. Etc.
//
// Example usage of BadRequestComplainer with pathmatch.Compile():
//
//	pattern, err := pathmatch.Compile("/fruits/{there_is_an_error_in_here/") // ← The uncompiled pattern there has an error in it.
//	if nil != err {
//		switch err.(type) { // ← Note that we are using a Go type-switch here.
//	
//		case pathmatch.BadRequestComplainer: // ← Here we are detecting if the error returned is a pathmatch.BadRequestComplainer.
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalErrorComplainer:
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
// use the pathmatch.Pattern.Match() method, which could also generate an error that fits the BadRequestComplainer
// interface.
//
// Example usage of BadRequestComplainer with pathmatch.Pattern.Match():
//
//	pattern, err := pathmatch.Compile("/users/{user_id}/cards/{fruit_id}")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.BadRequestComplainer:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalErrorComplainer:
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
//	didMatch, err := pattern.Match("/users/8sN.oP/cards/X3j_T4", userId, cardId)
//	if nil != err {
//		switch err.(type) { // ← Note that we are using a Go type-switch here.
//	
//		case pathmatch.BadRequestComplainer: // ← Here we are detecting if the error returned is a pathmatch.BadRequestComplainer.
//	
//			fmt.Printf("Something you did when you called pattern.Match() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalErrorComplainer:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pattern.Match() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
//
// Note that one can get more specific than just a BadRequestComplainer. For example:
// NotEnoughArgumentsComplainer, PatternSyntaxErrorComplainer, UnsupportedArgumentType,
// and StructFieldWrongTypeComplainer.
//
// To be able to detect those more specific error types, put them BEFORE the "case pathmatch.BadRequestComplainer:"
// in the type switch. For example:
//
//	pattern, err := pathmatch.Compile("/users/{user_id}/cards/{fruit_id}")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.PatternSyntaxErrorComplainer: // ← Here we are detecting if the error returned was due to a syntax error, in the uncompiled pattern. Also note that it comes BEFORE the 'pathmatch.BadRequestComplainer' case; THAT IS IMPORTANT!
//	
//			fmt.Printf("The uncompiled pattern passed to pathmatch.Compile() had a syntax error in it. The error message describing the syntax error is....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.BadRequestComplainer:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalErrorComplainer:
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
type BadRequestComplainer interface {
	error
	BadRequestComplainer()
}
