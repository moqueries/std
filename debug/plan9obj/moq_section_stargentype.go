// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package plan9obj

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that plan9obj.Section_starGenType is
// mocked completely
var _ Section_starGenType = (*MoqSection_starGenType_mock)(nil)

// Section_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Section_starGenType interface {
	Data() ([]byte, error)
	Open() io.ReadSeeker
}

// MoqSection_starGenType holds the state of a moq of the Section_starGenType
// type
type MoqSection_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSection_starGenType_mock

	ResultsByParams_Data []MoqSection_starGenType_Data_resultsByParams
	ResultsByParams_Open []MoqSection_starGenType_Open_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Data struct{}
			Open struct{}
		}
	}
}

// MoqSection_starGenType_mock isolates the mock interface of the
// Section_starGenType type
type MoqSection_starGenType_mock struct {
	Moq *MoqSection_starGenType
}

// MoqSection_starGenType_recorder isolates the recorder interface of the
// Section_starGenType type
type MoqSection_starGenType_recorder struct {
	Moq *MoqSection_starGenType
}

// MoqSection_starGenType_Data_params holds the params of the
// Section_starGenType type
type MoqSection_starGenType_Data_params struct{}

// MoqSection_starGenType_Data_paramsKey holds the map key params of the
// Section_starGenType type
type MoqSection_starGenType_Data_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSection_starGenType_Data_resultsByParams contains the results for a given
// set of parameters for the Section_starGenType type
type MoqSection_starGenType_Data_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSection_starGenType_Data_paramsKey]*MoqSection_starGenType_Data_results
}

// MoqSection_starGenType_Data_doFn defines the type of function needed when
// calling AndDo for the Section_starGenType type
type MoqSection_starGenType_Data_doFn func()

// MoqSection_starGenType_Data_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Section_starGenType type
type MoqSection_starGenType_Data_doReturnFn func() ([]byte, error)

// MoqSection_starGenType_Data_results holds the results of the
// Section_starGenType type
type MoqSection_starGenType_Data_results struct {
	Params  MoqSection_starGenType_Data_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Data_doFn
		DoReturnFn MoqSection_starGenType_Data_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSection_starGenType_Data_fnRecorder routes recorded function calls to the
// MoqSection_starGenType moq
type MoqSection_starGenType_Data_fnRecorder struct {
	Params    MoqSection_starGenType_Data_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSection_starGenType_Data_results
	Moq       *MoqSection_starGenType
}

// MoqSection_starGenType_Data_anyParams isolates the any params functions of
// the Section_starGenType type
type MoqSection_starGenType_Data_anyParams struct {
	Recorder *MoqSection_starGenType_Data_fnRecorder
}

// MoqSection_starGenType_Open_params holds the params of the
// Section_starGenType type
type MoqSection_starGenType_Open_params struct{}

// MoqSection_starGenType_Open_paramsKey holds the map key params of the
// Section_starGenType type
type MoqSection_starGenType_Open_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSection_starGenType_Open_resultsByParams contains the results for a given
// set of parameters for the Section_starGenType type
type MoqSection_starGenType_Open_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSection_starGenType_Open_paramsKey]*MoqSection_starGenType_Open_results
}

// MoqSection_starGenType_Open_doFn defines the type of function needed when
// calling AndDo for the Section_starGenType type
type MoqSection_starGenType_Open_doFn func()

// MoqSection_starGenType_Open_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Section_starGenType type
type MoqSection_starGenType_Open_doReturnFn func() io.ReadSeeker

// MoqSection_starGenType_Open_results holds the results of the
// Section_starGenType type
type MoqSection_starGenType_Open_results struct {
	Params  MoqSection_starGenType_Open_params
	Results []struct {
		Values *struct {
			Result1 io.ReadSeeker
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Open_doFn
		DoReturnFn MoqSection_starGenType_Open_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSection_starGenType_Open_fnRecorder routes recorded function calls to the
// MoqSection_starGenType moq
type MoqSection_starGenType_Open_fnRecorder struct {
	Params    MoqSection_starGenType_Open_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSection_starGenType_Open_results
	Moq       *MoqSection_starGenType
}

// MoqSection_starGenType_Open_anyParams isolates the any params functions of
// the Section_starGenType type
type MoqSection_starGenType_Open_anyParams struct {
	Recorder *MoqSection_starGenType_Open_fnRecorder
}

// NewMoqSection_starGenType creates a new moq of the Section_starGenType type
func NewMoqSection_starGenType(scene *moq.Scene, config *moq.Config) *MoqSection_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSection_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSection_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Data struct{}
				Open struct{}
			}
		}{ParameterIndexing: struct {
			Data struct{}
			Open struct{}
		}{
			Data: struct{}{},
			Open: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Section_starGenType type
func (m *MoqSection_starGenType) Mock() *MoqSection_starGenType_mock { return m.Moq }

func (m *MoqSection_starGenType_mock) Data() (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSection_starGenType_Data_params{}
	var results *MoqSection_starGenType_Data_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Data {
		paramsKey := m.Moq.ParamsKey_Data(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Data(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Data(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Data(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

func (m *MoqSection_starGenType_mock) Open() (result1 io.ReadSeeker) {
	m.Moq.Scene.T.Helper()
	params := MoqSection_starGenType_Open_params{}
	var results *MoqSection_starGenType_Open_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Open {
		paramsKey := m.Moq.ParamsKey_Open(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Open(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Open(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Open(params))
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

// OnCall returns the recorder implementation of the Section_starGenType type
func (m *MoqSection_starGenType) OnCall() *MoqSection_starGenType_recorder {
	return &MoqSection_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqSection_starGenType_recorder) Data() *MoqSection_starGenType_Data_fnRecorder {
	return &MoqSection_starGenType_Data_fnRecorder{
		Params:   MoqSection_starGenType_Data_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSection_starGenType_Data_fnRecorder) Any() *MoqSection_starGenType_Data_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Data(r.Params))
		return nil
	}
	return &MoqSection_starGenType_Data_anyParams{Recorder: r}
}

func (r *MoqSection_starGenType_Data_fnRecorder) Seq() *MoqSection_starGenType_Data_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Data(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSection_starGenType_Data_fnRecorder) NoSeq() *MoqSection_starGenType_Data_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Data(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSection_starGenType_Data_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqSection_starGenType_Data_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Data_doFn
		DoReturnFn MoqSection_starGenType_Data_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSection_starGenType_Data_fnRecorder) AndDo(fn MoqSection_starGenType_Data_doFn) *MoqSection_starGenType_Data_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSection_starGenType_Data_fnRecorder) DoReturnResults(fn MoqSection_starGenType_Data_doReturnFn) *MoqSection_starGenType_Data_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Data_doFn
		DoReturnFn MoqSection_starGenType_Data_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSection_starGenType_Data_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSection_starGenType_Data_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Data {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSection_starGenType_Data_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSection_starGenType_Data_paramsKey]*MoqSection_starGenType_Data_results{},
		}
		r.Moq.ResultsByParams_Data = append(r.Moq.ResultsByParams_Data, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Data) {
			copy(r.Moq.ResultsByParams_Data[insertAt+1:], r.Moq.ResultsByParams_Data[insertAt:0])
			r.Moq.ResultsByParams_Data[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Data(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSection_starGenType_Data_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSection_starGenType_Data_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSection_starGenType_Data_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqSection_starGenType_Data_doFn
				DoReturnFn MoqSection_starGenType_Data_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSection_starGenType) PrettyParams_Data(params MoqSection_starGenType_Data_params) string {
	return fmt.Sprintf("Data()")
}

func (m *MoqSection_starGenType) ParamsKey_Data(params MoqSection_starGenType_Data_params, anyParams uint64) MoqSection_starGenType_Data_paramsKey {
	m.Scene.T.Helper()
	return MoqSection_starGenType_Data_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqSection_starGenType_recorder) Open() *MoqSection_starGenType_Open_fnRecorder {
	return &MoqSection_starGenType_Open_fnRecorder{
		Params:   MoqSection_starGenType_Open_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSection_starGenType_Open_fnRecorder) Any() *MoqSection_starGenType_Open_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	return &MoqSection_starGenType_Open_anyParams{Recorder: r}
}

func (r *MoqSection_starGenType_Open_fnRecorder) Seq() *MoqSection_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSection_starGenType_Open_fnRecorder) NoSeq() *MoqSection_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSection_starGenType_Open_fnRecorder) ReturnResults(result1 io.ReadSeeker) *MoqSection_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadSeeker
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Open_doFn
		DoReturnFn MoqSection_starGenType_Open_doReturnFn
	}{
		Values: &struct {
			Result1 io.ReadSeeker
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSection_starGenType_Open_fnRecorder) AndDo(fn MoqSection_starGenType_Open_doFn) *MoqSection_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSection_starGenType_Open_fnRecorder) DoReturnResults(fn MoqSection_starGenType_Open_doReturnFn) *MoqSection_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadSeeker
		}
		Sequence   uint32
		DoFn       MoqSection_starGenType_Open_doFn
		DoReturnFn MoqSection_starGenType_Open_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSection_starGenType_Open_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSection_starGenType_Open_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Open {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSection_starGenType_Open_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSection_starGenType_Open_paramsKey]*MoqSection_starGenType_Open_results{},
		}
		r.Moq.ResultsByParams_Open = append(r.Moq.ResultsByParams_Open, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Open) {
			copy(r.Moq.ResultsByParams_Open[insertAt+1:], r.Moq.ResultsByParams_Open[insertAt:0])
			r.Moq.ResultsByParams_Open[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Open(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSection_starGenType_Open_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSection_starGenType_Open_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSection_starGenType_Open_fnRecorder {
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
					Result1 io.ReadSeeker
				}
				Sequence   uint32
				DoFn       MoqSection_starGenType_Open_doFn
				DoReturnFn MoqSection_starGenType_Open_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSection_starGenType) PrettyParams_Open(params MoqSection_starGenType_Open_params) string {
	return fmt.Sprintf("Open()")
}

func (m *MoqSection_starGenType) ParamsKey_Open(params MoqSection_starGenType_Open_params, anyParams uint64) MoqSection_starGenType_Open_paramsKey {
	m.Scene.T.Helper()
	return MoqSection_starGenType_Open_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqSection_starGenType) Reset() { m.ResultsByParams_Data = nil; m.ResultsByParams_Open = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSection_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Data {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Data(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Open {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Open(results.Params))
			}
		}
	}
}
