package rosetta

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	client http.Client
)

type Client struct {
	addr       string
	apiKey     string
	blockchain string
	network    string
}

func New(addr string, apiKey string) *Client {
	c := &Client{addr, apiKey, "", ""}
	return c
}

func (c *Client) SetBlockchain(chain string) {
	c.blockchain = chain
}
func (c *Client) SetNetwork(network string) {
	c.network = network
}

func (c *Client) callApi(path string, body interface{}) ([]byte, error) {
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	url := c.addr + path
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New("Failed request: " + resp.Status)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (c *Client) NetworkStatus() (*NetworkStatus, error) {
	rBody := RosettaRequest{
		NetworkIdentifier: NetworkIdentifier{
			Blockchain: c.blockchain,
			Network:    c.network,
		},
	}

	var result NetworkStatus
	b, err := c.callApi("/network/status", rBody)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Block(blockIdentifier BlockIdentifier) (*BlockResponse, error) {
	rBody := RosettaRequest{
		NetworkIdentifier: NetworkIdentifier{
			Blockchain: c.blockchain,
			Network:    c.network,
		},
		BlockIdentifier: blockIdentifier,
	}

	var result BlockResponse
	b, err := c.callApi("/block", rBody)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Transaction(blockIdentifier BlockIdentifier, txIdentifier TransactionIdentifier) (*BlockTransactionResponse, error) {
	rBody := RosettaRequest{
		NetworkIdentifier: NetworkIdentifier{
			Blockchain: c.blockchain,
			Network:    c.network,
		},
		BlockIdentifier:       blockIdentifier,
		TransactionIdentifier: txIdentifier,
	}

	var result BlockTransactionResponse
	b, err := c.callApi("/block/transaction", rBody)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
