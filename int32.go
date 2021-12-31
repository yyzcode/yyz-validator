package validator

import (
	"errors"
	"fmt"
)

type int32Validator struct {
	field    string
	value    int32
	handlers []func(int32) error
	errStr   string
	force    bool
}

func (validator *int32Validator) push(f func(int32) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *int32Validator) Require() *int32Validator {
	validator.force = true
	validator.push(func(i int32) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *int32Validator) Between(min, max int32) *int32Validator {
	validator.push(func(i int32) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *int32Validator) In(arr ...int32) *int32Validator {
	validator.push(func(i int32) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + int32sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *int32Validator) NotIn(arr ...int32) *int32Validator {
	validator.push(func(i int32) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + int32sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *int32Validator) Gt(min int32) *int32Validator {
	validator.push(func(i int32) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%d", min)
	})
	return validator
}

func (validator *int32Validator) Lt(max int32) *int32Validator {
	validator.push(func(i int32) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%d", max)
	})
	return validator
}

func (validator *int32Validator) Gte(min int32) *int32Validator {
	validator.push(func(i int32) error {
		if i < min {
			return fmt.Errorf("不能小于%d", min)
		}
		return nil
	})
	return validator
}

func (validator *int32Validator) Lte(max int32) *int32Validator {
	validator.push(func(i int32) error {
		if i > max {
			return fmt.Errorf("不能大于%d", max)
		}
		return nil
	})
	return validator
}

func (validator *int32Validator) AddRule(f func(int32) error) *int32Validator {
	validator.push(f)
	return validator
}

func (validator *int32Validator) Exec() (err error) {
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

func (validator *int32Validator) Err(s string) *int32Validator {
	validator.errStr = s
	return validator
}
