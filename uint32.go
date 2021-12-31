package validator

import (
	"errors"
	"fmt"
)

type uint32Validator struct {
	field    string
	value    uint32
	handlers []func(uint32) error
	errStr   string
	force    bool
}

func (validator *uint32Validator) push(f func(uint32) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *uint32Validator) Require() *uint32Validator {
	validator.force = true
	validator.push(func(i uint32) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *uint32Validator) Between(min, max uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *uint32Validator) In(arr ...uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + uint32sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *uint32Validator) NotIn(arr ...uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + uint32sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *uint32Validator) Gt(min uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *uint32Validator) Lt(max uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *uint32Validator) Gte(min uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *uint32Validator) Lte(max uint32) *uint32Validator {
	validator.push(func(i uint32) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *uint32Validator) AddRule(f func(uint32) error) *uint32Validator {
	validator.push(f)
	return validator
}

func (validator *uint32Validator) Exec() (err error) {
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

func (validator *uint32Validator) Err(s string) *uint32Validator {
	validator.errStr = s
	return validator
}
