// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Stack_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Stack_genType func(buf []byte, all bool) int

// MoqStack_genType holds the state of a moq of the Stack_genType type
type MoqStack_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStack_genType_mock

	ResultsByParams []MoqStack_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Buf moq.ParamIndexing
			All moq.ParamIndexing
		}
	}
}

// MoqStack_genType_mock isolates the mock interface of the Stack_genType type
type MoqStack_genType_mock struct {
	Moq *MoqStack_genType
}

// MoqStack_genType_params holds the params of the Stack_genType type
type MoqStack_genType_params struct {
	Buf []byte
	All bool
}

// MoqStack_genType_paramsKey holds the map key params of the Stack_genType
// type
type MoqStack_genType_paramsKey struct {
	Params struct{ All bool }
	Hashes struct {
		Buf hash.Hash
		All hash.Hash
	}
}

// MoqStack_genType_resultsByParams contains the results for a given set of
// parameters for the Stack_genType type
type MoqStack_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStack_genType_paramsKey]*MoqStack_genType_results
}

// MoqStack_genType_doFn defines the type of function needed when calling AndDo
// for the Stack_genType type
type MoqStack_genType_doFn func(buf []byte, all bool)

// MoqStack_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stack_genType type
type MoqStack_genType_doReturnFn func(buf []byte, all bool) int

// MoqStack_genType_results holds the results of the Stack_genType type
type MoqStack_genType_results struct {
	Params  MoqStack_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStack_genType_doFn
		DoReturnFn MoqStack_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStack_genType_fnRecorder routes recorded function calls to the
// MoqStack_genType moq
type MoqStack_genType_fnRecorder struct {
	Params    MoqStack_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStack_genType_results
	Moq       *MoqStack_genType
}

// MoqStack_genType_anyParams isolates the any params functions of the
// Stack_genType type
type MoqStack_genType_anyParams struct {
	Recorder *MoqStack_genType_fnRecorder
}

// NewMoqStack_genType creates a new moq of the Stack_genType type
func NewMoqStack_genType(scene *moq.Scene, config *moq.Config) *MoqStack_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStack_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStack_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Buf moq.ParamIndexing
				All moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Buf moq.ParamIndexing
			All moq.ParamIndexing
		}{
			Buf: moq.ParamIndexByHash,
			All: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Stack_genType type
func (m *MoqStack_genType) Mock() Stack_genType {
	return func(buf []byte, all bool) int {
		m.Scene.T.Helper()
		moq := &MoqStack_genType_mock{Moq: m}
		return moq.Fn(buf, all)
	}
}

func (m *MoqStack_genType_mock) Fn(buf []byte, all bool) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqStack_genType_params{
		Buf: buf,
		All: all,
	}
	var results *MoqStack_genType_results
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
		result.DoFn(buf, all)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(buf, all)
	}
	return
}

func (m *MoqStack_genType) OnCall(buf []byte, all bool) *MoqStack_genType_fnRecorder {
	return &MoqStack_genType_fnRecorder{
		Params: MoqStack_genType_params{
			Buf: buf,
			All: all,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStack_genType_fnRecorder) Any() *MoqStack_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStack_genType_anyParams{Recorder: r}
}

func (a *MoqStack_genType_anyParams) Buf() *MoqStack_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqStack_genType_anyParams) All() *MoqStack_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqStack_genType_fnRecorder) Seq() *MoqStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStack_genType_fnRecorder) NoSeq() *MoqStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStack_genType_fnRecorder) ReturnResults(result1 int) *MoqStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStack_genType_doFn
		DoReturnFn MoqStack_genType_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStack_genType_fnRecorder) AndDo(fn MoqStack_genType_doFn) *MoqStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStack_genType_fnRecorder) DoReturnResults(fn MoqStack_genType_doReturnFn) *MoqStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStack_genType_doFn
		DoReturnFn MoqStack_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStack_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStack_genType_resultsByParams
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
		results = &MoqStack_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStack_genType_paramsKey]*MoqStack_genType_results{},
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
		r.Results = &MoqStack_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStack_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStack_genType_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqStack_genType_doFn
				DoReturnFn MoqStack_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStack_genType) PrettyParams(params MoqStack_genType_params) string {
	return fmt.Sprintf("Stack_genType(%#v, %#v)", params.Buf, params.All)
}

func (m *MoqStack_genType) ParamsKey(params MoqStack_genType_params, anyParams uint64) MoqStack_genType_paramsKey {
	m.Scene.T.Helper()
	var bufUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Buf == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The buf parameter can't be indexed by value")
		}
		bufUsedHash = hash.DeepHash(params.Buf)
	}
	var allUsed bool
	var allUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.All == moq.ParamIndexByValue {
			allUsed = params.All
		} else {
			allUsedHash = hash.DeepHash(params.All)
		}
	}
	return MoqStack_genType_paramsKey{
		Params: struct{ All bool }{
			All: allUsed,
		},
		Hashes: struct {
			Buf hash.Hash
			All hash.Hash
		}{
			Buf: bufUsedHash,
			All: allUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStack_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStack_genType) AssertExpectationsMet() {
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
