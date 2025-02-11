package auth0

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	model2 "github.com/temesxgn/se6367-backend/auth/model"
	"github.com/temesxgn/se6367-backend/common/util/jsonutils"
	"github.com/temesxgn/se6367-backend/graph/model"
	"gopkg.in/auth0.v3"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/temesxgn/se6367-backend/config"
	"gopkg.in/auth0.v3/management"
)

var (
	service *auth0Service
	once    sync.Once
)

func initialize() error {
	var initErr error
	once.Do(func() {
		m, err := management.New(config.GetAuth0DomainName(), config.GetAuth0ClientID(), config.GetAuth0ClientSecret())
		if err != nil {
			initErr = err
			return
		}

		service = &auth0Service{
			client: m,
		}
	})

	return initErr
}

func NewService() (*auth0Service, error) {
	if err := initialize(); err != nil {
		return nil, err
	}

	return service, nil
}

type auth0Service struct {
	client *management.Management
}

func (s *auth0Service) GetToken() (string, error) {
	auth0Req, _ := jsonutils.Marshal(model2.NewAuth0TokenRequest(config.GetAuth0APIID(), config.GetAuth0ClientID(), config.GetAuth0ClientSecret()))
	url := fmt.Sprintf("%voauth/token", config.GetAuth0Domain())
	payload := strings.NewReader(auth0Req)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data model2.Auth0ClientCredentialsTokenResponse
	_ = jsonutils.Unmarshal(string(body), &data)

	return data.AccessToken, nil
}

func (s *auth0Service) GetUser(userID string) (*model.Auth0Profile, error) {
	token, err := s.GetToken()
	if err != nil {
		return nil, err
	}

	url := "https://fairbankz.auth0.com/api/v2/users/" + userID
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, errors.Wrap(err, "reading body")
	}

	var user model.Auth0Profile
	if err := jsonutils.Unmarshal(buf.String(), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *auth0Service) UpdateProfile(userID string, data *model.UpdateAuth0Profile) error {
	token, err := s.GetToken()
	if err != nil {
		return err
	}

	url := "https://fairbankz.auth0.com/api/v2/users/" + userID
	bdy, _ := jsonutils.Marshal(data)
	fmt.Println("Sending RQ: " + bdy)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(bdy)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return errors.Wrap(err, "reading body")
	}

	var user model.Auth0Profile
	fmt.Println("RS: " + buf.String())
	if err := jsonutils.Unmarshal(buf.String(), &user); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%v", user))
	return nil
}

func (s *auth0Service) CreateUser(connection, email string) error {
	usr := &management.User{
		Connection: auth0.String(connection),
		Email:      auth0.String(email),
	}

	return s.client.User.Create(usr)
}
