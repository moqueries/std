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

// ValueOf_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type ValueOf_genType func(i interface{}) reflect.Value

// MoqValueOf_genType holds the state of a moq of the ValueOf_genType type
type MoqValueOf_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqValueOf_genType_mock

	ResultsByParams []MoqValueOf_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
		}
	}
}

// MoqValueOf_genType_mock isolates the mock interface of the ValueOf_genType
// type
type MoqValueOf_genType_mock struct {
	Moq *MoqValueOf_genType
}

// MoqValueOf_genType_params holds the params of the ValueOf_genType type
type MoqValueOf_genType_params struct{ Param1 interface{} }

// MoqValueOf_genType_paramsKey holds the map key params of the ValueOf_genType
// type
type MoqValueOf_genType_paramsKey struct {
	Params struct{ Param1 interface{} }
	Hashes struct{ Param1 hash.Hash }
}

// MoqValueOf_genType_resultsByParams contains the results for a given set of
// parameters for the ValueOf_genType type
type MoqValueOf_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValueOf_genType_paramsKey]*MoqValueOf_genType_results
}

// MoqValueOf_genType_doFn defines the type of function needed when calling
// AndDo for the ValueOf_genType type
type MoqValueOf_genType_doFn func(i interface{})

// MoqValueOf_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ValueOf_genType type
type MoqValueOf_genType_doReturnFn func(i interface{}) reflect.Value

// MoqValueOf_genType_results holds the results of the ValueOf_genType type
type MoqValueOf_genType_results struct {
	Params  MoqValueOf_genType_params
	Results []struct {
		Values *struct {
			Result1 reflect.Value
		}
		Sequence   uint32
		DoFn       MoqValueOf_genType_doFn
		DoReturnFn MoqValueOf_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValueOf_genType_fnRecorder routes recorded function calls to the
// MoqValueOf_genType moq
type MoqValueOf_genType_fnRecorder struct {
	Params    MoqValueOf_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValueOf_genType_results
	Moq       *MoqValueOf_genType
}

// MoqValueOf_genType_anyParams isolates the any params functions of the
// ValueOf_genType type
type MoqValueOf_genType_anyParams struct {
	Recorder *MoqValueOf_genType_fnRecorder
}

// NewMoqValueOf_genType creates a new moq of the ValueOf_genType type
func NewMoqValueOf_genType(scene *moq.Scene, config *moq.Config) *MoqValueOf_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqValueOf_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqValueOf_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ValueOf_genType type
func (m *MoqValueOf_genType) Mock() ValueOf_genType {
	return func(param1 interface{}) reflect.Value {
		m.Scene.T.Helper()
		moq := &MoqValueOf_genType_mock{Moq: m}
		return moq.Fn(param1)
	}
}

func (m *MoqValueOf_genType_mock) Fn(param1 interface{}) (result1 reflect.Value) {
	m.Moq.Scene.T.Helper()
	params := MoqValueOf_genType_params{
		Param1: param1,
	}
	var results *MoqValueOf_genType_results
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
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqValueOf_genType) OnCall(param1 interface{}) *MoqValueOf_genType_fnRecorder {
	return &MoqValueOf_genType_fnRecorder{
		Params: MoqValueOf_genType_params{
			Param1: param1,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqValueOf_genType_fnRecorder) Any() *MoqValueOf_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqValueOf_genType_anyParams{Recorder: r}
}

func (a *MoqValueOf_genType_anyParams) Param1() *MoqValueOf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqValueOf_genType_fnRecorder) Seq() *MoqValueOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValueOf_genType_fnRecorder) NoSeq() *MoqValueOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValueOf_genType_fnRecorder) ReturnResults(result1 reflect.Value) *MoqValueOf_genType_fnRecorder {
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
		DoFn       MoqValueOf_genType_doFn
		DoReturnFn MoqValueOf_genType_doReturnFn
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

func (r *MoqValueOf_genType_fnRecorder) AndDo(fn MoqValueOf_genType_doFn) *MoqValueOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValueOf_genType_fnRecorder) DoReturnResults(fn MoqValueOf_genType_doReturnFn) *MoqValueOf_genType_fnRecorder {
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
		DoFn       MoqValueOf_genType_doFn
		DoReturnFn MoqValueOf_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValueOf_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValueOf_genType_resultsByParams
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
		results = &MoqValueOf_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValueOf_genType_paramsKey]*MoqValueOf_genType_results{},
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
		r.Results = &MoqValueOf_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValueOf_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValueOf_genType_fnRecorder {
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
				DoFn       MoqValueOf_genType_doFn
				DoReturnFn MoqValueOf_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValueOf_genType) PrettyParams(params MoqValueOf_genType_params) string {
	return fmt.Sprintf("ValueOf_genType(%#v)", params.Param1)
}

func (m *MoqValueOf_genType) ParamsKey(params MoqValueOf_genType_params, anyParams uint64) MoqValueOf_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used interface{}
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqValueOf_genType_paramsKey{
		Params: struct{ Param1 interface{} }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqValueOf_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqValueOf_genType) AssertExpectationsMet() {
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
