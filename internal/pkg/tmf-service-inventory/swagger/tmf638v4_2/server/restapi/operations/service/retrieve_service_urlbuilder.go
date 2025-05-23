// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RetrieveServiceURL generates an URL for the retrieve service operation
type RetrieveServiceURL struct {
	ID string

	Fields *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveServiceURL) WithBasePath(bp string) *RetrieveServiceURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RetrieveServiceURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RetrieveServiceURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/service/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on RetrieveServiceURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/tmf-api/serviceInventory/v4"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var fieldsQ string
	if o.Fields != nil {
		fieldsQ = *o.Fields
	}
	if fieldsQ != "" {
		qs.Set("fields", fieldsQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RetrieveServiceURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RetrieveServiceURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RetrieveServiceURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RetrieveServiceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RetrieveServiceURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *RetrieveServiceURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
