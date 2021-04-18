package entity

import (
	"encoding/json"

	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
)

type UserProfile user_managementpb.UserProfile

// Marshall ..
func (t *UserProfile) Marshall() ([]byte, error) {
	return json.Marshal(t)
}

// Unmarshall ..
func (t *UserProfile) Unmarshall(bytes []byte) error {
	return json.Unmarshal(bytes, t)
}
