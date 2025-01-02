package matrix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pazifical/onyx/logging"
)

type Service struct {
	credentials Credentials
	roomID      string
	accessToken string
}

func NewService(credentials Credentials, roomID string) Service {
	return Service{
		credentials: credentials,
		roomID:      roomID,
	}
}

func (s *Service) Authenticate() error {
	endpoint := "https://matrix.org/_matrix/client/v3/login"

	loginRequestData := LoginRequestData{
		Type:     "m.login.password",
		Password: s.credentials.Password,
		Identifier: UserData{
			Type: "m.id.user",
			User: s.credentials.Username,
		},
	}

	data, err := json.Marshal(loginRequestData)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		logging.Warning(fmt.Sprintf("response status %d from Matrix", response.StatusCode))

		var errorData interface{}
		err = json.NewDecoder(response.Body).Decode(&errorData)
		if err != nil {
			return err
		}

		jsonData, err := json.Marshal(errorData)
		if err != nil {
			return err
		}
		return errors.New(string(jsonData))
	}

	var loginResponseData LoginResponseData
	err = json.NewDecoder(response.Body).Decode(&loginResponseData)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	s.accessToken = loginResponseData.AccessToken

	return nil
}

func (s *Service) SendMessage(text string) error {
	endpoint := fmt.Sprintf("https://matrix.org/_matrix/client/v3/rooms/%s/send/m.room.message", s.roomID)

	message := MessageData{
		MsgType: "m.text",
		Body:    text,
	}

	data, err := json.Marshal(message)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		logging.Warning(fmt.Sprintf("response status %d from Matrix", response.StatusCode))

		var errorData interface{}
		err = json.NewDecoder(response.Body).Decode(&errorData)
		if err != nil {
			return err
		}

		jsonData, err := json.Marshal(errorData)
		if err != nil {
			return err
		}
		return errors.New(string(jsonData))
	}

	return nil
}
