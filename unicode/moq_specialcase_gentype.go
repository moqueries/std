// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package unicode

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that unicode.SpecialCase_genType is
// mocked completely
var _ SpecialCase_genType = (*MoqSpecialCase_genType_mock)(nil)

// SpecialCase_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type SpecialCase_genType interface {
	ToUpper(r rune) rune
	ToTitle(r rune) rune
	ToLower(r rune) rune
}

// MoqSpecialCase_genType holds the state of a moq of the SpecialCase_genType
// type
type MoqSpecialCase_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSpecialCase_genType_mock

	ResultsByParams_ToUpper []MoqSpecialCase_genType_ToUpper_resultsByParams
	ResultsByParams_ToTitle []MoqSpecialCase_genType_ToTitle_resultsByParams
	ResultsByParams_ToLower []MoqSpecialCase_genType_ToLower_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ToUpper struct {
				Param1 moq.ParamIndexing
			}
			ToTitle struct {
				Param1 moq.ParamIndexing
			}
			ToLower struct {
				Param1 moq.ParamIndexing
			}
		}
	}
	// MoqSpecialCase_genType_mock isolates the mock interface of the
}

// SpecialCase_genType type
type MoqSpecialCase_genType_mock struct {
	Moq *MoqSpecialCase_genType
}

// MoqSpecialCase_genType_recorder isolates the recorder interface of the
// SpecialCase_genType type
type MoqSpecialCase_genType_recorder struct {
	Moq *MoqSpecialCase_genType
}

// MoqSpecialCase_genType_ToUpper_params holds the params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_params struct{ Param1 rune }

// MoqSpecialCase_genType_ToUpper_paramsKey holds the map key params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_paramsKey struct {
	Params struct{ Param1 rune }
	Hashes struct{ Param1 hash.Hash }
}

// MoqSpecialCase_genType_ToUpper_resultsByParams contains the results for a
// given set of parameters for the SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSpecialCase_genType_ToUpper_paramsKey]*MoqSpecialCase_genType_ToUpper_results
}

// MoqSpecialCase_genType_ToUpper_doFn defines the type of function needed when
// calling AndDo for the SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_doFn func(r rune)

// MoqSpecialCase_genType_ToUpper_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_doReturnFn func(r rune) rune

// MoqSpecialCase_genType_ToUpper_results holds the results of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_results struct {
	Params  MoqSpecialCase_genType_ToUpper_params
	Results []struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToUpper_doFn
		DoReturnFn MoqSpecialCase_genType_ToUpper_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSpecialCase_genType_ToUpper_fnRecorder routes recorded function calls to
// the MoqSpecialCase_genType moq
type MoqSpecialCase_genType_ToUpper_fnRecorder struct {
	Params    MoqSpecialCase_genType_ToUpper_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSpecialCase_genType_ToUpper_results
	Moq       *MoqSpecialCase_genType
}

// MoqSpecialCase_genType_ToUpper_anyParams isolates the any params functions
// of the SpecialCase_genType type
type MoqSpecialCase_genType_ToUpper_anyParams struct {
	Recorder *MoqSpecialCase_genType_ToUpper_fnRecorder
}

// MoqSpecialCase_genType_ToTitle_params holds the params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_params struct{ Param1 rune }

// MoqSpecialCase_genType_ToTitle_paramsKey holds the map key params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_paramsKey struct {
	Params struct{ Param1 rune }
	Hashes struct{ Param1 hash.Hash }
}

// MoqSpecialCase_genType_ToTitle_resultsByParams contains the results for a
// given set of parameters for the SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSpecialCase_genType_ToTitle_paramsKey]*MoqSpecialCase_genType_ToTitle_results
}

// MoqSpecialCase_genType_ToTitle_doFn defines the type of function needed when
// calling AndDo for the SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_doFn func(r rune)

// MoqSpecialCase_genType_ToTitle_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_doReturnFn func(r rune) rune

// MoqSpecialCase_genType_ToTitle_results holds the results of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_results struct {
	Params  MoqSpecialCase_genType_ToTitle_params
	Results []struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToTitle_doFn
		DoReturnFn MoqSpecialCase_genType_ToTitle_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSpecialCase_genType_ToTitle_fnRecorder routes recorded function calls to
// the MoqSpecialCase_genType moq
type MoqSpecialCase_genType_ToTitle_fnRecorder struct {
	Params    MoqSpecialCase_genType_ToTitle_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSpecialCase_genType_ToTitle_results
	Moq       *MoqSpecialCase_genType
}

// MoqSpecialCase_genType_ToTitle_anyParams isolates the any params functions
// of the SpecialCase_genType type
type MoqSpecialCase_genType_ToTitle_anyParams struct {
	Recorder *MoqSpecialCase_genType_ToTitle_fnRecorder
}

// MoqSpecialCase_genType_ToLower_params holds the params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_params struct{ Param1 rune }

// MoqSpecialCase_genType_ToLower_paramsKey holds the map key params of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_paramsKey struct {
	Params struct{ Param1 rune }
	Hashes struct{ Param1 hash.Hash }
}

// MoqSpecialCase_genType_ToLower_resultsByParams contains the results for a
// given set of parameters for the SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSpecialCase_genType_ToLower_paramsKey]*MoqSpecialCase_genType_ToLower_results
}

// MoqSpecialCase_genType_ToLower_doFn defines the type of function needed when
// calling AndDo for the SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_doFn func(r rune)

// MoqSpecialCase_genType_ToLower_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_doReturnFn func(r rune) rune

// MoqSpecialCase_genType_ToLower_results holds the results of the
// SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_results struct {
	Params  MoqSpecialCase_genType_ToLower_params
	Results []struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToLower_doFn
		DoReturnFn MoqSpecialCase_genType_ToLower_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSpecialCase_genType_ToLower_fnRecorder routes recorded function calls to
// the MoqSpecialCase_genType moq
type MoqSpecialCase_genType_ToLower_fnRecorder struct {
	Params    MoqSpecialCase_genType_ToLower_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSpecialCase_genType_ToLower_results
	Moq       *MoqSpecialCase_genType
}

// MoqSpecialCase_genType_ToLower_anyParams isolates the any params functions
// of the SpecialCase_genType type
type MoqSpecialCase_genType_ToLower_anyParams struct {
	Recorder *MoqSpecialCase_genType_ToLower_fnRecorder
}

// NewMoqSpecialCase_genType creates a new moq of the SpecialCase_genType type
func NewMoqSpecialCase_genType(scene *moq.Scene, config *moq.Config) *MoqSpecialCase_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSpecialCase_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSpecialCase_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ToUpper struct {
					Param1 moq.ParamIndexing
				}
				ToTitle struct {
					Param1 moq.ParamIndexing
				}
				ToLower struct {
					Param1 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ToUpper struct {
				Param1 moq.ParamIndexing
			}
			ToTitle struct {
				Param1 moq.ParamIndexing
			}
			ToLower struct {
				Param1 moq.ParamIndexing
			}
		}{
			ToUpper: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByValue,
			},
			ToTitle: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByValue,
			},
			ToLower: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the SpecialCase_genType type
func (m *MoqSpecialCase_genType) Mock() *MoqSpecialCase_genType_mock { return m.Moq }

func (m *MoqSpecialCase_genType_mock) ToUpper(param1 rune) (result1 rune) {
	m.Moq.Scene.T.Helper()
	params := MoqSpecialCase_genType_ToUpper_params{
		Param1: param1,
	}
	var results *MoqSpecialCase_genType_ToUpper_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ToUpper {
		paramsKey := m.Moq.ParamsKey_ToUpper(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ToUpper(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ToUpper(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ToUpper(params))
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

func (m *MoqSpecialCase_genType_mock) ToTitle(param1 rune) (result1 rune) {
	m.Moq.Scene.T.Helper()
	params := MoqSpecialCase_genType_ToTitle_params{
		Param1: param1,
	}
	var results *MoqSpecialCase_genType_ToTitle_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ToTitle {
		paramsKey := m.Moq.ParamsKey_ToTitle(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ToTitle(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ToTitle(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ToTitle(params))
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

func (m *MoqSpecialCase_genType_mock) ToLower(param1 rune) (result1 rune) {
	m.Moq.Scene.T.Helper()
	params := MoqSpecialCase_genType_ToLower_params{
		Param1: param1,
	}
	var results *MoqSpecialCase_genType_ToLower_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ToLower {
		paramsKey := m.Moq.ParamsKey_ToLower(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ToLower(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ToLower(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ToLower(params))
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

// OnCall returns the recorder implementation of the SpecialCase_genType type
func (m *MoqSpecialCase_genType) OnCall() *MoqSpecialCase_genType_recorder {
	return &MoqSpecialCase_genType_recorder{
		Moq: m,
	}
}

func (m *MoqSpecialCase_genType_recorder) ToUpper(param1 rune) *MoqSpecialCase_genType_ToUpper_fnRecorder {
	return &MoqSpecialCase_genType_ToUpper_fnRecorder{
		Params: MoqSpecialCase_genType_ToUpper_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) Any() *MoqSpecialCase_genType_ToUpper_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToUpper(r.Params))
		return nil
	}
	return &MoqSpecialCase_genType_ToUpper_anyParams{Recorder: r}
}

func (a *MoqSpecialCase_genType_ToUpper_anyParams) Param1() *MoqSpecialCase_genType_ToUpper_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) Seq() *MoqSpecialCase_genType_ToUpper_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToUpper(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) NoSeq() *MoqSpecialCase_genType_ToUpper_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToUpper(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) ReturnResults(result1 rune) *MoqSpecialCase_genType_ToUpper_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToUpper_doFn
		DoReturnFn MoqSpecialCase_genType_ToUpper_doReturnFn
	}{
		Values: &struct {
			Result1 rune
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) AndDo(fn MoqSpecialCase_genType_ToUpper_doFn) *MoqSpecialCase_genType_ToUpper_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) DoReturnResults(fn MoqSpecialCase_genType_ToUpper_doReturnFn) *MoqSpecialCase_genType_ToUpper_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToUpper_doFn
		DoReturnFn MoqSpecialCase_genType_ToUpper_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSpecialCase_genType_ToUpper_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ToUpper {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSpecialCase_genType_ToUpper_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSpecialCase_genType_ToUpper_paramsKey]*MoqSpecialCase_genType_ToUpper_results{},
		}
		r.Moq.ResultsByParams_ToUpper = append(r.Moq.ResultsByParams_ToUpper, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ToUpper) {
			copy(r.Moq.ResultsByParams_ToUpper[insertAt+1:], r.Moq.ResultsByParams_ToUpper[insertAt:0])
			r.Moq.ResultsByParams_ToUpper[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ToUpper(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSpecialCase_genType_ToUpper_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSpecialCase_genType_ToUpper_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSpecialCase_genType_ToUpper_fnRecorder {
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
					Result1 rune
				}
				Sequence   uint32
				DoFn       MoqSpecialCase_genType_ToUpper_doFn
				DoReturnFn MoqSpecialCase_genType_ToUpper_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSpecialCase_genType) PrettyParams_ToUpper(params MoqSpecialCase_genType_ToUpper_params) string {
	return fmt.Sprintf("ToUpper(%#v)", params.Param1)
}

func (m *MoqSpecialCase_genType) ParamsKey_ToUpper(params MoqSpecialCase_genType_ToUpper_params, anyParams uint64) MoqSpecialCase_genType_ToUpper_paramsKey {
	m.Scene.T.Helper()
	var param1Used rune
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ToUpper.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqSpecialCase_genType_ToUpper_paramsKey{
		Params: struct{ Param1 rune }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

func (m *MoqSpecialCase_genType_recorder) ToTitle(param1 rune) *MoqSpecialCase_genType_ToTitle_fnRecorder {
	return &MoqSpecialCase_genType_ToTitle_fnRecorder{
		Params: MoqSpecialCase_genType_ToTitle_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) Any() *MoqSpecialCase_genType_ToTitle_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToTitle(r.Params))
		return nil
	}
	return &MoqSpecialCase_genType_ToTitle_anyParams{Recorder: r}
}

func (a *MoqSpecialCase_genType_ToTitle_anyParams) Param1() *MoqSpecialCase_genType_ToTitle_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) Seq() *MoqSpecialCase_genType_ToTitle_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToTitle(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) NoSeq() *MoqSpecialCase_genType_ToTitle_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToTitle(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) ReturnResults(result1 rune) *MoqSpecialCase_genType_ToTitle_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToTitle_doFn
		DoReturnFn MoqSpecialCase_genType_ToTitle_doReturnFn
	}{
		Values: &struct {
			Result1 rune
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) AndDo(fn MoqSpecialCase_genType_ToTitle_doFn) *MoqSpecialCase_genType_ToTitle_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) DoReturnResults(fn MoqSpecialCase_genType_ToTitle_doReturnFn) *MoqSpecialCase_genType_ToTitle_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToTitle_doFn
		DoReturnFn MoqSpecialCase_genType_ToTitle_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSpecialCase_genType_ToTitle_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ToTitle {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSpecialCase_genType_ToTitle_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSpecialCase_genType_ToTitle_paramsKey]*MoqSpecialCase_genType_ToTitle_results{},
		}
		r.Moq.ResultsByParams_ToTitle = append(r.Moq.ResultsByParams_ToTitle, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ToTitle) {
			copy(r.Moq.ResultsByParams_ToTitle[insertAt+1:], r.Moq.ResultsByParams_ToTitle[insertAt:0])
			r.Moq.ResultsByParams_ToTitle[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ToTitle(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSpecialCase_genType_ToTitle_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSpecialCase_genType_ToTitle_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSpecialCase_genType_ToTitle_fnRecorder {
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
					Result1 rune
				}
				Sequence   uint32
				DoFn       MoqSpecialCase_genType_ToTitle_doFn
				DoReturnFn MoqSpecialCase_genType_ToTitle_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSpecialCase_genType) PrettyParams_ToTitle(params MoqSpecialCase_genType_ToTitle_params) string {
	return fmt.Sprintf("ToTitle(%#v)", params.Param1)
}

func (m *MoqSpecialCase_genType) ParamsKey_ToTitle(params MoqSpecialCase_genType_ToTitle_params, anyParams uint64) MoqSpecialCase_genType_ToTitle_paramsKey {
	m.Scene.T.Helper()
	var param1Used rune
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ToTitle.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqSpecialCase_genType_ToTitle_paramsKey{
		Params: struct{ Param1 rune }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

func (m *MoqSpecialCase_genType_recorder) ToLower(param1 rune) *MoqSpecialCase_genType_ToLower_fnRecorder {
	return &MoqSpecialCase_genType_ToLower_fnRecorder{
		Params: MoqSpecialCase_genType_ToLower_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) Any() *MoqSpecialCase_genType_ToLower_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToLower(r.Params))
		return nil
	}
	return &MoqSpecialCase_genType_ToLower_anyParams{Recorder: r}
}

func (a *MoqSpecialCase_genType_ToLower_anyParams) Param1() *MoqSpecialCase_genType_ToLower_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) Seq() *MoqSpecialCase_genType_ToLower_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToLower(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) NoSeq() *MoqSpecialCase_genType_ToLower_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ToLower(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) ReturnResults(result1 rune) *MoqSpecialCase_genType_ToLower_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToLower_doFn
		DoReturnFn MoqSpecialCase_genType_ToLower_doReturnFn
	}{
		Values: &struct {
			Result1 rune
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) AndDo(fn MoqSpecialCase_genType_ToLower_doFn) *MoqSpecialCase_genType_ToLower_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) DoReturnResults(fn MoqSpecialCase_genType_ToLower_doReturnFn) *MoqSpecialCase_genType_ToLower_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rune
		}
		Sequence   uint32
		DoFn       MoqSpecialCase_genType_ToLower_doFn
		DoReturnFn MoqSpecialCase_genType_ToLower_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSpecialCase_genType_ToLower_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ToLower {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSpecialCase_genType_ToLower_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSpecialCase_genType_ToLower_paramsKey]*MoqSpecialCase_genType_ToLower_results{},
		}
		r.Moq.ResultsByParams_ToLower = append(r.Moq.ResultsByParams_ToLower, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ToLower) {
			copy(r.Moq.ResultsByParams_ToLower[insertAt+1:], r.Moq.ResultsByParams_ToLower[insertAt:0])
			r.Moq.ResultsByParams_ToLower[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ToLower(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSpecialCase_genType_ToLower_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSpecialCase_genType_ToLower_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSpecialCase_genType_ToLower_fnRecorder {
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
					Result1 rune
				}
				Sequence   uint32
				DoFn       MoqSpecialCase_genType_ToLower_doFn
				DoReturnFn MoqSpecialCase_genType_ToLower_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSpecialCase_genType) PrettyParams_ToLower(params MoqSpecialCase_genType_ToLower_params) string {
	return fmt.Sprintf("ToLower(%#v)", params.Param1)
}

func (m *MoqSpecialCase_genType) ParamsKey_ToLower(params MoqSpecialCase_genType_ToLower_params, anyParams uint64) MoqSpecialCase_genType_ToLower_paramsKey {
	m.Scene.T.Helper()
	var param1Used rune
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ToLower.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqSpecialCase_genType_ToLower_paramsKey{
		Params: struct{ Param1 rune }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSpecialCase_genType) Reset() {
	m.ResultsByParams_ToUpper = nil
	m.ResultsByParams_ToTitle = nil
	m.ResultsByParams_ToLower = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSpecialCase_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ToUpper {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ToUpper(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_ToTitle {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ToTitle(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_ToLower {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ToLower(results.Params))
			}
		}
	}
}