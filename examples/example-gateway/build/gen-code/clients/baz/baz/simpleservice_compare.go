// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// SimpleService_Compare_Args represents the arguments for the SimpleService.compare function.
//
// The arguments for compare are sent and received over the wire as this struct.
type SimpleService_Compare_Args struct {
	Arg1 *BazRequest `json:"arg1,required"`
	Arg2 *BazRequest `json:"arg2,required"`
}

// ToWire translates a SimpleService_Compare_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *SimpleService_Compare_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg1 == nil {
		return w, errors.New("field Arg1 of SimpleService_Compare_Args is required")
	}
	w, err = v.Arg1.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Arg2 == nil {
		return w, errors.New("field Arg2 of SimpleService_Compare_Args is required")
	}
	w, err = v.Arg2.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_Compare_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Compare_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Compare_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Compare_Args) FromWire(w wire.Value) error {
	var err error

	arg1IsSet := false
	arg2IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Arg1, err = _BazRequest_Read(field.Value)
				if err != nil {
					return err
				}
				arg1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Arg2, err = _BazRequest_Read(field.Value)
				if err != nil {
					return err
				}
				arg2IsSet = true
			}
		}
	}

	if !arg1IsSet {
		return errors.New("field Arg1 of SimpleService_Compare_Args is required")
	}

	if !arg2IsSet {
		return errors.New("field Arg2 of SimpleService_Compare_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Compare_Args
// struct.
func (v *SimpleService_Compare_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Arg1: %v", v.Arg1)
	i++
	fields[i] = fmt.Sprintf("Arg2: %v", v.Arg2)
	i++

	return fmt.Sprintf("SimpleService_Compare_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Compare_Args match the
// provided SimpleService_Compare_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_Compare_Args) Equals(rhs *SimpleService_Compare_Args) bool {
	if !v.Arg1.Equals(rhs.Arg1) {
		return false
	}
	if !v.Arg2.Equals(rhs.Arg2) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "compare" for this struct.
func (v *SimpleService_Compare_Args) MethodName() string {
	return "compare"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_Compare_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_Compare_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.compare
// function.
var SimpleService_Compare_Helper = struct {
	// Args accepts the parameters of compare in-order and returns
	// the arguments struct for the function.
	Args func(
		arg1 *BazRequest,
		arg2 *BazRequest,
	) *SimpleService_Compare_Args

	// IsException returns true if the given error can be thrown
	// by compare.
	//
	// An error can be thrown by compare only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for compare
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// compare into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by compare
	//
	//   value, err := compare(args)
	//   result, err := SimpleService_Compare_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from compare: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*base.BazResponse, error) (*SimpleService_Compare_Result, error)

	// UnwrapResponse takes the result struct for compare
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if compare threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SimpleService_Compare_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_Compare_Result) (*base.BazResponse, error)
}{}

func init() {
	SimpleService_Compare_Helper.Args = func(
		arg1 *BazRequest,
		arg2 *BazRequest,
	) *SimpleService_Compare_Args {
		return &SimpleService_Compare_Args{
			Arg1: arg1,
			Arg2: arg2,
		}
	}

	SimpleService_Compare_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *OtherAuthErr:
			return true
		default:
			return false
		}
	}

	SimpleService_Compare_Helper.WrapResponse = func(success *base.BazResponse, err error) (*SimpleService_Compare_Result, error) {
		if err == nil {
			return &SimpleService_Compare_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Compare_Result.AuthErr")
			}
			return &SimpleService_Compare_Result{AuthErr: e}, nil
		case *OtherAuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Compare_Result.OtherAuthErr")
			}
			return &SimpleService_Compare_Result{OtherAuthErr: e}, nil
		}

		return nil, err
	}
	SimpleService_Compare_Helper.UnwrapResponse = func(result *SimpleService_Compare_Result) (success *base.BazResponse, err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.OtherAuthErr != nil {
			err = result.OtherAuthErr
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SimpleService_Compare_Result represents the result of a SimpleService.compare function call.
//
// The result of a compare execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SimpleService_Compare_Result struct {
	// Value returned by compare after a successful execution.
	Success      *base.BazResponse `json:"success,omitempty"`
	AuthErr      *AuthErr          `json:"authErr,omitempty"`
	OtherAuthErr *OtherAuthErr     `json:"otherAuthErr,omitempty"`
}

// ToWire translates a SimpleService_Compare_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *SimpleService_Compare_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.OtherAuthErr != nil {
		w, err = v.OtherAuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_Compare_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _OtherAuthErr_Read(w wire.Value) (*OtherAuthErr, error) {
	var v OtherAuthErr
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a SimpleService_Compare_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Compare_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Compare_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Compare_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BazResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.OtherAuthErr, err = _OtherAuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.AuthErr != nil {
		count++
	}
	if v.OtherAuthErr != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_Compare_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Compare_Result
// struct.
func (v *SimpleService_Compare_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.OtherAuthErr != nil {
		fields[i] = fmt.Sprintf("OtherAuthErr: %v", v.OtherAuthErr)
		i++
	}

	return fmt.Sprintf("SimpleService_Compare_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Compare_Result match the
// provided SimpleService_Compare_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_Compare_Result) Equals(rhs *SimpleService_Compare_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.OtherAuthErr == nil && rhs.OtherAuthErr == nil) || (v.OtherAuthErr != nil && rhs.OtherAuthErr != nil && v.OtherAuthErr.Equals(rhs.OtherAuthErr))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "compare" for this struct.
func (v *SimpleService_Compare_Result) MethodName() string {
	return "compare"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_Compare_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
