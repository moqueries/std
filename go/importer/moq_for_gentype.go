// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package importer

import (
	"fmt"
	"go/importer"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// For_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type For_genType func(compiler string, lookup importer.Lookup) types.Importer

// MoqFor_genType holds the state of a moq of the For_genType type
type MoqFor_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFor_genType_mock

	ResultsByParams []MoqFor_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Compiler moq.ParamIndexing
			Lookup   moq.ParamIndexing
		}
	}
}

// MoqFor_genType_mock isolates the mock interface of the For_genType type
type MoqFor_genType_mock struct {
	Moq *MoqFor_genType
}

// MoqFor_genType_params holds the params of the For_genType type
type MoqFor_genType_params struct {
	Compiler string
	Lookup   importer.Lookup
}

// MoqFor_genType_paramsKey holds the map key params of the For_genType type
type MoqFor_genType_paramsKey struct {
	Params struct{ Compiler string }
	Hashes struct {
		Compiler hash.Hash
		Lookup   hash.Hash
	}
}

// MoqFor_genType_resultsByParams contains the results for a given set of
// parameters for the For_genType type
type MoqFor_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFor_genType_paramsKey]*MoqFor_genType_results
}

// MoqFor_genType_doFn defines the type of function needed when calling AndDo
// for the For_genType type
type MoqFor_genType_doFn func(compiler string, lookup importer.Lookup)

// MoqFor_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the For_genType type
type MoqFor_genType_doReturnFn func(compiler string, lookup importer.Lookup) types.Importer

// MoqFor_genType_results holds the results of the For_genType type
type MoqFor_genType_results struct {
	Params  MoqFor_genType_params
	Results []struct {
		Values *struct {
			Result1 types.Importer
		}
		Sequence   uint32
		DoFn       MoqFor_genType_doFn
		DoReturnFn MoqFor_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFor_genType_fnRecorder routes recorded function calls to the
// MoqFor_genType moq
type MoqFor_genType_fnRecorder struct {
	Params    MoqFor_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFor_genType_results
	Moq       *MoqFor_genType
}

// MoqFor_genType_anyParams isolates the any params functions of the
// For_genType type
type MoqFor_genType_anyParams struct {
	Recorder *MoqFor_genType_fnRecorder
}

// NewMoqFor_genType creates a new moq of the For_genType type
func NewMoqFor_genType(scene *moq.Scene, config *moq.Config) *MoqFor_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFor_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFor_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Compiler moq.ParamIndexing
				Lookup   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Compiler moq.ParamIndexing
			Lookup   moq.ParamIndexing
		}{
			Compiler: moq.ParamIndexByValue,
			Lookup:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the For_genType type
func (m *MoqFor_genType) Mock() For_genType {
	return func(compiler string, lookup importer.Lookup) types.Importer {
		m.Scene.T.Helper()
		moq := &MoqFor_genType_mock{Moq: m}
		return moq.Fn(compiler, lookup)
	}
}

func (m *MoqFor_genType_mock) Fn(compiler string, lookup importer.Lookup) (result1 types.Importer) {
	m.Moq.Scene.T.Helper()
	params := MoqFor_genType_params{
		Compiler: compiler,
		Lookup:   lookup,
	}
	var results *MoqFor_genType_results
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
		result.DoFn(compiler, lookup)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(compiler, lookup)
	}
	return
}

func (m *MoqFor_genType) OnCall(compiler string, lookup importer.Lookup) *MoqFor_genType_fnRecorder {
	return &MoqFor_genType_fnRecorder{
		Params: MoqFor_genType_params{
			Compiler: compiler,
			Lookup:   lookup,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFor_genType_fnRecorder) Any() *MoqFor_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFor_genType_anyParams{Recorder: r}
}

func (a *MoqFor_genType_anyParams) Compiler() *MoqFor_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFor_genType_anyParams) Lookup() *MoqFor_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFor_genType_fnRecorder) Seq() *MoqFor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFor_genType_fnRecorder) NoSeq() *MoqFor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFor_genType_fnRecorder) ReturnResults(result1 types.Importer) *MoqFor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Importer
		}
		Sequence   uint32
		DoFn       MoqFor_genType_doFn
		DoReturnFn MoqFor_genType_doReturnFn
	}{
		Values: &struct {
			Result1 types.Importer
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFor_genType_fnRecorder) AndDo(fn MoqFor_genType_doFn) *MoqFor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFor_genType_fnRecorder) DoReturnResults(fn MoqFor_genType_doReturnFn) *MoqFor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Importer
		}
		Sequence   uint32
		DoFn       MoqFor_genType_doFn
		DoReturnFn MoqFor_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFor_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFor_genType_resultsByParams
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
		results = &MoqFor_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFor_genType_paramsKey]*MoqFor_genType_results{},
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
		r.Results = &MoqFor_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFor_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFor_genType_fnRecorder {
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
					Result1 types.Importer
				}
				Sequence   uint32
				DoFn       MoqFor_genType_doFn
				DoReturnFn MoqFor_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFor_genType) PrettyParams(params MoqFor_genType_params) string {
	return fmt.Sprintf("For_genType(%#v, %#v)", params.Compiler, params.Lookup)
}

func (m *MoqFor_genType) ParamsKey(params MoqFor_genType_params, anyParams uint64) MoqFor_genType_paramsKey {
	m.Scene.T.Helper()
	var compilerUsed string
	var compilerUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Compiler == moq.ParamIndexByValue {
			compilerUsed = params.Compiler
		} else {
			compilerUsedHash = hash.DeepHash(params.Compiler)
		}
	}
	var lookupUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Lookup == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The lookup parameter can't be indexed by value")
		}
		lookupUsedHash = hash.DeepHash(params.Lookup)
	}
	return MoqFor_genType_paramsKey{
		Params: struct{ Compiler string }{
			Compiler: compilerUsed,
		},
		Hashes: struct {
			Compiler hash.Hash
			Lookup   hash.Hash
		}{
			Compiler: compilerUsedHash,
			Lookup:   lookupUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFor_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFor_genType) AssertExpectationsMet() {
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
