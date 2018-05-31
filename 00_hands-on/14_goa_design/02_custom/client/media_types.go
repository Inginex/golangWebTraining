// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/sampalm/goa/cellar/design
// --out=$(GOPATH)\src\github.com\sampalm\goa\cellar
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A bottle of distilled (default view)
//
// Identifier: application/vnd.bottle+json; view=default
type Bottle struct {
	// Bottle brand
	Brand string `form:"brand" json:"brand" xml:"brand"`
	// Bottle name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Brand == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "brand"))
	}
	return
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// BottleCollection is the media type for an array of Bottle (default view)
//
// Identifier: application/vnd.bottle+json; type=collection; view=default
type BottleCollection []*Bottle

// Validate validates the BottleCollection media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeBottleCollection decodes the BottleCollection instance encoded in resp body.
func (c *Client) DecodeBottleCollection(resp *http.Response) (BottleCollection, error) {
	var decoded BottleCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
