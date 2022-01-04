package validator

import (
	"errors"
	"fmt"
)

type intValidator struct {
	field    string
	value    int
	handlers []func(int) error
	errStr   string
	force    bool //是否强制验证，为false时遇到零值会直接跳过验证
}

func (validator *intValidator) push(f func(int) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *intValidator) Require() *intValidator {
	validator.force = true
	validator.push(func(i int) error {
		if i == 0 {
			return errors.New(validator.field + "不能为空")
		}
		return nil
	})
	return validator
}

func (validator *intValidator) Between(min, max int) *intValidator {
	validator.push(func(i int) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *intValidator) In(arr ...int) *intValidator {
	validator.push(func(i int) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New(validator.field + "必须在" + intsJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *intValidator) NotIn(arr ...int) *intValidator {
	validator.push(func(i int) error {
		for _, v := range arr {
			if v == i {
				return errors.New(validator.field + "不能在" + intsJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *intValidator) Gt(min int) *intValidator {
	validator.push(func(i int) error {
		if i > min {
			return nil
		}
		return fmt.Errorf(validator.field+"必须大于%d", min)
	})
	return validator
}

func (validator *intValidator) Lt(max int) *intValidator {
	validator.push(func(i int) error {
		if i < max {
			return nil
		}
		return fmt.Errorf(validator.field+"必须小于%d", max)
	})
	return validator
}

func (validator *intValidator) Gte(min int) *intValidator {
	validator.push(func(i int) error {
		if i < min {
			return fmt.Errorf(validator.field+"不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *intValidator) Lte(max int) *intValidator {
	validator.push(func(i int) error {
		if i > max {
			return fmt.Errorf(validator.field+"不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *intValidator) AddRule(f func(int) error) *intValidator {
	validator.push(f)
	return validator
}

func (validator *intValidator) Exec() (err error) {
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

func (validator *intValidator) Err(s string) *intValidator {
	validator.errStr = s
	return validator
}
