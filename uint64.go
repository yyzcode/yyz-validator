package validator

import (
	"errors"
	"fmt"
)

type uint64Validator struct {
	field    string
	value    uint64
	handlers []func(uint64) error
	errStr   string
	force    bool
}

func (validator *uint64Validator) push(f func(uint64) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *uint64Validator) Require() *uint64Validator {
	validator.force = true
	validator.push(func(i uint64) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *uint64Validator) Between(min, max uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *uint64Validator) In(arr ...uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + uint64sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *uint64Validator) NotIn(arr ...uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + uint64sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *uint64Validator) Gt(min uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *uint64Validator) Lt(max uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *uint64Validator) Gte(min uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *uint64Validator) Lte(max uint64) *uint64Validator {
	validator.push(func(i uint64) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *uint64Validator) AddRule(f func(uint64) error) *uint64Validator {
	validator.push(f)
	return validator
}

func (validator *uint64Validator) Exec() (err error) {
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

func (validator *uint64Validator) Err(s string) *uint64Validator {
	validator.errStr = s
	return validator
}
