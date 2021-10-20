package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidLen      = errors.New("invalid len")
	ErrInvalidTag      = errors.New("invalid tag")
	ErrValueNotInSet   = errors.New("value not in set")
	ErrInvalidRegexp   = errors.New("invalid regexp")
	ErrNoMatchRegexp   = errors.New("value not match regexp")
	ErrUnsupportedType = errors.New("field type unsupported")
	ErrValueLessMin    = errors.New("value is less than minimal")
	ErrValueGreatMax   = errors.New("value is greater than maximal")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) > 0 {
		s := "validation errors:\n"
		for _, e := range v {
			s = s + e.Field + ": " + e.Err.Error() + "\n"
		}
		return s
	}
	return ""
}

func Validate(v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Struct {
		err := errors.New("variable is not struct")
		return err
	}

	errs := ValidationErrors{}
	value := reflect.ValueOf(v)
	valueType := value.Type()

	for i := 0; i < valueType.NumField(); i++ {
		fieldTypeValue := valueType.Field(i)
		fieldValue := value.Field(i)
		tagBody := fieldTypeValue.Tag.Get("validate")

		if len(tagBody) == 0 {
			continue
		}

		tags := strings.Split(tagBody, "|")

		for _, tag := range tags {
			if fieldValue.Kind() != reflect.Slice {
				switch fieldValue.Kind() {
				case reflect.String:
					errs = validateString(fieldTypeValue.Name, fieldValue.String(), tag, errs)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					errs = validateInt(fieldTypeValue.Name, int(fieldValue.Int()), tag, errs)
				default:
					errs = append(errs, ValidationError{Field: fieldTypeValue.Name, Err: ErrUnsupportedType})
				}
			} else {
				switch fieldValue.Type().String() {
				case "[]string":
					errs = validateStrings(fieldTypeValue.Name, fieldValue.Interface().([]string), tag, errs)
				case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
					errs = validateInts(fieldTypeValue.Name, fieldValue.Interface().([]int), tag, errs)
				default:
					errs = append(errs, ValidationError{Field: fieldTypeValue.Name, Err: ErrUnsupportedType})
				}
			}
		}
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}

func validateString(fieldName string, fieldValue string, tag string, errs ValidationErrors) ValidationErrors {
	tagParams := strings.Split(tag, ":")
	if len(tagParams) != 2 {
		return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
	}

	switch tagParams[0] {
	case "len":
		{
			lenValue, err := strconv.Atoi(tagParams[1])
			if err != nil {
				return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
			}
			if len(fieldValue) != lenValue {
				errs = append(errs, ValidationError{Field: fieldName, Err: ErrInvalidLen})
			}
		}
	case "in":
		{
			if !strings.Contains(tag, fieldValue) {
				errs = append(errs, ValidationError{Field: fieldName, Err: ErrValueNotInSet})
			}
		}
	case "regexp":
		{
			re, err := regexp.Compile(tagParams[1])
			if err != nil {
				return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidRegexp})
			}

			if !re.MatchString(fieldValue) {
				errs = append(errs, ValidationError{Field: fieldName, Err: ErrNoMatchRegexp})
			}
		}
	default:
		return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
	}

	return errs
}

func validateStrings(fieldName string, fieldValue []string, tag string, errs ValidationErrors) ValidationErrors {
	for _, v := range fieldValue {
		errs = validateString(fieldName, v, tag, errs)
	}
	return errs
}

func validateInt(fieldName string, fieldValue int, tag string, errs ValidationErrors) ValidationErrors {
	tagParams := strings.Split(tag, ":")
	if len(tagParams) != 2 {
		return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
	}

	switch tagParams[0] {
	case "min":
		{
			minValue, err := strconv.Atoi(tagParams[1])
			if err != nil {
				return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
			}

			if fieldValue < minValue {
				errs = append(errs, ValidationError{Field: fieldName, Err: ErrValueLessMin})
			}
		}
	case "max":
		{
			maxValue, err := strconv.Atoi(tagParams[1])
			if err != nil {
				return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
			}

			if fieldValue > maxValue {
				errs = append(errs, ValidationError{Field: fieldName, Err: ErrValueGreatMax})
			}
		}
	case "in":
		{
			set := strings.Split(tagParams[1], ",")

			for _, val := range set {
				i, err := strconv.Atoi(val)
				if err != nil {
					return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
				}

				if fieldValue == i {
					return errs
				}
			}

			return append(errs, ValidationError{Field: fieldName, Err: ErrValueNotInSet})
		}
	default:
		return append(errs, ValidationError{Field: fieldName, Err: ErrInvalidTag})
	}

	return errs
}

func validateInts(fieldName string, fieldValue []int, tag string, errs ValidationErrors) ValidationErrors {
	for _, v := range fieldValue {
		errs = validateInt(fieldName, v, tag, errs)
	}
	return errs
}
