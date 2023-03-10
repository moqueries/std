// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package strings

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Replace_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Replace_genType func(s, old, new string, n int) string

// MoqReplace_genType holds the state of a moq of the Replace_genType type
type MoqReplace_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReplace_genType_mock

	ResultsByParams []MoqReplace_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S   moq.ParamIndexing
			Old moq.ParamIndexing
			New moq.ParamIndexing
			N   moq.ParamIndexing
		}
	}
}

// MoqReplace_genType_mock isolates the mock interface of the Replace_genType
// type
type MoqReplace_genType_mock struct {
	Moq *MoqReplace_genType
}

// MoqReplace_genType_params holds the params of the Replace_genType type
type MoqReplace_genType_params struct {
	S, Old, New string
	N           int
}

// MoqReplace_genType_paramsKey holds the map key params of the Replace_genType
// type
type MoqReplace_genType_paramsKey struct {
	Params struct {
		S, Old, New string
		N           int
	}
	Hashes struct {
		S, Old, New hash.Hash
		N           hash.Hash
	}
}

// MoqReplace_genType_resultsByParams contains the results for a given set of
// parameters for the Replace_genType type
type MoqReplace_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReplace_genType_paramsKey]*MoqReplace_genType_results
}

// MoqReplace_genType_doFn defines the type of function needed when calling
// AndDo for the Replace_genType type
type MoqReplace_genType_doFn func(s, old, new string, n int)

// MoqReplace_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Replace_genType type
type MoqReplace_genType_doReturnFn func(s, old, new string, n int) string

// MoqReplace_genType_results holds the results of the Replace_genType type
type MoqReplace_genType_results struct {
	Params  MoqReplace_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplace_genType_doFn
		DoReturnFn MoqReplace_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReplace_genType_fnRecorder routes recorded function calls to the
// MoqReplace_genType moq
type MoqReplace_genType_fnRecorder struct {
	Params    MoqReplace_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReplace_genType_results
	Moq       *MoqReplace_genType
}

// MoqReplace_genType_anyParams isolates the any params functions of the
// Replace_genType type
type MoqReplace_genType_anyParams struct {
	Recorder *MoqReplace_genType_fnRecorder
}

// NewMoqReplace_genType creates a new moq of the Replace_genType type
func NewMoqReplace_genType(scene *moq.Scene, config *moq.Config) *MoqReplace_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReplace_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReplace_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S   moq.ParamIndexing
				Old moq.ParamIndexing
				New moq.ParamIndexing
				N   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S   moq.ParamIndexing
			Old moq.ParamIndexing
			New moq.ParamIndexing
			N   moq.ParamIndexing
		}{
			S:   moq.ParamIndexByValue,
			Old: moq.ParamIndexByValue,
			New: moq.ParamIndexByValue,
			N:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Replace_genType type
func (m *MoqReplace_genType) Mock() Replace_genType {
	return func(s, old, new string, n int) string {
		m.Scene.T.Helper()
		moq := &MoqReplace_genType_mock{Moq: m}
		return moq.Fn(s, old, new, n)
	}
}

func (m *MoqReplace_genType_mock) Fn(s, old, new string, n int) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqReplace_genType_params{
		S:   s,
		Old: old,
		New: new,
		N:   n,
	}
	var results *MoqReplace_genType_results
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
		result.DoFn(s, old, new, n)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, old, new, n)
	}
	return
}

func (m *MoqReplace_genType) OnCall(s, old, new string, n int) *MoqReplace_genType_fnRecorder {
	return &MoqReplace_genType_fnRecorder{
		Params: MoqReplace_genType_params{
			S:   s,
			Old: old,
			New: new,
			N:   n,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReplace_genType_fnRecorder) Any() *MoqReplace_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReplace_genType_anyParams{Recorder: r}
}

func (a *MoqReplace_genType_anyParams) S() *MoqReplace_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReplace_genType_anyParams) Old() *MoqReplace_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqReplace_genType_anyParams) New() *MoqReplace_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqReplace_genType_anyParams) N() *MoqReplace_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqReplace_genType_fnRecorder) Seq() *MoqReplace_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReplace_genType_fnRecorder) NoSeq() *MoqReplace_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReplace_genType_fnRecorder) ReturnResults(result1 string) *MoqReplace_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplace_genType_doFn
		DoReturnFn MoqReplace_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReplace_genType_fnRecorder) AndDo(fn MoqReplace_genType_doFn) *MoqReplace_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReplace_genType_fnRecorder) DoReturnResults(fn MoqReplace_genType_doReturnFn) *MoqReplace_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplace_genType_doFn
		DoReturnFn MoqReplace_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReplace_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReplace_genType_resultsByParams
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
		results = &MoqReplace_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReplace_genType_paramsKey]*MoqReplace_genType_results{},
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
		r.Results = &MoqReplace_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReplace_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReplace_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqReplace_genType_doFn
				DoReturnFn MoqReplace_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReplace_genType) PrettyParams(params MoqReplace_genType_params) string {
	return fmt.Sprintf("Replace_genType(%#v, %#v, %#v, %#v)", params.S, params.Old, params.New, params.N)
}

func (m *MoqReplace_genType) ParamsKey(params MoqReplace_genType_params, anyParams uint64) MoqReplace_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	var oldUsed string
	var oldUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Old == moq.ParamIndexByValue {
			oldUsed = params.Old
		} else {
			oldUsedHash = hash.DeepHash(params.Old)
		}
	}
	var newUsed string
	var newUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.New == moq.ParamIndexByValue {
			newUsed = params.New
		} else {
			newUsedHash = hash.DeepHash(params.New)
		}
	}
	var nUsed int
	var nUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.N == moq.ParamIndexByValue {
			nUsed = params.N
		} else {
			nUsedHash = hash.DeepHash(params.N)
		}
	}
	return MoqReplace_genType_paramsKey{
		Params: struct {
			S, Old, New string
			N           int
		}{
			S:   sUsed,
			Old: oldUsed,
			New: newUsed,
			N:   nUsed,
		},
		Hashes: struct {
			S, Old, New hash.Hash
			N           hash.Hash
		}{
			S:   sUsedHash,
			Old: oldUsedHash,
			New: newUsedHash,
			N:   nUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReplace_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReplace_genType) AssertExpectationsMet() {
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
