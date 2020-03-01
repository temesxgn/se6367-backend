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
			fmt.Println("ERROR INIT auth0 management api: " + err.Error())
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
		fmt.Println("ERROR init auth0 service: " + err.Error())
		return nil, err
	}

	return service, nil
}

type auth0Service struct {
	client *management.Management
}

func (s *auth0Service) GetToken() (string, error) {
	body, _ := jsonutils.Marshal(model2.NewAuth0TokenRequest(config.GetAuth0Domain(), config.GetAuth0ClientID(), config.GetAuth0ClientSecret()))
	res, err := http.Post(fmt.Sprintf("%voauth/token", config.GetAuth0Domain()), "application/json", bytes.NewReader([]byte(body)))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response model2.Auth0ClientCredentialsTokenResponse
	if err := jsonutils.Unmarshal(string(respBody), response); err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func (s *auth0Service) GetUser(userID string) (*model.Auth0User, error) {
	url := "https://fairbankz.auth0.com/api/v2/users/" + userID
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlJVTTFPVFk0TTBVM09FSXdOREV4UlRkQ01ESXlPRFpGTXpjMlJUTkZRMFl5TnpBNU1ESTJNQSJ9.eyJpc3MiOiJodHRwczovL2ZhaXJiYW5rei5hdXRoMC5jb20vIiwic3ViIjoiSWdBUHRtUm9YUjRKR3Z0Ym5pTUdxNkZmZ3U4NTFiMXpAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vZmFpcmJhbmt6LmF1dGgwLmNvbS9hcGkvdjIvIiwiaWF0IjoxNTgyOTk3MDgwLCJleHAiOjE1ODMwODM0ODAsImF6cCI6IklnQVB0bVJvWFI0Skd2dGJuaU1HcTZGZmd1ODUxYjF6Iiwic2NvcGUiOiJyZWFkOmNsaWVudF9ncmFudHMgY3JlYXRlOmNsaWVudF9ncmFudHMgZGVsZXRlOmNsaWVudF9ncmFudHMgdXBkYXRlOmNsaWVudF9ncmFudHMgcmVhZDp1c2VycyB1cGRhdGU6dXNlcnMgZGVsZXRlOnVzZXJzIGNyZWF0ZTp1c2VycyByZWFkOnVzZXJzX2FwcF9tZXRhZGF0YSB1cGRhdGU6dXNlcnNfYXBwX21ldGFkYXRhIGRlbGV0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgY3JlYXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBjcmVhdGU6dXNlcl90aWNrZXRzIHJlYWQ6Y2xpZW50cyB1cGRhdGU6Y2xpZW50cyBkZWxldGU6Y2xpZW50cyBjcmVhdGU6Y2xpZW50cyByZWFkOmNsaWVudF9rZXlzIHVwZGF0ZTpjbGllbnRfa2V5cyBkZWxldGU6Y2xpZW50X2tleXMgY3JlYXRlOmNsaWVudF9rZXlzIHJlYWQ6Y29ubmVjdGlvbnMgdXBkYXRlOmNvbm5lY3Rpb25zIGRlbGV0ZTpjb25uZWN0aW9ucyBjcmVhdGU6Y29ubmVjdGlvbnMgcmVhZDpyZXNvdXJjZV9zZXJ2ZXJzIHVwZGF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGRlbGV0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGNyZWF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIHJlYWQ6ZGV2aWNlX2NyZWRlbnRpYWxzIHVwZGF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgZGVsZXRlOmRldmljZV9jcmVkZW50aWFscyBjcmVhdGU6ZGV2aWNlX2NyZWRlbnRpYWxzIHJlYWQ6cnVsZXMgdXBkYXRlOnJ1bGVzIGRlbGV0ZTpydWxlcyBjcmVhdGU6cnVsZXMgcmVhZDpydWxlc19jb25maWdzIHVwZGF0ZTpydWxlc19jb25maWdzIGRlbGV0ZTpydWxlc19jb25maWdzIHJlYWQ6aG9va3MgdXBkYXRlOmhvb2tzIGRlbGV0ZTpob29rcyBjcmVhdGU6aG9va3MgcmVhZDplbWFpbF9wcm92aWRlciB1cGRhdGU6ZW1haWxfcHJvdmlkZXIgZGVsZXRlOmVtYWlsX3Byb3ZpZGVyIGNyZWF0ZTplbWFpbF9wcm92aWRlciBibGFja2xpc3Q6dG9rZW5zIHJlYWQ6c3RhdHMgcmVhZDp0ZW5hbnRfc2V0dGluZ3MgdXBkYXRlOnRlbmFudF9zZXR0aW5ncyByZWFkOmxvZ3MgcmVhZDpzaGllbGRzIGNyZWF0ZTpzaGllbGRzIGRlbGV0ZTpzaGllbGRzIHJlYWQ6YW5vbWFseV9ibG9ja3MgZGVsZXRlOmFub21hbHlfYmxvY2tzIHVwZGF0ZTp0cmlnZ2VycyByZWFkOnRyaWdnZXJzIHJlYWQ6Z3JhbnRzIGRlbGV0ZTpncmFudHMgcmVhZDpndWFyZGlhbl9mYWN0b3JzIHVwZGF0ZTpndWFyZGlhbl9mYWN0b3JzIHJlYWQ6Z3VhcmRpYW5fZW5yb2xsbWVudHMgZGVsZXRlOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGNyZWF0ZTpndWFyZGlhbl9lbnJvbGxtZW50X3RpY2tldHMgcmVhZDp1c2VyX2lkcF90b2tlbnMgY3JlYXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgZGVsZXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgcmVhZDpjdXN0b21fZG9tYWlucyBkZWxldGU6Y3VzdG9tX2RvbWFpbnMgY3JlYXRlOmN1c3RvbV9kb21haW5zIHJlYWQ6ZW1haWxfdGVtcGxhdGVzIGNyZWF0ZTplbWFpbF90ZW1wbGF0ZXMgdXBkYXRlOmVtYWlsX3RlbXBsYXRlcyByZWFkOm1mYV9wb2xpY2llcyB1cGRhdGU6bWZhX3BvbGljaWVzIHJlYWQ6cm9sZXMgY3JlYXRlOnJvbGVzIGRlbGV0ZTpyb2xlcyB1cGRhdGU6cm9sZXMgcmVhZDpwcm9tcHRzIHVwZGF0ZTpwcm9tcHRzIHJlYWQ6YnJhbmRpbmcgdXBkYXRlOmJyYW5kaW5nIHJlYWQ6bG9nX3N0cmVhbXMgY3JlYXRlOmxvZ19zdHJlYW1zIGRlbGV0ZTpsb2dfc3RyZWFtcyB1cGRhdGU6bG9nX3N0cmVhbXMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.BJRv78wc9O1wqBa8AluauJhSYrR6eeRng12DlZfdOycPHJkHOmMtXwhbPKRUQ4-taeXIDk9Lhv2Iil655vgWlDJGa-3jEG-vcE3Zqp64IEo_Z5ra7UV0wY8UltpHya-lSo7vL2Y4uZcITAnPZOARbIXYmZ9ArwuiLTef0n9hyMBAeOrWZxlw-6ii2iR_nDdvNgvKsbR9XxkMdKuV8S0p6Y_DlJT1M3ADfDMEV4pRggUYFu8G2xer-3LpRz_SHF9pG2zOwMbwkEBIhqwTYR5P4IjuUJW0NuJq5zc9Pe4EsNbpmNRz1WrI53TMBuAvkkGco89yinrCzwvTKw5u9YeA8Q")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, errors.Wrap(err, "reading body")
	}

	var user model.Auth0User
	if err := jsonutils.Unmarshal(buf.String(), &user); err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *auth0Service) CreateUser(connection, email string) error {
	usr := &management.User{
		Connection: auth0.String(connection),
		Email:      auth0.String(email),
	}

	return s.client.User.Create(usr)
}
