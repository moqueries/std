// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bytes

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ReplaceAll_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ReplaceAll_genType func(s, old, new []byte) []byte

// MoqReplaceAll_genType holds the state of a moq of the ReplaceAll_genType
// type
type MoqReplaceAll_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReplaceAll_genType_mock

	ResultsByParams []MoqReplaceAll_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S   moq.ParamIndexing
			Old moq.ParamIndexing
			New moq.ParamIndexing
		}
	}
}

// MoqReplaceAll_genType_mock isolates the mock interface of the
// ReplaceAll_genType type
type MoqReplaceAll_genType_mock struct {
	Moq *MoqReplaceAll_genType
}

// MoqReplaceAll_genType_params holds the params of the ReplaceAll_genType type
type MoqReplaceAll_genType_params struct{ S, Old, New []byte }

// MoqReplaceAll_genType_paramsKey holds the map key params of the
// ReplaceAll_genType type
type MoqReplaceAll_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ S, Old, New hash.Hash }
}

// MoqReplaceAll_genType_resultsByParams contains the results for a given set
// of parameters for the ReplaceAll_genType type
type MoqReplaceAll_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReplaceAll_genType_paramsKey]*MoqReplaceAll_genType_results
}

// MoqReplaceAll_genType_doFn defines the type of function needed when calling
// AndDo for the ReplaceAll_genType type
type MoqReplaceAll_genType_doFn func(s, old, new []byte)

// MoqReplaceAll_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReplaceAll_genType type
type MoqReplaceAll_genType_doReturnFn func(s, old, new []byte) []byte

// MoqReplaceAll_genType_results holds the results of the ReplaceAll_genType
// type
type MoqReplaceAll_genType_results struct {
	Params  MoqReplaceAll_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqReplaceAll_genType_doFn
		DoReturnFn MoqReplaceAll_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReplaceAll_genType_fnRecorder routes recorded function calls to the
// MoqReplaceAll_genType moq
type MoqReplaceAll_genType_fnRecorder struct {
	Params    MoqReplaceAll_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReplaceAll_genType_results
	Moq       *MoqReplaceAll_genType
}

// MoqReplaceAll_genType_anyParams isolates the any params functions of the
// ReplaceAll_genType type
type MoqReplaceAll_genType_anyParams struct {
	Recorder *MoqReplaceAll_genType_fnRecorder
}

// NewMoqReplaceAll_genType creates a new moq of the ReplaceAll_genType type
func NewMoqReplaceAll_genType(scene *moq.Scene, config *moq.Config) *MoqReplaceAll_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReplaceAll_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReplaceAll_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S   moq.ParamIndexing
				Old moq.ParamIndexing
				New moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S   moq.ParamIndexing
			Old moq.ParamIndexing
			New moq.ParamIndexing
		}{
			S:   moq.ParamIndexByHash,
			Old: moq.ParamIndexByHash,
			New: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ReplaceAll_genType type
func (m *MoqReplaceAll_genType) Mock() ReplaceAll_genType {
	return func(s, old, new []byte) []byte {
		m.Scene.T.Helper()
		moq := &MoqReplaceAll_genType_mock{Moq: m}
		return moq.Fn(s, old, new)
	}
}

func (m *MoqReplaceAll_genType_mock) Fn(s, old, new []byte) (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqReplaceAll_genType_params{
		S:   s,
		Old: old,
		New: new,
	}
	var results *MoqReplaceAll_genType_results
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
		result.DoFn(s, old, new)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, old, new)
	}
	return
}

func (m *MoqReplaceAll_genType) OnCall(s, old, new []byte) *MoqReplaceAll_genType_fnRecorder {
	return &MoqReplaceAll_genType_fnRecorder{
		Params: MoqReplaceAll_genType_params{
			S:   s,
			Old: old,
			New: new,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReplaceAll_genType_fnRecorder) Any() *MoqReplaceAll_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReplaceAll_genType_anyParams{Recorder: r}
}

func (a *MoqReplaceAll_genType_anyParams) S() *MoqReplaceAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReplaceAll_genType_anyParams) Old() *MoqReplaceAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqReplaceAll_genType_anyParams) New() *MoqReplaceAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqReplaceAll_genType_fnRecorder) Seq() *MoqReplaceAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReplaceAll_genType_fnRecorder) NoSeq() *MoqReplaceAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReplaceAll_genType_fnRecorder) ReturnResults(result1 []byte) *MoqReplaceAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqReplaceAll_genType_doFn
		DoReturnFn MoqReplaceAll_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReplaceAll_genType_fnRecorder) AndDo(fn MoqReplaceAll_genType_doFn) *MoqReplaceAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReplaceAll_genType_fnRecorder) DoReturnResults(fn MoqReplaceAll_genType_doReturnFn) *MoqReplaceAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqReplaceAll_genType_doFn
		DoReturnFn MoqReplaceAll_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReplaceAll_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReplaceAll_genType_resultsByParams
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
		results = &MoqReplaceAll_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReplaceAll_genType_paramsKey]*MoqReplaceAll_genType_results{},
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
		r.Results = &MoqReplaceAll_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReplaceAll_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReplaceAll_genType_fnRecorder {
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
					Result1 []byte
				}
				Sequence   uint32
				DoFn       MoqReplaceAll_genType_doFn
				DoReturnFn MoqReplaceAll_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReplaceAll_genType) PrettyParams(params MoqReplaceAll_genType_params) string {
	return fmt.Sprintf("ReplaceAll_genType(%#v, %#v, %#v)", params.S, params.Old, params.New)
}

func (m *MoqReplaceAll_genType) ParamsKey(params MoqReplaceAll_genType_params, anyParams uint64) MoqReplaceAll_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The s parameter can't be indexed by value")
		}
		sUsedHash = hash.DeepHash(params.S)
	}
	var oldUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Old == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The old parameter can't be indexed by value")
		}
		oldUsedHash = hash.DeepHash(params.Old)
	}
	var newUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.New == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The new parameter can't be indexed by value")
		}
		newUsedHash = hash.DeepHash(params.New)
	}
	return MoqReplaceAll_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ S, Old, New hash.Hash }{
			S:   sUsedHash,
			Old: oldUsedHash,
			New: newUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReplaceAll_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReplaceAll_genType) AssertExpectationsMet() {
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
