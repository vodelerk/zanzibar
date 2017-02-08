// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gencode

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/thriftrw/compile"
)

// MethodSpec specifies all needed parts to generate code for a method in service.
type MethodSpec struct {
	Name       string
	HTTPMethod string
	// Used by edge gateway to generate endpoint.
	EndpointName string
	HTTPPath     string
	// Headers needed, generated from "zanzibar.http.headers"
	Headers             []string
	RequestType         string
	ResponseType        string
	OKStatusCode        []StatusCode
	ExceptionStatusCode []StatusCode
	// Additional struct generated from the bundle of request args.
	RequestStruct []StructSpec
}

// StructSpec specifies a Go struct to be generated.
type StructSpec struct {
	Type        string
	Name        string
	Annotations map[string]string
}

// StatusCode is for http status code with exception message.
type StatusCode struct {
	Code    int
	Message string
}

const (
	antHTTPMethod      = "zanzibar.http.method"
	antHTTPPath        = "zanzibar.http.path"
	antHTTPStatus      = "zanzibar.http.status"
	antHTTPReqDefBoxed = "zanzibar.http.req.def"
	antMeta            = "zanzibar.meta"
	antHandler         = "zanzibar.handler"
	antHTTPHeaders     = "zanzibar.http.headers"
)

// setRequestType sets the request type of the method specification. If the
// "zanzibar.http.req.def.boxed" is true, then the first parameter will be used as
// the request body; otherwise a new struct is generated to bundle the request
// parameters as http body and the name of the struct will be returned.
func (ms *MethodSpec) setRequestType(curThriftFile string, funcSpec *compile.FunctionSpec, packageHelper *PackageHelper) error {
	if len(funcSpec.ArgsSpec) == 0 {
		ms.RequestType = ""
		return nil
	}
	var err error
	if isRequestBoxed(funcSpec) {
		ms.RequestType, err = packageHelper.TypeFullName(curThriftFile, funcSpec.ArgsSpec[0].Type)
	} else {
		ms.RequestType, err = ms.newRequestType(curThriftFile, funcSpec, packageHelper)
	}
	if err != nil {
		return errors.Wrap(err, "failed to set request type")
	}
	return nil
}

func (ms *MethodSpec) newRequestType(curThriftFile string, f *compile.FunctionSpec, h *PackageHelper) (string, error) {
	requestType := strings.Title(f.Name) + "HTTPRequest"
	ms.RequestStruct = make([]StructSpec, len(f.ArgsSpec))
	for i, arg := range f.ArgsSpec {
		typeName, err := h.TypeFullName(curThriftFile, arg.Type)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate new request type")
		}
		ms.RequestStruct[i] = StructSpec{
			Type:        typeName,
			Name:        arg.Name,
			Annotations: arg.Annotations,
		}
	}
	return requestType, nil
}

func (ms *MethodSpec) setResponseType(curThriftFile string, respSpec *compile.ResultSpec, packageHelper *PackageHelper) error {
	if respSpec == nil {
		ms.ResponseType = ""
		return nil
	}
	typeName, err := packageHelper.TypeFullName(curThriftFile, respSpec.ReturnType)
	if err != nil {
		return errors.Wrap(err, "failed to get response type")
	}
	ms.ResponseType = typeName
	return nil
}
func (ms *MethodSpec) setOKStatusCode(statusCode string) error {
	if statusCode == "" {
		return errors.Errorf("no http OK status code set by annotation '%s' ", antHTTPStatus)
	}
	scode := strings.Split(statusCode, ",")
	ms.OKStatusCode = make([]StatusCode, len(scode))
	var err error
	for i, c := range scode {
		ms.OKStatusCode[i].Code, err = strconv.Atoi(c)
		if err != nil {
			return errors.Wrapf(err, "failed to parse the annotation %s for ok response status")
		}
	}
	return nil
}

func (ms *MethodSpec) setExceptionStatusCode(resultSpec *compile.ResultSpec) error {
	ms.ExceptionStatusCode = make([]StatusCode, len(resultSpec.Exceptions))
	for i, e := range resultSpec.Exceptions {
		code, err := strconv.Atoi(e.Annotations[antHTTPStatus])
		if err != nil {
			return errors.Wrapf(err, "cannot parse the annotation %s for exception %s", antHTTPStatus, e.Name)
		}
		ms.ExceptionStatusCode[i] = StatusCode{
			Code:    code,
			Message: e.Name,
		}
	}
	return nil
}

func isRequestBoxed(f *compile.FunctionSpec) bool {
	boxed, ok := f.Annotations[antHTTPReqDefBoxed]
	if ok && boxed == "true" {
		return true
	}
	return false
}

func headers(annotation string) []string {
	if annotation == "" {
		return nil
	}
	return strings.Split(annotation, ",")
}