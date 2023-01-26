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

// StartRegion_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type StartRegion_genType func(ctx context.Context, regionType string) *trace.Region

// MoqStartRegion_genType holds the state of a moq of the StartRegion_genType
// type
type MoqStartRegion_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStartRegion_genType_mock

	ResultsByParams []MoqStartRegion_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ctx        moq.ParamIndexing
			RegionType moq.ParamIndexing
		}
	}
}

// MoqStartRegion_genType_mock isolates the mock interface of the
// StartRegion_genType type
type MoqStartRegion_genType_mock struct {
	Moq *MoqStartRegion_genType
}

// MoqStartRegion_genType_params holds the params of the StartRegion_genType
// type
type MoqStartRegion_genType_params struct {
	Ctx        context.Context
	RegionType string
}

// MoqStartRegion_genType_paramsKey holds the map key params of the
// StartRegion_genType type
type MoqStartRegion_genType_paramsKey struct {
	Params struct {
		Ctx        context.Context
		RegionType string
	}
	Hashes struct {
		Ctx        hash.Hash
		RegionType hash.Hash
	}
}

// MoqStartRegion_genType_resultsByParams contains the results for a given set
// of parameters for the StartRegion_genType type
type MoqStartRegion_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStartRegion_genType_paramsKey]*MoqStartRegion_genType_results
}

// MoqStartRegion_genType_doFn defines the type of function needed when calling
// AndDo for the StartRegion_genType type
type MoqStartRegion_genType_doFn func(ctx context.Context, regionType string)

// MoqStartRegion_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the StartRegion_genType type
type MoqStartRegion_genType_doReturnFn func(ctx context.Context, regionType string) *trace.Region

// MoqStartRegion_genType_results holds the results of the StartRegion_genType
// type
type MoqStartRegion_genType_results struct {
	Params  MoqStartRegion_genType_params
	Results []struct {
		Values *struct {
			Result1 *trace.Region
		}
		Sequence   uint32
		DoFn       MoqStartRegion_genType_doFn
		DoReturnFn MoqStartRegion_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStartRegion_genType_fnRecorder routes recorded function calls to the
// MoqStartRegion_genType moq
type MoqStartRegion_genType_fnRecorder struct {
	Params    MoqStartRegion_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStartRegion_genType_results
	Moq       *MoqStartRegion_genType
}

// MoqStartRegion_genType_anyParams isolates the any params functions of the
// StartRegion_genType type
type MoqStartRegion_genType_anyParams struct {
	Recorder *MoqStartRegion_genType_fnRecorder
}

// NewMoqStartRegion_genType creates a new moq of the StartRegion_genType type
func NewMoqStartRegion_genType(scene *moq.Scene, config *moq.Config) *MoqStartRegion_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStartRegion_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStartRegion_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ctx        moq.ParamIndexing
				RegionType moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ctx        moq.ParamIndexing
			RegionType moq.ParamIndexing
		}{
			Ctx:        moq.ParamIndexByHash,
			RegionType: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the StartRegion_genType type
func (m *MoqStartRegion_genType) Mock() StartRegion_genType {
	return func(ctx context.Context, regionType string) *trace.Region {
		m.Scene.T.Helper()
		moq := &MoqStartRegion_genType_mock{Moq: m}
		return moq.Fn(ctx, regionType)
	}
}

func (m *MoqStartRegion_genType_mock) Fn(ctx context.Context, regionType string) (result1 *trace.Region) {
	m.Moq.Scene.T.Helper()
	params := MoqStartRegion_genType_params{
		Ctx:        ctx,
		RegionType: regionType,
	}
	var results *MoqStartRegion_genType_results
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
		result.DoFn(ctx, regionType)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ctx, regionType)
	}
	return
}

func (m *MoqStartRegion_genType) OnCall(ctx context.Context, regionType string) *MoqStartRegion_genType_fnRecorder {
	return &MoqStartRegion_genType_fnRecorder{
		Params: MoqStartRegion_genType_params{
			Ctx:        ctx,
			RegionType: regionType,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStartRegion_genType_fnRecorder) Any() *MoqStartRegion_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStartRegion_genType_anyParams{Recorder: r}
}

func (a *MoqStartRegion_genType_anyParams) Ctx() *MoqStartRegion_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqStartRegion_genType_anyParams) RegionType() *MoqStartRegion_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqStartRegion_genType_fnRecorder) Seq() *MoqStartRegion_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStartRegion_genType_fnRecorder) NoSeq() *MoqStartRegion_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStartRegion_genType_fnRecorder) ReturnResults(result1 *trace.Region) *MoqStartRegion_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *trace.Region
		}
		Sequence   uint32
		DoFn       MoqStartRegion_genType_doFn
		DoReturnFn MoqStartRegion_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *trace.Region
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStartRegion_genType_fnRecorder) AndDo(fn MoqStartRegion_genType_doFn) *MoqStartRegion_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStartRegion_genType_fnRecorder) DoReturnResults(fn MoqStartRegion_genType_doReturnFn) *MoqStartRegion_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *trace.Region
		}
		Sequence   uint32
		DoFn       MoqStartRegion_genType_doFn
		DoReturnFn MoqStartRegion_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStartRegion_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStartRegion_genType_resultsByParams
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
		results = &MoqStartRegion_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStartRegion_genType_paramsKey]*MoqStartRegion_genType_results{},
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
		r.Results = &MoqStartRegion_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStartRegion_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStartRegion_genType_fnRecorder {
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
					Result1 *trace.Region
				}
				Sequence   uint32
				DoFn       MoqStartRegion_genType_doFn
				DoReturnFn MoqStartRegion_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStartRegion_genType) PrettyParams(params MoqStartRegion_genType_params) string {
	return fmt.Sprintf("StartRegion_genType(%#v, %#v)", params.Ctx, params.RegionType)
}

func (m *MoqStartRegion_genType) ParamsKey(params MoqStartRegion_genType_params, anyParams uint64) MoqStartRegion_genType_paramsKey {
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
	var regionTypeUsed string
	var regionTypeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.RegionType == moq.ParamIndexByValue {
			regionTypeUsed = params.RegionType
		} else {
			regionTypeUsedHash = hash.DeepHash(params.RegionType)
		}
	}
	return MoqStartRegion_genType_paramsKey{
		Params: struct {
			Ctx        context.Context
			RegionType string
		}{
			Ctx:        ctxUsed,
			RegionType: regionTypeUsed,
		},
		Hashes: struct {
			Ctx        hash.Hash
			RegionType hash.Hash
		}{
			Ctx:        ctxUsedHash,
			RegionType: regionTypeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStartRegion_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStartRegion_genType) AssertExpectationsMet() {
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
