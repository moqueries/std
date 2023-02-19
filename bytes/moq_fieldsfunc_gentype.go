// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bytes

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// FieldsFunc_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type FieldsFunc_genType func(s []byte, f func(rune) bool) [][]byte

// MoqFieldsFunc_genType holds the state of a moq of the FieldsFunc_genType
// type
type MoqFieldsFunc_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFieldsFunc_genType_mock

	ResultsByParams []MoqFieldsFunc_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S moq.ParamIndexing
			F moq.ParamIndexing
		}
	}
}

// MoqFieldsFunc_genType_mock isolates the mock interface of the
// FieldsFunc_genType type
type MoqFieldsFunc_genType_mock struct {
	Moq *MoqFieldsFunc_genType
}

// MoqFieldsFunc_genType_params holds the params of the FieldsFunc_genType type
type MoqFieldsFunc_genType_params struct {
	S []byte
	F func(rune) bool
}

// MoqFieldsFunc_genType_paramsKey holds the map key params of the
// FieldsFunc_genType type
type MoqFieldsFunc_genType_paramsKey struct {
	Params struct{}
	Hashes struct {
		S hash.Hash
		F hash.Hash
	}
}

// MoqFieldsFunc_genType_resultsByParams contains the results for a given set
// of parameters for the FieldsFunc_genType type
type MoqFieldsFunc_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFieldsFunc_genType_paramsKey]*MoqFieldsFunc_genType_results
}

// MoqFieldsFunc_genType_doFn defines the type of function needed when calling
// AndDo for the FieldsFunc_genType type
type MoqFieldsFunc_genType_doFn func(s []byte, f func(rune) bool)

// MoqFieldsFunc_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the FieldsFunc_genType type
type MoqFieldsFunc_genType_doReturnFn func(s []byte, f func(rune) bool) [][]byte

// MoqFieldsFunc_genType_results holds the results of the FieldsFunc_genType
// type
type MoqFieldsFunc_genType_results struct {
	Params  MoqFieldsFunc_genType_params
	Results []struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqFieldsFunc_genType_doFn
		DoReturnFn MoqFieldsFunc_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFieldsFunc_genType_fnRecorder routes recorded function calls to the
// MoqFieldsFunc_genType moq
type MoqFieldsFunc_genType_fnRecorder struct {
	Params    MoqFieldsFunc_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFieldsFunc_genType_results
	Moq       *MoqFieldsFunc_genType
}

// MoqFieldsFunc_genType_anyParams isolates the any params functions of the
// FieldsFunc_genType type
type MoqFieldsFunc_genType_anyParams struct {
	Recorder *MoqFieldsFunc_genType_fnRecorder
}

// NewMoqFieldsFunc_genType creates a new moq of the FieldsFunc_genType type
func NewMoqFieldsFunc_genType(scene *moq.Scene, config *moq.Config) *MoqFieldsFunc_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFieldsFunc_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFieldsFunc_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S moq.ParamIndexing
				F moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S moq.ParamIndexing
			F moq.ParamIndexing
		}{
			S: moq.ParamIndexByHash,
			F: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the FieldsFunc_genType type
func (m *MoqFieldsFunc_genType) Mock() FieldsFunc_genType {
	return func(s []byte, f func(rune) bool) [][]byte {
		m.Scene.T.Helper()
		moq := &MoqFieldsFunc_genType_mock{Moq: m}
		return moq.Fn(s, f)
	}
}

func (m *MoqFieldsFunc_genType_mock) Fn(s []byte, f func(rune) bool) (result1 [][]byte) {
	m.Moq.Scene.T.Helper()
	params := MoqFieldsFunc_genType_params{
		S: s,
		F: f,
	}
	var results *MoqFieldsFunc_genType_results
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
		result.DoFn(s, f)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, f)
	}
	return
}

func (m *MoqFieldsFunc_genType) OnCall(s []byte, f func(rune) bool) *MoqFieldsFunc_genType_fnRecorder {
	return &MoqFieldsFunc_genType_fnRecorder{
		Params: MoqFieldsFunc_genType_params{
			S: s,
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFieldsFunc_genType_fnRecorder) Any() *MoqFieldsFunc_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFieldsFunc_genType_anyParams{Recorder: r}
}

func (a *MoqFieldsFunc_genType_anyParams) S() *MoqFieldsFunc_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFieldsFunc_genType_anyParams) F() *MoqFieldsFunc_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFieldsFunc_genType_fnRecorder) Seq() *MoqFieldsFunc_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFieldsFunc_genType_fnRecorder) NoSeq() *MoqFieldsFunc_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFieldsFunc_genType_fnRecorder) ReturnResults(result1 [][]byte) *MoqFieldsFunc_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqFieldsFunc_genType_doFn
		DoReturnFn MoqFieldsFunc_genType_doReturnFn
	}{
		Values: &struct {
			Result1 [][]byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFieldsFunc_genType_fnRecorder) AndDo(fn MoqFieldsFunc_genType_doFn) *MoqFieldsFunc_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFieldsFunc_genType_fnRecorder) DoReturnResults(fn MoqFieldsFunc_genType_doReturnFn) *MoqFieldsFunc_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqFieldsFunc_genType_doFn
		DoReturnFn MoqFieldsFunc_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFieldsFunc_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFieldsFunc_genType_resultsByParams
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
		results = &MoqFieldsFunc_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFieldsFunc_genType_paramsKey]*MoqFieldsFunc_genType_results{},
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
		r.Results = &MoqFieldsFunc_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFieldsFunc_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFieldsFunc_genType_fnRecorder {
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
					Result1 [][]byte
				}
				Sequence   uint32
				DoFn       MoqFieldsFunc_genType_doFn
				DoReturnFn MoqFieldsFunc_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFieldsFunc_genType) PrettyParams(params MoqFieldsFunc_genType_params) string {
	return fmt.Sprintf("FieldsFunc_genType(%#v, %#v)", params.S, moq.FnString(params.F))
}

func (m *MoqFieldsFunc_genType) ParamsKey(params MoqFieldsFunc_genType_params, anyParams uint64) MoqFieldsFunc_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The s parameter can't be indexed by value")
		}
		sUsedHash = hash.DeepHash(params.S)
	}
	var fUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The f parameter can't be indexed by value")
		}
		fUsedHash = hash.DeepHash(params.F)
	}
	return MoqFieldsFunc_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct {
			S hash.Hash
			F hash.Hash
		}{
			S: sUsedHash,
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFieldsFunc_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFieldsFunc_genType) AssertExpectationsMet() {
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
