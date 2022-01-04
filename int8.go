package validator

import (
	"errors"
	"fmt"
)

type int8Validator struct {
	field    string
	value    int8
	handlers []func(int8) error
	errStr   string
	force    bool
}

func (validator *int8Validator) push(f func(int8) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *int8Validator) Require() *int8Validator {
	validator.force = true
	validator.push(func(i int8) error {
		if i == 0 {
			return errors.New(validator.field + "不能为空")
		}
		return nil
	})
	return validator
}

func (validator *int8Validator) Between(min, max int8) *int8Validator {
	validator.push(func(i int8) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *int8Validator) In(arr ...int8) *int8Validator {
	validator.push(func(i int8) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New(validator.field + "必须在" + int8sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *int8Validator) NotIn(arr ...int8) *int8Validator {
	validator.push(func(i int8) error {
		for _, v := range arr {
			if v == i {
				return errors.New(validator.field + "不能在" + int8sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *int8Validator) Gt(min int8) *int8Validator {
	validator.push(func(i int8) error {
		if i > min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须大于%d", min)
	})
	return validator
}

func (validator *int8Validator) Lt(max int8) *int8Validator {
	validator.push(func(i int8) error {
		if i < max {
			return nil
		}
		return fmt.Errorf(validator.field+"必须小于%d", max)
	})
	return validator
}

func (validator *int8Validator) Gte(min int8) *int8Validator {
	validator.push(func(i int8) error {
		if i < min {
			return fmt.Errorf(validator.field+"不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *int8Validator) Lte(max int8) *int8Validator {
	validator.push(func(i int8) error {
		if i > max {
			return fmt.Errorf(validator.field+"不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *int8Validator) AddRule(f func(int8) error) *int8Validator {
	validator.push(f)
	return validator
}

func (validator *int8Validator) Exec() (err error) {
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

func (validator *int8Validator) Err(s string) *int8Validator {
	validator.errStr = s
	return validator
}
