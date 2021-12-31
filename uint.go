package validator

import (
	"errors"
	"fmt"
)

type uintValidator struct {
	field    string
	value    uint
	handlers []func(uint) error
	errStr   string
	force    bool
}

func (validator *uintValidator) push(f func(uint) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *uintValidator) Require() *uintValidator {
	validator.force = true
	validator.push(func(i uint) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *uintValidator) Between(min, max uint) *uintValidator {
	validator.push(func(i uint) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *uintValidator) In(arr ...uint) *uintValidator {
	validator.push(func(i uint) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + uintsJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *uintValidator) NotIn(arr ...uint) *uintValidator {
	validator.push(func(i uint) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + uintsJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *uintValidator) Gt(min uint) *uintValidator {
	validator.push(func(i uint) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *uintValidator) Lt(max uint) *uintValidator {
	validator.push(func(i uint) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *uintValidator) Gte(min uint) *uintValidator {
	validator.push(func(i uint) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *uintValidator) Lte(max uint) *uintValidator {
	validator.push(func(i uint) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *uintValidator) AddRule(f func(uint) error) *uintValidator {
	validator.push(f)
	return validator
}

func (validator *uintValidator) Exec() (err error) {
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

func (validator *uintValidator) Err(s string) *uintValidator {
	validator.errStr = s
	return validator
}
