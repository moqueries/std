// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package xml

import (
	"encoding/xml"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that xml.StartElement_genType is mocked
// completely
var _ StartElement_genType = (*MoqStartElement_genType_mock)(nil)

// StartElement_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type StartElement_genType interface {
	Copy() xml.StartElement
	End() xml.EndElement
}

// MoqStartElement_genType holds the state of a moq of the StartElement_genType
// type
type MoqStartElement_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStartElement_genType_mock

	ResultsByParams_Copy []MoqStartElement_genType_Copy_resultsByParams
	ResultsByParams_End  []MoqStartElement_genType_End_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Copy struct{}
			End  struct{}
		}
	}
}

// MoqStartElement_genType_mock isolates the mock interface of the
// StartElement_genType type
type MoqStartElement_genType_mock struct {
	Moq *MoqStartElement_genType
}

// MoqStartElement_genType_recorder isolates the recorder interface of the
// StartElement_genType type
type MoqStartElement_genType_recorder struct {
	Moq *MoqStartElement_genType
}

// MoqStartElement_genType_Copy_params holds the params of the
// StartElement_genType type
type MoqStartElement_genType_Copy_params struct{}

// MoqStartElement_genType_Copy_paramsKey holds the map key params of the
// StartElement_genType type
type MoqStartElement_genType_Copy_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStartElement_genType_Copy_resultsByParams contains the results for a
// given set of parameters for the StartElement_genType type
type MoqStartElement_genType_Copy_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStartElement_genType_Copy_paramsKey]*MoqStartElement_genType_Copy_results
}

// MoqStartElement_genType_Copy_doFn defines the type of function needed when
// calling AndDo for the StartElement_genType type
type MoqStartElement_genType_Copy_doFn func()

// MoqStartElement_genType_Copy_doReturnFn defines the type of function needed
// when calling DoReturnResults for the StartElement_genType type
type MoqStartElement_genType_Copy_doReturnFn func() xml.StartElement

// MoqStartElement_genType_Copy_results holds the results of the
// StartElement_genType type
type MoqStartElement_genType_Copy_results struct {
	Params  MoqStartElement_genType_Copy_params
	Results []struct {
		Values *struct {
			Result1 xml.StartElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_Copy_doFn
		DoReturnFn MoqStartElement_genType_Copy_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStartElement_genType_Copy_fnRecorder routes recorded function calls to
// the MoqStartElement_genType moq
type MoqStartElement_genType_Copy_fnRecorder struct {
	Params    MoqStartElement_genType_Copy_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStartElement_genType_Copy_results
	Moq       *MoqStartElement_genType
}

// MoqStartElement_genType_Copy_anyParams isolates the any params functions of
// the StartElement_genType type
type MoqStartElement_genType_Copy_anyParams struct {
	Recorder *MoqStartElement_genType_Copy_fnRecorder
}

// MoqStartElement_genType_End_params holds the params of the
// StartElement_genType type
type MoqStartElement_genType_End_params struct{}

// MoqStartElement_genType_End_paramsKey holds the map key params of the
// StartElement_genType type
type MoqStartElement_genType_End_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStartElement_genType_End_resultsByParams contains the results for a given
// set of parameters for the StartElement_genType type
type MoqStartElement_genType_End_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStartElement_genType_End_paramsKey]*MoqStartElement_genType_End_results
}

// MoqStartElement_genType_End_doFn defines the type of function needed when
// calling AndDo for the StartElement_genType type
type MoqStartElement_genType_End_doFn func()

// MoqStartElement_genType_End_doReturnFn defines the type of function needed
// when calling DoReturnResults for the StartElement_genType type
type MoqStartElement_genType_End_doReturnFn func() xml.EndElement

// MoqStartElement_genType_End_results holds the results of the
// StartElement_genType type
type MoqStartElement_genType_End_results struct {
	Params  MoqStartElement_genType_End_params
	Results []struct {
		Values *struct {
			Result1 xml.EndElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_End_doFn
		DoReturnFn MoqStartElement_genType_End_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStartElement_genType_End_fnRecorder routes recorded function calls to the
// MoqStartElement_genType moq
type MoqStartElement_genType_End_fnRecorder struct {
	Params    MoqStartElement_genType_End_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStartElement_genType_End_results
	Moq       *MoqStartElement_genType
}

// MoqStartElement_genType_End_anyParams isolates the any params functions of
// the StartElement_genType type
type MoqStartElement_genType_End_anyParams struct {
	Recorder *MoqStartElement_genType_End_fnRecorder
}

// NewMoqStartElement_genType creates a new moq of the StartElement_genType
// type
func NewMoqStartElement_genType(scene *moq.Scene, config *moq.Config) *MoqStartElement_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStartElement_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStartElement_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Copy struct{}
				End  struct{}
			}
		}{ParameterIndexing: struct {
			Copy struct{}
			End  struct{}
		}{
			Copy: struct{}{},
			End:  struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the StartElement_genType type
func (m *MoqStartElement_genType) Mock() *MoqStartElement_genType_mock { return m.Moq }

func (m *MoqStartElement_genType_mock) Copy() (result1 xml.StartElement) {
	m.Moq.Scene.T.Helper()
	params := MoqStartElement_genType_Copy_params{}
	var results *MoqStartElement_genType_Copy_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Copy {
		paramsKey := m.Moq.ParamsKey_Copy(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Copy(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Copy(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Copy(params))
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

func (m *MoqStartElement_genType_mock) End() (result1 xml.EndElement) {
	m.Moq.Scene.T.Helper()
	params := MoqStartElement_genType_End_params{}
	var results *MoqStartElement_genType_End_results
	for _, resultsByParams := range m.Moq.ResultsByParams_End {
		paramsKey := m.Moq.ParamsKey_End(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_End(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_End(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_End(params))
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

// OnCall returns the recorder implementation of the StartElement_genType type
func (m *MoqStartElement_genType) OnCall() *MoqStartElement_genType_recorder {
	return &MoqStartElement_genType_recorder{
		Moq: m,
	}
}

func (m *MoqStartElement_genType_recorder) Copy() *MoqStartElement_genType_Copy_fnRecorder {
	return &MoqStartElement_genType_Copy_fnRecorder{
		Params:   MoqStartElement_genType_Copy_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStartElement_genType_Copy_fnRecorder) Any() *MoqStartElement_genType_Copy_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	return &MoqStartElement_genType_Copy_anyParams{Recorder: r}
}

func (r *MoqStartElement_genType_Copy_fnRecorder) Seq() *MoqStartElement_genType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStartElement_genType_Copy_fnRecorder) NoSeq() *MoqStartElement_genType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStartElement_genType_Copy_fnRecorder) ReturnResults(result1 xml.StartElement) *MoqStartElement_genType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.StartElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_Copy_doFn
		DoReturnFn MoqStartElement_genType_Copy_doReturnFn
	}{
		Values: &struct {
			Result1 xml.StartElement
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStartElement_genType_Copy_fnRecorder) AndDo(fn MoqStartElement_genType_Copy_doFn) *MoqStartElement_genType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStartElement_genType_Copy_fnRecorder) DoReturnResults(fn MoqStartElement_genType_Copy_doReturnFn) *MoqStartElement_genType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.StartElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_Copy_doFn
		DoReturnFn MoqStartElement_genType_Copy_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStartElement_genType_Copy_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStartElement_genType_Copy_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Copy {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStartElement_genType_Copy_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStartElement_genType_Copy_paramsKey]*MoqStartElement_genType_Copy_results{},
		}
		r.Moq.ResultsByParams_Copy = append(r.Moq.ResultsByParams_Copy, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Copy) {
			copy(r.Moq.ResultsByParams_Copy[insertAt+1:], r.Moq.ResultsByParams_Copy[insertAt:0])
			r.Moq.ResultsByParams_Copy[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Copy(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStartElement_genType_Copy_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStartElement_genType_Copy_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStartElement_genType_Copy_fnRecorder {
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
					Result1 xml.StartElement
				}
				Sequence   uint32
				DoFn       MoqStartElement_genType_Copy_doFn
				DoReturnFn MoqStartElement_genType_Copy_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStartElement_genType) PrettyParams_Copy(params MoqStartElement_genType_Copy_params) string {
	return fmt.Sprintf("Copy()")
}

func (m *MoqStartElement_genType) ParamsKey_Copy(params MoqStartElement_genType_Copy_params, anyParams uint64) MoqStartElement_genType_Copy_paramsKey {
	m.Scene.T.Helper()
	return MoqStartElement_genType_Copy_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqStartElement_genType_recorder) End() *MoqStartElement_genType_End_fnRecorder {
	return &MoqStartElement_genType_End_fnRecorder{
		Params:   MoqStartElement_genType_End_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStartElement_genType_End_fnRecorder) Any() *MoqStartElement_genType_End_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	return &MoqStartElement_genType_End_anyParams{Recorder: r}
}

func (r *MoqStartElement_genType_End_fnRecorder) Seq() *MoqStartElement_genType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStartElement_genType_End_fnRecorder) NoSeq() *MoqStartElement_genType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStartElement_genType_End_fnRecorder) ReturnResults(result1 xml.EndElement) *MoqStartElement_genType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.EndElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_End_doFn
		DoReturnFn MoqStartElement_genType_End_doReturnFn
	}{
		Values: &struct {
			Result1 xml.EndElement
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStartElement_genType_End_fnRecorder) AndDo(fn MoqStartElement_genType_End_doFn) *MoqStartElement_genType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStartElement_genType_End_fnRecorder) DoReturnResults(fn MoqStartElement_genType_End_doReturnFn) *MoqStartElement_genType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.EndElement
		}
		Sequence   uint32
		DoFn       MoqStartElement_genType_End_doFn
		DoReturnFn MoqStartElement_genType_End_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStartElement_genType_End_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStartElement_genType_End_resultsByParams
	for n, res := range r.Moq.ResultsByParams_End {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStartElement_genType_End_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStartElement_genType_End_paramsKey]*MoqStartElement_genType_End_results{},
		}
		r.Moq.ResultsByParams_End = append(r.Moq.ResultsByParams_End, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_End) {
			copy(r.Moq.ResultsByParams_End[insertAt+1:], r.Moq.ResultsByParams_End[insertAt:0])
			r.Moq.ResultsByParams_End[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_End(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStartElement_genType_End_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStartElement_genType_End_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStartElement_genType_End_fnRecorder {
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
					Result1 xml.EndElement
				}
				Sequence   uint32
				DoFn       MoqStartElement_genType_End_doFn
				DoReturnFn MoqStartElement_genType_End_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStartElement_genType) PrettyParams_End(params MoqStartElement_genType_End_params) string {
	return fmt.Sprintf("End()")
}

func (m *MoqStartElement_genType) ParamsKey_End(params MoqStartElement_genType_End_params, anyParams uint64) MoqStartElement_genType_End_paramsKey {
	m.Scene.T.Helper()
	return MoqStartElement_genType_End_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqStartElement_genType) Reset() { m.ResultsByParams_Copy = nil; m.ResultsByParams_End = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStartElement_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Copy {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Copy(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_End {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_End(results.Params))
			}
		}
	}
}
