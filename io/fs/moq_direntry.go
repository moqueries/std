// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package fs

import (
	"fmt"
	"io/fs"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that fs.DirEntry is mocked completely
var _ fs.DirEntry = (*MoqDirEntry_mock)(nil)

// MoqDirEntry holds the state of a moq of the DirEntry type
type MoqDirEntry struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDirEntry_mock

	ResultsByParams_Name  []MoqDirEntry_Name_resultsByParams
	ResultsByParams_IsDir []MoqDirEntry_IsDir_resultsByParams
	ResultsByParams_Type  []MoqDirEntry_Type_resultsByParams
	ResultsByParams_Info  []MoqDirEntry_Info_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name  struct{}
			IsDir struct{}
			Type  struct{}
			Info  struct{}
		}
	}
}

// MoqDirEntry_mock isolates the mock interface of the DirEntry type
type MoqDirEntry_mock struct {
	Moq *MoqDirEntry
}

// MoqDirEntry_recorder isolates the recorder interface of the DirEntry type
type MoqDirEntry_recorder struct {
	Moq *MoqDirEntry
}

// MoqDirEntry_Name_params holds the params of the DirEntry type
type MoqDirEntry_Name_params struct{}

// MoqDirEntry_Name_paramsKey holds the map key params of the DirEntry type
type MoqDirEntry_Name_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDirEntry_Name_resultsByParams contains the results for a given set of
// parameters for the DirEntry type
type MoqDirEntry_Name_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDirEntry_Name_paramsKey]*MoqDirEntry_Name_results
}

// MoqDirEntry_Name_doFn defines the type of function needed when calling AndDo
// for the DirEntry type
type MoqDirEntry_Name_doFn func()

// MoqDirEntry_Name_doReturnFn defines the type of function needed when calling
// DoReturnResults for the DirEntry type
type MoqDirEntry_Name_doReturnFn func() string

// MoqDirEntry_Name_results holds the results of the DirEntry type
type MoqDirEntry_Name_results struct {
	Params  MoqDirEntry_Name_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Name_doFn
		DoReturnFn MoqDirEntry_Name_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDirEntry_Name_fnRecorder routes recorded function calls to the
// MoqDirEntry moq
type MoqDirEntry_Name_fnRecorder struct {
	Params    MoqDirEntry_Name_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDirEntry_Name_results
	Moq       *MoqDirEntry
}

// MoqDirEntry_Name_anyParams isolates the any params functions of the DirEntry
// type
type MoqDirEntry_Name_anyParams struct {
	Recorder *MoqDirEntry_Name_fnRecorder
}

// MoqDirEntry_IsDir_params holds the params of the DirEntry type
type MoqDirEntry_IsDir_params struct{}

// MoqDirEntry_IsDir_paramsKey holds the map key params of the DirEntry type
type MoqDirEntry_IsDir_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDirEntry_IsDir_resultsByParams contains the results for a given set of
// parameters for the DirEntry type
type MoqDirEntry_IsDir_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDirEntry_IsDir_paramsKey]*MoqDirEntry_IsDir_results
}

// MoqDirEntry_IsDir_doFn defines the type of function needed when calling
// AndDo for the DirEntry type
type MoqDirEntry_IsDir_doFn func()

// MoqDirEntry_IsDir_doReturnFn defines the type of function needed when
// calling DoReturnResults for the DirEntry type
type MoqDirEntry_IsDir_doReturnFn func() bool

// MoqDirEntry_IsDir_results holds the results of the DirEntry type
type MoqDirEntry_IsDir_results struct {
	Params  MoqDirEntry_IsDir_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqDirEntry_IsDir_doFn
		DoReturnFn MoqDirEntry_IsDir_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDirEntry_IsDir_fnRecorder routes recorded function calls to the
// MoqDirEntry moq
type MoqDirEntry_IsDir_fnRecorder struct {
	Params    MoqDirEntry_IsDir_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDirEntry_IsDir_results
	Moq       *MoqDirEntry
}

// MoqDirEntry_IsDir_anyParams isolates the any params functions of the
// DirEntry type
type MoqDirEntry_IsDir_anyParams struct {
	Recorder *MoqDirEntry_IsDir_fnRecorder
}

// MoqDirEntry_Type_params holds the params of the DirEntry type
type MoqDirEntry_Type_params struct{}

// MoqDirEntry_Type_paramsKey holds the map key params of the DirEntry type
type MoqDirEntry_Type_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDirEntry_Type_resultsByParams contains the results for a given set of
// parameters for the DirEntry type
type MoqDirEntry_Type_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDirEntry_Type_paramsKey]*MoqDirEntry_Type_results
}

// MoqDirEntry_Type_doFn defines the type of function needed when calling AndDo
// for the DirEntry type
type MoqDirEntry_Type_doFn func()

// MoqDirEntry_Type_doReturnFn defines the type of function needed when calling
// DoReturnResults for the DirEntry type
type MoqDirEntry_Type_doReturnFn func() fs.FileMode

// MoqDirEntry_Type_results holds the results of the DirEntry type
type MoqDirEntry_Type_results struct {
	Params  MoqDirEntry_Type_params
	Results []struct {
		Values *struct {
			Result1 fs.FileMode
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Type_doFn
		DoReturnFn MoqDirEntry_Type_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDirEntry_Type_fnRecorder routes recorded function calls to the
// MoqDirEntry moq
type MoqDirEntry_Type_fnRecorder struct {
	Params    MoqDirEntry_Type_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDirEntry_Type_results
	Moq       *MoqDirEntry
}

// MoqDirEntry_Type_anyParams isolates the any params functions of the DirEntry
// type
type MoqDirEntry_Type_anyParams struct {
	Recorder *MoqDirEntry_Type_fnRecorder
}

// MoqDirEntry_Info_params holds the params of the DirEntry type
type MoqDirEntry_Info_params struct{}

// MoqDirEntry_Info_paramsKey holds the map key params of the DirEntry type
type MoqDirEntry_Info_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDirEntry_Info_resultsByParams contains the results for a given set of
// parameters for the DirEntry type
type MoqDirEntry_Info_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDirEntry_Info_paramsKey]*MoqDirEntry_Info_results
}

// MoqDirEntry_Info_doFn defines the type of function needed when calling AndDo
// for the DirEntry type
type MoqDirEntry_Info_doFn func()

// MoqDirEntry_Info_doReturnFn defines the type of function needed when calling
// DoReturnResults for the DirEntry type
type MoqDirEntry_Info_doReturnFn func() (fs.FileInfo, error)

// MoqDirEntry_Info_results holds the results of the DirEntry type
type MoqDirEntry_Info_results struct {
	Params  MoqDirEntry_Info_params
	Results []struct {
		Values *struct {
			Result1 fs.FileInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Info_doFn
		DoReturnFn MoqDirEntry_Info_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDirEntry_Info_fnRecorder routes recorded function calls to the
// MoqDirEntry moq
type MoqDirEntry_Info_fnRecorder struct {
	Params    MoqDirEntry_Info_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDirEntry_Info_results
	Moq       *MoqDirEntry
}

// MoqDirEntry_Info_anyParams isolates the any params functions of the DirEntry
// type
type MoqDirEntry_Info_anyParams struct {
	Recorder *MoqDirEntry_Info_fnRecorder
}

// NewMoqDirEntry creates a new moq of the DirEntry type
func NewMoqDirEntry(scene *moq.Scene, config *moq.Config) *MoqDirEntry {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDirEntry{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDirEntry_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Name  struct{}
				IsDir struct{}
				Type  struct{}
				Info  struct{}
			}
		}{ParameterIndexing: struct {
			Name  struct{}
			IsDir struct{}
			Type  struct{}
			Info  struct{}
		}{
			Name:  struct{}{},
			IsDir: struct{}{},
			Type:  struct{}{},
			Info:  struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the DirEntry type
func (m *MoqDirEntry) Mock() *MoqDirEntry_mock { return m.Moq }

func (m *MoqDirEntry_mock) Name() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqDirEntry_Name_params{}
	var results *MoqDirEntry_Name_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Name {
		paramsKey := m.Moq.ParamsKey_Name(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Name(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Name(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Name(params))
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

func (m *MoqDirEntry_mock) IsDir() (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqDirEntry_IsDir_params{}
	var results *MoqDirEntry_IsDir_results
	for _, resultsByParams := range m.Moq.ResultsByParams_IsDir {
		paramsKey := m.Moq.ParamsKey_IsDir(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_IsDir(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_IsDir(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_IsDir(params))
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

func (m *MoqDirEntry_mock) Type() (result1 fs.FileMode) {
	m.Moq.Scene.T.Helper()
	params := MoqDirEntry_Type_params{}
	var results *MoqDirEntry_Type_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Type {
		paramsKey := m.Moq.ParamsKey_Type(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Type(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Type(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Type(params))
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

func (m *MoqDirEntry_mock) Info() (result1 fs.FileInfo, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDirEntry_Info_params{}
	var results *MoqDirEntry_Info_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Info {
		paramsKey := m.Moq.ParamsKey_Info(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Info(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Info(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Info(params))
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

// OnCall returns the recorder implementation of the DirEntry type
func (m *MoqDirEntry) OnCall() *MoqDirEntry_recorder {
	return &MoqDirEntry_recorder{
		Moq: m,
	}
}

func (m *MoqDirEntry_recorder) Name() *MoqDirEntry_Name_fnRecorder {
	return &MoqDirEntry_Name_fnRecorder{
		Params:   MoqDirEntry_Name_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDirEntry_Name_fnRecorder) Any() *MoqDirEntry_Name_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Name(r.Params))
		return nil
	}
	return &MoqDirEntry_Name_anyParams{Recorder: r}
}

func (r *MoqDirEntry_Name_fnRecorder) Seq() *MoqDirEntry_Name_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Name(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDirEntry_Name_fnRecorder) NoSeq() *MoqDirEntry_Name_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Name(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDirEntry_Name_fnRecorder) ReturnResults(result1 string) *MoqDirEntry_Name_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Name_doFn
		DoReturnFn MoqDirEntry_Name_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDirEntry_Name_fnRecorder) AndDo(fn MoqDirEntry_Name_doFn) *MoqDirEntry_Name_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDirEntry_Name_fnRecorder) DoReturnResults(fn MoqDirEntry_Name_doReturnFn) *MoqDirEntry_Name_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Name_doFn
		DoReturnFn MoqDirEntry_Name_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDirEntry_Name_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDirEntry_Name_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Name {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDirEntry_Name_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDirEntry_Name_paramsKey]*MoqDirEntry_Name_results{},
		}
		r.Moq.ResultsByParams_Name = append(r.Moq.ResultsByParams_Name, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Name) {
			copy(r.Moq.ResultsByParams_Name[insertAt+1:], r.Moq.ResultsByParams_Name[insertAt:0])
			r.Moq.ResultsByParams_Name[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Name(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDirEntry_Name_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDirEntry_Name_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDirEntry_Name_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqDirEntry_Name_doFn
				DoReturnFn MoqDirEntry_Name_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDirEntry) PrettyParams_Name(params MoqDirEntry_Name_params) string {
	return fmt.Sprintf("Name()")
}

func (m *MoqDirEntry) ParamsKey_Name(params MoqDirEntry_Name_params, anyParams uint64) MoqDirEntry_Name_paramsKey {
	m.Scene.T.Helper()
	return MoqDirEntry_Name_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqDirEntry_recorder) IsDir() *MoqDirEntry_IsDir_fnRecorder {
	return &MoqDirEntry_IsDir_fnRecorder{
		Params:   MoqDirEntry_IsDir_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDirEntry_IsDir_fnRecorder) Any() *MoqDirEntry_IsDir_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_IsDir(r.Params))
		return nil
	}
	return &MoqDirEntry_IsDir_anyParams{Recorder: r}
}

func (r *MoqDirEntry_IsDir_fnRecorder) Seq() *MoqDirEntry_IsDir_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_IsDir(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDirEntry_IsDir_fnRecorder) NoSeq() *MoqDirEntry_IsDir_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_IsDir(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDirEntry_IsDir_fnRecorder) ReturnResults(result1 bool) *MoqDirEntry_IsDir_fnRecorder {
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
		DoFn       MoqDirEntry_IsDir_doFn
		DoReturnFn MoqDirEntry_IsDir_doReturnFn
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

func (r *MoqDirEntry_IsDir_fnRecorder) AndDo(fn MoqDirEntry_IsDir_doFn) *MoqDirEntry_IsDir_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDirEntry_IsDir_fnRecorder) DoReturnResults(fn MoqDirEntry_IsDir_doReturnFn) *MoqDirEntry_IsDir_fnRecorder {
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
		DoFn       MoqDirEntry_IsDir_doFn
		DoReturnFn MoqDirEntry_IsDir_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDirEntry_IsDir_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDirEntry_IsDir_resultsByParams
	for n, res := range r.Moq.ResultsByParams_IsDir {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDirEntry_IsDir_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDirEntry_IsDir_paramsKey]*MoqDirEntry_IsDir_results{},
		}
		r.Moq.ResultsByParams_IsDir = append(r.Moq.ResultsByParams_IsDir, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_IsDir) {
			copy(r.Moq.ResultsByParams_IsDir[insertAt+1:], r.Moq.ResultsByParams_IsDir[insertAt:0])
			r.Moq.ResultsByParams_IsDir[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_IsDir(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDirEntry_IsDir_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDirEntry_IsDir_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDirEntry_IsDir_fnRecorder {
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
				DoFn       MoqDirEntry_IsDir_doFn
				DoReturnFn MoqDirEntry_IsDir_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDirEntry) PrettyParams_IsDir(params MoqDirEntry_IsDir_params) string {
	return fmt.Sprintf("IsDir()")
}

func (m *MoqDirEntry) ParamsKey_IsDir(params MoqDirEntry_IsDir_params, anyParams uint64) MoqDirEntry_IsDir_paramsKey {
	m.Scene.T.Helper()
	return MoqDirEntry_IsDir_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqDirEntry_recorder) Type() *MoqDirEntry_Type_fnRecorder {
	return &MoqDirEntry_Type_fnRecorder{
		Params:   MoqDirEntry_Type_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDirEntry_Type_fnRecorder) Any() *MoqDirEntry_Type_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Type(r.Params))
		return nil
	}
	return &MoqDirEntry_Type_anyParams{Recorder: r}
}

func (r *MoqDirEntry_Type_fnRecorder) Seq() *MoqDirEntry_Type_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Type(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDirEntry_Type_fnRecorder) NoSeq() *MoqDirEntry_Type_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Type(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDirEntry_Type_fnRecorder) ReturnResults(result1 fs.FileMode) *MoqDirEntry_Type_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FileMode
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Type_doFn
		DoReturnFn MoqDirEntry_Type_doReturnFn
	}{
		Values: &struct {
			Result1 fs.FileMode
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDirEntry_Type_fnRecorder) AndDo(fn MoqDirEntry_Type_doFn) *MoqDirEntry_Type_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDirEntry_Type_fnRecorder) DoReturnResults(fn MoqDirEntry_Type_doReturnFn) *MoqDirEntry_Type_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FileMode
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Type_doFn
		DoReturnFn MoqDirEntry_Type_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDirEntry_Type_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDirEntry_Type_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Type {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDirEntry_Type_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDirEntry_Type_paramsKey]*MoqDirEntry_Type_results{},
		}
		r.Moq.ResultsByParams_Type = append(r.Moq.ResultsByParams_Type, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Type) {
			copy(r.Moq.ResultsByParams_Type[insertAt+1:], r.Moq.ResultsByParams_Type[insertAt:0])
			r.Moq.ResultsByParams_Type[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Type(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDirEntry_Type_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDirEntry_Type_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDirEntry_Type_fnRecorder {
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
					Result1 fs.FileMode
				}
				Sequence   uint32
				DoFn       MoqDirEntry_Type_doFn
				DoReturnFn MoqDirEntry_Type_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDirEntry) PrettyParams_Type(params MoqDirEntry_Type_params) string {
	return fmt.Sprintf("Type()")
}

func (m *MoqDirEntry) ParamsKey_Type(params MoqDirEntry_Type_params, anyParams uint64) MoqDirEntry_Type_paramsKey {
	m.Scene.T.Helper()
	return MoqDirEntry_Type_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqDirEntry_recorder) Info() *MoqDirEntry_Info_fnRecorder {
	return &MoqDirEntry_Info_fnRecorder{
		Params:   MoqDirEntry_Info_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDirEntry_Info_fnRecorder) Any() *MoqDirEntry_Info_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Info(r.Params))
		return nil
	}
	return &MoqDirEntry_Info_anyParams{Recorder: r}
}

func (r *MoqDirEntry_Info_fnRecorder) Seq() *MoqDirEntry_Info_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Info(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDirEntry_Info_fnRecorder) NoSeq() *MoqDirEntry_Info_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Info(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDirEntry_Info_fnRecorder) ReturnResults(result1 fs.FileInfo, result2 error) *MoqDirEntry_Info_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FileInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Info_doFn
		DoReturnFn MoqDirEntry_Info_doReturnFn
	}{
		Values: &struct {
			Result1 fs.FileInfo
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDirEntry_Info_fnRecorder) AndDo(fn MoqDirEntry_Info_doFn) *MoqDirEntry_Info_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDirEntry_Info_fnRecorder) DoReturnResults(fn MoqDirEntry_Info_doReturnFn) *MoqDirEntry_Info_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FileInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDirEntry_Info_doFn
		DoReturnFn MoqDirEntry_Info_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDirEntry_Info_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDirEntry_Info_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Info {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDirEntry_Info_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDirEntry_Info_paramsKey]*MoqDirEntry_Info_results{},
		}
		r.Moq.ResultsByParams_Info = append(r.Moq.ResultsByParams_Info, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Info) {
			copy(r.Moq.ResultsByParams_Info[insertAt+1:], r.Moq.ResultsByParams_Info[insertAt:0])
			r.Moq.ResultsByParams_Info[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Info(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDirEntry_Info_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDirEntry_Info_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDirEntry_Info_fnRecorder {
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
					Result1 fs.FileInfo
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqDirEntry_Info_doFn
				DoReturnFn MoqDirEntry_Info_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDirEntry) PrettyParams_Info(params MoqDirEntry_Info_params) string {
	return fmt.Sprintf("Info()")
}

func (m *MoqDirEntry) ParamsKey_Info(params MoqDirEntry_Info_params, anyParams uint64) MoqDirEntry_Info_paramsKey {
	m.Scene.T.Helper()
	return MoqDirEntry_Info_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqDirEntry) Reset() {
	m.ResultsByParams_Name = nil
	m.ResultsByParams_IsDir = nil
	m.ResultsByParams_Type = nil
	m.ResultsByParams_Info = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDirEntry) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Name {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Name(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_IsDir {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_IsDir(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Type {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Type(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Info {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Info(results.Params))
			}
		}
	}
}
