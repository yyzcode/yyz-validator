package validator

import (
	"errors"
	"fmt"
)

type int16Validator struct {
	field    string
	value    int16
	handlers []func(int16) error
	errStr   string
	force    bool
}

func (validator *int16Validator) push(f func(int16) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *int16Validator) Require() *int16Validator {
	validator.force = true
	validator.push(func(i int16) error {
		if i == 0 {
			return errors.New(validator.field + "不能为空")
		}
		return nil
	})
	return validator
}

func (validator *int16Validator) Between(min, max int16) *int16Validator {
	validator.push(func(i int16) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *int16Validator) In(arr ...int16) *int16Validator {
	validator.push(func(i int16) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New(validator.field + "必须在" + int16sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *int16Validator) NotIn(arr ...int16) *int16Validator {
	validator.push(func(i int16) error {
		for _, v := range arr {
			if v == i {
				return errors.New(validator.field + "不能在" + int16sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *int16Validator) Gt(min int16) *int16Validator {
	validator.push(func(i int16) error {
		if i > min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须大于%d", min)
	})
	return validator
}

func (validator *int16Validator) Lt(max int16) *int16Validator {
	validator.push(func(i int16) error {
		if i < max {
			return nil
		}
		return fmt.Errorf(validator.field+"必须小于%d", max)
	})
	return validator
}

func (validator *int16Validator) Gte(min int16) *int16Validator {
	validator.push(func(i int16) error {
		if i < min {
			return fmt.Errorf(validator.field+"不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *int16Validator) Lte(max int16) *int16Validator {
	validator.push(func(i int16) error {
		if i > max {
			return fmt.Errorf(validator.field+"不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *int16Validator) AddRule(f func(int16) error) *int16Validator {
	validator.push(f)
	return validator
}

func (validator *int16Validator) Exec() (err error) {
	if !validator.force && validator.value == 0 {
		return nil
	}
	for i := 0; i < len(validator.handlers); i++ {
		err = validator.handlers[i](validator.value)
		if err != nil {
			if validator.errStr == "" {
				return
			}
			return errors.New(validator.errStr)
		}
	}
	return nil
}

func (validator *int16Validator) Err(s string) *int16Validator {
	validator.errStr = s
	return validator
}
