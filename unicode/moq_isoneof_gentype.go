// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package unicode

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"unicode"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// IsOneOf_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type IsOneOf_genType func(ranges []*unicode.RangeTable, r rune) bool

// MoqIsOneOf_genType holds the state of a moq of the IsOneOf_genType type
type MoqIsOneOf_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIsOneOf_genType_mock

	ResultsByParams []MoqIsOneOf_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ranges moq.ParamIndexing
			Param2 moq.ParamIndexing
		}
	}
}

// MoqIsOneOf_genType_mock isolates the mock interface of the IsOneOf_genType
// type
type MoqIsOneOf_genType_mock struct {
	Moq *MoqIsOneOf_genType
}

// MoqIsOneOf_genType_params holds the params of the IsOneOf_genType type
type MoqIsOneOf_genType_params struct {
	Ranges []*unicode.RangeTable
	Param2 rune
}

// MoqIsOneOf_genType_paramsKey holds the map key params of the IsOneOf_genType
// type
type MoqIsOneOf_genType_paramsKey struct {
	Params struct{ Param2 rune }
	Hashes struct {
		Ranges hash.Hash
		Param2 hash.Hash
	}
}

// MoqIsOneOf_genType_resultsByParams contains the results for a given set of
// parameters for the IsOneOf_genType type
type MoqIsOneOf_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIsOneOf_genType_paramsKey]*MoqIsOneOf_genType_results
}

// MoqIsOneOf_genType_doFn defines the type of function needed when calling
// AndDo for the IsOneOf_genType type
type MoqIsOneOf_genType_doFn func(ranges []*unicode.RangeTable, r rune)

// MoqIsOneOf_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IsOneOf_genType type
type MoqIsOneOf_genType_doReturnFn func(ranges []*unicode.RangeTable, r rune) bool

// MoqIsOneOf_genType_results holds the results of the IsOneOf_genType type
type MoqIsOneOf_genType_results struct {
	Params  MoqIsOneOf_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsOneOf_genType_doFn
		DoReturnFn MoqIsOneOf_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIsOneOf_genType_fnRecorder routes recorded function calls to the
// MoqIsOneOf_genType moq
type MoqIsOneOf_genType_fnRecorder struct {
	Params    MoqIsOneOf_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIsOneOf_genType_results
	Moq       *MoqIsOneOf_genType
}

// MoqIsOneOf_genType_anyParams isolates the any params functions of the
// IsOneOf_genType type
type MoqIsOneOf_genType_anyParams struct {
	Recorder *MoqIsOneOf_genType_fnRecorder
}

// NewMoqIsOneOf_genType creates a new moq of the IsOneOf_genType type
func NewMoqIsOneOf_genType(scene *moq.Scene, config *moq.Config) *MoqIsOneOf_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIsOneOf_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIsOneOf_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ranges moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ranges moq.ParamIndexing
			Param2 moq.ParamIndexing
		}{
			Ranges: moq.ParamIndexByHash,
			Param2: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IsOneOf_genType type
func (m *MoqIsOneOf_genType) Mock() IsOneOf_genType {
	return func(ranges []*unicode.RangeTable, param2 rune) bool {
		m.Scene.T.Helper()
		moq := &MoqIsOneOf_genType_mock{Moq: m}
		return moq.Fn(ranges, param2)
	}
}

func (m *MoqIsOneOf_genType_mock) Fn(ranges []*unicode.RangeTable, param2 rune) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIsOneOf_genType_params{
		Ranges: ranges,
		Param2: param2,
	}
	var results *MoqIsOneOf_genType_results
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
		result.DoFn(ranges, param2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ranges, param2)
	}
	return
}

func (m *MoqIsOneOf_genType) OnCall(ranges []*unicode.RangeTable, param2 rune) *MoqIsOneOf_genType_fnRecorder {
	return &MoqIsOneOf_genType_fnRecorder{
		Params: MoqIsOneOf_genType_params{
			Ranges: ranges,
			Param2: param2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIsOneOf_genType_fnRecorder) Any() *MoqIsOneOf_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIsOneOf_genType_anyParams{Recorder: r}
}

func (a *MoqIsOneOf_genType_anyParams) Ranges() *MoqIsOneOf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqIsOneOf_genType_anyParams) Param2() *MoqIsOneOf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqIsOneOf_genType_fnRecorder) Seq() *MoqIsOneOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIsOneOf_genType_fnRecorder) NoSeq() *MoqIsOneOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIsOneOf_genType_fnRecorder) ReturnResults(result1 bool) *MoqIsOneOf_genType_fnRecorder {
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
		DoFn       MoqIsOneOf_genType_doFn
		DoReturnFn MoqIsOneOf_genType_doReturnFn
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

func (r *MoqIsOneOf_genType_fnRecorder) AndDo(fn MoqIsOneOf_genType_doFn) *MoqIsOneOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIsOneOf_genType_fnRecorder) DoReturnResults(fn MoqIsOneOf_genType_doReturnFn) *MoqIsOneOf_genType_fnRecorder {
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
		DoFn       MoqIsOneOf_genType_doFn
		DoReturnFn MoqIsOneOf_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIsOneOf_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIsOneOf_genType_resultsByParams
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
		results = &MoqIsOneOf_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIsOneOf_genType_paramsKey]*MoqIsOneOf_genType_results{},
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
		r.Results = &MoqIsOneOf_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIsOneOf_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIsOneOf_genType_fnRecorder {
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
				DoFn       MoqIsOneOf_genType_doFn
				DoReturnFn MoqIsOneOf_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIsOneOf_genType) PrettyParams(params MoqIsOneOf_genType_params) string {
	return fmt.Sprintf("IsOneOf_genType(%#v, %#v)", params.Ranges, params.Param2)
}

func (m *MoqIsOneOf_genType) ParamsKey(params MoqIsOneOf_genType_params, anyParams uint64) MoqIsOneOf_genType_paramsKey {
	m.Scene.T.Helper()
	var rangesUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Ranges == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The ranges parameter can't be indexed by value")
		}
		rangesUsedHash = hash.DeepHash(params.Ranges)
	}
	var param2Used rune
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	return MoqIsOneOf_genType_paramsKey{
		Params: struct{ Param2 rune }{
			Param2: param2Used,
		},
		Hashes: struct {
			Ranges hash.Hash
			Param2 hash.Hash
		}{
			Ranges: rangesUsedHash,
			Param2: param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIsOneOf_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIsOneOf_genType) AssertExpectationsMet() {
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
