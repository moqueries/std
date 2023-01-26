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

// The following type assertion assures that driver.ExecerContext is mocked
// completely
var _ driver.ExecerContext = (*MoqExecerContext_mock)(nil)

// MoqExecerContext holds the state of a moq of the ExecerContext type
type MoqExecerContext struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqExecerContext_mock

	ResultsByParams_ExecContext []MoqExecerContext_ExecContext_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ExecContext struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}
		}
	}
	// MoqExecerContext_mock isolates the mock interface of the ExecerContext type
}

type MoqExecerContext_mock struct {
	Moq *MoqExecerContext
}

// MoqExecerContext_recorder isolates the recorder interface of the
// ExecerContext type
type MoqExecerContext_recorder struct {
	Moq *MoqExecerContext
}

// MoqExecerContext_ExecContext_params holds the params of the ExecerContext
// type
type MoqExecerContext_ExecContext_params struct {
	Ctx   context.Context
	Query string
	Args  []driver.NamedValue
}

// MoqExecerContext_ExecContext_paramsKey holds the map key params of the
// ExecerContext type
type MoqExecerContext_ExecContext_paramsKey struct {
	Params struct {
		Ctx   context.Context
		Query string
	}
	Hashes struct {
		Ctx   hash.Hash
		Query hash.Hash
		Args  hash.Hash
	}
}

// MoqExecerContext_ExecContext_resultsByParams contains the results for a
// given set of parameters for the ExecerContext type
type MoqExecerContext_ExecContext_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqExecerContext_ExecContext_paramsKey]*MoqExecerContext_ExecContext_results
}

// MoqExecerContext_ExecContext_doFn defines the type of function needed when
// calling AndDo for the ExecerContext type
type MoqExecerContext_ExecContext_doFn func(ctx context.Context, query string, args []driver.NamedValue)

// MoqExecerContext_ExecContext_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ExecerContext type
type MoqExecerContext_ExecContext_doReturnFn func(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)

// MoqExecerContext_ExecContext_results holds the results of the ExecerContext
// type
type MoqExecerContext_ExecContext_results struct {
	Params  MoqExecerContext_ExecContext_params
	Results []struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqExecerContext_ExecContext_doFn
		DoReturnFn MoqExecerContext_ExecContext_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqExecerContext_ExecContext_fnRecorder routes recorded function calls to
// the MoqExecerContext moq
type MoqExecerContext_ExecContext_fnRecorder struct {
	Params    MoqExecerContext_ExecContext_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqExecerContext_ExecContext_results
	Moq       *MoqExecerContext
}

// MoqExecerContext_ExecContext_anyParams isolates the any params functions of
// the ExecerContext type
type MoqExecerContext_ExecContext_anyParams struct {
	Recorder *MoqExecerContext_ExecContext_fnRecorder
}

// NewMoqExecerContext creates a new moq of the ExecerContext type
func NewMoqExecerContext(scene *moq.Scene, config *moq.Config) *MoqExecerContext {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqExecerContext{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqExecerContext_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ExecContext struct {
					Ctx   moq.ParamIndexing
					Query moq.ParamIndexing
					Args  moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ExecContext struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}
		}{
			ExecContext: struct {
				Ctx   moq.ParamIndexing
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}{
				Ctx:   moq.ParamIndexByHash,
				Query: moq.ParamIndexByValue,
				Args:  moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ExecerContext type
func (m *MoqExecerContext) Mock() *MoqExecerContext_mock { return m.Moq }

func (m *MoqExecerContext_mock) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (result1 driver.Result, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqExecerContext_ExecContext_params{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	var results *MoqExecerContext_ExecContext_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ExecContext {
		paramsKey := m.Moq.ParamsKey_ExecContext(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ExecContext(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ExecContext(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ExecContext(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(ctx, query, args)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(ctx, query, args)
	}
	return
}

// OnCall returns the recorder implementation of the ExecerContext type
func (m *MoqExecerContext) OnCall() *MoqExecerContext_recorder {
	return &MoqExecerContext_recorder{
		Moq: m,
	}
}

func (m *MoqExecerContext_recorder) ExecContext(ctx context.Context, query string, args []driver.NamedValue) *MoqExecerContext_ExecContext_fnRecorder {
	return &MoqExecerContext_ExecContext_fnRecorder{
		Params: MoqExecerContext_ExecContext_params{
			Ctx:   ctx,
			Query: query,
			Args:  args,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqExecerContext_ExecContext_fnRecorder) Any() *MoqExecerContext_ExecContext_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExecContext(r.Params))
		return nil
	}
	return &MoqExecerContext_ExecContext_anyParams{Recorder: r}
}

func (a *MoqExecerContext_ExecContext_anyParams) Ctx() *MoqExecerContext_ExecContext_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqExecerContext_ExecContext_anyParams) Query() *MoqExecerContext_ExecContext_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqExecerContext_ExecContext_anyParams) Args() *MoqExecerContext_ExecContext_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqExecerContext_ExecContext_fnRecorder) Seq() *MoqExecerContext_ExecContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExecContext(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqExecerContext_ExecContext_fnRecorder) NoSeq() *MoqExecerContext_ExecContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExecContext(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqExecerContext_ExecContext_fnRecorder) ReturnResults(result1 driver.Result, result2 error) *MoqExecerContext_ExecContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqExecerContext_ExecContext_doFn
		DoReturnFn MoqExecerContext_ExecContext_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Result
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqExecerContext_ExecContext_fnRecorder) AndDo(fn MoqExecerContext_ExecContext_doFn) *MoqExecerContext_ExecContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqExecerContext_ExecContext_fnRecorder) DoReturnResults(fn MoqExecerContext_ExecContext_doReturnFn) *MoqExecerContext_ExecContext_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqExecerContext_ExecContext_doFn
		DoReturnFn MoqExecerContext_ExecContext_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqExecerContext_ExecContext_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqExecerContext_ExecContext_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ExecContext {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqExecerContext_ExecContext_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqExecerContext_ExecContext_paramsKey]*MoqExecerContext_ExecContext_results{},
		}
		r.Moq.ResultsByParams_ExecContext = append(r.Moq.ResultsByParams_ExecContext, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ExecContext) {
			copy(r.Moq.ResultsByParams_ExecContext[insertAt+1:], r.Moq.ResultsByParams_ExecContext[insertAt:0])
			r.Moq.ResultsByParams_ExecContext[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ExecContext(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqExecerContext_ExecContext_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqExecerContext_ExecContext_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqExecerContext_ExecContext_fnRecorder {
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
					Result1 driver.Result
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqExecerContext_ExecContext_doFn
				DoReturnFn MoqExecerContext_ExecContext_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqExecerContext) PrettyParams_ExecContext(params MoqExecerContext_ExecContext_params) string {
	return fmt.Sprintf("ExecContext(%#v, %#v, %#v)", params.Ctx, params.Query, params.Args)
}

func (m *MoqExecerContext) ParamsKey_ExecContext(params MoqExecerContext_ExecContext_params, anyParams uint64) MoqExecerContext_ExecContext_paramsKey {
	m.Scene.T.Helper()
	var ctxUsed context.Context
	var ctxUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ExecContext.Ctx == moq.ParamIndexByValue {
			ctxUsed = params.Ctx
		} else {
			ctxUsedHash = hash.DeepHash(params.Ctx)
		}
	}
	var queryUsed string
	var queryUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.ExecContext.Query == moq.ParamIndexByValue {
			queryUsed = params.Query
		} else {
			queryUsedHash = hash.DeepHash(params.Query)
		}
	}
	var argsUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.ExecContext.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter of the ExecContext function can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqExecerContext_ExecContext_paramsKey{
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
			Args  hash.Hash
		}{
			Ctx:   ctxUsedHash,
			Query: queryUsedHash,
			Args:  argsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqExecerContext) Reset() { m.ResultsByParams_ExecContext = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqExecerContext) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ExecContext {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ExecContext(results.Params))
			}
		}
	}
}