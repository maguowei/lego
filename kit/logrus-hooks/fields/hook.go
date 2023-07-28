package fields

import (
	"github.com/sirupsen/logrus"
)

type Hook struct {
	logrus.Hook
	fields logrus.Fields
	levels []logrus.Level
}

func NewHook(fields logrus.Fields, levels []logrus.Level) *Hook {
	return &Hook{
		fields: fields,
		levels: levels,
	}
}

func (f *Hook) Levels() []logrus.Level {
	return f.levels
}

func (f *Hook) Fire(entry *logrus.Entry) error {
	for k, v := range f.fields {
		entry.Data[k] = v
	}
	return nil
}
