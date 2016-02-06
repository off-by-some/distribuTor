package test

import (
	req "github.com/parnurzeal/gorequest"
	. "gopkg.in/check.v1"
)

func (f *TestSuite) TestCreateNode(c *C) {
	resp, _, err := req.New().Post(f.Endpoint + "/node/create").End()

	if err != nil {
		panic(err)
	}

	c.Assert(resp.StatusCode, Equals, 201)
}
