package validator

import (
	"errors"
	"fmt"
)

type float64Validator struct {
	field    string
	value    float64
	handlers []func(float64) error
	errStr   string
}

func (validator *float64Validator) push(f func(float64) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *float64Validator) Require() *float64Validator {
	validator.push(func(i float64) error {
		if i == 0 {
			return errors.New("不能为空")
		}
		return nil
	})
	return validator
}

func (validator *float64Validator) Between(min, max float64) *float64Validator {
	validator.push(func(i float64) error {
		if i <= max && i >= min {
			return nil
		}
		return fmt.Errorf("必须在%g到%g之间", min, max)
	})
	return validator
}

func (validator *float64Validator) In(arr ...float64) *float64Validator {
	validator.push(func(i float64) error {
		for _, v := range arr {
			if v == i {
				return nil
			}
		}
		return errors.New("必须在" + float64sJoin(arr, ",") + "之中")
	})
	return validator
}

func (validator *float64Validator) NotIn(arr ...float64) *float64Validator {
	validator.push(func(i float64) error {
		for _, v := range arr {
			if v == i {
				return errors.New("不能在" + float64sJoin(arr, ",") + "之中")
			}
		}
		return nil
	})
	return validator
}

func (validator *float64Validator) Gt(min float64) *float64Validator {
	validator.push(func(i float64) error {
		if i > min {
			return nil
		}
		return fmt.Errorf("必须大于%g", min)
	})
	return validator
}

func (validator *float64Validator) Lt(max float64) *float64Validator {
	validator.push(func(i float64) error {
		if i < max {
			return nil
		}
		return fmt.Errorf("必须小于%g", max)
	})
	return validator
}

func (validator *float64Validator) Gte(min float64) *float64Validator {
	validator.push(func(i float64) error {
		if i < min {
			return fmt.Errorf("不能小于%g", min)
		}
		return nil
	})
	return validator
}

func (validator *float64Validator) Lte(max float64) *float64Validator {
	validator.push(func(i float64) error {
		if i > max {
			return fmt.Errorf("不能大于%g", max)
		}
		return nil
	})
	return validator
}

func (validator *float64Validator) AddRule(f func(float64) error) *float64Validator {
	validator.push(f)
	return validator
}

func (validator *float64Validator) Exec() (err error) {
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

func (validator *float64Validator) Err(s string) *float64Validator {
	validator.errStr = s
	return validator
}
