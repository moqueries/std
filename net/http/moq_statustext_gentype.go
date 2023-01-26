// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// StatusText_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type StatusText_genType func(code int) string

// MoqStatusText_genType holds the state of a moq of the StatusText_genType
// type
type MoqStatusText_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStatusText_genType_mock

	ResultsByParams []MoqStatusText_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Code moq.ParamIndexing
		}
	}
}

// MoqStatusText_genType_mock isolates the mock interface of the
// StatusText_genType type
type MoqStatusText_genType_mock struct {
	Moq *MoqStatusText_genType
}

// MoqStatusText_genType_params holds the params of the StatusText_genType type
type MoqStatusText_genType_params struct{ Code int }

// MoqStatusText_genType_paramsKey holds the map key params of the
// StatusText_genType type
type MoqStatusText_genType_paramsKey struct {
	Params struct{ Code int }
	Hashes struct{ Code hash.Hash }
}

// MoqStatusText_genType_resultsByParams contains the results for a given set
// of parameters for the StatusText_genType type
type MoqStatusText_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStatusText_genType_paramsKey]*MoqStatusText_genType_results
}

// MoqStatusText_genType_doFn defines the type of function needed when calling
// AndDo for the StatusText_genType type
type MoqStatusText_genType_doFn func(code int)

// MoqStatusText_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the StatusText_genType type
type MoqStatusText_genType_doReturnFn func(code int) string

// MoqStatusText_genType_results holds the results of the StatusText_genType
// type
type MoqStatusText_genType_results struct {
	Params  MoqStatusText_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqStatusText_genType_doFn
		DoReturnFn MoqStatusText_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStatusText_genType_fnRecorder routes recorded function calls to the
// MoqStatusText_genType moq
type MoqStatusText_genType_fnRecorder struct {
	Params    MoqStatusText_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStatusText_genType_results
	Moq       *MoqStatusText_genType
}

// MoqStatusText_genType_anyParams isolates the any params functions of the
// StatusText_genType type
type MoqStatusText_genType_anyParams struct {
	Recorder *MoqStatusText_genType_fnRecorder
}

// NewMoqStatusText_genType creates a new moq of the StatusText_genType type
func NewMoqStatusText_genType(scene *moq.Scene, config *moq.Config) *MoqStatusText_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStatusText_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStatusText_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Code moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Code moq.ParamIndexing
		}{
			Code: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the StatusText_genType type
func (m *MoqStatusText_genType) Mock() StatusText_genType {
	return func(code int) string {
		m.Scene.T.Helper()
		moq := &MoqStatusText_genType_mock{Moq: m}
		return moq.Fn(code)
	}
}

func (m *MoqStatusText_genType_mock) Fn(code int) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqStatusText_genType_params{
		Code: code,
	}
	var results *MoqStatusText_genType_results
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
		result.DoFn(code)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(code)
	}
	return
}

func (m *MoqStatusText_genType) OnCall(code int) *MoqStatusText_genType_fnRecorder {
	return &MoqStatusText_genType_fnRecorder{
		Params: MoqStatusText_genType_params{
			Code: code,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStatusText_genType_fnRecorder) Any() *MoqStatusText_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStatusText_genType_anyParams{Recorder: r}
}

func (a *MoqStatusText_genType_anyParams) Code() *MoqStatusText_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqStatusText_genType_fnRecorder) Seq() *MoqStatusText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStatusText_genType_fnRecorder) NoSeq() *MoqStatusText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStatusText_genType_fnRecorder) ReturnResults(result1 string) *MoqStatusText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqStatusText_genType_doFn
		DoReturnFn MoqStatusText_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStatusText_genType_fnRecorder) AndDo(fn MoqStatusText_genType_doFn) *MoqStatusText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStatusText_genType_fnRecorder) DoReturnResults(fn MoqStatusText_genType_doReturnFn) *MoqStatusText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqStatusText_genType_doFn
		DoReturnFn MoqStatusText_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStatusText_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStatusText_genType_resultsByParams
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
		results = &MoqStatusText_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStatusText_genType_paramsKey]*MoqStatusText_genType_results{},
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
		r.Results = &MoqStatusText_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStatusText_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStatusText_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqStatusText_genType_doFn
				DoReturnFn MoqStatusText_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStatusText_genType) PrettyParams(params MoqStatusText_genType_params) string {
	return fmt.Sprintf("StatusText_genType(%#v)", params.Code)
}

func (m *MoqStatusText_genType) ParamsKey(params MoqStatusText_genType_params, anyParams uint64) MoqStatusText_genType_paramsKey {
	m.Scene.T.Helper()
	var codeUsed int
	var codeUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Code == moq.ParamIndexByValue {
			codeUsed = params.Code
		} else {
			codeUsedHash = hash.DeepHash(params.Code)
		}
	}
	return MoqStatusText_genType_paramsKey{
		Params: struct{ Code int }{
			Code: codeUsed,
		},
		Hashes: struct{ Code hash.Hash }{
			Code: codeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStatusText_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStatusText_genType) AssertExpectationsMet() {
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