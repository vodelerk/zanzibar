// Code generated by zanzibar
// @generated

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

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgWithManyQueryParamsHandler is the handler for "/bar/argWithManyQueryParams"
type BarArgWithManyQueryParamsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarArgWithManyQueryParamsHandler creates a handler
func NewBarArgWithManyQueryParamsHandler(
	g *zanzibar.Gateway,
	deps *module.Dependencies,
) *BarArgWithManyQueryParamsHandler {
	handler := &BarArgWithManyQueryParamsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "argWithManyQueryParams",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarArgWithManyQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/argWithManyQueryParams",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/argWithManyQueryParams".
func (h *BarArgWithManyQueryParamsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgWithManyQueryParams_Args

	aStrOk := req.CheckQueryValue("aStr")
	if !aStrOk {
		return
	}
	aStrQuery, ok := req.GetQueryValue("aStr")
	if !ok {
		return
	}
	requestBody.AStr = aStrQuery

	anOptStrOk := req.HasQueryValue("anOptStr")
	if anOptStrOk {
		anOptStrQuery, ok := req.GetQueryValue("anOptStr")
		if !ok {
			return
		}
		requestBody.AnOptStr = ptr.String(anOptStrQuery)
	}

	aBoolOk := req.CheckQueryValue("aBool")
	if !aBoolOk {
		return
	}
	aBoolQuery, ok := req.GetQueryBool("aBool")
	if !ok {
		return
	}
	requestBody.ABool = aBoolQuery

	anOptBoolOk := req.HasQueryValue("anOptBool")
	if anOptBoolOk {
		anOptBoolQuery, ok := req.GetQueryBool("anOptBool")
		if !ok {
			return
		}
		requestBody.AnOptBool = ptr.Bool(anOptBoolQuery)
	}

	aInt8Ok := req.CheckQueryValue("aInt8")
	if !aInt8Ok {
		return
	}
	aInt8Query, ok := req.GetQueryInt8("aInt8")
	if !ok {
		return
	}
	requestBody.AInt8 = aInt8Query

	anOptInt8Ok := req.HasQueryValue("anOptInt8")
	if anOptInt8Ok {
		anOptInt8Query, ok := req.GetQueryInt8("anOptInt8")
		if !ok {
			return
		}
		requestBody.AnOptInt8 = ptr.Int8(anOptInt8Query)
	}

	aInt16Ok := req.CheckQueryValue("aInt16")
	if !aInt16Ok {
		return
	}
	aInt16Query, ok := req.GetQueryInt16("aInt16")
	if !ok {
		return
	}
	requestBody.AInt16 = aInt16Query

	anOptInt16Ok := req.HasQueryValue("anOptInt16")
	if anOptInt16Ok {
		anOptInt16Query, ok := req.GetQueryInt16("anOptInt16")
		if !ok {
			return
		}
		requestBody.AnOptInt16 = ptr.Int16(anOptInt16Query)
	}

	aInt32Ok := req.CheckQueryValue("aInt32")
	if !aInt32Ok {
		return
	}
	aInt32Query, ok := req.GetQueryInt32("aInt32")
	if !ok {
		return
	}
	requestBody.AInt32 = aInt32Query

	anOptInt32Ok := req.HasQueryValue("anOptInt32")
	if anOptInt32Ok {
		anOptInt32Query, ok := req.GetQueryInt32("anOptInt32")
		if !ok {
			return
		}
		requestBody.AnOptInt32 = ptr.Int32(anOptInt32Query)
	}

	aInt64Ok := req.CheckQueryValue("aInt64")
	if !aInt64Ok {
		return
	}
	aInt64Query, ok := req.GetQueryInt64("aInt64")
	if !ok {
		return
	}
	requestBody.AInt64 = aInt64Query

	anOptInt64Ok := req.HasQueryValue("anOptInt64")
	if anOptInt64Ok {
		anOptInt64Query, ok := req.GetQueryInt64("anOptInt64")
		if !ok {
			return
		}
		requestBody.AnOptInt64 = ptr.Int64(anOptInt64Query)
	}

	aFloat64Ok := req.CheckQueryValue("aFloat64")
	if !aFloat64Ok {
		return
	}
	aFloat64Query, ok := req.GetQueryFloat64("aFloat64")
	if !ok {
		return
	}
	requestBody.AFloat64 = aFloat64Query

	anOptFloat64Ok := req.HasQueryValue("anOptFloat64")
	if anOptFloat64Ok {
		anOptFloat64Query, ok := req.GetQueryFloat64("anOptFloat64")
		if !ok {
			return
		}
		requestBody.AnOptFloat64 = ptr.Float64(anOptFloat64Query)
	}

	workflow := ArgWithManyQueryParamsEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// ArgWithManyQueryParamsEndpoint calls thrift client Bar.ArgWithManyQueryParams
type ArgWithManyQueryParamsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ArgWithManyQueryParamsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithManyQueryParams_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithManyQueryParamsClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.ArgWithManyQueryParams(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertArgWithManyQueryParamsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithManyQueryParamsClientRequest(in *endpointsBarBar.Bar_ArgWithManyQueryParams_Args) *clientsBarBar.Bar_ArgWithManyQueryParams_Args {
	out := &clientsBarBar.Bar_ArgWithManyQueryParams_Args{}

	out.AStr = string(in.AStr)
	out.AnOptStr = (*string)(in.AnOptStr)
	out.ABool = bool(in.ABool)
	out.AnOptBool = (*bool)(in.AnOptBool)
	out.AInt8 = int8(in.AInt8)
	out.AnOptInt8 = (*int8)(in.AnOptInt8)
	out.AInt16 = int16(in.AInt16)
	out.AnOptInt16 = (*int16)(in.AnOptInt16)
	out.AInt32 = int32(in.AInt32)
	out.AnOptInt32 = (*int32)(in.AnOptInt32)
	out.AInt64 = int64(in.AInt64)
	out.AnOptInt64 = (*int64)(in.AnOptInt64)
	out.AFloat64 = float64(in.AFloat64)
	out.AnOptFloat64 = (*float64)(in.AnOptFloat64)

	return out
}

func convertArgWithManyQueryParamsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[string]int32, len(in.MapIntWithRange))
	for key, value := range in.MapIntWithRange {
		out.MapIntWithRange[key] = int32(value)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key, value := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key] = int32(value)
	}
	out.BinaryField = []byte(in.BinaryField)

	return out
}
