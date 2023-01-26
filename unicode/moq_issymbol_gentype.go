// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package unicode

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// IsSymbol_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type IsSymbol_genType func(r rune) bool

// MoqIsSymbol_genType holds the state of a moq of the IsSymbol_genType type
type MoqIsSymbol_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIsSymbol_genType_mock

	ResultsByParams []MoqIsSymbol_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
		}
	}
}

// MoqIsSymbol_genType_mock isolates the mock interface of the IsSymbol_genType
// type
type MoqIsSymbol_genType_mock struct {
	Moq *MoqIsSymbol_genType
}

// MoqIsSymbol_genType_params holds the params of the IsSymbol_genType type
type MoqIsSymbol_genType_params struct{ Param1 rune }

// MoqIsSymbol_genType_paramsKey holds the map key params of the
// IsSymbol_genType type
type MoqIsSymbol_genType_paramsKey struct {
	Params struct{ Param1 rune }
	Hashes struct{ Param1 hash.Hash }
}

// MoqIsSymbol_genType_resultsByParams contains the results for a given set of
// parameters for the IsSymbol_genType type
type MoqIsSymbol_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIsSymbol_genType_paramsKey]*MoqIsSymbol_genType_results
}

// MoqIsSymbol_genType_doFn defines the type of function needed when calling
// AndDo for the IsSymbol_genType type
type MoqIsSymbol_genType_doFn func(r rune)

// MoqIsSymbol_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IsSymbol_genType type
type MoqIsSymbol_genType_doReturnFn func(r rune) bool

// MoqIsSymbol_genType_results holds the results of the IsSymbol_genType type
type MoqIsSymbol_genType_results struct {
	Params  MoqIsSymbol_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSymbol_genType_doFn
		DoReturnFn MoqIsSymbol_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIsSymbol_genType_fnRecorder routes recorded function calls to the
// MoqIsSymbol_genType moq
type MoqIsSymbol_genType_fnRecorder struct {
	Params    MoqIsSymbol_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIsSymbol_genType_results
	Moq       *MoqIsSymbol_genType
}

// MoqIsSymbol_genType_anyParams isolates the any params functions of the
// IsSymbol_genType type
type MoqIsSymbol_genType_anyParams struct {
	Recorder *MoqIsSymbol_genType_fnRecorder
}

// NewMoqIsSymbol_genType creates a new moq of the IsSymbol_genType type
func NewMoqIsSymbol_genType(scene *moq.Scene, config *moq.Config) *MoqIsSymbol_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIsSymbol_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIsSymbol_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IsSymbol_genType type
func (m *MoqIsSymbol_genType) Mock() IsSymbol_genType {
	return func(param1 rune) bool {
		m.Scene.T.Helper()
		moq := &MoqIsSymbol_genType_mock{Moq: m}
		return moq.Fn(param1)
	}
}

func (m *MoqIsSymbol_genType_mock) Fn(param1 rune) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIsSymbol_genType_params{
		Param1: param1,
	}
	var results *MoqIsSymbol_genType_results
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
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqIsSymbol_genType) OnCall(param1 rune) *MoqIsSymbol_genType_fnRecorder {
	return &MoqIsSymbol_genType_fnRecorder{
		Params: MoqIsSymbol_genType_params{
			Param1: param1,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIsSymbol_genType_fnRecorder) Any() *MoqIsSymbol_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIsSymbol_genType_anyParams{Recorder: r}
}

func (a *MoqIsSymbol_genType_anyParams) Param1() *MoqIsSymbol_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIsSymbol_genType_fnRecorder) Seq() *MoqIsSymbol_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIsSymbol_genType_fnRecorder) NoSeq() *MoqIsSymbol_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIsSymbol_genType_fnRecorder) ReturnResults(result1 bool) *MoqIsSymbol_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSymbol_genType_doFn
		DoReturnFn MoqIsSymbol_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIsSymbol_genType_fnRecorder) AndDo(fn MoqIsSymbol_genType_doFn) *MoqIsSymbol_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIsSymbol_genType_fnRecorder) DoReturnResults(fn MoqIsSymbol_genType_doReturnFn) *MoqIsSymbol_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsSymbol_genType_doFn
		DoReturnFn MoqIsSymbol_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIsSymbol_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIsSymbol_genType_resultsByParams
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
		results = &MoqIsSymbol_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIsSymbol_genType_paramsKey]*MoqIsSymbol_genType_results{},
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
		r.Results = &MoqIsSymbol_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIsSymbol_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIsSymbol_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqIsSymbol_genType_doFn
				DoReturnFn MoqIsSymbol_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIsSymbol_genType) PrettyParams(params MoqIsSymbol_genType_params) string {
	return fmt.Sprintf("IsSymbol_genType(%#v)", params.Param1)
}

func (m *MoqIsSymbol_genType) ParamsKey(params MoqIsSymbol_genType_params, anyParams uint64) MoqIsSymbol_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used rune
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqIsSymbol_genType_paramsKey{
		Params: struct{ Param1 rune }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIsSymbol_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIsSymbol_genType) AssertExpectationsMet() {
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
