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

// SetGoroutineLabels_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type SetGoroutineLabels_genType func(ctx context.Context)

// MoqSetGoroutineLabels_genType holds the state of a moq of the
// SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetGoroutineLabels_genType_mock

	ResultsByParams []MoqSetGoroutineLabels_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ctx moq.ParamIndexing
		}
	}
}

// MoqSetGoroutineLabels_genType_mock isolates the mock interface of the
// SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_mock struct {
	Moq *MoqSetGoroutineLabels_genType
}

// MoqSetGoroutineLabels_genType_params holds the params of the
// SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_params struct{ Ctx context.Context }

// MoqSetGoroutineLabels_genType_paramsKey holds the map key params of the
// SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_paramsKey struct {
	Params struct{ Ctx context.Context }
	Hashes struct{ Ctx hash.Hash }
}

// MoqSetGoroutineLabels_genType_resultsByParams contains the results for a
// given set of parameters for the SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetGoroutineLabels_genType_paramsKey]*MoqSetGoroutineLabels_genType_results
}

// MoqSetGoroutineLabels_genType_doFn defines the type of function needed when
// calling AndDo for the SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_doFn func(ctx context.Context)

// MoqSetGoroutineLabels_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_doReturnFn func(ctx context.Context)

// MoqSetGoroutineLabels_genType_results holds the results of the
// SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_results struct {
	Params  MoqSetGoroutineLabels_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetGoroutineLabels_genType_doFn
		DoReturnFn MoqSetGoroutineLabels_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetGoroutineLabels_genType_fnRecorder routes recorded function calls to
// the MoqSetGoroutineLabels_genType moq
type MoqSetGoroutineLabels_genType_fnRecorder struct {
	Params    MoqSetGoroutineLabels_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetGoroutineLabels_genType_results
	Moq       *MoqSetGoroutineLabels_genType
}

// MoqSetGoroutineLabels_genType_anyParams isolates the any params functions of
// the SetGoroutineLabels_genType type
type MoqSetGoroutineLabels_genType_anyParams struct {
	Recorder *MoqSetGoroutineLabels_genType_fnRecorder
}

// NewMoqSetGoroutineLabels_genType creates a new moq of the
// SetGoroutineLabels_genType type
func NewMoqSetGoroutineLabels_genType(scene *moq.Scene, config *moq.Config) *MoqSetGoroutineLabels_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetGoroutineLabels_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetGoroutineLabels_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ctx moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ctx moq.ParamIndexing
		}{
			Ctx: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SetGoroutineLabels_genType type
func (m *MoqSetGoroutineLabels_genType) Mock() SetGoroutineLabels_genType {
	return func(ctx context.Context) {
		m.Scene.T.Helper()
		moq := &MoqSetGoroutineLabels_genType_mock{Moq: m}
		moq.Fn(ctx)
	}
}

func (m *MoqSetGoroutineLabels_genType_mock) Fn(ctx context.Context) {
	m.Moq.Scene.T.Helper()
	params := MoqSetGoroutineLabels_genType_params{
		Ctx: ctx,
	}
	var results *MoqSetGoroutineLabels_genType_results
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
		result.DoFn(ctx)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(ctx)
	}
	return
}

func (m *MoqSetGoroutineLabels_genType) OnCall(ctx context.Context) *MoqSetGoroutineLabels_genType_fnRecorder {
	return &MoqSetGoroutineLabels_genType_fnRecorder{
		Params: MoqSetGoroutineLabels_genType_params{
			Ctx: ctx,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) Any() *MoqSetGoroutineLabels_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetGoroutineLabels_genType_anyParams{Recorder: r}
}

func (a *MoqSetGoroutineLabels_genType_anyParams) Ctx() *MoqSetGoroutineLabels_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) Seq() *MoqSetGoroutineLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) NoSeq() *MoqSetGoroutineLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) ReturnResults() *MoqSetGoroutineLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetGoroutineLabels_genType_doFn
		DoReturnFn MoqSetGoroutineLabels_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) AndDo(fn MoqSetGoroutineLabels_genType_doFn) *MoqSetGoroutineLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) DoReturnResults(fn MoqSetGoroutineLabels_genType_doReturnFn) *MoqSetGoroutineLabels_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetGoroutineLabels_genType_doFn
		DoReturnFn MoqSetGoroutineLabels_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetGoroutineLabels_genType_resultsByParams
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
		results = &MoqSetGoroutineLabels_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetGoroutineLabels_genType_paramsKey]*MoqSetGoroutineLabels_genType_results{},
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
		r.Results = &MoqSetGoroutineLabels_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetGoroutineLabels_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetGoroutineLabels_genType_fnRecorder {
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
				DoFn       MoqSetGoroutineLabels_genType_doFn
				DoReturnFn MoqSetGoroutineLabels_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetGoroutineLabels_genType) PrettyParams(params MoqSetGoroutineLabels_genType_params) string {
	return fmt.Sprintf("SetGoroutineLabels_genType(%#v)", params.Ctx)
}

func (m *MoqSetGoroutineLabels_genType) ParamsKey(params MoqSetGoroutineLabels_genType_params, anyParams uint64) MoqSetGoroutineLabels_genType_paramsKey {
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
	return MoqSetGoroutineLabels_genType_paramsKey{
		Params: struct{ Ctx context.Context }{
			Ctx: ctxUsed,
		},
		Hashes: struct{ Ctx hash.Hash }{
			Ctx: ctxUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetGoroutineLabels_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetGoroutineLabels_genType) AssertExpectationsMet() {
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
