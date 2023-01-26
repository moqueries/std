// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package trace

import (
	"context"
	"fmt"
	"math/bits"
	"runtime/trace"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewTask_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type NewTask_genType func(pctx context.Context, taskType string) (ctx context.Context, task *trace.Task)

// MoqNewTask_genType holds the state of a moq of the NewTask_genType type
type MoqNewTask_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewTask_genType_mock

	ResultsByParams []MoqNewTask_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pctx     moq.ParamIndexing
			TaskType moq.ParamIndexing
		}
	}
}

// MoqNewTask_genType_mock isolates the mock interface of the NewTask_genType
// type
type MoqNewTask_genType_mock struct {
	Moq *MoqNewTask_genType
}

// MoqNewTask_genType_params holds the params of the NewTask_genType type
type MoqNewTask_genType_params struct {
	Pctx     context.Context
	TaskType string
}

// MoqNewTask_genType_paramsKey holds the map key params of the NewTask_genType
// type
type MoqNewTask_genType_paramsKey struct {
	Params struct {
		Pctx     context.Context
		TaskType string
	}
	Hashes struct {
		Pctx     hash.Hash
		TaskType hash.Hash
	}
}

// MoqNewTask_genType_resultsByParams contains the results for a given set of
// parameters for the NewTask_genType type
type MoqNewTask_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewTask_genType_paramsKey]*MoqNewTask_genType_results
}

// MoqNewTask_genType_doFn defines the type of function needed when calling
// AndDo for the NewTask_genType type
type MoqNewTask_genType_doFn func(pctx context.Context, taskType string)

// MoqNewTask_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewTask_genType type
type MoqNewTask_genType_doReturnFn func(pctx context.Context, taskType string) (ctx context.Context, task *trace.Task)

// MoqNewTask_genType_results holds the results of the NewTask_genType type
type MoqNewTask_genType_results struct {
	Params  MoqNewTask_genType_params
	Results []struct {
		Values *struct {
			Ctx  context.Context
			Task *trace.Task
		}
		Sequence   uint32
		DoFn       MoqNewTask_genType_doFn
		DoReturnFn MoqNewTask_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewTask_genType_fnRecorder routes recorded function calls to the
// MoqNewTask_genType moq
type MoqNewTask_genType_fnRecorder struct {
	Params    MoqNewTask_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewTask_genType_results
	Moq       *MoqNewTask_genType
}

// MoqNewTask_genType_anyParams isolates the any params functions of the
// NewTask_genType type
type MoqNewTask_genType_anyParams struct {
	Recorder *MoqNewTask_genType_fnRecorder
}

// NewMoqNewTask_genType creates a new moq of the NewTask_genType type
func NewMoqNewTask_genType(scene *moq.Scene, config *moq.Config) *MoqNewTask_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewTask_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewTask_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pctx     moq.ParamIndexing
				TaskType moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pctx     moq.ParamIndexing
			TaskType moq.ParamIndexing
		}{
			Pctx:     moq.ParamIndexByHash,
			TaskType: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewTask_genType type
func (m *MoqNewTask_genType) Mock() NewTask_genType {
	return func(pctx context.Context, taskType string) (_ context.Context, _ *trace.Task) {
		m.Scene.T.Helper()
		moq := &MoqNewTask_genType_mock{Moq: m}
		return moq.Fn(pctx, taskType)
	}
}

func (m *MoqNewTask_genType_mock) Fn(pctx context.Context, taskType string) (ctx context.Context, task *trace.Task) {
	m.Moq.Scene.T.Helper()
	params := MoqNewTask_genType_params{
		Pctx:     pctx,
		TaskType: taskType,
	}
	var results *MoqNewTask_genType_results
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
		result.DoFn(pctx, taskType)
	}

	if result.Values != nil {
		ctx = result.Values.Ctx
		task = result.Values.Task
	}
	if result.DoReturnFn != nil {
		ctx, task = result.DoReturnFn(pctx, taskType)
	}
	return
}

func (m *MoqNewTask_genType) OnCall(pctx context.Context, taskType string) *MoqNewTask_genType_fnRecorder {
	return &MoqNewTask_genType_fnRecorder{
		Params: MoqNewTask_genType_params{
			Pctx:     pctx,
			TaskType: taskType,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewTask_genType_fnRecorder) Any() *MoqNewTask_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewTask_genType_anyParams{Recorder: r}
}

func (a *MoqNewTask_genType_anyParams) Pctx() *MoqNewTask_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewTask_genType_anyParams) TaskType() *MoqNewTask_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewTask_genType_fnRecorder) Seq() *MoqNewTask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewTask_genType_fnRecorder) NoSeq() *MoqNewTask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewTask_genType_fnRecorder) ReturnResults(ctx context.Context, task *trace.Task) *MoqNewTask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ctx  context.Context
			Task *trace.Task
		}
		Sequence   uint32
		DoFn       MoqNewTask_genType_doFn
		DoReturnFn MoqNewTask_genType_doReturnFn
	}{
		Values: &struct {
			Ctx  context.Context
			Task *trace.Task
		}{
			Ctx:  ctx,
			Task: task,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewTask_genType_fnRecorder) AndDo(fn MoqNewTask_genType_doFn) *MoqNewTask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewTask_genType_fnRecorder) DoReturnResults(fn MoqNewTask_genType_doReturnFn) *MoqNewTask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ctx  context.Context
			Task *trace.Task
		}
		Sequence   uint32
		DoFn       MoqNewTask_genType_doFn
		DoReturnFn MoqNewTask_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewTask_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewTask_genType_resultsByParams
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
		results = &MoqNewTask_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewTask_genType_paramsKey]*MoqNewTask_genType_results{},
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
		r.Results = &MoqNewTask_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewTask_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewTask_genType_fnRecorder {
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
					Ctx  context.Context
					Task *trace.Task
				}
				Sequence   uint32
				DoFn       MoqNewTask_genType_doFn
				DoReturnFn MoqNewTask_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewTask_genType) PrettyParams(params MoqNewTask_genType_params) string {
	return fmt.Sprintf("NewTask_genType(%#v, %#v)", params.Pctx, params.TaskType)
}

func (m *MoqNewTask_genType) ParamsKey(params MoqNewTask_genType_params, anyParams uint64) MoqNewTask_genType_paramsKey {
	m.Scene.T.Helper()
	var pctxUsed context.Context
	var pctxUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pctx == moq.ParamIndexByValue {
			pctxUsed = params.Pctx
		} else {
			pctxUsedHash = hash.DeepHash(params.Pctx)
		}
	}
	var taskTypeUsed string
	var taskTypeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.TaskType == moq.ParamIndexByValue {
			taskTypeUsed = params.TaskType
		} else {
			taskTypeUsedHash = hash.DeepHash(params.TaskType)
		}
	}
	return MoqNewTask_genType_paramsKey{
		Params: struct {
			Pctx     context.Context
			TaskType string
		}{
			Pctx:     pctxUsed,
			TaskType: taskTypeUsed,
		},
		Hashes: struct {
			Pctx     hash.Hash
			TaskType hash.Hash
		}{
			Pctx:     pctxUsedHash,
			TaskType: taskTypeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewTask_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewTask_genType) AssertExpectationsMet() {
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
