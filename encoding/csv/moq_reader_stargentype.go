// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package csv

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that csv.Reader_starGenType is mocked
// completely
var _ Reader_starGenType = (*MoqReader_starGenType_mock)(nil)

// Reader_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Reader_starGenType interface {
	Read() (record []string, err error)
	FieldPos(field int) (line, column int)
	ReadAll() (records [][]string, err error)
}

// MoqReader_starGenType holds the state of a moq of the Reader_starGenType
// type
type MoqReader_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReader_starGenType_mock

	ResultsByParams_Read     []MoqReader_starGenType_Read_resultsByParams
	ResultsByParams_FieldPos []MoqReader_starGenType_FieldPos_resultsByParams
	ResultsByParams_ReadAll  []MoqReader_starGenType_ReadAll_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Read     struct{}
			FieldPos struct {
				Field moq.ParamIndexing
			}
			ReadAll struct{}
		}
	}
}

// MoqReader_starGenType_mock isolates the mock interface of the
// Reader_starGenType type
type MoqReader_starGenType_mock struct {
	Moq *MoqReader_starGenType
}

// MoqReader_starGenType_recorder isolates the recorder interface of the
// Reader_starGenType type
type MoqReader_starGenType_recorder struct {
	Moq *MoqReader_starGenType
}

// MoqReader_starGenType_Read_params holds the params of the Reader_starGenType
// type
type MoqReader_starGenType_Read_params struct{}

// MoqReader_starGenType_Read_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_Read_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReader_starGenType_Read_resultsByParams contains the results for a given
// set of parameters for the Reader_starGenType type
type MoqReader_starGenType_Read_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_Read_paramsKey]*MoqReader_starGenType_Read_results
}

// MoqReader_starGenType_Read_doFn defines the type of function needed when
// calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_Read_doFn func()

// MoqReader_starGenType_Read_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_Read_doReturnFn func() (record []string, err error)

// MoqReader_starGenType_Read_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_Read_results struct {
	Params  MoqReader_starGenType_Read_params
	Results []struct {
		Values *struct {
			Record []string
			Err    error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_Read_doFn
		DoReturnFn MoqReader_starGenType_Read_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_Read_fnRecorder routes recorded function calls to the
// MoqReader_starGenType moq
type MoqReader_starGenType_Read_fnRecorder struct {
	Params    MoqReader_starGenType_Read_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_Read_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_Read_anyParams isolates the any params functions of
// the Reader_starGenType type
type MoqReader_starGenType_Read_anyParams struct {
	Recorder *MoqReader_starGenType_Read_fnRecorder
}

// MoqReader_starGenType_FieldPos_params holds the params of the
// Reader_starGenType type
type MoqReader_starGenType_FieldPos_params struct{ Field int }

// MoqReader_starGenType_FieldPos_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_FieldPos_paramsKey struct {
	Params struct{ Field int }
	Hashes struct{ Field hash.Hash }
}

// MoqReader_starGenType_FieldPos_resultsByParams contains the results for a
// given set of parameters for the Reader_starGenType type
type MoqReader_starGenType_FieldPos_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_FieldPos_paramsKey]*MoqReader_starGenType_FieldPos_results
}

// MoqReader_starGenType_FieldPos_doFn defines the type of function needed when
// calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_FieldPos_doFn func(field int)

// MoqReader_starGenType_FieldPos_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_FieldPos_doReturnFn func(field int) (line, column int)

// MoqReader_starGenType_FieldPos_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_FieldPos_results struct {
	Params  MoqReader_starGenType_FieldPos_params
	Results []struct {
		Values     *struct{ Line, Column int }
		Sequence   uint32
		DoFn       MoqReader_starGenType_FieldPos_doFn
		DoReturnFn MoqReader_starGenType_FieldPos_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_FieldPos_fnRecorder routes recorded function calls to
// the MoqReader_starGenType moq
type MoqReader_starGenType_FieldPos_fnRecorder struct {
	Params    MoqReader_starGenType_FieldPos_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_FieldPos_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_FieldPos_anyParams isolates the any params functions
// of the Reader_starGenType type
type MoqReader_starGenType_FieldPos_anyParams struct {
	Recorder *MoqReader_starGenType_FieldPos_fnRecorder
}

// MoqReader_starGenType_ReadAll_params holds the params of the
// Reader_starGenType type
type MoqReader_starGenType_ReadAll_params struct{}

// MoqReader_starGenType_ReadAll_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_ReadAll_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReader_starGenType_ReadAll_resultsByParams contains the results for a
// given set of parameters for the Reader_starGenType type
type MoqReader_starGenType_ReadAll_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_ReadAll_paramsKey]*MoqReader_starGenType_ReadAll_results
}

// MoqReader_starGenType_ReadAll_doFn defines the type of function needed when
// calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_ReadAll_doFn func()

// MoqReader_starGenType_ReadAll_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_ReadAll_doReturnFn func() (records [][]string, err error)

// MoqReader_starGenType_ReadAll_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_ReadAll_results struct {
	Params  MoqReader_starGenType_ReadAll_params
	Results []struct {
		Values *struct {
			Records [][]string
			Err     error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadAll_doFn
		DoReturnFn MoqReader_starGenType_ReadAll_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_ReadAll_fnRecorder routes recorded function calls to
// the MoqReader_starGenType moq
type MoqReader_starGenType_ReadAll_fnRecorder struct {
	Params    MoqReader_starGenType_ReadAll_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_ReadAll_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_ReadAll_anyParams isolates the any params functions of
// the Reader_starGenType type
type MoqReader_starGenType_ReadAll_anyParams struct {
	Recorder *MoqReader_starGenType_ReadAll_fnRecorder
}

// NewMoqReader_starGenType creates a new moq of the Reader_starGenType type
func NewMoqReader_starGenType(scene *moq.Scene, config *moq.Config) *MoqReader_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReader_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReader_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Read     struct{}
				FieldPos struct {
					Field moq.ParamIndexing
				}
				ReadAll struct{}
			}
		}{ParameterIndexing: struct {
			Read     struct{}
			FieldPos struct {
				Field moq.ParamIndexing
			}
			ReadAll struct{}
		}{
			Read: struct{}{},
			FieldPos: struct {
				Field moq.ParamIndexing
			}{
				Field: moq.ParamIndexByValue,
			},
			ReadAll: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Reader_starGenType type
func (m *MoqReader_starGenType) Mock() *MoqReader_starGenType_mock { return m.Moq }

func (m *MoqReader_starGenType_mock) Read() (record []string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_Read_params{}
	var results *MoqReader_starGenType_Read_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Read {
		paramsKey := m.Moq.ParamsKey_Read(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Read(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Read(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Read(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		record = result.Values.Record
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		record, err = result.DoReturnFn()
	}
	return
}

func (m *MoqReader_starGenType_mock) FieldPos(field int) (line, column int) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_FieldPos_params{
		Field: field,
	}
	var results *MoqReader_starGenType_FieldPos_results
	for _, resultsByParams := range m.Moq.ResultsByParams_FieldPos {
		paramsKey := m.Moq.ParamsKey_FieldPos(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_FieldPos(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_FieldPos(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_FieldPos(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(field)
	}

	if result.Values != nil {
		line = result.Values.Line
		column = result.Values.Column
	}
	if result.DoReturnFn != nil {
		line, column = result.DoReturnFn(field)
	}
	return
}

func (m *MoqReader_starGenType_mock) ReadAll() (records [][]string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_ReadAll_params{}
	var results *MoqReader_starGenType_ReadAll_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ReadAll {
		paramsKey := m.Moq.ParamsKey_ReadAll(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ReadAll(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ReadAll(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ReadAll(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		records = result.Values.Records
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		records, err = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Reader_starGenType type
func (m *MoqReader_starGenType) OnCall() *MoqReader_starGenType_recorder {
	return &MoqReader_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqReader_starGenType_recorder) Read() *MoqReader_starGenType_Read_fnRecorder {
	return &MoqReader_starGenType_Read_fnRecorder{
		Params:   MoqReader_starGenType_Read_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_Read_fnRecorder) Any() *MoqReader_starGenType_Read_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	return &MoqReader_starGenType_Read_anyParams{Recorder: r}
}

func (r *MoqReader_starGenType_Read_fnRecorder) Seq() *MoqReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_Read_fnRecorder) NoSeq() *MoqReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_Read_fnRecorder) ReturnResults(record []string, err error) *MoqReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Record []string
			Err    error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_Read_doFn
		DoReturnFn MoqReader_starGenType_Read_doReturnFn
	}{
		Values: &struct {
			Record []string
			Err    error
		}{
			Record: record,
			Err:    err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_Read_fnRecorder) AndDo(fn MoqReader_starGenType_Read_doFn) *MoqReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_Read_fnRecorder) DoReturnResults(fn MoqReader_starGenType_Read_doReturnFn) *MoqReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Record []string
			Err    error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_Read_doFn
		DoReturnFn MoqReader_starGenType_Read_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_Read_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_Read_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Read {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_Read_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_Read_paramsKey]*MoqReader_starGenType_Read_results{},
		}
		r.Moq.ResultsByParams_Read = append(r.Moq.ResultsByParams_Read, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Read) {
			copy(r.Moq.ResultsByParams_Read[insertAt+1:], r.Moq.ResultsByParams_Read[insertAt:0])
			r.Moq.ResultsByParams_Read[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Read(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_Read_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_Read_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_Read_fnRecorder {
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
					Record []string
					Err    error
				}
				Sequence   uint32
				DoFn       MoqReader_starGenType_Read_doFn
				DoReturnFn MoqReader_starGenType_Read_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_Read(params MoqReader_starGenType_Read_params) string {
	return fmt.Sprintf("Read()")
}

func (m *MoqReader_starGenType) ParamsKey_Read(params MoqReader_starGenType_Read_params, anyParams uint64) MoqReader_starGenType_Read_paramsKey {
	m.Scene.T.Helper()
	return MoqReader_starGenType_Read_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqReader_starGenType_recorder) FieldPos(field int) *MoqReader_starGenType_FieldPos_fnRecorder {
	return &MoqReader_starGenType_FieldPos_fnRecorder{
		Params: MoqReader_starGenType_FieldPos_params{
			Field: field,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) Any() *MoqReader_starGenType_FieldPos_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FieldPos(r.Params))
		return nil
	}
	return &MoqReader_starGenType_FieldPos_anyParams{Recorder: r}
}

func (a *MoqReader_starGenType_FieldPos_anyParams) Field() *MoqReader_starGenType_FieldPos_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) Seq() *MoqReader_starGenType_FieldPos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FieldPos(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) NoSeq() *MoqReader_starGenType_FieldPos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FieldPos(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) ReturnResults(line, column int) *MoqReader_starGenType_FieldPos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Line, Column int }
		Sequence   uint32
		DoFn       MoqReader_starGenType_FieldPos_doFn
		DoReturnFn MoqReader_starGenType_FieldPos_doReturnFn
	}{
		Values: &struct{ Line, Column int }{
			Line:   line,
			Column: column,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) AndDo(fn MoqReader_starGenType_FieldPos_doFn) *MoqReader_starGenType_FieldPos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) DoReturnResults(fn MoqReader_starGenType_FieldPos_doReturnFn) *MoqReader_starGenType_FieldPos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Line, Column int }
		Sequence   uint32
		DoFn       MoqReader_starGenType_FieldPos_doFn
		DoReturnFn MoqReader_starGenType_FieldPos_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_FieldPos_resultsByParams
	for n, res := range r.Moq.ResultsByParams_FieldPos {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_FieldPos_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_FieldPos_paramsKey]*MoqReader_starGenType_FieldPos_results{},
		}
		r.Moq.ResultsByParams_FieldPos = append(r.Moq.ResultsByParams_FieldPos, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_FieldPos) {
			copy(r.Moq.ResultsByParams_FieldPos[insertAt+1:], r.Moq.ResultsByParams_FieldPos[insertAt:0])
			r.Moq.ResultsByParams_FieldPos[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_FieldPos(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_FieldPos_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_FieldPos_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_FieldPos_fnRecorder {
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
				Values     *struct{ Line, Column int }
				Sequence   uint32
				DoFn       MoqReader_starGenType_FieldPos_doFn
				DoReturnFn MoqReader_starGenType_FieldPos_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_FieldPos(params MoqReader_starGenType_FieldPos_params) string {
	return fmt.Sprintf("FieldPos(%#v)", params.Field)
}

func (m *MoqReader_starGenType) ParamsKey_FieldPos(params MoqReader_starGenType_FieldPos_params, anyParams uint64) MoqReader_starGenType_FieldPos_paramsKey {
	m.Scene.T.Helper()
	var fieldUsed int
	var fieldUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.FieldPos.Field == moq.ParamIndexByValue {
			fieldUsed = params.Field
		} else {
			fieldUsedHash = hash.DeepHash(params.Field)
		}
	}
	return MoqReader_starGenType_FieldPos_paramsKey{
		Params: struct{ Field int }{
			Field: fieldUsed,
		},
		Hashes: struct{ Field hash.Hash }{
			Field: fieldUsedHash,
		},
	}
}

func (m *MoqReader_starGenType_recorder) ReadAll() *MoqReader_starGenType_ReadAll_fnRecorder {
	return &MoqReader_starGenType_ReadAll_fnRecorder{
		Params:   MoqReader_starGenType_ReadAll_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) Any() *MoqReader_starGenType_ReadAll_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAll(r.Params))
		return nil
	}
	return &MoqReader_starGenType_ReadAll_anyParams{Recorder: r}
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) Seq() *MoqReader_starGenType_ReadAll_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAll(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) NoSeq() *MoqReader_starGenType_ReadAll_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAll(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) ReturnResults(records [][]string, err error) *MoqReader_starGenType_ReadAll_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Records [][]string
			Err     error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadAll_doFn
		DoReturnFn MoqReader_starGenType_ReadAll_doReturnFn
	}{
		Values: &struct {
			Records [][]string
			Err     error
		}{
			Records: records,
			Err:     err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) AndDo(fn MoqReader_starGenType_ReadAll_doFn) *MoqReader_starGenType_ReadAll_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) DoReturnResults(fn MoqReader_starGenType_ReadAll_doReturnFn) *MoqReader_starGenType_ReadAll_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Records [][]string
			Err     error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadAll_doFn
		DoReturnFn MoqReader_starGenType_ReadAll_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_ReadAll_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ReadAll {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_ReadAll_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_ReadAll_paramsKey]*MoqReader_starGenType_ReadAll_results{},
		}
		r.Moq.ResultsByParams_ReadAll = append(r.Moq.ResultsByParams_ReadAll, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ReadAll) {
			copy(r.Moq.ResultsByParams_ReadAll[insertAt+1:], r.Moq.ResultsByParams_ReadAll[insertAt:0])
			r.Moq.ResultsByParams_ReadAll[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ReadAll(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_ReadAll_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_ReadAll_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_ReadAll_fnRecorder {
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
					Records [][]string
					Err     error
				}
				Sequence   uint32
				DoFn       MoqReader_starGenType_ReadAll_doFn
				DoReturnFn MoqReader_starGenType_ReadAll_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_ReadAll(params MoqReader_starGenType_ReadAll_params) string {
	return fmt.Sprintf("ReadAll()")
}

func (m *MoqReader_starGenType) ParamsKey_ReadAll(params MoqReader_starGenType_ReadAll_params, anyParams uint64) MoqReader_starGenType_ReadAll_paramsKey {
	m.Scene.T.Helper()
	return MoqReader_starGenType_ReadAll_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqReader_starGenType) Reset() {
	m.ResultsByParams_Read = nil
	m.ResultsByParams_FieldPos = nil
	m.ResultsByParams_ReadAll = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReader_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Read {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Read(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_FieldPos {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_FieldPos(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_ReadAll {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ReadAll(results.Params))
			}
		}
	}
}
