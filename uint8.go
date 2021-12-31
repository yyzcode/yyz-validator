package validator

import (
	"errors"
	"fmt"
)

type uint8Validator struct {
	field    string
	value    uint8
	handlers []func(uint8) error
	errStr   string
}

func (validator *uint8Validator) push(f func(uint8) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *uint8Validator) Require() *uint8Validator {
	validator.push(func(i uint8) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *uint8Validator) Between(min, max uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *uint8Validator) In(arr ...uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + uint8sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *uint8Validator) NotIn(arr ...uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + uint8sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *uint8Validator) Gt(min uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *uint8Validator) Lt(max uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *uint8Validator) Gte(min uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *uint8Validator) Lte(max uint8) *uint8Validator {
	validator.push(func(i uint8) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *uint8Validator) AddRule(f func(uint8) error) *uint8Validator {
	validator.push(f)
	return validator
}

func (validator *uint8Validator) Exec() (err error) {
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

func (validator *uint8Validator) Err(s string) *uint8Validator {
	validator.errStr = s
	return validator
}
