// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package gob

import (
	"fmt"
	"math/bits"
	"reflect"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that gob.Encoder_starGenType is mocked
// completely
var _ Encoder_starGenType = (*MoqEncoder_starGenType_mock)(nil)

// Encoder_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Encoder_starGenType interface {
	Encode(e any) error
	EncodeValue(value reflect.Value) error
}

// MoqEncoder_starGenType holds the state of a moq of the Encoder_starGenType
// type
type MoqEncoder_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncoder_starGenType_mock

	ResultsByParams_Encode      []MoqEncoder_starGenType_Encode_resultsByParams
	ResultsByParams_EncodeValue []MoqEncoder_starGenType_EncodeValue_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Encode struct {
				E moq.ParamIndexing
			}
			EncodeValue struct {
				Value moq.ParamIndexing
			}
		}
	}
	// MoqEncoder_starGenType_mock isolates the mock interface of the
}

// Encoder_starGenType type
type MoqEncoder_starGenType_mock struct {
	Moq *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_recorder isolates the recorder interface of the
// Encoder_starGenType type
type MoqEncoder_starGenType_recorder struct {
	Moq *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_Encode_params holds the params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_params struct{ E any }

// MoqEncoder_starGenType_Encode_paramsKey holds the map key params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_paramsKey struct {
	Params struct{ E any }
	Hashes struct{ E hash.Hash }
}

// MoqEncoder_starGenType_Encode_resultsByParams contains the results for a
// given set of parameters for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoder_starGenType_Encode_paramsKey]*MoqEncoder_starGenType_Encode_results
}

// MoqEncoder_starGenType_Encode_doFn defines the type of function needed when
// calling AndDo for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_doFn func(e any)

// MoqEncoder_starGenType_Encode_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_doReturnFn func(e any) error

// MoqEncoder_starGenType_Encode_results holds the results of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_results struct {
	Params  MoqEncoder_starGenType_Encode_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoder_starGenType_Encode_fnRecorder routes recorded function calls to
// the MoqEncoder_starGenType moq
type MoqEncoder_starGenType_Encode_fnRecorder struct {
	Params    MoqEncoder_starGenType_Encode_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoder_starGenType_Encode_results
	Moq       *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_Encode_anyParams isolates the any params functions of
// the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_anyParams struct {
	Recorder *MoqEncoder_starGenType_Encode_fnRecorder
}

// MoqEncoder_starGenType_EncodeValue_params holds the params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_params struct{ Value reflect.Value }

// MoqEncoder_starGenType_EncodeValue_paramsKey holds the map key params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_paramsKey struct {
	Params struct{ Value reflect.Value }
	Hashes struct{ Value hash.Hash }
}

// MoqEncoder_starGenType_EncodeValue_resultsByParams contains the results for
// a given set of parameters for the Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoder_starGenType_EncodeValue_paramsKey]*MoqEncoder_starGenType_EncodeValue_results
}

// MoqEncoder_starGenType_EncodeValue_doFn defines the type of function needed
// when calling AndDo for the Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_doFn func(value reflect.Value)

// MoqEncoder_starGenType_EncodeValue_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_doReturnFn func(value reflect.Value) error

// MoqEncoder_starGenType_EncodeValue_results holds the results of the
// Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_results struct {
	Params  MoqEncoder_starGenType_EncodeValue_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_EncodeValue_doFn
		DoReturnFn MoqEncoder_starGenType_EncodeValue_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoder_starGenType_EncodeValue_fnRecorder routes recorded function calls
// to the MoqEncoder_starGenType moq
type MoqEncoder_starGenType_EncodeValue_fnRecorder struct {
	Params    MoqEncoder_starGenType_EncodeValue_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoder_starGenType_EncodeValue_results
	Moq       *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_EncodeValue_anyParams isolates the any params
// functions of the Encoder_starGenType type
type MoqEncoder_starGenType_EncodeValue_anyParams struct {
	Recorder *MoqEncoder_starGenType_EncodeValue_fnRecorder
}

// NewMoqEncoder_starGenType creates a new moq of the Encoder_starGenType type
func NewMoqEncoder_starGenType(scene *moq.Scene, config *moq.Config) *MoqEncoder_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncoder_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncoder_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Encode struct {
					E moq.ParamIndexing
				}
				EncodeValue struct {
					Value moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Encode struct {
				E moq.ParamIndexing
			}
			EncodeValue struct {
				Value moq.ParamIndexing
			}
		}{
			Encode: struct {
				E moq.ParamIndexing
			}{
				E: moq.ParamIndexByValue,
			},
			EncodeValue: struct {
				Value moq.ParamIndexing
			}{
				Value: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Encoder_starGenType type
func (m *MoqEncoder_starGenType) Mock() *MoqEncoder_starGenType_mock { return m.Moq }

func (m *MoqEncoder_starGenType_mock) Encode(e any) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoder_starGenType_Encode_params{
		E: e,
	}
	var results *MoqEncoder_starGenType_Encode_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Encode {
		paramsKey := m.Moq.ParamsKey_Encode(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Encode(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Encode(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Encode(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(e)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(e)
	}
	return
}

func (m *MoqEncoder_starGenType_mock) EncodeValue(value reflect.Value) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoder_starGenType_EncodeValue_params{
		Value: value,
	}
	var results *MoqEncoder_starGenType_EncodeValue_results
	for _, resultsByParams := range m.Moq.ResultsByParams_EncodeValue {
		paramsKey := m.Moq.ParamsKey_EncodeValue(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_EncodeValue(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_EncodeValue(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_EncodeValue(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(value)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(value)
	}
	return
}

// OnCall returns the recorder implementation of the Encoder_starGenType type
func (m *MoqEncoder_starGenType) OnCall() *MoqEncoder_starGenType_recorder {
	return &MoqEncoder_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqEncoder_starGenType_recorder) Encode(e any) *MoqEncoder_starGenType_Encode_fnRecorder {
	return &MoqEncoder_starGenType_Encode_fnRecorder{
		Params: MoqEncoder_starGenType_Encode_params{
			E: e,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Any() *MoqEncoder_starGenType_Encode_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	return &MoqEncoder_starGenType_Encode_anyParams{Recorder: r}
}

func (a *MoqEncoder_starGenType_Encode_anyParams) E() *MoqEncoder_starGenType_Encode_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Seq() *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) NoSeq() *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) ReturnResults(result1 error) *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) AndDo(fn MoqEncoder_starGenType_Encode_doFn) *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) DoReturnResults(fn MoqEncoder_starGenType_Encode_doReturnFn) *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoder_starGenType_Encode_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Encode {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoder_starGenType_Encode_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoder_starGenType_Encode_paramsKey]*MoqEncoder_starGenType_Encode_results{},
		}
		r.Moq.ResultsByParams_Encode = append(r.Moq.ResultsByParams_Encode, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Encode) {
			copy(r.Moq.ResultsByParams_Encode[insertAt+1:], r.Moq.ResultsByParams_Encode[insertAt:0])
			r.Moq.ResultsByParams_Encode[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Encode(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoder_starGenType_Encode_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqEncoder_starGenType_Encode_doFn
				DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoder_starGenType) PrettyParams_Encode(params MoqEncoder_starGenType_Encode_params) string {
	return fmt.Sprintf("Encode(%#v)", params.E)
}

func (m *MoqEncoder_starGenType) ParamsKey_Encode(params MoqEncoder_starGenType_Encode_params, anyParams uint64) MoqEncoder_starGenType_Encode_paramsKey {
	m.Scene.T.Helper()
	var eUsed any
	var eUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Encode.E == moq.ParamIndexByValue {
			eUsed = params.E
		} else {
			eUsedHash = hash.DeepHash(params.E)
		}
	}
	return MoqEncoder_starGenType_Encode_paramsKey{
		Params: struct{ E any }{
			E: eUsed,
		},
		Hashes: struct{ E hash.Hash }{
			E: eUsedHash,
		},
	}
}

func (m *MoqEncoder_starGenType_recorder) EncodeValue(value reflect.Value) *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	return &MoqEncoder_starGenType_EncodeValue_fnRecorder{
		Params: MoqEncoder_starGenType_EncodeValue_params{
			Value: value,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) Any() *MoqEncoder_starGenType_EncodeValue_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_EncodeValue(r.Params))
		return nil
	}
	return &MoqEncoder_starGenType_EncodeValue_anyParams{Recorder: r}
}

func (a *MoqEncoder_starGenType_EncodeValue_anyParams) Value() *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) Seq() *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_EncodeValue(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) NoSeq() *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_EncodeValue(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) ReturnResults(result1 error) *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_EncodeValue_doFn
		DoReturnFn MoqEncoder_starGenType_EncodeValue_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) AndDo(fn MoqEncoder_starGenType_EncodeValue_doFn) *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) DoReturnResults(fn MoqEncoder_starGenType_EncodeValue_doReturnFn) *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_EncodeValue_doFn
		DoReturnFn MoqEncoder_starGenType_EncodeValue_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoder_starGenType_EncodeValue_resultsByParams
	for n, res := range r.Moq.ResultsByParams_EncodeValue {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoder_starGenType_EncodeValue_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoder_starGenType_EncodeValue_paramsKey]*MoqEncoder_starGenType_EncodeValue_results{},
		}
		r.Moq.ResultsByParams_EncodeValue = append(r.Moq.ResultsByParams_EncodeValue, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_EncodeValue) {
			copy(r.Moq.ResultsByParams_EncodeValue[insertAt+1:], r.Moq.ResultsByParams_EncodeValue[insertAt:0])
			r.Moq.ResultsByParams_EncodeValue[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_EncodeValue(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoder_starGenType_EncodeValue_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoder_starGenType_EncodeValue_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoder_starGenType_EncodeValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqEncoder_starGenType_EncodeValue_doFn
				DoReturnFn MoqEncoder_starGenType_EncodeValue_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoder_starGenType) PrettyParams_EncodeValue(params MoqEncoder_starGenType_EncodeValue_params) string {
	return fmt.Sprintf("EncodeValue(%#v)", params.Value)
}

func (m *MoqEncoder_starGenType) ParamsKey_EncodeValue(params MoqEncoder_starGenType_EncodeValue_params, anyParams uint64) MoqEncoder_starGenType_EncodeValue_paramsKey {
	m.Scene.T.Helper()
	var valueUsed reflect.Value
	var valueUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.EncodeValue.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	return MoqEncoder_starGenType_EncodeValue_paramsKey{
		Params: struct{ Value reflect.Value }{
			Value: valueUsed,
		},
		Hashes: struct{ Value hash.Hash }{
			Value: valueUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEncoder_starGenType) Reset() {
	m.ResultsByParams_Encode = nil
	m.ResultsByParams_EncodeValue = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncoder_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Encode {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Encode(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_EncodeValue {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_EncodeValue(results.Params))
			}
		}
	}
}
