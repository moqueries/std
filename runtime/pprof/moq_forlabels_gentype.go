// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pprof

import (
	"context"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ForLabels_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ForLabels_genType func(ctx context.Context, f func(key, value string) bool)

// MoqForLabels_genType holds the state of a moq of the ForLabels_genType type
type MoqForLabels_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqForLabels_genType_mock

	ResultsByParams []MoqForLabels_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ctx moq.ParamIndexing
			F   moq.ParamIndexing
		}
	}
}

// MoqForLabels_genType_mock isolates the mock interface of the
// ForLabels_genType type
type MoqForLabels_genType_mock struct {
	Moq *MoqForLabels_genType
}

// MoqForLabels_genType_params holds the params of the ForLabels_genType type
type MoqForLabels_genType_params struct {
	Ctx context.Context
	F   func(key, value string) bool
}

// MoqForLabels_genType_paramsKey holds the map key params of the
// ForLabels_genType type
type MoqForLabels_genType_paramsKey struct {
	Params struct{ Ctx context.Context }
	Hashes struct {
		Ctx hash.Hash
		F   hash.Hash
	}
}

// MoqForLabels_genType_resultsByParams contains the results for a given set of
// parameters for the ForLabels_genType type
type MoqForLabels_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqForLabels_genType_paramsKey]*MoqForLabels_genType_results
}

// MoqForLabels_genType_doFn defines the type of function needed when calling
// AndDo for the ForLabels_genType type
type MoqForLabels_genType_doFn func(ctx context.Context, f func(key, value string) bool)

// MoqForLabels_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ForLabels_genType type
type MoqForLabels_genType_doReturnFn func(ctx context.Context, f func(key, value string) bool)

// MoqForLabels_genType_results holds the results of the ForLabels_genType type
type MoqForLabels_genType_results struct {
	Params  MoqForLabels_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqForLabels_genType_doFn
		DoReturnFn MoqForLabels_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqForLabels_genType_fnRecorder routes recorded function calls to the
// MoqForLabels_genType moq
type MoqForLabels_genType_fnRecorder struct {
	Params    MoqForLabels_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqForLabels_genType_results
	Moq       *MoqForLabels_genType
}

// MoqForLabels_genType_anyParams isolates the any params functions of the
// ForLabels_genType type
type MoqForLabels_genType_anyParams struct {
	Recorder *MoqForLabels_genType_fnRecorder
}

// NewMoqForLabels_genType creates a new moq of the ForLabels_genType type
func NewMoqForLabels_genType(scene *moq.Scene, config *moq.Config) *MoqForLabels_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqForLabels_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqForLabels_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ctx moq.ParamIndexing
				F   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ctx moq.ParamIndexing
			F   moq.ParamIndexing
		}{
			Ctx: moq.ParamIndexByHash,
			F:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ForLabels_genType type
func (m *MoqForLabels_genType) Mock() ForLabels_genType {
	return func(ctx context.Context, f func(key, value string) bool) {
		m.Scene.T.Helper()
		moq := &MoqForLabels_genType_mock{Moq: m}
		moq.Fn(ctx, f)
	}
}

func (m *MoqForLabels_genType_mock) Fn(ctx context.Context, f func(key, value string) bool) {
	m.Moq.Scene.T.Helper()
	params := MoqForLabels_genType_params{
		Ctx: ctx,
		F:   f,
	}
	var results *MoqForLabels_genType_results
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
		result.DoFn(ctx, f)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(ctx, f)
	}
	return
}

func (m *MoqForLabels_genType) OnCall(ctx context.Context, f func(key, value string) bool) *MoqForLabels_genType_fnRecorder {
	return &MoqForLabels_genType_fnRecorder{
		Params: MoqForLabels_genType_params{
			Ctx: ctx,
			F:   f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqForLabels_genType_fnRecorder) Any() *MoqForLabels_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqForLabels_genType_anyParams{Recorder: r}
}

func (a *MoqForLabels_genType_anyParams) Ctx() *MoqForLabels_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqForLabels_genType_anyParams) F() *MoqForLabels_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqForLabels_genType_fnRecorder) Seq() *MoqForLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqForLabels_genType_fnRecorder) NoSeq() *MoqForLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqForLabels_genType_fnRecorder) ReturnResults() *MoqForLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqForLabels_genType_doFn
		DoReturnFn MoqForLabels_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqForLabels_genType_fnRecorder) AndDo(fn MoqForLabels_genType_doFn) *MoqForLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqForLabels_genType_fnRecorder) DoReturnResults(fn MoqForLabels_genType_doReturnFn) *MoqForLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqForLabels_genType_doFn
		DoReturnFn MoqForLabels_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqForLabels_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqForLabels_genType_resultsByParams
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
		results = &MoqForLabels_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqForLabels_genType_paramsKey]*MoqForLabels_genType_results{},
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
		r.Results = &MoqForLabels_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqForLabels_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqForLabels_genType_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqForLabels_genType_doFn
				DoReturnFn MoqForLabels_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqForLabels_genType) PrettyParams(params MoqForLabels_genType_params) string {
	return fmt.Sprintf("ForLabels_genType(%#v, %#v)", params.Ctx, moq.FnString(params.F))
}

func (m *MoqForLabels_genType) ParamsKey(params MoqForLabels_genType_params, anyParams uint64) MoqForLabels_genType_paramsKey {
	m.Scene.T.Helper()
	var ctxUsed context.Context
	var ctxUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Ctx == moq.ParamIndexByValue {
			ctxUsed = params.Ctx
		} else {
			ctxUsedHash = hash.DeepHash(params.Ctx)
		}
	}
	var fUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The f parameter can't be indexed by value")
		}
		fUsedHash = hash.DeepHash(params.F)
	}
	return MoqForLabels_genType_paramsKey{
		Params: struct{ Ctx context.Context }{
			Ctx: ctxUsed,
		},
		Hashes: struct {
			Ctx hash.Hash
			F   hash.Hash
		}{
			Ctx: ctxUsedHash,
			F:   fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqForLabels_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqForLabels_genType) AssertExpectationsMet() {
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
