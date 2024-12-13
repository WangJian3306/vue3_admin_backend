package timestamp

import (
	"time"
)

type TimeStamp time.Time

var TimeFormats = []string{"2006-01-02 15:04:05", "20060102150405"}

func (t *TimeStamp) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = TimeStamp(time.Time{})
		return
	}

	var now time.Time
	for _, format := range TimeFormats {
		// 指定解析的格式
		if now, err = time.ParseInLocation(format, string(data), time.Local); err == nil {
			*t = TimeStamp(now)
			return
		}
		// 指定解析的格式
		if now, err = time.ParseInLocation(`"`+format+`"`, string(data), time.Local); err == nil {
			*t = TimeStamp(now)
			return
		}
	}
	return
}
func (t TimeStamp) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormats[0])+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormats[0])
	b = append(b, '"')
	return b, nil
}

func (t TimeStamp) String() string {
	return time.Time(t).Format(TimeFormats[0])
}
