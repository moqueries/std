// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package dwarf

import (
	"debug/dwarf"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that dwarf.CommonType_starGenType is
// mocked completely
var _ CommonType_starGenType = (*MoqCommonType_starGenType_mock)(nil)

// CommonType_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type CommonType_starGenType interface {
	Common() *dwarf.CommonType
	Size() int64
}

// MoqCommonType_starGenType holds the state of a moq of the
// CommonType_starGenType type
type MoqCommonType_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCommonType_starGenType_mock

	ResultsByParams_Common []MoqCommonType_starGenType_Common_resultsByParams
	ResultsByParams_Size   []MoqCommonType_starGenType_Size_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Common struct{}
			Size   struct{}
		}
	}
}

// MoqCommonType_starGenType_mock isolates the mock interface of the
// CommonType_starGenType type
type MoqCommonType_starGenType_mock struct {
	Moq *MoqCommonType_starGenType
}

// MoqCommonType_starGenType_recorder isolates the recorder interface of the
// CommonType_starGenType type
type MoqCommonType_starGenType_recorder struct {
	Moq *MoqCommonType_starGenType
}

// MoqCommonType_starGenType_Common_params holds the params of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Common_params struct{}

// MoqCommonType_starGenType_Common_paramsKey holds the map key params of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Common_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCommonType_starGenType_Common_resultsByParams contains the results for a
// given set of parameters for the CommonType_starGenType type
type MoqCommonType_starGenType_Common_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCommonType_starGenType_Common_paramsKey]*MoqCommonType_starGenType_Common_results
}

// MoqCommonType_starGenType_Common_doFn defines the type of function needed
// when calling AndDo for the CommonType_starGenType type
type MoqCommonType_starGenType_Common_doFn func()

// MoqCommonType_starGenType_Common_doReturnFn defines the type of function
// needed when calling DoReturnResults for the CommonType_starGenType type
type MoqCommonType_starGenType_Common_doReturnFn func() *dwarf.CommonType

// MoqCommonType_starGenType_Common_results holds the results of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Common_results struct {
	Params  MoqCommonType_starGenType_Common_params
	Results []struct {
		Values *struct {
			Result1 *dwarf.CommonType
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Common_doFn
		DoReturnFn MoqCommonType_starGenType_Common_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCommonType_starGenType_Common_fnRecorder routes recorded function calls
// to the MoqCommonType_starGenType moq
type MoqCommonType_starGenType_Common_fnRecorder struct {
	Params    MoqCommonType_starGenType_Common_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCommonType_starGenType_Common_results
	Moq       *MoqCommonType_starGenType
}

// MoqCommonType_starGenType_Common_anyParams isolates the any params functions
// of the CommonType_starGenType type
type MoqCommonType_starGenType_Common_anyParams struct {
	Recorder *MoqCommonType_starGenType_Common_fnRecorder
}

// MoqCommonType_starGenType_Size_params holds the params of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Size_params struct{}

// MoqCommonType_starGenType_Size_paramsKey holds the map key params of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Size_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCommonType_starGenType_Size_resultsByParams contains the results for a
// given set of parameters for the CommonType_starGenType type
type MoqCommonType_starGenType_Size_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCommonType_starGenType_Size_paramsKey]*MoqCommonType_starGenType_Size_results
}

// MoqCommonType_starGenType_Size_doFn defines the type of function needed when
// calling AndDo for the CommonType_starGenType type
type MoqCommonType_starGenType_Size_doFn func()

// MoqCommonType_starGenType_Size_doReturnFn defines the type of function
// needed when calling DoReturnResults for the CommonType_starGenType type
type MoqCommonType_starGenType_Size_doReturnFn func() int64

// MoqCommonType_starGenType_Size_results holds the results of the
// CommonType_starGenType type
type MoqCommonType_starGenType_Size_results struct {
	Params  MoqCommonType_starGenType_Size_params
	Results []struct {
		Values *struct {
			Result1 int64
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Size_doFn
		DoReturnFn MoqCommonType_starGenType_Size_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCommonType_starGenType_Size_fnRecorder routes recorded function calls to
// the MoqCommonType_starGenType moq
type MoqCommonType_starGenType_Size_fnRecorder struct {
	Params    MoqCommonType_starGenType_Size_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCommonType_starGenType_Size_results
	Moq       *MoqCommonType_starGenType
}

// MoqCommonType_starGenType_Size_anyParams isolates the any params functions
// of the CommonType_starGenType type
type MoqCommonType_starGenType_Size_anyParams struct {
	Recorder *MoqCommonType_starGenType_Size_fnRecorder
}

// NewMoqCommonType_starGenType creates a new moq of the CommonType_starGenType
// type
func NewMoqCommonType_starGenType(scene *moq.Scene, config *moq.Config) *MoqCommonType_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCommonType_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCommonType_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Common struct{}
				Size   struct{}
			}
		}{ParameterIndexing: struct {
			Common struct{}
			Size   struct{}
		}{
			Common: struct{}{},
			Size:   struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the CommonType_starGenType type
func (m *MoqCommonType_starGenType) Mock() *MoqCommonType_starGenType_mock { return m.Moq }

func (m *MoqCommonType_starGenType_mock) Common() (result1 *dwarf.CommonType) {
	m.Moq.Scene.T.Helper()
	params := MoqCommonType_starGenType_Common_params{}
	var results *MoqCommonType_starGenType_Common_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Common {
		paramsKey := m.Moq.ParamsKey_Common(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Common(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Common(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Common(params))
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

func (m *MoqCommonType_starGenType_mock) Size() (result1 int64) {
	m.Moq.Scene.T.Helper()
	params := MoqCommonType_starGenType_Size_params{}
	var results *MoqCommonType_starGenType_Size_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Size {
		paramsKey := m.Moq.ParamsKey_Size(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Size(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Size(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Size(params))
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

// OnCall returns the recorder implementation of the CommonType_starGenType
// type
func (m *MoqCommonType_starGenType) OnCall() *MoqCommonType_starGenType_recorder {
	return &MoqCommonType_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqCommonType_starGenType_recorder) Common() *MoqCommonType_starGenType_Common_fnRecorder {
	return &MoqCommonType_starGenType_Common_fnRecorder{
		Params:   MoqCommonType_starGenType_Common_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) Any() *MoqCommonType_starGenType_Common_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Common(r.Params))
		return nil
	}
	return &MoqCommonType_starGenType_Common_anyParams{Recorder: r}
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) Seq() *MoqCommonType_starGenType_Common_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Common(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) NoSeq() *MoqCommonType_starGenType_Common_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Common(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) ReturnResults(result1 *dwarf.CommonType) *MoqCommonType_starGenType_Common_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *dwarf.CommonType
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Common_doFn
		DoReturnFn MoqCommonType_starGenType_Common_doReturnFn
	}{
		Values: &struct {
			Result1 *dwarf.CommonType
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) AndDo(fn MoqCommonType_starGenType_Common_doFn) *MoqCommonType_starGenType_Common_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) DoReturnResults(fn MoqCommonType_starGenType_Common_doReturnFn) *MoqCommonType_starGenType_Common_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *dwarf.CommonType
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Common_doFn
		DoReturnFn MoqCommonType_starGenType_Common_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCommonType_starGenType_Common_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Common {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqCommonType_starGenType_Common_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCommonType_starGenType_Common_paramsKey]*MoqCommonType_starGenType_Common_results{},
		}
		r.Moq.ResultsByParams_Common = append(r.Moq.ResultsByParams_Common, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Common) {
			copy(r.Moq.ResultsByParams_Common[insertAt+1:], r.Moq.ResultsByParams_Common[insertAt:0])
			r.Moq.ResultsByParams_Common[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Common(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqCommonType_starGenType_Common_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCommonType_starGenType_Common_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCommonType_starGenType_Common_fnRecorder {
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
					Result1 *dwarf.CommonType
				}
				Sequence   uint32
				DoFn       MoqCommonType_starGenType_Common_doFn
				DoReturnFn MoqCommonType_starGenType_Common_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCommonType_starGenType) PrettyParams_Common(params MoqCommonType_starGenType_Common_params) string {
	return fmt.Sprintf("Common()")
}

func (m *MoqCommonType_starGenType) ParamsKey_Common(params MoqCommonType_starGenType_Common_params, anyParams uint64) MoqCommonType_starGenType_Common_paramsKey {
	m.Scene.T.Helper()
	return MoqCommonType_starGenType_Common_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqCommonType_starGenType_recorder) Size() *MoqCommonType_starGenType_Size_fnRecorder {
	return &MoqCommonType_starGenType_Size_fnRecorder{
		Params:   MoqCommonType_starGenType_Size_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) Any() *MoqCommonType_starGenType_Size_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	return &MoqCommonType_starGenType_Size_anyParams{Recorder: r}
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) Seq() *MoqCommonType_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) NoSeq() *MoqCommonType_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) ReturnResults(result1 int64) *MoqCommonType_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Size_doFn
		DoReturnFn MoqCommonType_starGenType_Size_doReturnFn
	}{
		Values: &struct {
			Result1 int64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) AndDo(fn MoqCommonType_starGenType_Size_doFn) *MoqCommonType_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) DoReturnResults(fn MoqCommonType_starGenType_Size_doReturnFn) *MoqCommonType_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
		}
		Sequence   uint32
		DoFn       MoqCommonType_starGenType_Size_doFn
		DoReturnFn MoqCommonType_starGenType_Size_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCommonType_starGenType_Size_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Size {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqCommonType_starGenType_Size_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCommonType_starGenType_Size_paramsKey]*MoqCommonType_starGenType_Size_results{},
		}
		r.Moq.ResultsByParams_Size = append(r.Moq.ResultsByParams_Size, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Size) {
			copy(r.Moq.ResultsByParams_Size[insertAt+1:], r.Moq.ResultsByParams_Size[insertAt:0])
			r.Moq.ResultsByParams_Size[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Size(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqCommonType_starGenType_Size_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCommonType_starGenType_Size_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCommonType_starGenType_Size_fnRecorder {
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
					Result1 int64
				}
				Sequence   uint32
				DoFn       MoqCommonType_starGenType_Size_doFn
				DoReturnFn MoqCommonType_starGenType_Size_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCommonType_starGenType) PrettyParams_Size(params MoqCommonType_starGenType_Size_params) string {
	return fmt.Sprintf("Size()")
}

func (m *MoqCommonType_starGenType) ParamsKey_Size(params MoqCommonType_starGenType_Size_params, anyParams uint64) MoqCommonType_starGenType_Size_paramsKey {
	m.Scene.T.Helper()
	return MoqCommonType_starGenType_Size_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqCommonType_starGenType) Reset() {
	m.ResultsByParams_Common = nil
	m.ResultsByParams_Size = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCommonType_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Common {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Common(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Size {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Size(results.Params))
			}
		}
	}
}
