package pro

import (
	"testing"

	"github.com/photoprism/photoprism/internal/form"
	"github.com/stretchr/testify/assert"
)

func TestNewFeedback(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		feedback := NewFeedback("xxx")
		assert.Equal(t, "xxx", feedback.ClientVersion)
	})
}

func TestSendFeedback(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := NewConfig("0.0.0", "testdata/new.yml")

		feedback := Feedback{
			Category:      "Bug Report",
			Subject:       "",
			Message:       "I found a new bug",
			UserName:      "Test User",
			UserEmail:     "test@example.com",
			UserAgent:     "",
			ApiKey:        "123456",
			ClientVersion: "0.0.0",
			ClientOS:      "linux",
			ClientArch:    "amd64",
			ClientCPU:     2,
		}

		feedbackForm, err := form.NewFeedback(feedback)

		if err != nil {
			t.Fatal(err)
		}

		err2 := c.SendFeedback(feedbackForm)
		assert.Contains(t, err2.Error(), "failed")
	})
}
