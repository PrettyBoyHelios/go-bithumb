package go_bithumb

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const(
	BASE_URL = "https://global-openapi.bithumb.pro/openapi/v1"
)

type Client struct {
	apiKey string
	secretKey string
	client http.Client
	url string
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) getSha256HashCode(preSign string) string {
	h := hmac.New(sha256.New, []byte(c.secretKey))
	h.Write([]byte(preSign))
	hashCode := hex.EncodeToString(h.Sum(nil))
	return hashCode
}

func (c *Client) sign(preMap map[string]string) string {
	var keys []string
	for k := range preMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var preSign string
	for _, k := range keys {
		preSign += k + "=" + preMap[k] + "&"
	}
	preSign = strings.TrimSuffix(preSign, "&")
	fmt.Println("prepare signature string >======= ", preSign)
	signature := c.getSha256HashCode(preSign)
	fmt.Println("signature string >====== ", signature)
	return signature
}

func (c *Client) post(url string, params interface{}, result interface{}) error {
	preMap := c.struct2map(params)
	preMap["apiKey"] = c.apiKey
	preMap["timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	preMap["signature"] = c.sign(preMap)
	err := post(url, preMap, result)
	handleErr(err)
	return err
}

func (c *Client) struct2map(params interface{}) map[string]string {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	}
	return data
}

func (c *Client) get(url string, r interface{}) error {
	resp := doGet(url)
	err := doParse(resp, r)
	return err
}

func doGet(url string) []byte {
	resp, err := http.Get(url)
	return handleResp(resp, err)
}

func post(url string, params interface{}, r interface{}) error {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp := doPost(url, jsonBytes)
	nil := doParse(resp, r)
	return nil
}

func doParse(resp []byte, in interface{}) error {
	err := json.Unmarshal(resp, in)
	if err != nil {
		return err
	}
	return nil
}

func doPost(url string, data []byte) []byte {
	body := bytes.NewReader(data)
	resp, err := http.Post(url, "application/json", body)
	return handleResp(resp, err)
}

func handleResp(resp *http.Response, err error) []byte {
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return r
}
