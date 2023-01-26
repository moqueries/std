// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bits

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ReverseBytes32_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ReverseBytes32_genType func(x uint32) uint32

// MoqReverseBytes32_genType holds the state of a moq of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReverseBytes32_genType_mock

	ResultsByParams []MoqReverseBytes32_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqReverseBytes32_genType_mock isolates the mock interface of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType_mock struct {
	Moq *MoqReverseBytes32_genType
}

// MoqReverseBytes32_genType_params holds the params of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType_params struct{ X uint32 }

// MoqReverseBytes32_genType_paramsKey holds the map key params of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType_paramsKey struct {
	Params struct{ X uint32 }
	Hashes struct{ X hash.Hash }
}

// MoqReverseBytes32_genType_resultsByParams contains the results for a given
// set of parameters for the ReverseBytes32_genType type
type MoqReverseBytes32_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReverseBytes32_genType_paramsKey]*MoqReverseBytes32_genType_results
}

// MoqReverseBytes32_genType_doFn defines the type of function needed when
// calling AndDo for the ReverseBytes32_genType type
type MoqReverseBytes32_genType_doFn func(x uint32)

// MoqReverseBytes32_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ReverseBytes32_genType type
type MoqReverseBytes32_genType_doReturnFn func(x uint32) uint32

// MoqReverseBytes32_genType_results holds the results of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType_results struct {
	Params  MoqReverseBytes32_genType_params
	Results []struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqReverseBytes32_genType_doFn
		DoReturnFn MoqReverseBytes32_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReverseBytes32_genType_fnRecorder routes recorded function calls to the
// MoqReverseBytes32_genType moq
type MoqReverseBytes32_genType_fnRecorder struct {
	Params    MoqReverseBytes32_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReverseBytes32_genType_results
	Moq       *MoqReverseBytes32_genType
}

// MoqReverseBytes32_genType_anyParams isolates the any params functions of the
// ReverseBytes32_genType type
type MoqReverseBytes32_genType_anyParams struct {
	Recorder *MoqReverseBytes32_genType_fnRecorder
}

// NewMoqReverseBytes32_genType creates a new moq of the ReverseBytes32_genType
// type
func NewMoqReverseBytes32_genType(scene *moq.Scene, config *moq.Config) *MoqReverseBytes32_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReverseBytes32_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReverseBytes32_genType_mock{},

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

// Mock returns the moq implementation of the ReverseBytes32_genType type
func (m *MoqReverseBytes32_genType) Mock() ReverseBytes32_genType {
	return func(x uint32) uint32 {
		m.Scene.T.Helper()
		moq := &MoqReverseBytes32_genType_mock{Moq: m}
		return moq.Fn(x)
	}
}

func (m *MoqReverseBytes32_genType_mock) Fn(x uint32) (result1 uint32) {
	m.Moq.Scene.T.Helper()
	params := MoqReverseBytes32_genType_params{
		X: x,
	}
	var results *MoqReverseBytes32_genType_results
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

func (m *MoqReverseBytes32_genType) OnCall(x uint32) *MoqReverseBytes32_genType_fnRecorder {
	return &MoqReverseBytes32_genType_fnRecorder{
		Params: MoqReverseBytes32_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReverseBytes32_genType_fnRecorder) Any() *MoqReverseBytes32_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReverseBytes32_genType_anyParams{Recorder: r}
}

func (a *MoqReverseBytes32_genType_anyParams) X() *MoqReverseBytes32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReverseBytes32_genType_fnRecorder) Seq() *MoqReverseBytes32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReverseBytes32_genType_fnRecorder) NoSeq() *MoqReverseBytes32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReverseBytes32_genType_fnRecorder) ReturnResults(result1 uint32) *MoqReverseBytes32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqReverseBytes32_genType_doFn
		DoReturnFn MoqReverseBytes32_genType_doReturnFn
	}{
		Values: &struct {
			Result1 uint32
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReverseBytes32_genType_fnRecorder) AndDo(fn MoqReverseBytes32_genType_doFn) *MoqReverseBytes32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReverseBytes32_genType_fnRecorder) DoReturnResults(fn MoqReverseBytes32_genType_doReturnFn) *MoqReverseBytes32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqReverseBytes32_genType_doFn
		DoReturnFn MoqReverseBytes32_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReverseBytes32_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReverseBytes32_genType_resultsByParams
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
		results = &MoqReverseBytes32_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReverseBytes32_genType_paramsKey]*MoqReverseBytes32_genType_results{},
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
		r.Results = &MoqReverseBytes32_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReverseBytes32_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReverseBytes32_genType_fnRecorder {
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
					Result1 uint32
				}
				Sequence   uint32
				DoFn       MoqReverseBytes32_genType_doFn
				DoReturnFn MoqReverseBytes32_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReverseBytes32_genType) PrettyParams(params MoqReverseBytes32_genType_params) string {
	return fmt.Sprintf("ReverseBytes32_genType(%#v)", params.X)
}

func (m *MoqReverseBytes32_genType) ParamsKey(params MoqReverseBytes32_genType_params, anyParams uint64) MoqReverseBytes32_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed uint32
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqReverseBytes32_genType_paramsKey{
		Params: struct{ X uint32 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReverseBytes32_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReverseBytes32_genType) AssertExpectationsMet() {
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
