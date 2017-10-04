// Code generated by zanzibar
// @generated
// Checksum : OdWhxU6ZkipI5YUcgUzp2Q==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	foo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(in *jlexer.Lexer, out *Bar_TooManyArgs_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				(*out.BarException).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(out *jwriter.Writer, in Bar_TooManyArgs_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			(*in.Success).MarshalEasyJSON(out)
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			(*in.BarException).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(l, v)
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(in *jlexer.Lexer, out *Bar_TooManyArgs_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				(*out.Request).UnmarshalEasyJSON(in)
			}
			RequestSet = true
		case "foo":
			if in.IsNull() {
				in.Skip()
				out.Foo = nil
			} else {
				if out.Foo == nil {
					out.Foo = new(foo.FooStruct)
				}
				(*out.Foo).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(out *jwriter.Writer, in Bar_TooManyArgs_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		(*in.Request).MarshalEasyJSON(out)
	}
	if in.Foo != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"foo\":")
		if in.Foo == nil {
			out.RawString("null")
		} else {
			(*in.Foo).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(l, v)
}
