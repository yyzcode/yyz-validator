package validator

import (
	"errors"
	"fmt"
)

type float32Validator struct {
	field    string
	value    float32
	handlers []func(float32) error
	errStr   string
}

func (validator *float32Validator) push(f func(float32) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *float32Validator) Require() *float32Validator {
	validator.push(func(i float32) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *float32Validator) Between(min, max float32) *float32Validator {
	validator.push(func(i float32) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%g到%g之间", min, max)
	})
	return validator
}

func (validator *float32Validator) In(arr ...float32) *float32Validator {
	validator.push(func(i float32) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + float32sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *float32Validator) NotIn(arr ...float32) *float32Validator {
	validator.push(func(i float32) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + float32sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *float32Validator) Gt(min float32) *float32Validator {
	validator.push(func(i float32) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%g", min)
	})
	return validator
}

func (validator *float32Validator) Lt(max float32) *float32Validator {
	validator.push(func(i float32) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%g", max)
	})
	return validator
}

func (validator *float32Validator) Gte(min float32) *float32Validator {
	validator.push(func(i float32) error {
		if i < min {
			return fmt.Errorf("不能小于%g", min)
		}
		return nil
	})
	return validator
}

func (validator *float32Validator) Lte(max float32) *float32Validator {
	validator.push(func(i float32) error {
		if i > max {
			return fmt.Errorf("不能大于%g", max)
		}
		return nil
	})
	return validator
}

func (validator *float32Validator) AddRule(f func(float32) error) *float32Validator {
	validator.push(f)
	return validator
}

func (validator *float32Validator) Exec() (err error) {
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

func (validator *float32Validator) Err(s string) *float32Validator {
	validator.errStr = s
	return validator
}
