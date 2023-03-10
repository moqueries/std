// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cmplx

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Rect_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Rect_genType func(r, θ float64) complex128

// MoqRect_genType holds the state of a moq of the Rect_genType type
type MoqRect_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRect_genType_mock

	ResultsByParams []MoqRect_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
			Θ      moq.ParamIndexing
		}
	}
}

// MoqRect_genType_mock isolates the mock interface of the Rect_genType type
type MoqRect_genType_mock struct {
	Moq *MoqRect_genType
}

// MoqRect_genType_params holds the params of the Rect_genType type
type MoqRect_genType_params struct{ Param1, Θ float64 }

// MoqRect_genType_paramsKey holds the map key params of the Rect_genType type
type MoqRect_genType_paramsKey struct {
	Params struct{ Param1, Θ float64 }
	Hashes struct{ Param1, Θ hash.Hash }
}

// MoqRect_genType_resultsByParams contains the results for a given set of
// parameters for the Rect_genType type
type MoqRect_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRect_genType_paramsKey]*MoqRect_genType_results
}

// MoqRect_genType_doFn defines the type of function needed when calling AndDo
// for the Rect_genType type
type MoqRect_genType_doFn func(r, θ float64)

// MoqRect_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Rect_genType type
type MoqRect_genType_doReturnFn func(r, θ float64) complex128

// MoqRect_genType_results holds the results of the Rect_genType type
type MoqRect_genType_results struct {
	Params  MoqRect_genType_params
	Results []struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqRect_genType_doFn
		DoReturnFn MoqRect_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRect_genType_fnRecorder routes recorded function calls to the
// MoqRect_genType moq
type MoqRect_genType_fnRecorder struct {
	Params    MoqRect_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRect_genType_results
	Moq       *MoqRect_genType
}

// MoqRect_genType_anyParams isolates the any params functions of the
// Rect_genType type
type MoqRect_genType_anyParams struct {
	Recorder *MoqRect_genType_fnRecorder
}

// NewMoqRect_genType creates a new moq of the Rect_genType type
func NewMoqRect_genType(scene *moq.Scene, config *moq.Config) *MoqRect_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRect_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRect_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
				Θ      moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
			Θ      moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByValue,
			Θ:      moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Rect_genType type
func (m *MoqRect_genType) Mock() Rect_genType {
	return func(param1, θ float64) complex128 {
		m.Scene.T.Helper()
		moq := &MoqRect_genType_mock{Moq: m}
		return moq.Fn(param1, θ)
	}
}

func (m *MoqRect_genType_mock) Fn(param1, θ float64) (result1 complex128) {
	m.Moq.Scene.T.Helper()
	params := MoqRect_genType_params{
		Param1: param1,
		Θ:      θ,
	}
	var results *MoqRect_genType_results
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
		result.DoFn(param1, θ)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1, θ)
	}
	return
}

func (m *MoqRect_genType) OnCall(param1, θ float64) *MoqRect_genType_fnRecorder {
	return &MoqRect_genType_fnRecorder{
		Params: MoqRect_genType_params{
			Param1: param1,
			Θ:      θ,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRect_genType_fnRecorder) Any() *MoqRect_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRect_genType_anyParams{Recorder: r}
}

func (a *MoqRect_genType_anyParams) Param1() *MoqRect_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqRect_genType_anyParams) Θ() *MoqRect_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqRect_genType_fnRecorder) Seq() *MoqRect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRect_genType_fnRecorder) NoSeq() *MoqRect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRect_genType_fnRecorder) ReturnResults(result1 complex128) *MoqRect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqRect_genType_doFn
		DoReturnFn MoqRect_genType_doReturnFn
	}{
		Values: &struct {
			Result1 complex128
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRect_genType_fnRecorder) AndDo(fn MoqRect_genType_doFn) *MoqRect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRect_genType_fnRecorder) DoReturnResults(fn MoqRect_genType_doReturnFn) *MoqRect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqRect_genType_doFn
		DoReturnFn MoqRect_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRect_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRect_genType_resultsByParams
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
		results = &MoqRect_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRect_genType_paramsKey]*MoqRect_genType_results{},
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
		r.Results = &MoqRect_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRect_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRect_genType_fnRecorder {
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
					Result1 complex128
				}
				Sequence   uint32
				DoFn       MoqRect_genType_doFn
				DoReturnFn MoqRect_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRect_genType) PrettyParams(params MoqRect_genType_params) string {
	return fmt.Sprintf("Rect_genType(%#v, %#v)", params.Param1, params.Θ)
}

func (m *MoqRect_genType) ParamsKey(params MoqRect_genType_params, anyParams uint64) MoqRect_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used float64
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	var θUsed float64
	var θUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Θ == moq.ParamIndexByValue {
			θUsed = params.Θ
		} else {
			θUsedHash = hash.DeepHash(params.Θ)
		}
	}
	return MoqRect_genType_paramsKey{
		Params: struct{ Param1, Θ float64 }{
			Param1: param1Used,
			Θ:      θUsed,
		},
		Hashes: struct{ Param1, Θ hash.Hash }{
			Param1: param1UsedHash,
			Θ:      θUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRect_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRect_genType) AssertExpectationsMet() {
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
