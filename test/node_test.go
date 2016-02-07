package test

import (
	"encoding/json"
	"strconv"

	req "github.com/parnurzeal/gorequest"
	. "gopkg.in/check.v1"
)

// http://stackoverflow.com/a/9573928
type TorConnection struct {
	Port        int `json:"port,int"`
	ControlPort int `json:"control_port,int"`
}

func createNode(endpoint string) (req.Response, TorConnection) {
	resp, body, err := req.New().Post(endpoint + "/node/create").End()

	if err != nil {
		panic(err)
	}

	var torConn TorConnection
	marshErr := json.Unmarshal([]byte(body), &torConn)

	if marshErr != nil {
		panic(err)
	}

	return resp, torConn
}

func (f *TestSuite) TestCreateNode(c *C) {
	response, _ := createNode(f.Endpoint)

	c.Assert(response.StatusCode, Equals, 201)
}

func (f *TestSuite) TestDeleteNode(c *C) {
	response, connection := createNode(f.Endpoint)

	c.Assert(response.StatusCode, Equals, 201)

	// Send our delete request
	resp, _, err := req.New().
		Delete(f.Endpoint + "/node/" + strconv.Itoa(connection.ControlPort)).
		End()

	if err != nil {
		panic(err)
	}

	c.Assert(resp.StatusCode, Equals, 204)

}
