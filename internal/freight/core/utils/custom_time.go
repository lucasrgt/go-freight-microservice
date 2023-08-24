package utils

import "time"

type CustomTime time.Time

const layout = "2006-01-02T15:04"

func (customTime *CustomTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(layout, string(b))
	if err != nil {
		return err
	}
	*customTime = CustomTime(t)
	return nil
}
