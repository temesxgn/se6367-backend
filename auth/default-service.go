package auth

import (
	"bytes"
	"fmt"
	"github.com/temesxgn/se6367-backend/auth/models"
	"github.com/temesxgn/se6367-backend/util/jsonutils"
	"gopkg.in/auth0.v3"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/temesxgn/se6367-backend/config"
	"gopkg.in/auth0.v3/management"
)

var (
	service *auth0Service
	once    sync.Once
)

func initialize() {
	once.Do(func() {
		m, err := management.New(config.GetAuth0Domain(), config.GetAuth0ClientID(), config.GetAuth0ClientSecret())
		if err != nil {

		}

		service = &auth0Service{
			client: m,
		}
	})
}

func NewDefaultService() *auth0Service {
	initialize()
	return service
}

type auth0Service struct {
	client *management.Management
}

func (s *auth0Service) GetToken() (string, error) {
	body, _ := jsonutils.Marshal(models.NewAuth0TokenRequest(config.GetAuth0Domain(), config.GetAuth0ClientID(), config.GetAuth0ClientSecret()))
	res, err := http.Post(fmt.Sprintf("%voauth/token", config.GetAuth0Domain()), "application/json", bytes.NewReader([]byte(body)))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response models.Auth0ClientCredentialsTokenResponse
	if err := jsonutils.Unmarshal(string(respBody), response); err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func (s *auth0Service) GetUser(userID string) (*management.User, error) {
	return s.client.User.Read(userID)
}

func (s *auth0Service) CreateUser(connection, email string) error {
	usr := &management.User{
		Connection: auth0.String(connection),
		Email:      auth0.String(email),
	}

	return s.client.User.Create(usr)
}
