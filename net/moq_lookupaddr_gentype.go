// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LookupAddr_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LookupAddr_genType func(addr string) (names []string, err error)

// MoqLookupAddr_genType holds the state of a moq of the LookupAddr_genType
// type
type MoqLookupAddr_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLookupAddr_genType_mock

	ResultsByParams []MoqLookupAddr_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr moq.ParamIndexing
		}
	}
}

// MoqLookupAddr_genType_mock isolates the mock interface of the
// LookupAddr_genType type
type MoqLookupAddr_genType_mock struct {
	Moq *MoqLookupAddr_genType
}

// MoqLookupAddr_genType_params holds the params of the LookupAddr_genType type
type MoqLookupAddr_genType_params struct{ Addr string }

// MoqLookupAddr_genType_paramsKey holds the map key params of the
// LookupAddr_genType type
type MoqLookupAddr_genType_paramsKey struct {
	Params struct{ Addr string }
	Hashes struct{ Addr hash.Hash }
}

// MoqLookupAddr_genType_resultsByParams contains the results for a given set
// of parameters for the LookupAddr_genType type
type MoqLookupAddr_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLookupAddr_genType_paramsKey]*MoqLookupAddr_genType_results
}

// MoqLookupAddr_genType_doFn defines the type of function needed when calling
// AndDo for the LookupAddr_genType type
type MoqLookupAddr_genType_doFn func(addr string)

// MoqLookupAddr_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LookupAddr_genType type
type MoqLookupAddr_genType_doReturnFn func(addr string) (names []string, err error)

// MoqLookupAddr_genType_results holds the results of the LookupAddr_genType
// type
type MoqLookupAddr_genType_results struct {
	Params  MoqLookupAddr_genType_params
	Results []struct {
		Values *struct {
			Names []string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupAddr_genType_doFn
		DoReturnFn MoqLookupAddr_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLookupAddr_genType_fnRecorder routes recorded function calls to the
// MoqLookupAddr_genType moq
type MoqLookupAddr_genType_fnRecorder struct {
	Params    MoqLookupAddr_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLookupAddr_genType_results
	Moq       *MoqLookupAddr_genType
}

// MoqLookupAddr_genType_anyParams isolates the any params functions of the
// LookupAddr_genType type
type MoqLookupAddr_genType_anyParams struct {
	Recorder *MoqLookupAddr_genType_fnRecorder
}

// NewMoqLookupAddr_genType creates a new moq of the LookupAddr_genType type
func NewMoqLookupAddr_genType(scene *moq.Scene, config *moq.Config) *MoqLookupAddr_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLookupAddr_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLookupAddr_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr moq.ParamIndexing
		}{
			Addr: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the LookupAddr_genType type
func (m *MoqLookupAddr_genType) Mock() LookupAddr_genType {
	return func(addr string) (_ []string, _ error) {
		m.Scene.T.Helper()
		moq := &MoqLookupAddr_genType_mock{Moq: m}
		return moq.Fn(addr)
	}
}

func (m *MoqLookupAddr_genType_mock) Fn(addr string) (names []string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqLookupAddr_genType_params{
		Addr: addr,
	}
	var results *MoqLookupAddr_genType_results
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
		result.DoFn(addr)
	}

	if result.Values != nil {
		names = result.Values.Names
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		names, err = result.DoReturnFn(addr)
	}
	return
}

func (m *MoqLookupAddr_genType) OnCall(addr string) *MoqLookupAddr_genType_fnRecorder {
	return &MoqLookupAddr_genType_fnRecorder{
		Params: MoqLookupAddr_genType_params{
			Addr: addr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLookupAddr_genType_fnRecorder) Any() *MoqLookupAddr_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLookupAddr_genType_anyParams{Recorder: r}
}

func (a *MoqLookupAddr_genType_anyParams) Addr() *MoqLookupAddr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLookupAddr_genType_fnRecorder) Seq() *MoqLookupAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLookupAddr_genType_fnRecorder) NoSeq() *MoqLookupAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLookupAddr_genType_fnRecorder) ReturnResults(names []string, err error) *MoqLookupAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Names []string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupAddr_genType_doFn
		DoReturnFn MoqLookupAddr_genType_doReturnFn
	}{
		Values: &struct {
			Names []string
			Err   error
		}{
			Names: names,
			Err:   err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLookupAddr_genType_fnRecorder) AndDo(fn MoqLookupAddr_genType_doFn) *MoqLookupAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLookupAddr_genType_fnRecorder) DoReturnResults(fn MoqLookupAddr_genType_doReturnFn) *MoqLookupAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Names []string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupAddr_genType_doFn
		DoReturnFn MoqLookupAddr_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLookupAddr_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLookupAddr_genType_resultsByParams
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
		results = &MoqLookupAddr_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLookupAddr_genType_paramsKey]*MoqLookupAddr_genType_results{},
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
		r.Results = &MoqLookupAddr_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLookupAddr_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLookupAddr_genType_fnRecorder {
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
					Names []string
					Err   error
				}
				Sequence   uint32
				DoFn       MoqLookupAddr_genType_doFn
				DoReturnFn MoqLookupAddr_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLookupAddr_genType) PrettyParams(params MoqLookupAddr_genType_params) string {
	return fmt.Sprintf("LookupAddr_genType(%#v)", params.Addr)
}

func (m *MoqLookupAddr_genType) ParamsKey(params MoqLookupAddr_genType_params, anyParams uint64) MoqLookupAddr_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed string
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	return MoqLookupAddr_genType_paramsKey{
		Params: struct{ Addr string }{
			Addr: addrUsed,
		},
		Hashes: struct{ Addr hash.Hash }{
			Addr: addrUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLookupAddr_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLookupAddr_genType) AssertExpectationsMet() {
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
