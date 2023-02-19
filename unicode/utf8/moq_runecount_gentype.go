// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package utf8

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// RuneCount_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type RuneCount_genType func(p []byte) int

// MoqRuneCount_genType holds the state of a moq of the RuneCount_genType type
type MoqRuneCount_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRuneCount_genType_mock

	ResultsByParams []MoqRuneCount_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			P moq.ParamIndexing
		}
	}
}

// MoqRuneCount_genType_mock isolates the mock interface of the
// RuneCount_genType type
type MoqRuneCount_genType_mock struct {
	Moq *MoqRuneCount_genType
}

// MoqRuneCount_genType_params holds the params of the RuneCount_genType type
type MoqRuneCount_genType_params struct{ P []byte }

// MoqRuneCount_genType_paramsKey holds the map key params of the
// RuneCount_genType type
type MoqRuneCount_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqRuneCount_genType_resultsByParams contains the results for a given set of
// parameters for the RuneCount_genType type
type MoqRuneCount_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRuneCount_genType_paramsKey]*MoqRuneCount_genType_results
}

// MoqRuneCount_genType_doFn defines the type of function needed when calling
// AndDo for the RuneCount_genType type
type MoqRuneCount_genType_doFn func(p []byte)

// MoqRuneCount_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the RuneCount_genType type
type MoqRuneCount_genType_doReturnFn func(p []byte) int

// MoqRuneCount_genType_results holds the results of the RuneCount_genType type
type MoqRuneCount_genType_results struct {
	Params  MoqRuneCount_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqRuneCount_genType_doFn
		DoReturnFn MoqRuneCount_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRuneCount_genType_fnRecorder routes recorded function calls to the
// MoqRuneCount_genType moq
type MoqRuneCount_genType_fnRecorder struct {
	Params    MoqRuneCount_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRuneCount_genType_results
	Moq       *MoqRuneCount_genType
}

// MoqRuneCount_genType_anyParams isolates the any params functions of the
// RuneCount_genType type
type MoqRuneCount_genType_anyParams struct {
	Recorder *MoqRuneCount_genType_fnRecorder
}

// NewMoqRuneCount_genType creates a new moq of the RuneCount_genType type
func NewMoqRuneCount_genType(scene *moq.Scene, config *moq.Config) *MoqRuneCount_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRuneCount_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRuneCount_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				P moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			P moq.ParamIndexing
		}{
			P: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the RuneCount_genType type
func (m *MoqRuneCount_genType) Mock() RuneCount_genType {
	return func(p []byte) int { m.Scene.T.Helper(); moq := &MoqRuneCount_genType_mock{Moq: m}; return moq.Fn(p) }
}

func (m *MoqRuneCount_genType_mock) Fn(p []byte) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqRuneCount_genType_params{
		P: p,
	}
	var results *MoqRuneCount_genType_results
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
		result.DoFn(p)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(p)
	}
	return
}

func (m *MoqRuneCount_genType) OnCall(p []byte) *MoqRuneCount_genType_fnRecorder {
	return &MoqRuneCount_genType_fnRecorder{
		Params: MoqRuneCount_genType_params{
			P: p,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRuneCount_genType_fnRecorder) Any() *MoqRuneCount_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRuneCount_genType_anyParams{Recorder: r}
}

func (a *MoqRuneCount_genType_anyParams) P() *MoqRuneCount_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqRuneCount_genType_fnRecorder) Seq() *MoqRuneCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRuneCount_genType_fnRecorder) NoSeq() *MoqRuneCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRuneCount_genType_fnRecorder) ReturnResults(result1 int) *MoqRuneCount_genType_fnRecorder {
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
		DoFn       MoqRuneCount_genType_doFn
		DoReturnFn MoqRuneCount_genType_doReturnFn
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

func (r *MoqRuneCount_genType_fnRecorder) AndDo(fn MoqRuneCount_genType_doFn) *MoqRuneCount_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRuneCount_genType_fnRecorder) DoReturnResults(fn MoqRuneCount_genType_doReturnFn) *MoqRuneCount_genType_fnRecorder {
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
		DoFn       MoqRuneCount_genType_doFn
		DoReturnFn MoqRuneCount_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRuneCount_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRuneCount_genType_resultsByParams
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
		results = &MoqRuneCount_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRuneCount_genType_paramsKey]*MoqRuneCount_genType_results{},
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
		r.Results = &MoqRuneCount_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRuneCount_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRuneCount_genType_fnRecorder {
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
				DoFn       MoqRuneCount_genType_doFn
				DoReturnFn MoqRuneCount_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRuneCount_genType) PrettyParams(params MoqRuneCount_genType_params) string {
	return fmt.Sprintf("RuneCount_genType(%#v)", params.P)
}

func (m *MoqRuneCount_genType) ParamsKey(params MoqRuneCount_genType_params, anyParams uint64) MoqRuneCount_genType_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqRuneCount_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRuneCount_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRuneCount_genType) AssertExpectationsMet() {
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