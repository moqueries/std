// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package subtle

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ConstantTimeEq_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ConstantTimeEq_genType func(x, y int32) int

// MoqConstantTimeEq_genType holds the state of a moq of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqConstantTimeEq_genType_mock

	ResultsByParams []MoqConstantTimeEq_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
			Y moq.ParamIndexing
		}
	}
}

// MoqConstantTimeEq_genType_mock isolates the mock interface of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_mock struct {
	Moq *MoqConstantTimeEq_genType
}

// MoqConstantTimeEq_genType_params holds the params of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_params struct{ X, Y int32 }

// MoqConstantTimeEq_genType_paramsKey holds the map key params of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_paramsKey struct {
	Params struct{ X, Y int32 }
	Hashes struct{ X, Y hash.Hash }
}

// MoqConstantTimeEq_genType_resultsByParams contains the results for a given
// set of parameters for the ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqConstantTimeEq_genType_paramsKey]*MoqConstantTimeEq_genType_results
}

// MoqConstantTimeEq_genType_doFn defines the type of function needed when
// calling AndDo for the ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_doFn func(x, y int32)

// MoqConstantTimeEq_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_doReturnFn func(x, y int32) int

// MoqConstantTimeEq_genType_results holds the results of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_results struct {
	Params  MoqConstantTimeEq_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqConstantTimeEq_genType_doFn
		DoReturnFn MoqConstantTimeEq_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqConstantTimeEq_genType_fnRecorder routes recorded function calls to the
// MoqConstantTimeEq_genType moq
type MoqConstantTimeEq_genType_fnRecorder struct {
	Params    MoqConstantTimeEq_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqConstantTimeEq_genType_results
	Moq       *MoqConstantTimeEq_genType
}

// MoqConstantTimeEq_genType_anyParams isolates the any params functions of the
// ConstantTimeEq_genType type
type MoqConstantTimeEq_genType_anyParams struct {
	Recorder *MoqConstantTimeEq_genType_fnRecorder
}

// NewMoqConstantTimeEq_genType creates a new moq of the ConstantTimeEq_genType
// type
func NewMoqConstantTimeEq_genType(scene *moq.Scene, config *moq.Config) *MoqConstantTimeEq_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqConstantTimeEq_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqConstantTimeEq_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X moq.ParamIndexing
			Y moq.ParamIndexing
		}{
			X: moq.ParamIndexByValue,
			Y: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ConstantTimeEq_genType type
func (m *MoqConstantTimeEq_genType) Mock() ConstantTimeEq_genType {
	return func(x, y int32) int {
		m.Scene.T.Helper()
		moq := &MoqConstantTimeEq_genType_mock{Moq: m}
		return moq.Fn(x, y)
	}
}

func (m *MoqConstantTimeEq_genType_mock) Fn(x, y int32) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqConstantTimeEq_genType_params{
		X: x,
		Y: y,
	}
	var results *MoqConstantTimeEq_genType_results
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
		result.DoFn(x, y)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x, y)
	}
	return
}

func (m *MoqConstantTimeEq_genType) OnCall(x, y int32) *MoqConstantTimeEq_genType_fnRecorder {
	return &MoqConstantTimeEq_genType_fnRecorder{
		Params: MoqConstantTimeEq_genType_params{
			X: x,
			Y: y,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqConstantTimeEq_genType_fnRecorder) Any() *MoqConstantTimeEq_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqConstantTimeEq_genType_anyParams{Recorder: r}
}

func (a *MoqConstantTimeEq_genType_anyParams) X() *MoqConstantTimeEq_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqConstantTimeEq_genType_anyParams) Y() *MoqConstantTimeEq_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqConstantTimeEq_genType_fnRecorder) Seq() *MoqConstantTimeEq_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqConstantTimeEq_genType_fnRecorder) NoSeq() *MoqConstantTimeEq_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqConstantTimeEq_genType_fnRecorder) ReturnResults(result1 int) *MoqConstantTimeEq_genType_fnRecorder {
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
		DoFn       MoqConstantTimeEq_genType_doFn
		DoReturnFn MoqConstantTimeEq_genType_doReturnFn
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

func (r *MoqConstantTimeEq_genType_fnRecorder) AndDo(fn MoqConstantTimeEq_genType_doFn) *MoqConstantTimeEq_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqConstantTimeEq_genType_fnRecorder) DoReturnResults(fn MoqConstantTimeEq_genType_doReturnFn) *MoqConstantTimeEq_genType_fnRecorder {
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
		DoFn       MoqConstantTimeEq_genType_doFn
		DoReturnFn MoqConstantTimeEq_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqConstantTimeEq_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqConstantTimeEq_genType_resultsByParams
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
		results = &MoqConstantTimeEq_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqConstantTimeEq_genType_paramsKey]*MoqConstantTimeEq_genType_results{},
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
		r.Results = &MoqConstantTimeEq_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqConstantTimeEq_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqConstantTimeEq_genType_fnRecorder {
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
				DoFn       MoqConstantTimeEq_genType_doFn
				DoReturnFn MoqConstantTimeEq_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqConstantTimeEq_genType) PrettyParams(params MoqConstantTimeEq_genType_params) string {
	return fmt.Sprintf("ConstantTimeEq_genType(%#v, %#v)", params.X, params.Y)
}

func (m *MoqConstantTimeEq_genType) ParamsKey(params MoqConstantTimeEq_genType_params, anyParams uint64) MoqConstantTimeEq_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed int32
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	var yUsed int32
	var yUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	return MoqConstantTimeEq_genType_paramsKey{
		Params: struct{ X, Y int32 }{
			X: xUsed,
			Y: yUsed,
		},
		Hashes: struct{ X, Y hash.Hash }{
			X: xUsedHash,
			Y: yUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqConstantTimeEq_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqConstantTimeEq_genType) AssertExpectationsMet() {
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
