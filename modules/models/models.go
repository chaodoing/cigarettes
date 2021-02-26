package models

import (
	"errors"
	"github.com/chaodoing/cigarettes/providers/utils"
	"github.com/gookit/validate"
	"gorm.io/gorm"
	"strings"
	"time"
)

// Datum [...]
type Datum struct {
	ID       uint       `gorm:"primaryKey;column:id;type:int unsigned;not null" json:"-"` // 主键
	UUID     string     `gorm:"column:uuid;type:varchar(128);not null" json:"uuid"`       // 唯一uuid
	Data     string     `gorm:"column:data;type:json" json:"data"`                        // 数据
	CreateAt utils.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"` // 创建时间
	UpdateAt utils.Time `gorm:"column:update_at;type:datetime" json:"update_at"`          // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Datum) TableName() string {
	return "datum"
}

type _Datum struct {
	Db        *gorm.DB
	TableName string
}

// Validate 数据验证
//  @param data 要验证的结构体
//  @param scene 验证场景 login|password|update|create|delete|disable
//  @return hasErr 是否有错误
//  @return messages 错误消息
func (this *_Datum) Validate(data *Datum, scene string) (hasErr bool, messages validate.Errors) {
	if strings.EqualFold(scene, "create") {
		data.CreateAt = utils.Time(time.Now())
	}
	data.UpdateAt = utils.Time(time.Now())
	valid := validate.Struct(data)
	valid.WithScenarios(map[string][]string{
		"create": {"UUID", "Data"},
		"delete": {"ID"},
	})
	return !valid.AtScene(scene).Validate(), valid.Errors
}

// Add 添加数据
//  @param data 表单数据
//  @return data 添加完成后有主键的数据
//  @return error 是否有错误
func (this *_Datum) Add(data Datum) (Datum, int64, error) {
	var find Datum
	err := this.Db.Table(this.TableName).Where("`uuid`=?", data.UUID).Find(&find).Error
	if err == nil {
		if find.ID > 0 {
			return Datum{}, 0, errors.New("用户手机号码已经存在")
		}
	}
	create := this.Db.Table(this.TableName).Create(&data)
	return data, create.RowsAffected, create.Error
}

// Delete 删除数据
//  @param data 表单数据
//  @return error 是否有错误
func (this *_Datum) Delete(data Datum) error {
	return this.Db.Delete(&data).Error
}
// FindByUUID 通过uuid查找
func (this *_Datum) FindByUUID(uuid string) (err error, find *Datum) {
	err = this.Db.Table(this.TableName).Where("`uuid`=?", uuid).Find(&find).Error
	return err, find
}