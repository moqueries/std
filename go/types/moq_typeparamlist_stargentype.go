// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that types.TypeParamList_starGenType is
// mocked completely
var _ TypeParamList_starGenType = (*MoqTypeParamList_starGenType_mock)(nil)

// TypeParamList_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type TypeParamList_starGenType interface {
	Len() int
	At(i int) *types.TypeParam
}

// MoqTypeParamList_starGenType holds the state of a moq of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTypeParamList_starGenType_mock

	ResultsByParams_Len []MoqTypeParamList_starGenType_Len_resultsByParams
	ResultsByParams_At  []MoqTypeParamList_starGenType_At_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Len struct{}
			At  struct {
				Param1 moq.ParamIndexing
			}
		}
	}
	// MoqTypeParamList_starGenType_mock isolates the mock interface of the
}

// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_mock struct {
	Moq *MoqTypeParamList_starGenType
}

// MoqTypeParamList_starGenType_recorder isolates the recorder interface of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_recorder struct {
	Moq *MoqTypeParamList_starGenType
}

// MoqTypeParamList_starGenType_Len_params holds the params of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_params struct{}

// MoqTypeParamList_starGenType_Len_paramsKey holds the map key params of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqTypeParamList_starGenType_Len_resultsByParams contains the results for a
// given set of parameters for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTypeParamList_starGenType_Len_paramsKey]*MoqTypeParamList_starGenType_Len_results
}

// MoqTypeParamList_starGenType_Len_doFn defines the type of function needed
// when calling AndDo for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_doFn func()

// MoqTypeParamList_starGenType_Len_doReturnFn defines the type of function
// needed when calling DoReturnResults for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_doReturnFn func() int

// MoqTypeParamList_starGenType_Len_results holds the results of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_results struct {
	Params  MoqTypeParamList_starGenType_Len_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqTypeParamList_starGenType_Len_doFn
		DoReturnFn MoqTypeParamList_starGenType_Len_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTypeParamList_starGenType_Len_fnRecorder routes recorded function calls
// to the MoqTypeParamList_starGenType moq
type MoqTypeParamList_starGenType_Len_fnRecorder struct {
	Params    MoqTypeParamList_starGenType_Len_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTypeParamList_starGenType_Len_results
	Moq       *MoqTypeParamList_starGenType
}

// MoqTypeParamList_starGenType_Len_anyParams isolates the any params functions
// of the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_Len_anyParams struct {
	Recorder *MoqTypeParamList_starGenType_Len_fnRecorder
}

// MoqTypeParamList_starGenType_At_params holds the params of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_params struct{ Param1 int }

// MoqTypeParamList_starGenType_At_paramsKey holds the map key params of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_paramsKey struct {
	Params struct{ Param1 int }
	Hashes struct{ Param1 hash.Hash }
}

// MoqTypeParamList_starGenType_At_resultsByParams contains the results for a
// given set of parameters for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTypeParamList_starGenType_At_paramsKey]*MoqTypeParamList_starGenType_At_results
}

// MoqTypeParamList_starGenType_At_doFn defines the type of function needed
// when calling AndDo for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_doFn func(i int)

// MoqTypeParamList_starGenType_At_doReturnFn defines the type of function
// needed when calling DoReturnResults for the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_doReturnFn func(i int) *types.TypeParam

// MoqTypeParamList_starGenType_At_results holds the results of the
// TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_results struct {
	Params  MoqTypeParamList_starGenType_At_params
	Results []struct {
		Values *struct {
			Result1 *types.TypeParam
		}
		Sequence   uint32
		DoFn       MoqTypeParamList_starGenType_At_doFn
		DoReturnFn MoqTypeParamList_starGenType_At_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTypeParamList_starGenType_At_fnRecorder routes recorded function calls to
// the MoqTypeParamList_starGenType moq
type MoqTypeParamList_starGenType_At_fnRecorder struct {
	Params    MoqTypeParamList_starGenType_At_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTypeParamList_starGenType_At_results
	Moq       *MoqTypeParamList_starGenType
}

// MoqTypeParamList_starGenType_At_anyParams isolates the any params functions
// of the TypeParamList_starGenType type
type MoqTypeParamList_starGenType_At_anyParams struct {
	Recorder *MoqTypeParamList_starGenType_At_fnRecorder
}

// NewMoqTypeParamList_starGenType creates a new moq of the
// TypeParamList_starGenType type
func NewMoqTypeParamList_starGenType(scene *moq.Scene, config *moq.Config) *MoqTypeParamList_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTypeParamList_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTypeParamList_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Len struct{}
				At  struct {
					Param1 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Len struct{}
			At  struct {
				Param1 moq.ParamIndexing
			}
		}{
			Len: struct{}{},
			At: struct {
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

// Mock returns the mock implementation of the TypeParamList_starGenType type
func (m *MoqTypeParamList_starGenType) Mock() *MoqTypeParamList_starGenType_mock { return m.Moq }

func (m *MoqTypeParamList_starGenType_mock) Len() (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqTypeParamList_starGenType_Len_params{}
	var results *MoqTypeParamList_starGenType_Len_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Len {
		paramsKey := m.Moq.ParamsKey_Len(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Len(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Len(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Len(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqTypeParamList_starGenType_mock) At(param1 int) (result1 *types.TypeParam) {
	m.Moq.Scene.T.Helper()
	params := MoqTypeParamList_starGenType_At_params{
		Param1: param1,
	}
	var results *MoqTypeParamList_starGenType_At_results
	for _, resultsByParams := range m.Moq.ResultsByParams_At {
		paramsKey := m.Moq.ParamsKey_At(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_At(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_At(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_At(params))
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

// OnCall returns the recorder implementation of the TypeParamList_starGenType
// type
func (m *MoqTypeParamList_starGenType) OnCall() *MoqTypeParamList_starGenType_recorder {
	return &MoqTypeParamList_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqTypeParamList_starGenType_recorder) Len() *MoqTypeParamList_starGenType_Len_fnRecorder {
	return &MoqTypeParamList_starGenType_Len_fnRecorder{
		Params:   MoqTypeParamList_starGenType_Len_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) Any() *MoqTypeParamList_starGenType_Len_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	return &MoqTypeParamList_starGenType_Len_anyParams{Recorder: r}
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) Seq() *MoqTypeParamList_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) NoSeq() *MoqTypeParamList_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) ReturnResults(result1 int) *MoqTypeParamList_starGenType_Len_fnRecorder {
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
		DoFn       MoqTypeParamList_starGenType_Len_doFn
		DoReturnFn MoqTypeParamList_starGenType_Len_doReturnFn
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

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) AndDo(fn MoqTypeParamList_starGenType_Len_doFn) *MoqTypeParamList_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) DoReturnResults(fn MoqTypeParamList_starGenType_Len_doReturnFn) *MoqTypeParamList_starGenType_Len_fnRecorder {
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
		DoFn       MoqTypeParamList_starGenType_Len_doFn
		DoReturnFn MoqTypeParamList_starGenType_Len_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTypeParamList_starGenType_Len_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Len {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqTypeParamList_starGenType_Len_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTypeParamList_starGenType_Len_paramsKey]*MoqTypeParamList_starGenType_Len_results{},
		}
		r.Moq.ResultsByParams_Len = append(r.Moq.ResultsByParams_Len, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Len) {
			copy(r.Moq.ResultsByParams_Len[insertAt+1:], r.Moq.ResultsByParams_Len[insertAt:0])
			r.Moq.ResultsByParams_Len[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Len(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqTypeParamList_starGenType_Len_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTypeParamList_starGenType_Len_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTypeParamList_starGenType_Len_fnRecorder {
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
				DoFn       MoqTypeParamList_starGenType_Len_doFn
				DoReturnFn MoqTypeParamList_starGenType_Len_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTypeParamList_starGenType) PrettyParams_Len(params MoqTypeParamList_starGenType_Len_params) string {
	return fmt.Sprintf("Len()")
}

func (m *MoqTypeParamList_starGenType) ParamsKey_Len(params MoqTypeParamList_starGenType_Len_params, anyParams uint64) MoqTypeParamList_starGenType_Len_paramsKey {
	m.Scene.T.Helper()
	return MoqTypeParamList_starGenType_Len_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqTypeParamList_starGenType_recorder) At(param1 int) *MoqTypeParamList_starGenType_At_fnRecorder {
	return &MoqTypeParamList_starGenType_At_fnRecorder{
		Params: MoqTypeParamList_starGenType_At_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) Any() *MoqTypeParamList_starGenType_At_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	return &MoqTypeParamList_starGenType_At_anyParams{Recorder: r}
}

func (a *MoqTypeParamList_starGenType_At_anyParams) Param1() *MoqTypeParamList_starGenType_At_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) Seq() *MoqTypeParamList_starGenType_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) NoSeq() *MoqTypeParamList_starGenType_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) ReturnResults(result1 *types.TypeParam) *MoqTypeParamList_starGenType_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.TypeParam
		}
		Sequence   uint32
		DoFn       MoqTypeParamList_starGenType_At_doFn
		DoReturnFn MoqTypeParamList_starGenType_At_doReturnFn
	}{
		Values: &struct {
			Result1 *types.TypeParam
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) AndDo(fn MoqTypeParamList_starGenType_At_doFn) *MoqTypeParamList_starGenType_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) DoReturnResults(fn MoqTypeParamList_starGenType_At_doReturnFn) *MoqTypeParamList_starGenType_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.TypeParam
		}
		Sequence   uint32
		DoFn       MoqTypeParamList_starGenType_At_doFn
		DoReturnFn MoqTypeParamList_starGenType_At_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTypeParamList_starGenType_At_resultsByParams
	for n, res := range r.Moq.ResultsByParams_At {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqTypeParamList_starGenType_At_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTypeParamList_starGenType_At_paramsKey]*MoqTypeParamList_starGenType_At_results{},
		}
		r.Moq.ResultsByParams_At = append(r.Moq.ResultsByParams_At, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_At) {
			copy(r.Moq.ResultsByParams_At[insertAt+1:], r.Moq.ResultsByParams_At[insertAt:0])
			r.Moq.ResultsByParams_At[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_At(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqTypeParamList_starGenType_At_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTypeParamList_starGenType_At_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTypeParamList_starGenType_At_fnRecorder {
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
					Result1 *types.TypeParam
				}
				Sequence   uint32
				DoFn       MoqTypeParamList_starGenType_At_doFn
				DoReturnFn MoqTypeParamList_starGenType_At_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTypeParamList_starGenType) PrettyParams_At(params MoqTypeParamList_starGenType_At_params) string {
	return fmt.Sprintf("At(%#v)", params.Param1)
}

func (m *MoqTypeParamList_starGenType) ParamsKey_At(params MoqTypeParamList_starGenType_At_params, anyParams uint64) MoqTypeParamList_starGenType_At_paramsKey {
	m.Scene.T.Helper()
	var param1Used int
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.At.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqTypeParamList_starGenType_At_paramsKey{
		Params: struct{ Param1 int }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqTypeParamList_starGenType) Reset() {
	m.ResultsByParams_Len = nil
	m.ResultsByParams_At = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTypeParamList_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Len {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Len(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_At {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_At(results.Params))
			}
		}
	}
}
