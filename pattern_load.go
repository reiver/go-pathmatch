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


func (pattern *internalPattern) MatchAndLoad(path string, strct interface{}) (bool, error) {

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

	reflectedValue := reflect.ValueOf(strct)
	if reflect.Ptr != reflectedValue.Kind() {
		return doesNotMatter, errExpectedAPointerToAStruct
	}

	reflectedValueElem := reflectedValue.Elem()

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
