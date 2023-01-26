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

// Is_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Is_genType func(rangeTab *unicode.RangeTable, r rune) bool

// MoqIs_genType holds the state of a moq of the Is_genType type
type MoqIs_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIs_genType_mock

	ResultsByParams []MoqIs_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			RangeTab moq.ParamIndexing
			Param2   moq.ParamIndexing
		}
	}
}

// MoqIs_genType_mock isolates the mock interface of the Is_genType type
type MoqIs_genType_mock struct {
	Moq *MoqIs_genType
}

// MoqIs_genType_params holds the params of the Is_genType type
type MoqIs_genType_params struct {
	RangeTab *unicode.RangeTable
	Param2   rune
}

// MoqIs_genType_paramsKey holds the map key params of the Is_genType type
type MoqIs_genType_paramsKey struct {
	Params struct {
		RangeTab *unicode.RangeTable
		Param2   rune
	}
	Hashes struct {
		RangeTab hash.Hash
		Param2   hash.Hash
	}
}

// MoqIs_genType_resultsByParams contains the results for a given set of
// parameters for the Is_genType type
type MoqIs_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIs_genType_paramsKey]*MoqIs_genType_results
}

// MoqIs_genType_doFn defines the type of function needed when calling AndDo
// for the Is_genType type
type MoqIs_genType_doFn func(rangeTab *unicode.RangeTable, r rune)

// MoqIs_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Is_genType type
type MoqIs_genType_doReturnFn func(rangeTab *unicode.RangeTable, r rune) bool

// MoqIs_genType_results holds the results of the Is_genType type
type MoqIs_genType_results struct {
	Params  MoqIs_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIs_genType_doFn
		DoReturnFn MoqIs_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIs_genType_fnRecorder routes recorded function calls to the MoqIs_genType
// moq
type MoqIs_genType_fnRecorder struct {
	Params    MoqIs_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIs_genType_results
	Moq       *MoqIs_genType
}

// MoqIs_genType_anyParams isolates the any params functions of the Is_genType
// type
type MoqIs_genType_anyParams struct {
	Recorder *MoqIs_genType_fnRecorder
}

// NewMoqIs_genType creates a new moq of the Is_genType type
func NewMoqIs_genType(scene *moq.Scene, config *moq.Config) *MoqIs_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIs_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIs_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				RangeTab moq.ParamIndexing
				Param2   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			RangeTab moq.ParamIndexing
			Param2   moq.ParamIndexing
		}{
			RangeTab: moq.ParamIndexByHash,
			Param2:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Is_genType type
func (m *MoqIs_genType) Mock() Is_genType {
	return func(rangeTab *unicode.RangeTable, param2 rune) bool {
		m.Scene.T.Helper()
		moq := &MoqIs_genType_mock{Moq: m}
		return moq.Fn(rangeTab, param2)
	}
}

func (m *MoqIs_genType_mock) Fn(rangeTab *unicode.RangeTable, param2 rune) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIs_genType_params{
		RangeTab: rangeTab,
		Param2:   param2,
	}
	var results *MoqIs_genType_results
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
		result.DoFn(rangeTab, param2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(rangeTab, param2)
	}
	return
}

func (m *MoqIs_genType) OnCall(rangeTab *unicode.RangeTable, param2 rune) *MoqIs_genType_fnRecorder {
	return &MoqIs_genType_fnRecorder{
		Params: MoqIs_genType_params{
			RangeTab: rangeTab,
			Param2:   param2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIs_genType_fnRecorder) Any() *MoqIs_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIs_genType_anyParams{Recorder: r}
}

func (a *MoqIs_genType_anyParams) RangeTab() *MoqIs_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqIs_genType_anyParams) Param2() *MoqIs_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqIs_genType_fnRecorder) Seq() *MoqIs_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIs_genType_fnRecorder) NoSeq() *MoqIs_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIs_genType_fnRecorder) ReturnResults(result1 bool) *MoqIs_genType_fnRecorder {
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
		DoFn       MoqIs_genType_doFn
		DoReturnFn MoqIs_genType_doReturnFn
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

func (r *MoqIs_genType_fnRecorder) AndDo(fn MoqIs_genType_doFn) *MoqIs_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIs_genType_fnRecorder) DoReturnResults(fn MoqIs_genType_doReturnFn) *MoqIs_genType_fnRecorder {
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
		DoFn       MoqIs_genType_doFn
		DoReturnFn MoqIs_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIs_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIs_genType_resultsByParams
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
		results = &MoqIs_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIs_genType_paramsKey]*MoqIs_genType_results{},
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
		r.Results = &MoqIs_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIs_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIs_genType_fnRecorder {
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
				DoFn       MoqIs_genType_doFn
				DoReturnFn MoqIs_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIs_genType) PrettyParams(params MoqIs_genType_params) string {
	return fmt.Sprintf("Is_genType(%#v, %#v)", params.RangeTab, params.Param2)
}

func (m *MoqIs_genType) ParamsKey(params MoqIs_genType_params, anyParams uint64) MoqIs_genType_paramsKey {
	m.Scene.T.Helper()
	var rangeTabUsed *unicode.RangeTable
	var rangeTabUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.RangeTab == moq.ParamIndexByValue {
			rangeTabUsed = params.RangeTab
		} else {
			rangeTabUsedHash = hash.DeepHash(params.RangeTab)
		}
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
	return MoqIs_genType_paramsKey{
		Params: struct {
			RangeTab *unicode.RangeTable
			Param2   rune
		}{
			RangeTab: rangeTabUsed,
			Param2:   param2Used,
		},
		Hashes: struct {
			RangeTab hash.Hash
			Param2   hash.Hash
		}{
			RangeTab: rangeTabUsedHash,
			Param2:   param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIs_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIs_genType) AssertExpectationsMet() {
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
