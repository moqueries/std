// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package path

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Join_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Join_genType func(elem ...string) string

// MoqJoin_genType holds the state of a moq of the Join_genType type
type MoqJoin_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqJoin_genType_mock

	ResultsByParams []MoqJoin_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Elem moq.ParamIndexing
		}
	}
}

// MoqJoin_genType_mock isolates the mock interface of the Join_genType type
type MoqJoin_genType_mock struct {
	Moq *MoqJoin_genType
}

// MoqJoin_genType_params holds the params of the Join_genType type
type MoqJoin_genType_params struct{ Elem []string }

// MoqJoin_genType_paramsKey holds the map key params of the Join_genType type
type MoqJoin_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Elem hash.Hash }
}

// MoqJoin_genType_resultsByParams contains the results for a given set of
// parameters for the Join_genType type
type MoqJoin_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqJoin_genType_paramsKey]*MoqJoin_genType_results
}

// MoqJoin_genType_doFn defines the type of function needed when calling AndDo
// for the Join_genType type
type MoqJoin_genType_doFn func(elem ...string)

// MoqJoin_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Join_genType type
type MoqJoin_genType_doReturnFn func(elem ...string) string

// MoqJoin_genType_results holds the results of the Join_genType type
type MoqJoin_genType_results struct {
	Params  MoqJoin_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqJoin_genType_doFn
		DoReturnFn MoqJoin_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqJoin_genType_fnRecorder routes recorded function calls to the
// MoqJoin_genType moq
type MoqJoin_genType_fnRecorder struct {
	Params    MoqJoin_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqJoin_genType_results
	Moq       *MoqJoin_genType
}

// MoqJoin_genType_anyParams isolates the any params functions of the
// Join_genType type
type MoqJoin_genType_anyParams struct {
	Recorder *MoqJoin_genType_fnRecorder
}

// NewMoqJoin_genType creates a new moq of the Join_genType type
func NewMoqJoin_genType(scene *moq.Scene, config *moq.Config) *MoqJoin_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqJoin_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqJoin_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Elem moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Elem moq.ParamIndexing
		}{
			Elem: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Join_genType type
func (m *MoqJoin_genType) Mock() Join_genType {
	return func(elem ...string) string {
		m.Scene.T.Helper()
		moq := &MoqJoin_genType_mock{Moq: m}
		return moq.Fn(elem...)
	}
}

func (m *MoqJoin_genType_mock) Fn(elem ...string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqJoin_genType_params{
		Elem: elem,
	}
	var results *MoqJoin_genType_results
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
		result.DoFn(elem...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(elem...)
	}
	return
}

func (m *MoqJoin_genType) OnCall(elem ...string) *MoqJoin_genType_fnRecorder {
	return &MoqJoin_genType_fnRecorder{
		Params: MoqJoin_genType_params{
			Elem: elem,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqJoin_genType_fnRecorder) Any() *MoqJoin_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqJoin_genType_anyParams{Recorder: r}
}

func (a *MoqJoin_genType_anyParams) Elem() *MoqJoin_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqJoin_genType_fnRecorder) Seq() *MoqJoin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqJoin_genType_fnRecorder) NoSeq() *MoqJoin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqJoin_genType_fnRecorder) ReturnResults(result1 string) *MoqJoin_genType_fnRecorder {
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
		DoFn       MoqJoin_genType_doFn
		DoReturnFn MoqJoin_genType_doReturnFn
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

func (r *MoqJoin_genType_fnRecorder) AndDo(fn MoqJoin_genType_doFn) *MoqJoin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqJoin_genType_fnRecorder) DoReturnResults(fn MoqJoin_genType_doReturnFn) *MoqJoin_genType_fnRecorder {
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
		DoFn       MoqJoin_genType_doFn
		DoReturnFn MoqJoin_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqJoin_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqJoin_genType_resultsByParams
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
		results = &MoqJoin_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqJoin_genType_paramsKey]*MoqJoin_genType_results{},
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
		r.Results = &MoqJoin_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqJoin_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqJoin_genType_fnRecorder {
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
				DoFn       MoqJoin_genType_doFn
				DoReturnFn MoqJoin_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqJoin_genType) PrettyParams(params MoqJoin_genType_params) string {
	return fmt.Sprintf("Join_genType(%#v)", params.Elem)
}

func (m *MoqJoin_genType) ParamsKey(params MoqJoin_genType_params, anyParams uint64) MoqJoin_genType_paramsKey {
	m.Scene.T.Helper()
	var elemUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Elem == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The elem parameter can't be indexed by value")
		}
		elemUsedHash = hash.DeepHash(params.Elem)
	}
	return MoqJoin_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Elem hash.Hash }{
			Elem: elemUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqJoin_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqJoin_genType) AssertExpectationsMet() {
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