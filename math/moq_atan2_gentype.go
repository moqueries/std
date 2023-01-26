// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Atan2_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Atan2_genType func(y, x float64) float64

// MoqAtan2_genType holds the state of a moq of the Atan2_genType type
type MoqAtan2_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAtan2_genType_mock

	ResultsByParams []MoqAtan2_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Y moq.ParamIndexing
			X moq.ParamIndexing
		}
	}
}

// MoqAtan2_genType_mock isolates the mock interface of the Atan2_genType type
type MoqAtan2_genType_mock struct {
	Moq *MoqAtan2_genType
}

// MoqAtan2_genType_params holds the params of the Atan2_genType type
type MoqAtan2_genType_params struct{ Y, X float64 }

// MoqAtan2_genType_paramsKey holds the map key params of the Atan2_genType
// type
type MoqAtan2_genType_paramsKey struct {
	Params struct{ Y, X float64 }
	Hashes struct{ Y, X hash.Hash }
}

// MoqAtan2_genType_resultsByParams contains the results for a given set of
// parameters for the Atan2_genType type
type MoqAtan2_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAtan2_genType_paramsKey]*MoqAtan2_genType_results
}

// MoqAtan2_genType_doFn defines the type of function needed when calling AndDo
// for the Atan2_genType type
type MoqAtan2_genType_doFn func(y, x float64)

// MoqAtan2_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Atan2_genType type
type MoqAtan2_genType_doReturnFn func(y, x float64) float64

// MoqAtan2_genType_results holds the results of the Atan2_genType type
type MoqAtan2_genType_results struct {
	Params  MoqAtan2_genType_params
	Results []struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqAtan2_genType_doFn
		DoReturnFn MoqAtan2_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAtan2_genType_fnRecorder routes recorded function calls to the
// MoqAtan2_genType moq
type MoqAtan2_genType_fnRecorder struct {
	Params    MoqAtan2_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAtan2_genType_results
	Moq       *MoqAtan2_genType
}

// MoqAtan2_genType_anyParams isolates the any params functions of the
// Atan2_genType type
type MoqAtan2_genType_anyParams struct {
	Recorder *MoqAtan2_genType_fnRecorder
}

// NewMoqAtan2_genType creates a new moq of the Atan2_genType type
func NewMoqAtan2_genType(scene *moq.Scene, config *moq.Config) *MoqAtan2_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAtan2_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAtan2_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Y moq.ParamIndexing
				X moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Y moq.ParamIndexing
			X moq.ParamIndexing
		}{
			Y: moq.ParamIndexByValue,
			X: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Atan2_genType type
func (m *MoqAtan2_genType) Mock() Atan2_genType {
	return func(y, x float64) float64 {
		m.Scene.T.Helper()
		moq := &MoqAtan2_genType_mock{Moq: m}
		return moq.Fn(y, x)
	}
}

func (m *MoqAtan2_genType_mock) Fn(y, x float64) (result1 float64) {
	m.Moq.Scene.T.Helper()
	params := MoqAtan2_genType_params{
		Y: y,
		X: x,
	}
	var results *MoqAtan2_genType_results
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
		result.DoFn(y, x)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(y, x)
	}
	return
}

func (m *MoqAtan2_genType) OnCall(y, x float64) *MoqAtan2_genType_fnRecorder {
	return &MoqAtan2_genType_fnRecorder{
		Params: MoqAtan2_genType_params{
			Y: y,
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAtan2_genType_fnRecorder) Any() *MoqAtan2_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAtan2_genType_anyParams{Recorder: r}
}

func (a *MoqAtan2_genType_anyParams) Y() *MoqAtan2_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAtan2_genType_anyParams) X() *MoqAtan2_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqAtan2_genType_fnRecorder) Seq() *MoqAtan2_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAtan2_genType_fnRecorder) NoSeq() *MoqAtan2_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAtan2_genType_fnRecorder) ReturnResults(result1 float64) *MoqAtan2_genType_fnRecorder {
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
		DoFn       MoqAtan2_genType_doFn
		DoReturnFn MoqAtan2_genType_doReturnFn
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

func (r *MoqAtan2_genType_fnRecorder) AndDo(fn MoqAtan2_genType_doFn) *MoqAtan2_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAtan2_genType_fnRecorder) DoReturnResults(fn MoqAtan2_genType_doReturnFn) *MoqAtan2_genType_fnRecorder {
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
		DoFn       MoqAtan2_genType_doFn
		DoReturnFn MoqAtan2_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAtan2_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAtan2_genType_resultsByParams
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
		results = &MoqAtan2_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAtan2_genType_paramsKey]*MoqAtan2_genType_results{},
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
		r.Results = &MoqAtan2_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAtan2_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAtan2_genType_fnRecorder {
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
				DoFn       MoqAtan2_genType_doFn
				DoReturnFn MoqAtan2_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAtan2_genType) PrettyParams(params MoqAtan2_genType_params) string {
	return fmt.Sprintf("Atan2_genType(%#v, %#v)", params.Y, params.X)
}

func (m *MoqAtan2_genType) ParamsKey(params MoqAtan2_genType_params, anyParams uint64) MoqAtan2_genType_paramsKey {
	m.Scene.T.Helper()
	var yUsed float64
	var yUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	var xUsed float64
	var xUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqAtan2_genType_paramsKey{
		Params: struct{ Y, X float64 }{
			Y: yUsed,
			X: xUsed,
		},
		Hashes: struct{ Y, X hash.Hash }{
			Y: yUsedHash,
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAtan2_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAtan2_genType) AssertExpectationsMet() {
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