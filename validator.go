package validator

type validator interface {
	Exec() error
}

type InternalError struct {
	ErrInfo string
}

func (err InternalError) Error() string {
	return err.ErrInfo
}

func Validate(validators ...validator) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = InternalError{e.(string)}
		}
	}()
	for i := 0; i < len(validators); i++ {
		err = validators[i].Exec()
		if err != nil {
			return
		}
	}
	return nil
}

func Int(value int, field string) *intValidator {
	return &intValidator{
		field: field,
		value: value,
	}
}

func Uint(value uint, field string) *uintValidator {
	return &uintValidator{
		field: field,
		value: value,
	}
}

func Int8(value int8, field string) *int8Validator {
	return &int8Validator{
		field: field,
		value: value,
	}
}

func Uint8(value uint8, field string) *uint8Validator {
	return &uint8Validator{
		field: field,
		value: value,
	}
}

func Int16(value int16, field string) *int16Validator {
	return &int16Validator{
		field: field,
		value: value,
	}
}

func Uint16(value uint16, field string) *uint16Validator {
	return &uint16Validator{
		field: field,
		value: value,
	}
}

func Int32(value int32, field string) *int32Validator {
	return &int32Validator{
		field: field,
		value: value,
	}
}

func Uint32(value uint32, field string) *uint32Validator {
	return &uint32Validator{
		field: field,
		value: value,
	}
}

func Int64(value int64, field string) *int64Validator {
	return &int64Validator{
		field: field,
		value: value,
	}
}

func Uint64(value uint64, field string) *uint64Validator {
	return &uint64Validator{
		field: field,
		value: value,
	}
}

func Float32(value float32, field string) *float32Validator {
	return &float32Validator{
		field: field,
		value: value,
	}
}

func Float64(value float64, field string) *float64Validator {
	return &float64Validator{
		field: field,
		value: value,
	}
}

func String(value string, field string) *stringValidator {
	return &stringValidator{
		field: field,
		value: value,
	}
}

func Func(field string) *fnValidator {
	return &fnValidator{
		field: field,
	}
}
