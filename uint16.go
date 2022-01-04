package validator

import (
	"errors"
	"fmt"
)

type uint16Validator struct {
	field    string
	value    uint16
	handlers []func(uint16) error
	errStr   string
	force    bool
}

func (validator *uint16Validator) push(f func(uint16) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *uint16Validator) Require() *uint16Validator {
	validator.force = true
	validator.push(func(i uint16) error {
		if i == 0 {
			return errors.New(validator.field + "不能为空")
		}
		return nil
	})
	return validator
}

func (validator *uint16Validator) Between(min, max uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *uint16Validator) In(arr ...uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New(validator.field + "必须在" + uint16sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *uint16Validator) NotIn(arr ...uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		for _, v := range arr {
			if v == i {
				return errors.New(validator.field + "不能在" + uint16sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *uint16Validator) Gt(min uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		if i > min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须大于%d", min)
	})
	return validator
}

func (validator *uint16Validator) Lt(max uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		if i < max {
			return nil
		}
		return fmt.Errorf(validator.field+"必须小于%d", max)
	})
	return validator
}

func (validator *uint16Validator) Gte(min uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		if i < min {
			return fmt.Errorf(validator.field+"不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *uint16Validator) Lte(max uint16) *uint16Validator {
	validator.push(func(i uint16) error {
		if i > max {
			return fmt.Errorf(validator.field+"不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *uint16Validator) AddRule(f func(uint16) error) *uint16Validator {
	validator.push(f)
	return validator
}

func (validator *uint16Validator) Exec() (err error) {
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

func (validator *uint16Validator) Err(s string) *uint16Validator {
	validator.errStr = s
	return validator
}
