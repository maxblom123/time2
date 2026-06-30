package time2

import "time"

func (t T) InLocation(name string) (T, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return t, err
	}
	return T{t.Time.In(loc)}, nil
}

func (t T) MustInLocation(name string) T {
	result, err := t.InLocation(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t T) UTC() T           { return T{t.Time.UTC()} }
func (t T) Local() T         { return T{t.Time.Local()} }
func (t T) OffsetHours() int { _, offset := t.Zone(); return offset / 3600 }