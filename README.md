## 自用库go数据验证器
数据在验证前应该已经是确定的数据类型了，所以不必使用interface后做大量的反射或是断言。每种基础数据类型int，int8，int32,uint,uint8，float32，string等等都提供了专用的数据类型验证器接口。虽然这样做可能会产生较多的冗余代码，但效率高比反射高，理论上代码也更健壮。

### 安装
    go get -u github.com/yyzcoder/yyz-validator

### 示例

```go
package main

import (
	validator "github.com/yyzcoder/yyz-validator"
	"gorm.io/gorm"
	"time"
	"fmt"
)

type SysRole struct {
	Id               int
	Name             string `json:"name"`
	Desc             string
	Creator          int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt      `json:"-"`
	SysPermissions   []SysRolePermission `json:"SysPermissions,omitempty"`
	SysPermissionIds []int               `gorm:"-" json:"SysPermissionIds,omitempty"`
}

func (s *SysRole) Valid() error {
	err := validator.Validate(
		//验证int类型数据，必填（数据的零值会被拦截）且大于等于1
		validator.Int(s.Creator, "创建人").Require().Gte(1),
		//验证string类型数据，必填且长度大于2且为汉字和字母数字，自定义错误信息
		validator.String(s.Name, "角色名").Require().Length(2, 0).ChsAlphaNum().Err("角色名称不合法"),
		//验证string类型数据，非必填（如果为零值直接跳过验证）且长度小于255且为纯字母
		validator.String(s.Desc, "角色描述").Length(0, 255).Alpha(),
		//使用自定义函数进行验证
		validator.Func("角色权限").AddRule(s.permissionValid),
	)
	return err
}
/**
 *可以使用自定义函数可以完成嵌套规则验证，或者完成自定义数据类型验证
 */
func (s *SysRole) permissionValid() error {
	if len(s.SysPermissionIds) == 0 {
		return fmt.Errorf("角色权限不能为空")
	}
	return nil
}

func main(){
	sysRole := &sysRole{
		Id:1,
		Name:"角色名",
		Desc:"角色功能职责描述",
		Creator:1,
	}
	err := sysRole.Valid()
	if err != nil {
		if e, ok := err.(validator.InternalError); ok {
			//验证器内部错误（runtime错误）
			panic(e.Error())
		}
		fmt.Println(err) //数据验证未通过
	}
}
```
