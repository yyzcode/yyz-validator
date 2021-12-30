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
