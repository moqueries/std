// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rpc

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Register_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Register_genType func(rcvr interface{}) error

// MoqRegister_genType holds the state of a moq of the Register_genType type
type MoqRegister_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRegister_genType_mock

	ResultsByParams []MoqRegister_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rcvr moq.ParamIndexing
		}
	}
}

// MoqRegister_genType_mock isolates the mock interface of the Register_genType
// type
type MoqRegister_genType_mock struct {
	Moq *MoqRegister_genType
}

// MoqRegister_genType_params holds the params of the Register_genType type
type MoqRegister_genType_params struct{ Rcvr interface{} }

// MoqRegister_genType_paramsKey holds the map key params of the
// Register_genType type
type MoqRegister_genType_paramsKey struct {
	Params struct{ Rcvr interface{} }
	Hashes struct{ Rcvr hash.Hash }
}

// MoqRegister_genType_resultsByParams contains the results for a given set of
// parameters for the Register_genType type
type MoqRegister_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRegister_genType_paramsKey]*MoqRegister_genType_results
}

// MoqRegister_genType_doFn defines the type of function needed when calling
// AndDo for the Register_genType type
type MoqRegister_genType_doFn func(rcvr interface{})

// MoqRegister_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Register_genType type
type MoqRegister_genType_doReturnFn func(rcvr interface{}) error

// MoqRegister_genType_results holds the results of the Register_genType type
type MoqRegister_genType_results struct {
	Params  MoqRegister_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqRegister_genType_doFn
		DoReturnFn MoqRegister_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRegister_genType_fnRecorder routes recorded function calls to the
// MoqRegister_genType moq
type MoqRegister_genType_fnRecorder struct {
	Params    MoqRegister_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRegister_genType_results
	Moq       *MoqRegister_genType
}

// MoqRegister_genType_anyParams isolates the any params functions of the
// Register_genType type
type MoqRegister_genType_anyParams struct {
	Recorder *MoqRegister_genType_fnRecorder
}

// NewMoqRegister_genType creates a new moq of the Register_genType type
func NewMoqRegister_genType(scene *moq.Scene, config *moq.Config) *MoqRegister_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRegister_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRegister_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rcvr moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rcvr moq.ParamIndexing
		}{
			Rcvr: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Register_genType type
func (m *MoqRegister_genType) Mock() Register_genType {
	return func(rcvr interface{}) error {
		m.Scene.T.Helper()
		moq := &MoqRegister_genType_mock{Moq: m}
		return moq.Fn(rcvr)
	}
}

func (m *MoqRegister_genType_mock) Fn(rcvr interface{}) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRegister_genType_params{
		Rcvr: rcvr,
	}
	var results *MoqRegister_genType_results
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
		result.DoFn(rcvr)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(rcvr)
	}
	return
}

func (m *MoqRegister_genType) OnCall(rcvr interface{}) *MoqRegister_genType_fnRecorder {
	return &MoqRegister_genType_fnRecorder{
		Params: MoqRegister_genType_params{
			Rcvr: rcvr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRegister_genType_fnRecorder) Any() *MoqRegister_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRegister_genType_anyParams{Recorder: r}
}

func (a *MoqRegister_genType_anyParams) Rcvr() *MoqRegister_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqRegister_genType_fnRecorder) Seq() *MoqRegister_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRegister_genType_fnRecorder) NoSeq() *MoqRegister_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRegister_genType_fnRecorder) ReturnResults(result1 error) *MoqRegister_genType_fnRecorder {
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
		DoFn       MoqRegister_genType_doFn
		DoReturnFn MoqRegister_genType_doReturnFn
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

func (r *MoqRegister_genType_fnRecorder) AndDo(fn MoqRegister_genType_doFn) *MoqRegister_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRegister_genType_fnRecorder) DoReturnResults(fn MoqRegister_genType_doReturnFn) *MoqRegister_genType_fnRecorder {
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
		DoFn       MoqRegister_genType_doFn
		DoReturnFn MoqRegister_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRegister_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRegister_genType_resultsByParams
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
		results = &MoqRegister_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRegister_genType_paramsKey]*MoqRegister_genType_results{},
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
		r.Results = &MoqRegister_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRegister_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRegister_genType_fnRecorder {
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
				DoFn       MoqRegister_genType_doFn
				DoReturnFn MoqRegister_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRegister_genType) PrettyParams(params MoqRegister_genType_params) string {
	return fmt.Sprintf("Register_genType(%#v)", params.Rcvr)
}

func (m *MoqRegister_genType) ParamsKey(params MoqRegister_genType_params, anyParams uint64) MoqRegister_genType_paramsKey {
	m.Scene.T.Helper()
	var rcvrUsed interface{}
	var rcvrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Rcvr == moq.ParamIndexByValue {
			rcvrUsed = params.Rcvr
		} else {
			rcvrUsedHash = hash.DeepHash(params.Rcvr)
		}
	}
	return MoqRegister_genType_paramsKey{
		Params: struct{ Rcvr interface{} }{
			Rcvr: rcvrUsed,
		},
		Hashes: struct{ Rcvr hash.Hash }{
			Rcvr: rcvrUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRegister_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRegister_genType) AssertExpectationsMet() {
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