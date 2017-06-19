package fluent

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
)

const (
	tagField     = "tag"
	messageField = "message"
)

// Hook implements a Logrus hook for Fluent.
type Hook struct {
	Fluent     fluent.Fluent
	Tag        string
	DefaultTag string
}

// Fire is invoked by logrus and sends logs to Fluentd.
func (h *Hook) Fire(entry *logrus.Entry) error {
	tag := h.getTag(entry)

	data := entry.Data

	if _, ok := data[messageField]; !ok {
		data[messageField] = entry.Message
	}

	return h.Fluent.PostWithTime(tag, entry.Time, data)
}

// getTag finds the appropriate tag.
// Order of detection:
// 1. Value of Tag field
// 2. Tag defined in context
// 3. Value of DefaultTag field
func (h *Hook) getTag(entry *logrus.Entry) string {
	if h.Tag != "" {
		return h.Tag
	}

	if tag, ok := entry.Data[tagField].(string); ok {
		delete(entry.Data, tagField)

		return tag
	}

	return h.DefaultTag
}
