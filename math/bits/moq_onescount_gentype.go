// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bits

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// OnesCount_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type OnesCount_genType func(x uint) int

// MoqOnesCount_genType holds the state of a moq of the OnesCount_genType type
type MoqOnesCount_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqOnesCount_genType_mock

	ResultsByParams []MoqOnesCount_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqOnesCount_genType_mock isolates the mock interface of the
// OnesCount_genType type
type MoqOnesCount_genType_mock struct {
	Moq *MoqOnesCount_genType
}

// MoqOnesCount_genType_params holds the params of the OnesCount_genType type
type MoqOnesCount_genType_params struct{ X uint }

// MoqOnesCount_genType_paramsKey holds the map key params of the
// OnesCount_genType type
type MoqOnesCount_genType_paramsKey struct {
	Params struct{ X uint }
	Hashes struct{ X hash.Hash }
}

// MoqOnesCount_genType_resultsByParams contains the results for a given set of
// parameters for the OnesCount_genType type
type MoqOnesCount_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOnesCount_genType_paramsKey]*MoqOnesCount_genType_results
}

// MoqOnesCount_genType_doFn defines the type of function needed when calling
// AndDo for the OnesCount_genType type
type MoqOnesCount_genType_doFn func(x uint)

// MoqOnesCount_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the OnesCount_genType type
type MoqOnesCount_genType_doReturnFn func(x uint) int

// MoqOnesCount_genType_results holds the results of the OnesCount_genType type
type MoqOnesCount_genType_results struct {
	Params  MoqOnesCount_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqOnesCount_genType_doFn
		DoReturnFn MoqOnesCount_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOnesCount_genType_fnRecorder routes recorded function calls to the
// MoqOnesCount_genType moq
type MoqOnesCount_genType_fnRecorder struct {
	Params    MoqOnesCount_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOnesCount_genType_results
	Moq       *MoqOnesCount_genType
}

// MoqOnesCount_genType_anyParams isolates the any params functions of the
// OnesCount_genType type
type MoqOnesCount_genType_anyParams struct {
	Recorder *MoqOnesCount_genType_fnRecorder
}

// NewMoqOnesCount_genType creates a new moq of the OnesCount_genType type
func NewMoqOnesCount_genType(scene *moq.Scene, config *moq.Config) *MoqOnesCount_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqOnesCount_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqOnesCount_genType_mock{},

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

// Mock returns the moq implementation of the OnesCount_genType type
func (m *MoqOnesCount_genType) Mock() OnesCount_genType {
	return func(x uint) int { m.Scene.T.Helper(); moq := &MoqOnesCount_genType_mock{Moq: m}; return moq.Fn(x) }
}

func (m *MoqOnesCount_genType_mock) Fn(x uint) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqOnesCount_genType_params{
		X: x,
	}
	var results *MoqOnesCount_genType_results
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

func (m *MoqOnesCount_genType) OnCall(x uint) *MoqOnesCount_genType_fnRecorder {
	return &MoqOnesCount_genType_fnRecorder{
		Params: MoqOnesCount_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqOnesCount_genType_fnRecorder) Any() *MoqOnesCount_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqOnesCount_genType_anyParams{Recorder: r}
}

func (a *MoqOnesCount_genType_anyParams) X() *MoqOnesCount_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqOnesCount_genType_fnRecorder) Seq() *MoqOnesCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOnesCount_genType_fnRecorder) NoSeq() *MoqOnesCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOnesCount_genType_fnRecorder) ReturnResults(result1 int) *MoqOnesCount_genType_fnRecorder {
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
		DoFn       MoqOnesCount_genType_doFn
		DoReturnFn MoqOnesCount_genType_doReturnFn
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

func (r *MoqOnesCount_genType_fnRecorder) AndDo(fn MoqOnesCount_genType_doFn) *MoqOnesCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOnesCount_genType_fnRecorder) DoReturnResults(fn MoqOnesCount_genType_doReturnFn) *MoqOnesCount_genType_fnRecorder {
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
		DoFn       MoqOnesCount_genType_doFn
		DoReturnFn MoqOnesCount_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOnesCount_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOnesCount_genType_resultsByParams
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
		results = &MoqOnesCount_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOnesCount_genType_paramsKey]*MoqOnesCount_genType_results{},
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
		r.Results = &MoqOnesCount_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOnesCount_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOnesCount_genType_fnRecorder {
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
				DoFn       MoqOnesCount_genType_doFn
				DoReturnFn MoqOnesCount_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOnesCount_genType) PrettyParams(params MoqOnesCount_genType_params) string {
	return fmt.Sprintf("OnesCount_genType(%#v)", params.X)
}

func (m *MoqOnesCount_genType) ParamsKey(params MoqOnesCount_genType_params, anyParams uint64) MoqOnesCount_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed uint
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqOnesCount_genType_paramsKey{
		Params: struct{ X uint }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqOnesCount_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqOnesCount_genType) AssertExpectationsMet() {
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
