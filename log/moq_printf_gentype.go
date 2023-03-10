// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package log

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Printf_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Printf_genType func(format string, v ...any)

// MoqPrintf_genType holds the state of a moq of the Printf_genType type
type MoqPrintf_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPrintf_genType_mock

	ResultsByParams []MoqPrintf_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Format moq.ParamIndexing
			V      moq.ParamIndexing
		}
	}
}

// MoqPrintf_genType_mock isolates the mock interface of the Printf_genType
// type
type MoqPrintf_genType_mock struct {
	Moq *MoqPrintf_genType
}

// MoqPrintf_genType_params holds the params of the Printf_genType type
type MoqPrintf_genType_params struct {
	Format string
	V      []any
}

// MoqPrintf_genType_paramsKey holds the map key params of the Printf_genType
// type
type MoqPrintf_genType_paramsKey struct {
	Params struct{ Format string }
	Hashes struct {
		Format hash.Hash
		V      hash.Hash
	}
}

// MoqPrintf_genType_resultsByParams contains the results for a given set of
// parameters for the Printf_genType type
type MoqPrintf_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPrintf_genType_paramsKey]*MoqPrintf_genType_results
}

// MoqPrintf_genType_doFn defines the type of function needed when calling
// AndDo for the Printf_genType type
type MoqPrintf_genType_doFn func(format string, v ...any)

// MoqPrintf_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Printf_genType type
type MoqPrintf_genType_doReturnFn func(format string, v ...any)

// MoqPrintf_genType_results holds the results of the Printf_genType type
type MoqPrintf_genType_results struct {
	Params  MoqPrintf_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintf_genType_doFn
		DoReturnFn MoqPrintf_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPrintf_genType_fnRecorder routes recorded function calls to the
// MoqPrintf_genType moq
type MoqPrintf_genType_fnRecorder struct {
	Params    MoqPrintf_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPrintf_genType_results
	Moq       *MoqPrintf_genType
}

// MoqPrintf_genType_anyParams isolates the any params functions of the
// Printf_genType type
type MoqPrintf_genType_anyParams struct {
	Recorder *MoqPrintf_genType_fnRecorder
}

// NewMoqPrintf_genType creates a new moq of the Printf_genType type
func NewMoqPrintf_genType(scene *moq.Scene, config *moq.Config) *MoqPrintf_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPrintf_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPrintf_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Format moq.ParamIndexing
				V      moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Format moq.ParamIndexing
			V      moq.ParamIndexing
		}{
			Format: moq.ParamIndexByValue,
			V:      moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Printf_genType type
func (m *MoqPrintf_genType) Mock() Printf_genType {
	return func(format string, v ...any) {
		m.Scene.T.Helper()
		moq := &MoqPrintf_genType_mock{Moq: m}
		moq.Fn(format, v...)
	}
}

func (m *MoqPrintf_genType_mock) Fn(format string, v ...any) {
	m.Moq.Scene.T.Helper()
	params := MoqPrintf_genType_params{
		Format: format,
		V:      v,
	}
	var results *MoqPrintf_genType_results
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
		result.DoFn(format, v...)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(format, v...)
	}
	return
}

func (m *MoqPrintf_genType) OnCall(format string, v ...any) *MoqPrintf_genType_fnRecorder {
	return &MoqPrintf_genType_fnRecorder{
		Params: MoqPrintf_genType_params{
			Format: format,
			V:      v,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPrintf_genType_fnRecorder) Any() *MoqPrintf_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPrintf_genType_anyParams{Recorder: r}
}

func (a *MoqPrintf_genType_anyParams) Format() *MoqPrintf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqPrintf_genType_anyParams) V() *MoqPrintf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqPrintf_genType_fnRecorder) Seq() *MoqPrintf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPrintf_genType_fnRecorder) NoSeq() *MoqPrintf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPrintf_genType_fnRecorder) ReturnResults() *MoqPrintf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintf_genType_doFn
		DoReturnFn MoqPrintf_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPrintf_genType_fnRecorder) AndDo(fn MoqPrintf_genType_doFn) *MoqPrintf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPrintf_genType_fnRecorder) DoReturnResults(fn MoqPrintf_genType_doReturnFn) *MoqPrintf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintf_genType_doFn
		DoReturnFn MoqPrintf_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPrintf_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPrintf_genType_resultsByParams
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
		results = &MoqPrintf_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPrintf_genType_paramsKey]*MoqPrintf_genType_results{},
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
		r.Results = &MoqPrintf_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPrintf_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPrintf_genType_fnRecorder {
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
				DoFn       MoqPrintf_genType_doFn
				DoReturnFn MoqPrintf_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPrintf_genType) PrettyParams(params MoqPrintf_genType_params) string {
	return fmt.Sprintf("Printf_genType(%#v, %#v)", params.Format, params.V)
}

func (m *MoqPrintf_genType) ParamsKey(params MoqPrintf_genType_params, anyParams uint64) MoqPrintf_genType_paramsKey {
	m.Scene.T.Helper()
	var formatUsed string
	var formatUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Format == moq.ParamIndexByValue {
			formatUsed = params.Format
		} else {
			formatUsedHash = hash.DeepHash(params.Format)
		}
	}
	var vUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The v parameter can't be indexed by value")
		}
		vUsedHash = hash.DeepHash(params.V)
	}
	return MoqPrintf_genType_paramsKey{
		Params: struct{ Format string }{
			Format: formatUsed,
		},
		Hashes: struct {
			Format hash.Hash
			V      hash.Hash
		}{
			Format: formatUsedHash,
			V:      vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPrintf_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPrintf_genType) AssertExpectationsMet() {
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
