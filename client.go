package cloudsign

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

var (
	// ErrorNoClientID is a error indicating given client id is empty.
	ErrorNoClientID = fmt.Errorf("client should have a client id")
)

const (
	apiBasePath        = "api.cloudsign.jp"
	apiSandBoxBasePath = "api-sandbox.cloudsign.jp"
)

const reqProtocol = "https"

// Client represents a api client for cloudsign api.
type Client struct {
	clientID    string
	baseURL     string
	httpClient  *http.Client
	accessToken *AccessToken
	logger      *log.Logger
	isSandBox   bool
}

// NewClient returns a new api client.
func NewClient(clientID string, logger *log.Logger, isSandbox bool) (*Client, error) {
	if len(clientID) <= 0 {
		return nil, ErrorNoClientID
	}

	c := new(Client)
	c.clientID = clientID
	if logger == nil {
		c.logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	c.httpClient = http.DefaultClient
	c.isSandBox = isSandbox
	if c.isSandBox {
		c.baseURL = apiSandBoxBasePath
	} else {
		c.baseURL = apiBasePath
	}

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method, endpoint string, postForm url.Values) (*http.Request, error) {
	urlPath := path.Join(c.baseURL, endpoint)
	url := fmt.Sprintf("%s://%s", reqProtocol, urlPath)

	req, err := http.NewRequest(method, url, strings.NewReader(postForm.Encode()))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if c.accessToken != nil {
		headerValue := fmt.Sprintf("Bearer %s", c.accessToken.AccessToken)
		req.Header.Set("Authorization", headerValue)
	}

	return req, nil
}

func (c *Client) decodeResponse(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// IssueAccessToken makes a post request to access token issurance api.
func (c *Client) IssueAccessToken(ctx context.Context) (*AccessToken, error) {
	reqForm := url.Values{}
	reqForm.Add("client_id", c.clientID)
	req, err := c.newRequest(ctx, "POST", "/token", reqForm)
	if err != nil {
		return nil, fmt.Errorf("failed to create requerst object for POST /token")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request POST /token: response:%s, %w", "aaa", err)
	}
	if resp.StatusCode > http.StatusSeeOther {
		return nil, fmt.Errorf("failld to get access token status code : %s", resp.Status)
	}

	accessToken := new(AccessToken)
	err = c.decodeResponse(resp, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respnse of POST /token %w", err)
	}
	accessToken.CreatedAt = time.Now()
	c.accessToken = accessToken

	return accessToken, nil
}

// PostDocument is endpoint for create cloudsign document see: https://app.swaggerhub.com/apis/CloudSign/cloudsign-web_api/0.13.0#/default/post_documents
func (c *Client) PostDocument(ctx context.Context, postDocumentReq *PostDocumentRequest) (*Document, error) {
	return nil, nil
}
