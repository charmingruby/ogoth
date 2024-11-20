package model

import "encoding/json"

type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func JSONToUserModel(data []byte) (*User, error) {
	var user User

	err := json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UserModelToJSON(user *User) ([]byte, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	return data, nil
}
