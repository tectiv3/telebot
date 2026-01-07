package telebot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_HasTopicsEnabled(t *testing.T) {
	data := `{
		"id": 123456789,
		"first_name": "Test",
		"username": "testuser",
		"has_topics_enabled": true
	}`

	var user User
	err := json.Unmarshal([]byte(data), &user)
	assert.NoError(t, err)
	assert.True(t, user.HasTopicsEnabled)
	assert.Equal(t, int64(123456789), user.ID)
	assert.Equal(t, "Test", user.FirstName)
	assert.Equal(t, "testuser", user.Username)
}

func TestUser_HasTopicsEnabledFalse(t *testing.T) {
	data := `{
		"id": 987654321,
		"first_name": "Test2",
		"has_topics_enabled": false
	}`

	var user User
	err := json.Unmarshal([]byte(data), &user)
	assert.NoError(t, err)
	assert.False(t, user.HasTopicsEnabled)
}

func TestUser_HasTopicsEnabledOmitted(t *testing.T) {
	data := `{
		"id": 111222333,
		"first_name": "Test3"
	}`

	var user User
	err := json.Unmarshal([]byte(data), &user)
	assert.NoError(t, err)
	assert.False(t, user.HasTopicsEnabled)
}

func TestMessage_ThreadIDInPrivateChat(t *testing.T) {
	data := `{
		"message_id": 1,
		"date": 1234567890,
		"chat": {"id": 123, "type": "private"},
		"from": {"id": 456, "first_name": "Test"},
		"message_thread_id": 42,
		"is_topic_message": true,
		"text": "Hello in topic"
	}`

	var msg Message
	err := json.Unmarshal([]byte(data), &msg)
	assert.NoError(t, err)
	assert.Equal(t, 42, msg.ThreadID)
	assert.True(t, msg.TopicMessage)
	assert.Equal(t, "Hello in topic", msg.Text)
}
