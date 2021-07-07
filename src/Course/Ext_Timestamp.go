package Course

import (
	"database/sql/driver"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"
)

func(this *Timestamp) Scan(value interface{}) error {
	switch t:=value.(type) {
	case time.Time:
		var err error
		this.Timestamp,err = ptypes.TimestampProto(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("timestamp error")
	}
	return nil
}

func (this Timestamp) Value() (driver.Value,error) {
	return ptypes.Timestamp(this.Timestamp)
}
