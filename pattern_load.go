package pathmatch


import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)


var (
	errExpectedAPointerToAStruct = newUnsupportedArgumentType("Expected a pointer to a struct, but wasn't.")
)


// FindAndLoad compares ‘path’ against its (compiled) pattern template; if it matches
// it loads the matches into ‘dest’, and then returns true.
//
// ‘dest’ can be a pointer struct, or a pointer to a []string.
//
// Find may set some, or all of the items or fields in ‘dest’ even if it returns false, and even if it returns an error.
func (pattern *Pattern) FindAndLoad(path string, dest interface{}) (bool, error) {
	if nil == pattern {
		return false, errNilReceiver
	}

	pattern.mutex.RLock()
	defer pattern.mutex.RUnlock()

//@TODO: Is it a good idea to be dynamically creating this?
//@TODO: Also, can the struct fields be put in here directly instead?
	args := []interface{}{}
	numNames := len(pattern.MatchNames())
	for i:=0; i<numNames; i++ {
		args = append(args, new(string))
	}

	didMatch, err := pattern.Find(path, args...)
	if nil != err {
		return doesNotMatter, err
	}

	if !didMatch {
		return false, nil
	}

	reflectedValue := reflect.ValueOf(dest)
	if reflect.Ptr != reflectedValue.Kind() {
//@TODO: change error
		return doesNotMatter, errExpectedAPointerToAStruct
	}

	reflectedValueElem := reflectedValue.Elem()
	switch reflectedValueElem.Kind() {
	case reflect.Slice:
		var a []string = make([]string, len(args))
		for i, arg := range args {
			a[i] = *(arg.(*string))
		}

		return loadSlice(dest, a...)
	case reflect.Struct:
		return pattern.loadStruct(reflectedValueElem, args)
	default:
//@TODO: change error
		return doesNotMatter, errExpectedAPointerToAStruct
	}
}

func loadSlice(dest interface{}, matches ...string) (bool, error) {
	if nil == dest {
		return false, errNilTarget
	}

	target, casted := dest.(*[]string)
	if !casted {
//@TODO: CHANGE ERROR! ============================
		return false, errExpectedAPointerToAStruct
	}
	if nil == target {
		return false, errNilTarget
	}

	*target = (*target)[:0]
	for _, match := range matches {
		*target = append(*target, match)
	}

	return true, nil
}

func (pattern *Pattern) loadStruct(reflectedValueElem reflect.Value, args []interface{}) (bool, error) {
	if nil == pattern {
		return false, errNilReceiver
	}

	if reflect.Struct != reflectedValueElem.Kind() {
		return doesNotMatter, errExpectedAPointerToAStruct
	}

	reflectedValueElemType := reflectedValueElem.Type()

	numFields := reflectedValueElemType.NumField()
	for fieldNumber:=0; fieldNumber<numFields; fieldNumber++ {
		//field := reflectedValueElemType.Field(fieldNumber)

		//fieldTag := field.Tag

		//name := fieldTag.Get(pattern.fieldTagName)

		value := *(args[fieldNumber].(*string))

		err := func(rValue reflect.Value, value string, matchName string) (err error) {

			defer func() {

				if r := recover(); nil != r {
					// See if we received a message of the form:
					//
					//	reflect.Set: value of type ??? is not assignable to type ???
					//
					// If we did then we interpret this as the programmer using this
					// trying to load into a struct field of the wrong type.
					//
					// We return a special error for that.
					if s, ok := r.(string); ok {
						needle := "reflect.Set: value of type "

						if strings.HasPrefix(s, needle) {
							needle = " is not assignable to type "

							if strings.Contains(s, needle) {
								err = newStructFieldWrongType(matchName)
								return
							}
						}
					}

					msg := fmt.Sprintf("%T %v", r, r)

					err = errors.New( msg )
					return
				}

			}()

			rValue.Set( reflect.ValueOf(value) )

			return nil
		}(reflectedValueElem.Field(fieldNumber), value, pattern.fieldTagName)
		if nil != err {
			return doesNotMatter, err
		}
	}

	return true, nil
}
