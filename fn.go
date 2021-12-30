package validator

import "errors"

type fnValidator struct {
	field    string
	handlers []func() error
	errStr   string
}

func (validator *fnValidator) push(f func() error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *fnValidator) AddRule(f func() error) *fnValidator {
	validator.push(f)
	return validator
}

func (validator *fnValidator) Exec() (err error) {
	for i := 0; i < len(validator.handlers); i++ {
		err = validator.handlers[i]()
		if err != nil {
			if validator.errStr == "" {
				return
			}
			return errors.New(validator.errStr)
		}
	}
	return nil
}

func (validator *fnValidator) Err(s string) *fnValidator {
	validator.errStr = s
	return validator
}
