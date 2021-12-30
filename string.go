package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexpRules = map[string]*regexp.Regexp{
	"number":      regexp.MustCompile("^[\\-0-9][0-9]+(.[0-9]+)?$"),                                                                       //正数、负数、小数
	"integer":     regexp.MustCompile("^\\-?[0-9]+$"),                                                                                     //整数
	"alpha":       regexp.MustCompile("^[A-Za-z]+$"),                                                                                      //纯字母
	"alphaNum":    regexp.MustCompile("^[A-Za-z0-9]+$"),                                                                                   //字母和数字
	"alphaDash":   regexp.MustCompile("^[A-Za-z0-9\\-\\_]+$"),                                                                             //字母和数字，下划线_及破折号-
	"chs":         regexp.MustCompile("^[\\p{Han}]+$"),                                                                                    //汉字
	"chsAlpha":    regexp.MustCompile("^[\\p{Han}a-zA-Z]+$"),                                                                              //汉字、字母
	"chsAlphaNum": regexp.MustCompile("^[\\p{Han}a-zA-Z0-9]+$"),                                                                           //汉字、字母和数字
	"chsDash":     regexp.MustCompile("^[\\p{Han}a-zA-Z0-9\\-\\_]+$"),                                                                     //汉字、字母、数字和下划线_及破折号-
	"phone":       regexp.MustCompile("^1[3-9]\\d{9}$"),                                                                                   //手机号
	"idCard":      regexp.MustCompile("^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"), //身份证
	"zip":         regexp.MustCompile("\\d{6}"),                                                                                           //邮编
	"ip":          regexp.MustCompile("^((([01]{0,1}\\d{0,1}\\d|2[0-4]\\d|25[0-5])\\.){3}([01]{0,1}\\d{0,1}\\d|2[0-4]\\d|25[0-5]))$"),     //ip地址
}

type stringValidator struct {
	field    string
	value    string
	handlers []func(string) error
	errStr   string
}

func (validator *stringValidator) push(f func(string) error) {
	validator.handlers = append(validator.handlers, f)
}

func (validator *stringValidator) Require() *stringValidator {
	validator.push(func(s string) error {
		if s != "" {
			return nil
		}
		return errors.New("不能为空")
	})
	return validator
}

func (validator *stringValidator) Length(min, max int) *stringValidator {
	if max <= 0 {
		validator.push(func(s string) error {
			if len(s) > min {
				return nil
			}
			return fmt.Errorf("长度不能小于%d", min)
		})
	}
	if min <= 0 {
		validator.push(func(s string) error {
			if len(s) < max {
				return nil
			}
			return fmt.Errorf("长度不能超过%d", min)
		})
	}
	validator.push(func(s string) error {
		if len(s) >= min && len(s) <= max {
			return nil
		}
		return fmt.Errorf("长度必须在%d到%d之间", min, max)
	})
	return validator
}

func (validator *stringValidator) In(arr ...string) *stringValidator {
	validator.push(func(s string) error {
		for i := 0; i < len(arr); i++ {
			if arr[i] == s {
				return nil
			}
		}
		return fmt.Errorf("必须在%s之间", strings.Join(arr, ","))
	})
	return validator
}

func (validator *stringValidator) NotIn(arr ...string) *stringValidator {
	validator.push(func(s string) error {
		for i := 0; i < len(arr); i++ {
			if arr[i] == s {
				return fmt.Errorf("不能在%s之间", strings.Join(arr, ","))
			}
		}
		return nil
	})
	return validator
}

func (validator *stringValidator) Number() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["number"].MatchString(s) {
			return nil
		}
		return errors.New("必须是数字")
	})
	return validator
}

func (validator *stringValidator) Integer() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["integer"].MatchString(s) {
			return nil
		}
		return errors.New("必须是整数")
	})
	return validator
}

func (validator *stringValidator) Gt(min float64) *stringValidator {
	validator.push(func(s string) error {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err.Error())
		}
		if f > min {
			return nil
		}
		return fmt.Errorf("必须大于%g", min)
	})
	return validator
}

func (validator *stringValidator) Lt(max float64) *stringValidator {
	validator.push(func(s string) error {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err.Error())
		}
		if f < max {
			return nil
		}
		return fmt.Errorf("必须小于%g", max)
	})
	return validator
}

func (validator *stringValidator) Gte(min float64) *stringValidator {
	validator.push(func(s string) error {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err.Error())
		}
		if f > min {
			return nil
		}
		return fmt.Errorf("不能小于%g", min)
	})
	return validator
}

func (validator *stringValidator) Lte(max float64) *stringValidator {
	validator.push(func(s string) error {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err.Error())
		}
		if f < max {
			return nil
		}
		return fmt.Errorf("不能大于%g", max)
	})
	return validator
}

func (validator *stringValidator) Alpha() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["alpha"].MatchString(s) {
			return nil
		}
		return errors.New("只能包含字母")
	})
	return validator
}

func (validator *stringValidator) AlphaNum() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["alphaNum"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("只能包含字母和数字")
	})
	return validator
}

func (validator *stringValidator) AlphaDash() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["alphaDash"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("只能包含字母、数字和-_")
	})
	return validator
}

func (validator *stringValidator) Chs() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["chs"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("只能包含中文")
	})
	return validator
}

func (validator *stringValidator) ChsAlpha() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["chsAlpha"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("只能包含中文和字母")
	})
	return validator
}

func (validator *stringValidator) ChsAlphaNum() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["chsAlphaNum"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("只能包含中文、字母和数字")
	})
	return validator
}

func (validator *stringValidator) ChsDash() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["chsDash"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("不能包含特殊字符")
	})
	return validator
}

func (validator *stringValidator) Phone() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["phone"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("不是一个有效的手机号")
	})
	return validator
}

func (validator *stringValidator) Zip() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["zip"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("不是一个有效的邮编")
	})
	return validator
}

func (validator *stringValidator) Ip() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["ip"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("不是一个有效的ip")
	})
	return validator
}

func (validator *stringValidator) IdCard() *stringValidator {
	validator.push(func(s string) error {
		if regexpRules["idCard"].MatchString(s) {
			return nil
		}
		return fmt.Errorf("不是一个有效的身份证号")
	})
	return validator
}

func (validator *stringValidator) AddRule(f func(string) error) *stringValidator {
	validator.push(f)
	return validator
}

func (validator *stringValidator) Exec() (err error) {
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

func (validator *stringValidator) Err(s string) *stringValidator {
	validator.errStr = s
	return validator
}
