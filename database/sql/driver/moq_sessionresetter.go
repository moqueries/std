// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package driver

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that driver.SessionResetter is mocked
// completely
var _ driver.SessionResetter = (*MoqSessionResetter_mock)(nil)

// MoqSessionResetter holds the state of a moq of the SessionResetter type
type MoqSessionResetter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSessionResetter_mock

	ResultsByParams_ResetSession []MoqSessionResetter_ResetSession_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ResetSession struct {
				Ctx moq.ParamIndexing
			}
		}
	}
	// MoqSessionResetter_mock isolates the mock interface of the SessionResetter
}

// type
type MoqSessionResetter_mock struct {
	Moq *MoqSessionResetter
}

// MoqSessionResetter_recorder isolates the recorder interface of the
// SessionResetter type
type MoqSessionResetter_recorder struct {
	Moq *MoqSessionResetter
}

// MoqSessionResetter_ResetSession_params holds the params of the
// SessionResetter type
type MoqSessionResetter_ResetSession_params struct{ Ctx context.Context }

// MoqSessionResetter_ResetSession_paramsKey holds the map key params of the
// SessionResetter type
type MoqSessionResetter_ResetSession_paramsKey struct {
	Params struct{ Ctx context.Context }
	Hashes struct{ Ctx hash.Hash }
}

// MoqSessionResetter_ResetSession_resultsByParams contains the results for a
// given set of parameters for the SessionResetter type
type MoqSessionResetter_ResetSession_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSessionResetter_ResetSession_paramsKey]*MoqSessionResetter_ResetSession_results
}

// MoqSessionResetter_ResetSession_doFn defines the type of function needed
// when calling AndDo for the SessionResetter type
type MoqSessionResetter_ResetSession_doFn func(ctx context.Context)

// MoqSessionResetter_ResetSession_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SessionResetter type
type MoqSessionResetter_ResetSession_doReturnFn func(ctx context.Context) error

// MoqSessionResetter_ResetSession_results holds the results of the
// SessionResetter type
type MoqSessionResetter_ResetSession_results struct {
	Params  MoqSessionResetter_ResetSession_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSessionResetter_ResetSession_doFn
		DoReturnFn MoqSessionResetter_ResetSession_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSessionResetter_ResetSession_fnRecorder routes recorded function calls to
// the MoqSessionResetter moq
type MoqSessionResetter_ResetSession_fnRecorder struct {
	Params    MoqSessionResetter_ResetSession_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSessionResetter_ResetSession_results
	Moq       *MoqSessionResetter
}

// MoqSessionResetter_ResetSession_anyParams isolates the any params functions
// of the SessionResetter type
type MoqSessionResetter_ResetSession_anyParams struct {
	Recorder *MoqSessionResetter_ResetSession_fnRecorder
}

// NewMoqSessionResetter creates a new moq of the SessionResetter type
func NewMoqSessionResetter(scene *moq.Scene, config *moq.Config) *MoqSessionResetter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSessionResetter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSessionResetter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ResetSession struct {
					Ctx moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ResetSession struct {
				Ctx moq.ParamIndexing
			}
		}{
			ResetSession: struct {
				Ctx moq.ParamIndexing
			}{
				Ctx: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the SessionResetter type
func (m *MoqSessionResetter) Mock() *MoqSessionResetter_mock { return m.Moq }

func (m *MoqSessionResetter_mock) ResetSession(ctx context.Context) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSessionResetter_ResetSession_params{
		Ctx: ctx,
	}
	var results *MoqSessionResetter_ResetSession_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ResetSession {
		paramsKey := m.Moq.ParamsKey_ResetSession(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ResetSession(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ResetSession(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ResetSession(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(ctx)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ctx)
	}
	return
}

// OnCall returns the recorder implementation of the SessionResetter type
func (m *MoqSessionResetter) OnCall() *MoqSessionResetter_recorder {
	return &MoqSessionResetter_recorder{
		Moq: m,
	}
}

func (m *MoqSessionResetter_recorder) ResetSession(ctx context.Context) *MoqSessionResetter_ResetSession_fnRecorder {
	return &MoqSessionResetter_ResetSession_fnRecorder{
		Params: MoqSessionResetter_ResetSession_params{
			Ctx: ctx,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) Any() *MoqSessionResetter_ResetSession_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ResetSession(r.Params))
		return nil
	}
	return &MoqSessionResetter_ResetSession_anyParams{Recorder: r}
}

func (a *MoqSessionResetter_ResetSession_anyParams) Ctx() *MoqSessionResetter_ResetSession_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) Seq() *MoqSessionResetter_ResetSession_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ResetSession(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) NoSeq() *MoqSessionResetter_ResetSession_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ResetSession(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) ReturnResults(result1 error) *MoqSessionResetter_ResetSession_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSessionResetter_ResetSession_doFn
		DoReturnFn MoqSessionResetter_ResetSession_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) AndDo(fn MoqSessionResetter_ResetSession_doFn) *MoqSessionResetter_ResetSession_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) DoReturnResults(fn MoqSessionResetter_ResetSession_doReturnFn) *MoqSessionResetter_ResetSession_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSessionResetter_ResetSession_doFn
		DoReturnFn MoqSessionResetter_ResetSession_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSessionResetter_ResetSession_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ResetSession {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSessionResetter_ResetSession_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSessionResetter_ResetSession_paramsKey]*MoqSessionResetter_ResetSession_results{},
		}
		r.Moq.ResultsByParams_ResetSession = append(r.Moq.ResultsByParams_ResetSession, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ResetSession) {
			copy(r.Moq.ResultsByParams_ResetSession[insertAt+1:], r.Moq.ResultsByParams_ResetSession[insertAt:0])
			r.Moq.ResultsByParams_ResetSession[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ResetSession(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSessionResetter_ResetSession_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSessionResetter_ResetSession_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSessionResetter_ResetSession_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqSessionResetter_ResetSession_doFn
				DoReturnFn MoqSessionResetter_ResetSession_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSessionResetter) PrettyParams_ResetSession(params MoqSessionResetter_ResetSession_params) string {
	return fmt.Sprintf("ResetSession(%#v)", params.Ctx)
}

func (m *MoqSessionResetter) ParamsKey_ResetSession(params MoqSessionResetter_ResetSession_params, anyParams uint64) MoqSessionResetter_ResetSession_paramsKey {
	m.Scene.T.Helper()
	var ctxUsed context.Context
	var ctxUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ResetSession.Ctx == moq.ParamIndexByValue {
			ctxUsed = params.Ctx
		} else {
			ctxUsedHash = hash.DeepHash(params.Ctx)
		}
	}
	return MoqSessionResetter_ResetSession_paramsKey{
		Params: struct{ Ctx context.Context }{
			Ctx: ctxUsed,
		},
		Hashes: struct{ Ctx hash.Hash }{
			Ctx: ctxUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSessionResetter) Reset() { m.ResultsByParams_ResetSession = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSessionResetter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ResetSession {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ResetSession(results.Params))
			}
		}
	}
}
