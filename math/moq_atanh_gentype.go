// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Atanh_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Atanh_genType func(x float64) float64

// MoqAtanh_genType holds the state of a moq of the Atanh_genType type
type MoqAtanh_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAtanh_genType_mock

	ResultsByParams []MoqAtanh_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqAtanh_genType_mock isolates the mock interface of the Atanh_genType type
type MoqAtanh_genType_mock struct {
	Moq *MoqAtanh_genType
}

// MoqAtanh_genType_params holds the params of the Atanh_genType type
type MoqAtanh_genType_params struct{ X float64 }

// MoqAtanh_genType_paramsKey holds the map key params of the Atanh_genType
// type
type MoqAtanh_genType_paramsKey struct {
	Params struct{ X float64 }
	Hashes struct{ X hash.Hash }
}

// MoqAtanh_genType_resultsByParams contains the results for a given set of
// parameters for the Atanh_genType type
type MoqAtanh_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAtanh_genType_paramsKey]*MoqAtanh_genType_results
}

// MoqAtanh_genType_doFn defines the type of function needed when calling AndDo
// for the Atanh_genType type
type MoqAtanh_genType_doFn func(x float64)

// MoqAtanh_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Atanh_genType type
type MoqAtanh_genType_doReturnFn func(x float64) float64

// MoqAtanh_genType_results holds the results of the Atanh_genType type
type MoqAtanh_genType_results struct {
	Params  MoqAtanh_genType_params
	Results []struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqAtanh_genType_doFn
		DoReturnFn MoqAtanh_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAtanh_genType_fnRecorder routes recorded function calls to the
// MoqAtanh_genType moq
type MoqAtanh_genType_fnRecorder struct {
	Params    MoqAtanh_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAtanh_genType_results
	Moq       *MoqAtanh_genType
}

// MoqAtanh_genType_anyParams isolates the any params functions of the
// Atanh_genType type
type MoqAtanh_genType_anyParams struct {
	Recorder *MoqAtanh_genType_fnRecorder
}

// NewMoqAtanh_genType creates a new moq of the Atanh_genType type
func NewMoqAtanh_genType(scene *moq.Scene, config *moq.Config) *MoqAtanh_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAtanh_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAtanh_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X moq.ParamIndexing
		}{
			X: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Atanh_genType type
func (m *MoqAtanh_genType) Mock() Atanh_genType {
	return func(x float64) float64 { m.Scene.T.Helper(); moq := &MoqAtanh_genType_mock{Moq: m}; return moq.Fn(x) }
}

func (m *MoqAtanh_genType_mock) Fn(x float64) (result1 float64) {
	m.Moq.Scene.T.Helper()
	params := MoqAtanh_genType_params{
		X: x,
	}
	var results *MoqAtanh_genType_results
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
		result.DoFn(x)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x)
	}
	return
}

func (m *MoqAtanh_genType) OnCall(x float64) *MoqAtanh_genType_fnRecorder {
	return &MoqAtanh_genType_fnRecorder{
		Params: MoqAtanh_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAtanh_genType_fnRecorder) Any() *MoqAtanh_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAtanh_genType_anyParams{Recorder: r}
}

func (a *MoqAtanh_genType_anyParams) X() *MoqAtanh_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqAtanh_genType_fnRecorder) Seq() *MoqAtanh_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAtanh_genType_fnRecorder) NoSeq() *MoqAtanh_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAtanh_genType_fnRecorder) ReturnResults(result1 float64) *MoqAtanh_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqAtanh_genType_doFn
		DoReturnFn MoqAtanh_genType_doReturnFn
	}{
		Values: &struct {
			Result1 float64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAtanh_genType_fnRecorder) AndDo(fn MoqAtanh_genType_doFn) *MoqAtanh_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAtanh_genType_fnRecorder) DoReturnResults(fn MoqAtanh_genType_doReturnFn) *MoqAtanh_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqAtanh_genType_doFn
		DoReturnFn MoqAtanh_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAtanh_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAtanh_genType_resultsByParams
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
		results = &MoqAtanh_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAtanh_genType_paramsKey]*MoqAtanh_genType_results{},
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
		r.Results = &MoqAtanh_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAtanh_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAtanh_genType_fnRecorder {
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
					Result1 float64
				}
				Sequence   uint32
				DoFn       MoqAtanh_genType_doFn
				DoReturnFn MoqAtanh_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAtanh_genType) PrettyParams(params MoqAtanh_genType_params) string {
	return fmt.Sprintf("Atanh_genType(%#v)", params.X)
}

func (m *MoqAtanh_genType) ParamsKey(params MoqAtanh_genType_params, anyParams uint64) MoqAtanh_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed float64
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqAtanh_genType_paramsKey{
		Params: struct{ X float64 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAtanh_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAtanh_genType) AssertExpectationsMet() {
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
