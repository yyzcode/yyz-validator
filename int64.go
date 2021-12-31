package validator

import (
	"errors"
	"fmt"
)

type int64Validator struct {
	field    string
	value    int64
	handlers []func(int64) error
	errStr   string
	force    bool
}

func (validator *int64Validator) push(f func(int64) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *int64Validator) Require() *int64Validator {
	validator.force = true
	validator.push(func(i int64) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *int64Validator) Between(min, max int64) *int64Validator {
	validator.push(func(i int64) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *int64Validator) In(arr ...int64) *int64Validator {
	validator.push(func(i int64) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + int64sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *int64Validator) NotIn(arr ...int64) *int64Validator {
	validator.push(func(i int64) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + int64sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *int64Validator) Gt(min int64) *int64Validator {
	validator.push(func(i int64) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *int64Validator) Lt(max int64) *int64Validator {
	validator.push(func(i int64) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *int64Validator) Gte(min int64) *int64Validator {
	validator.push(func(i int64) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *int64Validator) Lte(max int64) *int64Validator {
	validator.push(func(i int64) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *int64Validator) AddRule(f func(int64) error) *int64Validator {
	validator.push(f)
	return validator
}

func (validator *int64Validator) Exec() (err error) {
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

func (validator *int64Validator) Err(s string) *int64Validator {
	validator.errStr = s
	return validator
}
