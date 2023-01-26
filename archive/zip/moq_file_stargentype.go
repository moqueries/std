// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package zip

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that zip.File_starGenType is mocked
// completely
var _ File_starGenType = (*MoqFile_starGenType_mock)(nil)

// File_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type File_starGenType interface {
	DataOffset() (offset int64, err error)
	Open() (io.ReadCloser, error)
}

// MoqFile_starGenType holds the state of a moq of the File_starGenType type
type MoqFile_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFile_starGenType_mock

	ResultsByParams_DataOffset []MoqFile_starGenType_DataOffset_resultsByParams
	ResultsByParams_Open       []MoqFile_starGenType_Open_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			DataOffset struct{}
			Open       struct{}
		}
	}
}

// MoqFile_starGenType_mock isolates the mock interface of the File_starGenType
// type
type MoqFile_starGenType_mock struct {
	Moq *MoqFile_starGenType
}

// MoqFile_starGenType_recorder isolates the recorder interface of the
// File_starGenType type
type MoqFile_starGenType_recorder struct {
	Moq *MoqFile_starGenType
}

// MoqFile_starGenType_DataOffset_params holds the params of the
// File_starGenType type
type MoqFile_starGenType_DataOffset_params struct{}

// MoqFile_starGenType_DataOffset_paramsKey holds the map key params of the
// File_starGenType type
type MoqFile_starGenType_DataOffset_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFile_starGenType_DataOffset_resultsByParams contains the results for a
// given set of parameters for the File_starGenType type
type MoqFile_starGenType_DataOffset_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_starGenType_DataOffset_paramsKey]*MoqFile_starGenType_DataOffset_results
}

// MoqFile_starGenType_DataOffset_doFn defines the type of function needed when
// calling AndDo for the File_starGenType type
type MoqFile_starGenType_DataOffset_doFn func()

// MoqFile_starGenType_DataOffset_doReturnFn defines the type of function
// needed when calling DoReturnResults for the File_starGenType type
type MoqFile_starGenType_DataOffset_doReturnFn func() (offset int64, err error)

// MoqFile_starGenType_DataOffset_results holds the results of the
// File_starGenType type
type MoqFile_starGenType_DataOffset_results struct {
	Params  MoqFile_starGenType_DataOffset_params
	Results []struct {
		Values *struct {
			Offset int64
			Err    error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_DataOffset_doFn
		DoReturnFn MoqFile_starGenType_DataOffset_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_starGenType_DataOffset_fnRecorder routes recorded function calls to
// the MoqFile_starGenType moq
type MoqFile_starGenType_DataOffset_fnRecorder struct {
	Params    MoqFile_starGenType_DataOffset_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_starGenType_DataOffset_results
	Moq       *MoqFile_starGenType
}

// MoqFile_starGenType_DataOffset_anyParams isolates the any params functions
// of the File_starGenType type
type MoqFile_starGenType_DataOffset_anyParams struct {
	Recorder *MoqFile_starGenType_DataOffset_fnRecorder
}

// MoqFile_starGenType_Open_params holds the params of the File_starGenType
// type
type MoqFile_starGenType_Open_params struct{}

// MoqFile_starGenType_Open_paramsKey holds the map key params of the
// File_starGenType type
type MoqFile_starGenType_Open_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFile_starGenType_Open_resultsByParams contains the results for a given
// set of parameters for the File_starGenType type
type MoqFile_starGenType_Open_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_starGenType_Open_paramsKey]*MoqFile_starGenType_Open_results
}

// MoqFile_starGenType_Open_doFn defines the type of function needed when
// calling AndDo for the File_starGenType type
type MoqFile_starGenType_Open_doFn func()

// MoqFile_starGenType_Open_doReturnFn defines the type of function needed when
// calling DoReturnResults for the File_starGenType type
type MoqFile_starGenType_Open_doReturnFn func() (io.ReadCloser, error)

// MoqFile_starGenType_Open_results holds the results of the File_starGenType
// type
type MoqFile_starGenType_Open_results struct {
	Params  MoqFile_starGenType_Open_params
	Results []struct {
		Values *struct {
			Result1 io.ReadCloser
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Open_doFn
		DoReturnFn MoqFile_starGenType_Open_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_starGenType_Open_fnRecorder routes recorded function calls to the
// MoqFile_starGenType moq
type MoqFile_starGenType_Open_fnRecorder struct {
	Params    MoqFile_starGenType_Open_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_starGenType_Open_results
	Moq       *MoqFile_starGenType
}

// MoqFile_starGenType_Open_anyParams isolates the any params functions of the
// File_starGenType type
type MoqFile_starGenType_Open_anyParams struct {
	Recorder *MoqFile_starGenType_Open_fnRecorder
}

// NewMoqFile_starGenType creates a new moq of the File_starGenType type
func NewMoqFile_starGenType(scene *moq.Scene, config *moq.Config) *MoqFile_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFile_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFile_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				DataOffset struct{}
				Open       struct{}
			}
		}{ParameterIndexing: struct {
			DataOffset struct{}
			Open       struct{}
		}{
			DataOffset: struct{}{},
			Open:       struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the File_starGenType type
func (m *MoqFile_starGenType) Mock() *MoqFile_starGenType_mock { return m.Moq }

func (m *MoqFile_starGenType_mock) DataOffset() (offset int64, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_starGenType_DataOffset_params{}
	var results *MoqFile_starGenType_DataOffset_results
	for _, resultsByParams := range m.Moq.ResultsByParams_DataOffset {
		paramsKey := m.Moq.ParamsKey_DataOffset(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_DataOffset(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_DataOffset(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_DataOffset(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		offset = result.Values.Offset
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		offset, err = result.DoReturnFn()
	}
	return
}

func (m *MoqFile_starGenType_mock) Open() (result1 io.ReadCloser, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_starGenType_Open_params{}
	var results *MoqFile_starGenType_Open_results
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
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the File_starGenType type
func (m *MoqFile_starGenType) OnCall() *MoqFile_starGenType_recorder {
	return &MoqFile_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqFile_starGenType_recorder) DataOffset() *MoqFile_starGenType_DataOffset_fnRecorder {
	return &MoqFile_starGenType_DataOffset_fnRecorder{
		Params:   MoqFile_starGenType_DataOffset_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) Any() *MoqFile_starGenType_DataOffset_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DataOffset(r.Params))
		return nil
	}
	return &MoqFile_starGenType_DataOffset_anyParams{Recorder: r}
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) Seq() *MoqFile_starGenType_DataOffset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DataOffset(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) NoSeq() *MoqFile_starGenType_DataOffset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DataOffset(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) ReturnResults(offset int64, err error) *MoqFile_starGenType_DataOffset_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Offset int64
			Err    error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_DataOffset_doFn
		DoReturnFn MoqFile_starGenType_DataOffset_doReturnFn
	}{
		Values: &struct {
			Offset int64
			Err    error
		}{
			Offset: offset,
			Err:    err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) AndDo(fn MoqFile_starGenType_DataOffset_doFn) *MoqFile_starGenType_DataOffset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) DoReturnResults(fn MoqFile_starGenType_DataOffset_doReturnFn) *MoqFile_starGenType_DataOffset_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Offset int64
			Err    error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_DataOffset_doFn
		DoReturnFn MoqFile_starGenType_DataOffset_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_starGenType_DataOffset_resultsByParams
	for n, res := range r.Moq.ResultsByParams_DataOffset {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFile_starGenType_DataOffset_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_starGenType_DataOffset_paramsKey]*MoqFile_starGenType_DataOffset_results{},
		}
		r.Moq.ResultsByParams_DataOffset = append(r.Moq.ResultsByParams_DataOffset, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_DataOffset) {
			copy(r.Moq.ResultsByParams_DataOffset[insertAt+1:], r.Moq.ResultsByParams_DataOffset[insertAt:0])
			r.Moq.ResultsByParams_DataOffset[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_DataOffset(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFile_starGenType_DataOffset_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_starGenType_DataOffset_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_starGenType_DataOffset_fnRecorder {
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
					Offset int64
					Err    error
				}
				Sequence   uint32
				DoFn       MoqFile_starGenType_DataOffset_doFn
				DoReturnFn MoqFile_starGenType_DataOffset_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile_starGenType) PrettyParams_DataOffset(params MoqFile_starGenType_DataOffset_params) string {
	return fmt.Sprintf("DataOffset()")
}

func (m *MoqFile_starGenType) ParamsKey_DataOffset(params MoqFile_starGenType_DataOffset_params, anyParams uint64) MoqFile_starGenType_DataOffset_paramsKey {
	m.Scene.T.Helper()
	return MoqFile_starGenType_DataOffset_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqFile_starGenType_recorder) Open() *MoqFile_starGenType_Open_fnRecorder {
	return &MoqFile_starGenType_Open_fnRecorder{
		Params:   MoqFile_starGenType_Open_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_starGenType_Open_fnRecorder) Any() *MoqFile_starGenType_Open_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	return &MoqFile_starGenType_Open_anyParams{Recorder: r}
}

func (r *MoqFile_starGenType_Open_fnRecorder) Seq() *MoqFile_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_starGenType_Open_fnRecorder) NoSeq() *MoqFile_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_starGenType_Open_fnRecorder) ReturnResults(result1 io.ReadCloser, result2 error) *MoqFile_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Open_doFn
		DoReturnFn MoqFile_starGenType_Open_doReturnFn
	}{
		Values: &struct {
			Result1 io.ReadCloser
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFile_starGenType_Open_fnRecorder) AndDo(fn MoqFile_starGenType_Open_doFn) *MoqFile_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_starGenType_Open_fnRecorder) DoReturnResults(fn MoqFile_starGenType_Open_doReturnFn) *MoqFile_starGenType_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Open_doFn
		DoReturnFn MoqFile_starGenType_Open_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_starGenType_Open_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_starGenType_Open_resultsByParams
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
		results = &MoqFile_starGenType_Open_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_starGenType_Open_paramsKey]*MoqFile_starGenType_Open_results{},
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
		r.Results = &MoqFile_starGenType_Open_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_starGenType_Open_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_starGenType_Open_fnRecorder {
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
					Result1 io.ReadCloser
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqFile_starGenType_Open_doFn
				DoReturnFn MoqFile_starGenType_Open_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile_starGenType) PrettyParams_Open(params MoqFile_starGenType_Open_params) string {
	return fmt.Sprintf("Open()")
}

func (m *MoqFile_starGenType) ParamsKey_Open(params MoqFile_starGenType_Open_params, anyParams uint64) MoqFile_starGenType_Open_paramsKey {
	m.Scene.T.Helper()
	return MoqFile_starGenType_Open_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqFile_starGenType) Reset() {
	m.ResultsByParams_DataOffset = nil
	m.ResultsByParams_Open = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFile_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_DataOffset {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_DataOffset(results.Params))
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
