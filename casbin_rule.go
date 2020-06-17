package goroseadapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const CRUD string = "/api/casbin_rule"

type Client struct {
	ApiKey string
	Host   string
	Path   string
	Legion string
}

type CasbinRule struct {
	Id     int64
	P_type string
	V0     string
	V1     string
	V2     string
	V3     string
	V4     string
	V5     string
	Legion string
}

func New(apiKey, apiHost, apiLegion string) *Client {
	if apiKey == "" {
		log.Fatal("API_KEY is not specified")
	}
	if apiHost == "" {
		log.Fatal("API_HOST is not specified")
	}
	if apiLegion == "" {
		log.Fatal("API_LEGION is not specified")
	}
	return &Client{
		ApiKey: apiKey,
		Host:   apiHost,
		Path:   apiHost + CRUD,
		Legion: apiLegion,
	}
}

func NewFromEnvionment() *Client {
	return New(
		os.Getenv("API_KEY"),
		os.Getenv("API_HOST"),
		os.Getenv("API_LEGION"),
	)
}

func LoadParams(rule *CasbinRule) string {
	params := url.Values{}
	if rule.Id != 0 {
		id := strconv.FormatInt(rule.Id, 10)
		params.Set("id", id)
	}
	if rule.P_type != "" {
		params.Set("p_type", rule.P_type)
	}
	if rule.V0 != "" {
		params.Set("v0", rule.V0)
	}
	if rule.V1 != "" {
		params.Set("v1", rule.V1)
	}
	if rule.V2 != "" {
		params.Set("v2", rule.V2)
	}
	if rule.V3 != "" {
		params.Set("v3", rule.V3)
	}
	if rule.V4 != "" {
		params.Set("v4", rule.V4)
	}
	if rule.V5 != "" {
		params.Set("v5", rule.V5)
	}
	return params.Encode()
}

func LoadParamsWhere(rule *CasbinRule) string {
	_where := make([]string, 0)
	if rule.Id != 0 {
		id := strconv.FormatInt(rule.Id, 10)
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "id", id))
	}
	if rule.P_type != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "p_type", rule.P_type))
	}
	if rule.V0 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v0", rule.V0))
	}
	if rule.V1 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v1", rule.V1))
	}
	if rule.V2 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v2", rule.V2))
	}
	if rule.V3 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v3", rule.V3))
	}
	if rule.V4 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v4", rule.V4))
	}
	if rule.V5 != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "v5", rule.V5))
	}
	if rule.Legion != "" {
		_where = append(_where, fmt.Sprintf("(%s,eq,%s)", "legion", rule.Legion))
	}
	if len(_where) > 0 {
		return "_where=" + strings.Join(_where, "~and")
	}
	return ""
}

func (c *Client) Create(rule *CasbinRule) ([]byte, error) {
	rule.Legion = c.Legion
	marshal, err2 := json.Marshal(rule)
	if err2 != nil {
		return nil, err2
	}
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, c.Path, strings.NewReader(string(marshal)))
	if err != nil {
		return nil, err
	}
	reqest.Header.Add("api-key", c.ApiKey)
	reqest.Header.Add("Content-Type", "application/json")
	response, err := client.Do(reqest)
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (c *Client) Retrieve(rule *CasbinRule) ([]byte, error) {
	rule.Legion = c.Legion
	Url, err := url.Parse(c.Path)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = LoadParamsWhere(rule)
	urlPath := Url.String()
	fmt.Println(urlPath)
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, err
	}
	reqest.Header.Add("api-key", c.ApiKey)
	reqest.Header.Add("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (c *Client) Update(rule *CasbinRule) {

}

func (c *Client) Delete(ids string) ([]byte, error) {
	urlPath := c.Path + "/bulk?_ids=" + ids
	fmt.Println(urlPath)
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}
	reqest.Header.Add("api-key", c.ApiKey)
	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
