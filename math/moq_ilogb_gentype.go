// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Ilogb_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Ilogb_genType func(x float64) int

// MoqIlogb_genType holds the state of a moq of the Ilogb_genType type
type MoqIlogb_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIlogb_genType_mock

	ResultsByParams []MoqIlogb_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqIlogb_genType_mock isolates the mock interface of the Ilogb_genType type
type MoqIlogb_genType_mock struct {
	Moq *MoqIlogb_genType
}

// MoqIlogb_genType_params holds the params of the Ilogb_genType type
type MoqIlogb_genType_params struct{ X float64 }

// MoqIlogb_genType_paramsKey holds the map key params of the Ilogb_genType
// type
type MoqIlogb_genType_paramsKey struct {
	Params struct{ X float64 }
	Hashes struct{ X hash.Hash }
}

// MoqIlogb_genType_resultsByParams contains the results for a given set of
// parameters for the Ilogb_genType type
type MoqIlogb_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIlogb_genType_paramsKey]*MoqIlogb_genType_results
}

// MoqIlogb_genType_doFn defines the type of function needed when calling AndDo
// for the Ilogb_genType type
type MoqIlogb_genType_doFn func(x float64)

// MoqIlogb_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Ilogb_genType type
type MoqIlogb_genType_doReturnFn func(x float64) int

// MoqIlogb_genType_results holds the results of the Ilogb_genType type
type MoqIlogb_genType_results struct {
	Params  MoqIlogb_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqIlogb_genType_doFn
		DoReturnFn MoqIlogb_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIlogb_genType_fnRecorder routes recorded function calls to the
// MoqIlogb_genType moq
type MoqIlogb_genType_fnRecorder struct {
	Params    MoqIlogb_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIlogb_genType_results
	Moq       *MoqIlogb_genType
}

// MoqIlogb_genType_anyParams isolates the any params functions of the
// Ilogb_genType type
type MoqIlogb_genType_anyParams struct {
	Recorder *MoqIlogb_genType_fnRecorder
}

// NewMoqIlogb_genType creates a new moq of the Ilogb_genType type
func NewMoqIlogb_genType(scene *moq.Scene, config *moq.Config) *MoqIlogb_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIlogb_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIlogb_genType_mock{},

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

// Mock returns the moq implementation of the Ilogb_genType type
func (m *MoqIlogb_genType) Mock() Ilogb_genType {
	return func(x float64) int { m.Scene.T.Helper(); moq := &MoqIlogb_genType_mock{Moq: m}; return moq.Fn(x) }
}

func (m *MoqIlogb_genType_mock) Fn(x float64) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqIlogb_genType_params{
		X: x,
	}
	var results *MoqIlogb_genType_results
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

func (m *MoqIlogb_genType) OnCall(x float64) *MoqIlogb_genType_fnRecorder {
	return &MoqIlogb_genType_fnRecorder{
		Params: MoqIlogb_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIlogb_genType_fnRecorder) Any() *MoqIlogb_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIlogb_genType_anyParams{Recorder: r}
}

func (a *MoqIlogb_genType_anyParams) X() *MoqIlogb_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIlogb_genType_fnRecorder) Seq() *MoqIlogb_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIlogb_genType_fnRecorder) NoSeq() *MoqIlogb_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIlogb_genType_fnRecorder) ReturnResults(result1 int) *MoqIlogb_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqIlogb_genType_doFn
		DoReturnFn MoqIlogb_genType_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIlogb_genType_fnRecorder) AndDo(fn MoqIlogb_genType_doFn) *MoqIlogb_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIlogb_genType_fnRecorder) DoReturnResults(fn MoqIlogb_genType_doReturnFn) *MoqIlogb_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqIlogb_genType_doFn
		DoReturnFn MoqIlogb_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIlogb_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIlogb_genType_resultsByParams
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
		results = &MoqIlogb_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIlogb_genType_paramsKey]*MoqIlogb_genType_results{},
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
		r.Results = &MoqIlogb_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIlogb_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIlogb_genType_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqIlogb_genType_doFn
				DoReturnFn MoqIlogb_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIlogb_genType) PrettyParams(params MoqIlogb_genType_params) string {
	return fmt.Sprintf("Ilogb_genType(%#v)", params.X)
}

func (m *MoqIlogb_genType) ParamsKey(params MoqIlogb_genType_params, anyParams uint64) MoqIlogb_genType_paramsKey {
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
	return MoqIlogb_genType_paramsKey{
		Params: struct{ X float64 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIlogb_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIlogb_genType) AssertExpectationsMet() {
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