// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"runtime"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ThreadCreateProfile_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type ThreadCreateProfile_genType func(p []runtime.StackRecord) (n int, ok bool)

// MoqThreadCreateProfile_genType holds the state of a moq of the
// ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqThreadCreateProfile_genType_mock

	ResultsByParams []MoqThreadCreateProfile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			P moq.ParamIndexing
		}
	}
}

// MoqThreadCreateProfile_genType_mock isolates the mock interface of the
// ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_mock struct {
	Moq *MoqThreadCreateProfile_genType
}

// MoqThreadCreateProfile_genType_params holds the params of the
// ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_params struct{ P []runtime.StackRecord }

// MoqThreadCreateProfile_genType_paramsKey holds the map key params of the
// ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqThreadCreateProfile_genType_resultsByParams contains the results for a
// given set of parameters for the ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqThreadCreateProfile_genType_paramsKey]*MoqThreadCreateProfile_genType_results
}

// MoqThreadCreateProfile_genType_doFn defines the type of function needed when
// calling AndDo for the ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_doFn func(p []runtime.StackRecord)

// MoqThreadCreateProfile_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_doReturnFn func(p []runtime.StackRecord) (n int, ok bool)

// MoqThreadCreateProfile_genType_results holds the results of the
// ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_results struct {
	Params  MoqThreadCreateProfile_genType_params
	Results []struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqThreadCreateProfile_genType_doFn
		DoReturnFn MoqThreadCreateProfile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqThreadCreateProfile_genType_fnRecorder routes recorded function calls to
// the MoqThreadCreateProfile_genType moq
type MoqThreadCreateProfile_genType_fnRecorder struct {
	Params    MoqThreadCreateProfile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqThreadCreateProfile_genType_results
	Moq       *MoqThreadCreateProfile_genType
}

// MoqThreadCreateProfile_genType_anyParams isolates the any params functions
// of the ThreadCreateProfile_genType type
type MoqThreadCreateProfile_genType_anyParams struct {
	Recorder *MoqThreadCreateProfile_genType_fnRecorder
}

// NewMoqThreadCreateProfile_genType creates a new moq of the
// ThreadCreateProfile_genType type
func NewMoqThreadCreateProfile_genType(scene *moq.Scene, config *moq.Config) *MoqThreadCreateProfile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqThreadCreateProfile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqThreadCreateProfile_genType_mock{},

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

// Mock returns the moq implementation of the ThreadCreateProfile_genType type
func (m *MoqThreadCreateProfile_genType) Mock() ThreadCreateProfile_genType {
	return func(p []runtime.StackRecord) (_ int, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqThreadCreateProfile_genType_mock{Moq: m}
		return moq.Fn(p)
	}
}

func (m *MoqThreadCreateProfile_genType_mock) Fn(p []runtime.StackRecord) (n int, ok bool) {
	m.Moq.Scene.T.Helper()
	params := MoqThreadCreateProfile_genType_params{
		P: p,
	}
	var results *MoqThreadCreateProfile_genType_results
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
		n = result.Values.N
		ok = result.Values.Ok
	}
	if result.DoReturnFn != nil {
		n, ok = result.DoReturnFn(p)
	}
	return
}

func (m *MoqThreadCreateProfile_genType) OnCall(p []runtime.StackRecord) *MoqThreadCreateProfile_genType_fnRecorder {
	return &MoqThreadCreateProfile_genType_fnRecorder{
		Params: MoqThreadCreateProfile_genType_params{
			P: p,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) Any() *MoqThreadCreateProfile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqThreadCreateProfile_genType_anyParams{Recorder: r}
}

func (a *MoqThreadCreateProfile_genType_anyParams) P() *MoqThreadCreateProfile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) Seq() *MoqThreadCreateProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) NoSeq() *MoqThreadCreateProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) ReturnResults(n int, ok bool) *MoqThreadCreateProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqThreadCreateProfile_genType_doFn
		DoReturnFn MoqThreadCreateProfile_genType_doReturnFn
	}{
		Values: &struct {
			N  int
			Ok bool
		}{
			N:  n,
			Ok: ok,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) AndDo(fn MoqThreadCreateProfile_genType_doFn) *MoqThreadCreateProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) DoReturnResults(fn MoqThreadCreateProfile_genType_doReturnFn) *MoqThreadCreateProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqThreadCreateProfile_genType_doFn
		DoReturnFn MoqThreadCreateProfile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqThreadCreateProfile_genType_resultsByParams
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
		results = &MoqThreadCreateProfile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqThreadCreateProfile_genType_paramsKey]*MoqThreadCreateProfile_genType_results{},
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
		r.Results = &MoqThreadCreateProfile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqThreadCreateProfile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqThreadCreateProfile_genType_fnRecorder {
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
					N  int
					Ok bool
				}
				Sequence   uint32
				DoFn       MoqThreadCreateProfile_genType_doFn
				DoReturnFn MoqThreadCreateProfile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqThreadCreateProfile_genType) PrettyParams(params MoqThreadCreateProfile_genType_params) string {
	return fmt.Sprintf("ThreadCreateProfile_genType(%#v)", params.P)
}

func (m *MoqThreadCreateProfile_genType) ParamsKey(params MoqThreadCreateProfile_genType_params, anyParams uint64) MoqThreadCreateProfile_genType_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqThreadCreateProfile_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqThreadCreateProfile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqThreadCreateProfile_genType) AssertExpectationsMet() {
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
