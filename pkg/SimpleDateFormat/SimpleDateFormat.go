package SimpleDateFormat

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// SimpleDateFormat 自定义一个类似Java的时间类型
// 2023-01-01 12:00:00
type SimpleDateFormat struct {
	time.Time
}

// MarshalJSON 实现 JSON 序列化
func (t *SimpleDateFormat) MarshalJSON() ([]byte, error) {
	if t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现 JSON 反序列化
func (t *SimpleDateFormat) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	parsedTime, err := time.Parse("\"2006-01-02 15:04:05\"", string(data))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// Value 实现数据库驱动接口
func (t SimpleDateFormat) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现数据库扫描接口
func (t *SimpleDateFormat) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		t.Time = v
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}
