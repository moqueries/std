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

// The following type assertion assures that driver.ConnPrepareContext is
// mocked completely
var _ driver.ConnPrepareContext = (*MoqConnPrepareContext_mock)(nil)

// MoqConnPrepareContext holds the state of a moq of the ConnPrepareContext
// type
type MoqConnPrepareContext struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqConnPrepareContext_mock

	ResultsByParams_PrepareContext []MoqConnPrepareContext_PrepareContext_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			PrepareContext struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
			}
		}
	}
	// MoqConnPrepareContext_mock isolates the mock interface of the
}

// ConnPrepareContext type
type MoqConnPrepareContext_mock struct {
	Moq *MoqConnPrepareContext
}

// MoqConnPrepareContext_recorder isolates the recorder interface of the
// ConnPrepareContext type
type MoqConnPrepareContext_recorder struct {
	Moq *MoqConnPrepareContext
}

// MoqConnPrepareContext_PrepareContext_params holds the params of the
// ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_params struct {
	Ctx   context.Context
	Query string
}

// MoqConnPrepareContext_PrepareContext_paramsKey holds the map key params of
// the ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_paramsKey struct {
	Params struct {
		Ctx   context.Context
		Query string
	}
	Hashes struct {
		Ctx   hash.Hash
		Query hash.Hash
	}
}

// MoqConnPrepareContext_PrepareContext_resultsByParams contains the results
// for a given set of parameters for the ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqConnPrepareContext_PrepareContext_paramsKey]*MoqConnPrepareContext_PrepareContext_results
}

// MoqConnPrepareContext_PrepareContext_doFn defines the type of function
// needed when calling AndDo for the ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_doFn func(ctx context.Context, query string)

// MoqConnPrepareContext_PrepareContext_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_doReturnFn func(ctx context.Context, query string) (driver.Stmt, error)

// MoqConnPrepareContext_PrepareContext_results holds the results of the
// ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_results struct {
	Params  MoqConnPrepareContext_PrepareContext_params
	Results []struct {
		Values *struct {
			Result1 driver.Stmt
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqConnPrepareContext_PrepareContext_doFn
		DoReturnFn MoqConnPrepareContext_PrepareContext_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqConnPrepareContext_PrepareContext_fnRecorder routes recorded function
// calls to the MoqConnPrepareContext moq
type MoqConnPrepareContext_PrepareContext_fnRecorder struct {
	Params    MoqConnPrepareContext_PrepareContext_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqConnPrepareContext_PrepareContext_results
	Moq       *MoqConnPrepareContext
}

// MoqConnPrepareContext_PrepareContext_anyParams isolates the any params
// functions of the ConnPrepareContext type
type MoqConnPrepareContext_PrepareContext_anyParams struct {
	Recorder *MoqConnPrepareContext_PrepareContext_fnRecorder
}

// NewMoqConnPrepareContext creates a new moq of the ConnPrepareContext type
func NewMoqConnPrepareContext(scene *moq.Scene, config *moq.Config) *MoqConnPrepareContext {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqConnPrepareContext{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqConnPrepareContext_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				PrepareContext struct {
					Ctx   moq.ParamIndexing
					Query moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			PrepareContext struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
			}
		}{
			PrepareContext: struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
			}{
				Ctx:   moq.ParamIndexByHash,
				Query: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ConnPrepareContext type
func (m *MoqConnPrepareContext) Mock() *MoqConnPrepareContext_mock { return m.Moq }

func (m *MoqConnPrepareContext_mock) PrepareContext(ctx context.Context, query string) (result1 driver.Stmt, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqConnPrepareContext_PrepareContext_params{
		Ctx:   ctx,
		Query: query,
	}
	var results *MoqConnPrepareContext_PrepareContext_results
	for _, resultsByParams := range m.Moq.ResultsByParams_PrepareContext {
		paramsKey := m.Moq.ParamsKey_PrepareContext(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_PrepareContext(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_PrepareContext(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_PrepareContext(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(ctx, query)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(ctx, query)
	}
	return
}

// OnCall returns the recorder implementation of the ConnPrepareContext type
func (m *MoqConnPrepareContext) OnCall() *MoqConnPrepareContext_recorder {
	return &MoqConnPrepareContext_recorder{
		Moq: m,
	}
}

func (m *MoqConnPrepareContext_recorder) PrepareContext(ctx context.Context, query string) *MoqConnPrepareContext_PrepareContext_fnRecorder {
	return &MoqConnPrepareContext_PrepareContext_fnRecorder{
		Params: MoqConnPrepareContext_PrepareContext_params{
			Ctx:   ctx,
			Query: query,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) Any() *MoqConnPrepareContext_PrepareContext_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrepareContext(r.Params))
		return nil
	}
	return &MoqConnPrepareContext_PrepareContext_anyParams{Recorder: r}
}

func (a *MoqConnPrepareContext_PrepareContext_anyParams) Ctx() *MoqConnPrepareContext_PrepareContext_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqConnPrepareContext_PrepareContext_anyParams) Query() *MoqConnPrepareContext_PrepareContext_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) Seq() *MoqConnPrepareContext_PrepareContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrepareContext(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) NoSeq() *MoqConnPrepareContext_PrepareContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrepareContext(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) ReturnResults(result1 driver.Stmt, result2 error) *MoqConnPrepareContext_PrepareContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Stmt
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqConnPrepareContext_PrepareContext_doFn
		DoReturnFn MoqConnPrepareContext_PrepareContext_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Stmt
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) AndDo(fn MoqConnPrepareContext_PrepareContext_doFn) *MoqConnPrepareContext_PrepareContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) DoReturnResults(fn MoqConnPrepareContext_PrepareContext_doReturnFn) *MoqConnPrepareContext_PrepareContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Stmt
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqConnPrepareContext_PrepareContext_doFn
		DoReturnFn MoqConnPrepareContext_PrepareContext_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqConnPrepareContext_PrepareContext_resultsByParams
	for n, res := range r.Moq.ResultsByParams_PrepareContext {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqConnPrepareContext_PrepareContext_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqConnPrepareContext_PrepareContext_paramsKey]*MoqConnPrepareContext_PrepareContext_results{},
		}
		r.Moq.ResultsByParams_PrepareContext = append(r.Moq.ResultsByParams_PrepareContext, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_PrepareContext) {
			copy(r.Moq.ResultsByParams_PrepareContext[insertAt+1:], r.Moq.ResultsByParams_PrepareContext[insertAt:0])
			r.Moq.ResultsByParams_PrepareContext[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_PrepareContext(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqConnPrepareContext_PrepareContext_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqConnPrepareContext_PrepareContext_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqConnPrepareContext_PrepareContext_fnRecorder {
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
					Result1 driver.Stmt
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqConnPrepareContext_PrepareContext_doFn
				DoReturnFn MoqConnPrepareContext_PrepareContext_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqConnPrepareContext) PrettyParams_PrepareContext(params MoqConnPrepareContext_PrepareContext_params) string {
	return fmt.Sprintf("PrepareContext(%#v, %#v)", params.Ctx, params.Query)
}

func (m *MoqConnPrepareContext) ParamsKey_PrepareContext(params MoqConnPrepareContext_PrepareContext_params, anyParams uint64) MoqConnPrepareContext_PrepareContext_paramsKey {
	m.Scene.T.Helper()
	var ctxUsed context.Context
	var ctxUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.PrepareContext.Ctx == moq.ParamIndexByValue {
			ctxUsed = params.Ctx
		} else {
			ctxUsedHash = hash.DeepHash(params.Ctx)
		}
	}
	var queryUsed string
	var queryUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.PrepareContext.Query == moq.ParamIndexByValue {
			queryUsed = params.Query
		} else {
			queryUsedHash = hash.DeepHash(params.Query)
		}
	}
	return MoqConnPrepareContext_PrepareContext_paramsKey{
		Params: struct {
			Ctx   context.Context
			Query string
		}{
			Ctx:   ctxUsed,
			Query: queryUsed,
		},
		Hashes: struct {
			Ctx   hash.Hash
			Query hash.Hash
		}{
			Ctx:   ctxUsedHash,
			Query: queryUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqConnPrepareContext) Reset() { m.ResultsByParams_PrepareContext = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqConnPrepareContext) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_PrepareContext {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_PrepareContext(results.Params))
			}
		}
	}
}
