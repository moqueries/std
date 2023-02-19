// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package reflect

import (
	"fmt"
	"math/bits"
	"reflect"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Append_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Append_genType func(s reflect.Value, x ...reflect.Value) reflect.Value

// MoqAppend_genType holds the state of a moq of the Append_genType type
type MoqAppend_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAppend_genType_mock

	ResultsByParams []MoqAppend_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S moq.ParamIndexing
			X moq.ParamIndexing
		}
	}
}

// MoqAppend_genType_mock isolates the mock interface of the Append_genType
// type
type MoqAppend_genType_mock struct {
	Moq *MoqAppend_genType
}

// MoqAppend_genType_params holds the params of the Append_genType type
type MoqAppend_genType_params struct {
	S reflect.Value
	X []reflect.Value
}

// MoqAppend_genType_paramsKey holds the map key params of the Append_genType
// type
type MoqAppend_genType_paramsKey struct {
	Params struct{ S reflect.Value }
	Hashes struct {
		S hash.Hash
		X hash.Hash
	}
}

// MoqAppend_genType_resultsByParams contains the results for a given set of
// parameters for the Append_genType type
type MoqAppend_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAppend_genType_paramsKey]*MoqAppend_genType_results
}

// MoqAppend_genType_doFn defines the type of function needed when calling
// AndDo for the Append_genType type
type MoqAppend_genType_doFn func(s reflect.Value, x ...reflect.Value)

// MoqAppend_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Append_genType type
type MoqAppend_genType_doReturnFn func(s reflect.Value, x ...reflect.Value) reflect.Value

// MoqAppend_genType_results holds the results of the Append_genType type
type MoqAppend_genType_results struct {
	Params  MoqAppend_genType_params
	Results []struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqAppend_genType_doFn
		DoReturnFn MoqAppend_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAppend_genType_fnRecorder routes recorded function calls to the
// MoqAppend_genType moq
type MoqAppend_genType_fnRecorder struct {
	Params    MoqAppend_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAppend_genType_results
	Moq       *MoqAppend_genType
}

// MoqAppend_genType_anyParams isolates the any params functions of the
// Append_genType type
type MoqAppend_genType_anyParams struct {
	Recorder *MoqAppend_genType_fnRecorder
}

// NewMoqAppend_genType creates a new moq of the Append_genType type
func NewMoqAppend_genType(scene *moq.Scene, config *moq.Config) *MoqAppend_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAppend_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAppend_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S moq.ParamIndexing
				X moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S moq.ParamIndexing
			X moq.ParamIndexing
		}{
			S: moq.ParamIndexByHash,
			X: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Append_genType type
func (m *MoqAppend_genType) Mock() Append_genType {
	return func(s reflect.Value, x ...reflect.Value) reflect.Value {
		m.Scene.T.Helper()
		moq := &MoqAppend_genType_mock{Moq: m}
		return moq.Fn(s, x...)
	}
}

func (m *MoqAppend_genType_mock) Fn(s reflect.Value, x ...reflect.Value) (result1 reflect.Value) {
	m.Moq.Scene.T.Helper()
	params := MoqAppend_genType_params{
		S: s,
		X: x,
	}
	var results *MoqAppend_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(s, x...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, x...)
	}
	return
}

func (m *MoqAppend_genType) OnCall(s reflect.Value, x ...reflect.Value) *MoqAppend_genType_fnRecorder {
	return &MoqAppend_genType_fnRecorder{
		Params: MoqAppend_genType_params{
			S: s,
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAppend_genType_fnRecorder) Any() *MoqAppend_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAppend_genType_anyParams{Recorder: r}
}

func (a *MoqAppend_genType_anyParams) S() *MoqAppend_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAppend_genType_anyParams) X() *MoqAppend_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqAppend_genType_fnRecorder) Seq() *MoqAppend_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAppend_genType_fnRecorder) NoSeq() *MoqAppend_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAppend_genType_fnRecorder) ReturnResults(result1 reflect.Value) *MoqAppend_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqAppend_genType_doFn
		DoReturnFn MoqAppend_genType_doReturnFn
	}{
		Values: &struct {
			Result1 reflect.Value
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAppend_genType_fnRecorder) AndDo(fn MoqAppend_genType_doFn) *MoqAppend_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAppend_genType_fnRecorder) DoReturnResults(fn MoqAppend_genType_doReturnFn) *MoqAppend_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqAppend_genType_doFn
		DoReturnFn MoqAppend_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAppend_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAppend_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqAppend_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAppend_genType_paramsKey]*MoqAppend_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqAppend_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAppend_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAppend_genType_fnRecorder {
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
					Result1 reflect.Value
				}
				Sequence   uint32
				DoFn       MoqAppend_genType_doFn
				DoReturnFn MoqAppend_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAppend_genType) PrettyParams(params MoqAppend_genType_params) string {
	return fmt.Sprintf("Append_genType(%#v, %#v)", params.S, params.X)
}

func (m *MoqAppend_genType) ParamsKey(params MoqAppend_genType_params, anyParams uint64) MoqAppend_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsed reflect.Value
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	var xUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The x parameter can't be indexed by value")
		}
		xUsedHash = hash.DeepHash(params.X)
	}
	return MoqAppend_genType_paramsKey{
		Params: struct{ S reflect.Value }{
			S: sUsed,
		},
		Hashes: struct {
			S hash.Hash
			X hash.Hash
		}{
			S: sUsedHash,
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAppend_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAppend_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}
