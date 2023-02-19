// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package plan9obj

import (
	"debug/plan9obj"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that plan9obj.File_starGenType is
// mocked completely
var _ File_starGenType = (*MoqFile_starGenType_mock)(nil)

// File_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type File_starGenType interface {
	Close() error
	Symbols() ([]plan9obj.Sym, error)
	Section(name string) *plan9obj.Section
}

// MoqFile_starGenType holds the state of a moq of the File_starGenType type
type MoqFile_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFile_starGenType_mock

	ResultsByParams_Close   []MoqFile_starGenType_Close_resultsByParams
	ResultsByParams_Symbols []MoqFile_starGenType_Symbols_resultsByParams
	ResultsByParams_Section []MoqFile_starGenType_Section_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Close   struct{}
			Symbols struct{}
			Section struct {
				Name moq.ParamIndexing
			}
		}
	}
	// MoqFile_starGenType_mock isolates the mock interface of the File_starGenType
}

// type
type MoqFile_starGenType_mock struct {
	Moq *MoqFile_starGenType
}

// MoqFile_starGenType_recorder isolates the recorder interface of the
// File_starGenType type
type MoqFile_starGenType_recorder struct {
	Moq *MoqFile_starGenType
}

// MoqFile_starGenType_Close_params holds the params of the File_starGenType
// type
type MoqFile_starGenType_Close_params struct{}

// MoqFile_starGenType_Close_paramsKey holds the map key params of the
// File_starGenType type
type MoqFile_starGenType_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFile_starGenType_Close_resultsByParams contains the results for a given
// set of parameters for the File_starGenType type
type MoqFile_starGenType_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_starGenType_Close_paramsKey]*MoqFile_starGenType_Close_results
}

// MoqFile_starGenType_Close_doFn defines the type of function needed when
// calling AndDo for the File_starGenType type
type MoqFile_starGenType_Close_doFn func()

// MoqFile_starGenType_Close_doReturnFn defines the type of function needed
// when calling DoReturnResults for the File_starGenType type
type MoqFile_starGenType_Close_doReturnFn func() error

// MoqFile_starGenType_Close_results holds the results of the File_starGenType
// type
type MoqFile_starGenType_Close_results struct {
	Params  MoqFile_starGenType_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Close_doFn
		DoReturnFn MoqFile_starGenType_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_starGenType_Close_fnRecorder routes recorded function calls to the
// MoqFile_starGenType moq
type MoqFile_starGenType_Close_fnRecorder struct {
	Params    MoqFile_starGenType_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_starGenType_Close_results
	Moq       *MoqFile_starGenType
}

// MoqFile_starGenType_Close_anyParams isolates the any params functions of the
// File_starGenType type
type MoqFile_starGenType_Close_anyParams struct {
	Recorder *MoqFile_starGenType_Close_fnRecorder
}

// MoqFile_starGenType_Symbols_params holds the params of the File_starGenType
// type
type MoqFile_starGenType_Symbols_params struct{}

// MoqFile_starGenType_Symbols_paramsKey holds the map key params of the
// File_starGenType type
type MoqFile_starGenType_Symbols_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFile_starGenType_Symbols_resultsByParams contains the results for a given
// set of parameters for the File_starGenType type
type MoqFile_starGenType_Symbols_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_starGenType_Symbols_paramsKey]*MoqFile_starGenType_Symbols_results
}

// MoqFile_starGenType_Symbols_doFn defines the type of function needed when
// calling AndDo for the File_starGenType type
type MoqFile_starGenType_Symbols_doFn func()

// MoqFile_starGenType_Symbols_doReturnFn defines the type of function needed
// when calling DoReturnResults for the File_starGenType type
type MoqFile_starGenType_Symbols_doReturnFn func() ([]plan9obj.Sym, error)

// MoqFile_starGenType_Symbols_results holds the results of the
// File_starGenType type
type MoqFile_starGenType_Symbols_results struct {
	Params  MoqFile_starGenType_Symbols_params
	Results []struct {
		Values *struct {
			Result1 []plan9obj.Sym
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Symbols_doFn
		DoReturnFn MoqFile_starGenType_Symbols_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_starGenType_Symbols_fnRecorder routes recorded function calls to the
// MoqFile_starGenType moq
type MoqFile_starGenType_Symbols_fnRecorder struct {
	Params    MoqFile_starGenType_Symbols_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_starGenType_Symbols_results
	Moq       *MoqFile_starGenType
}

// MoqFile_starGenType_Symbols_anyParams isolates the any params functions of
// the File_starGenType type
type MoqFile_starGenType_Symbols_anyParams struct {
	Recorder *MoqFile_starGenType_Symbols_fnRecorder
}

// MoqFile_starGenType_Section_params holds the params of the File_starGenType
// type
type MoqFile_starGenType_Section_params struct{ Name string }

// MoqFile_starGenType_Section_paramsKey holds the map key params of the
// File_starGenType type
type MoqFile_starGenType_Section_paramsKey struct {
	Params struct{ Name string }
	Hashes struct{ Name hash.Hash }
}

// MoqFile_starGenType_Section_resultsByParams contains the results for a given
// set of parameters for the File_starGenType type
type MoqFile_starGenType_Section_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_starGenType_Section_paramsKey]*MoqFile_starGenType_Section_results
}

// MoqFile_starGenType_Section_doFn defines the type of function needed when
// calling AndDo for the File_starGenType type
type MoqFile_starGenType_Section_doFn func(name string)

// MoqFile_starGenType_Section_doReturnFn defines the type of function needed
// when calling DoReturnResults for the File_starGenType type
type MoqFile_starGenType_Section_doReturnFn func(name string) *plan9obj.Section

// MoqFile_starGenType_Section_results holds the results of the
// File_starGenType type
type MoqFile_starGenType_Section_results struct {
	Params  MoqFile_starGenType_Section_params
	Results []struct {
		Values *struct {
			Result1 *plan9obj.Section
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Section_doFn
		DoReturnFn MoqFile_starGenType_Section_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_starGenType_Section_fnRecorder routes recorded function calls to the
// MoqFile_starGenType moq
type MoqFile_starGenType_Section_fnRecorder struct {
	Params    MoqFile_starGenType_Section_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_starGenType_Section_results
	Moq       *MoqFile_starGenType
}

// MoqFile_starGenType_Section_anyParams isolates the any params functions of
// the File_starGenType type
type MoqFile_starGenType_Section_anyParams struct {
	Recorder *MoqFile_starGenType_Section_fnRecorder
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
				Close   struct{}
				Symbols struct{}
				Section struct {
					Name moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Close   struct{}
			Symbols struct{}
			Section struct {
				Name moq.ParamIndexing
			}
		}{
			Close:   struct{}{},
			Symbols: struct{}{},
			Section: struct {
				Name moq.ParamIndexing
			}{
				Name: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the File_starGenType type
func (m *MoqFile_starGenType) Mock() *MoqFile_starGenType_mock { return m.Moq }

func (m *MoqFile_starGenType_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_starGenType_Close_params{}
	var results *MoqFile_starGenType_Close_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Close {
		paramsKey := m.Moq.ParamsKey_Close(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Close(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Close(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Close(params))
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

func (m *MoqFile_starGenType_mock) Symbols() (result1 []plan9obj.Sym, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_starGenType_Symbols_params{}
	var results *MoqFile_starGenType_Symbols_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Symbols {
		paramsKey := m.Moq.ParamsKey_Symbols(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Symbols(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Symbols(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Symbols(params))
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

func (m *MoqFile_starGenType_mock) Section(name string) (result1 *plan9obj.Section) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_starGenType_Section_params{
		Name: name,
	}
	var results *MoqFile_starGenType_Section_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Section {
		paramsKey := m.Moq.ParamsKey_Section(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Section(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Section(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Section(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(name)
	}
	return
}

// OnCall returns the recorder implementation of the File_starGenType type
func (m *MoqFile_starGenType) OnCall() *MoqFile_starGenType_recorder {
	return &MoqFile_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqFile_starGenType_recorder) Close() *MoqFile_starGenType_Close_fnRecorder {
	return &MoqFile_starGenType_Close_fnRecorder{
		Params:   MoqFile_starGenType_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_starGenType_Close_fnRecorder) Any() *MoqFile_starGenType_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqFile_starGenType_Close_anyParams{Recorder: r}
}

func (r *MoqFile_starGenType_Close_fnRecorder) Seq() *MoqFile_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_starGenType_Close_fnRecorder) NoSeq() *MoqFile_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_starGenType_Close_fnRecorder) ReturnResults(result1 error) *MoqFile_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Close_doFn
		DoReturnFn MoqFile_starGenType_Close_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFile_starGenType_Close_fnRecorder) AndDo(fn MoqFile_starGenType_Close_doFn) *MoqFile_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_starGenType_Close_fnRecorder) DoReturnResults(fn MoqFile_starGenType_Close_doReturnFn) *MoqFile_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Close_doFn
		DoReturnFn MoqFile_starGenType_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_starGenType_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_starGenType_Close_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Close {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFile_starGenType_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_starGenType_Close_paramsKey]*MoqFile_starGenType_Close_results{},
		}
		r.Moq.ResultsByParams_Close = append(r.Moq.ResultsByParams_Close, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Close) {
			copy(r.Moq.ResultsByParams_Close[insertAt+1:], r.Moq.ResultsByParams_Close[insertAt:0])
			r.Moq.ResultsByParams_Close[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Close(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFile_starGenType_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_starGenType_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_starGenType_Close_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqFile_starGenType_Close_doFn
				DoReturnFn MoqFile_starGenType_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile_starGenType) PrettyParams_Close(params MoqFile_starGenType_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqFile_starGenType) ParamsKey_Close(params MoqFile_starGenType_Close_params, anyParams uint64) MoqFile_starGenType_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqFile_starGenType_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqFile_starGenType_recorder) Symbols() *MoqFile_starGenType_Symbols_fnRecorder {
	return &MoqFile_starGenType_Symbols_fnRecorder{
		Params:   MoqFile_starGenType_Symbols_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) Any() *MoqFile_starGenType_Symbols_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Symbols(r.Params))
		return nil
	}
	return &MoqFile_starGenType_Symbols_anyParams{Recorder: r}
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) Seq() *MoqFile_starGenType_Symbols_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Symbols(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) NoSeq() *MoqFile_starGenType_Symbols_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Symbols(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) ReturnResults(result1 []plan9obj.Sym, result2 error) *MoqFile_starGenType_Symbols_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []plan9obj.Sym
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Symbols_doFn
		DoReturnFn MoqFile_starGenType_Symbols_doReturnFn
	}{
		Values: &struct {
			Result1 []plan9obj.Sym
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) AndDo(fn MoqFile_starGenType_Symbols_doFn) *MoqFile_starGenType_Symbols_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) DoReturnResults(fn MoqFile_starGenType_Symbols_doReturnFn) *MoqFile_starGenType_Symbols_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []plan9obj.Sym
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Symbols_doFn
		DoReturnFn MoqFile_starGenType_Symbols_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_starGenType_Symbols_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Symbols {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFile_starGenType_Symbols_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_starGenType_Symbols_paramsKey]*MoqFile_starGenType_Symbols_results{},
		}
		r.Moq.ResultsByParams_Symbols = append(r.Moq.ResultsByParams_Symbols, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Symbols) {
			copy(r.Moq.ResultsByParams_Symbols[insertAt+1:], r.Moq.ResultsByParams_Symbols[insertAt:0])
			r.Moq.ResultsByParams_Symbols[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Symbols(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFile_starGenType_Symbols_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_starGenType_Symbols_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_starGenType_Symbols_fnRecorder {
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
					Result1 []plan9obj.Sym
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqFile_starGenType_Symbols_doFn
				DoReturnFn MoqFile_starGenType_Symbols_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile_starGenType) PrettyParams_Symbols(params MoqFile_starGenType_Symbols_params) string {
	return fmt.Sprintf("Symbols()")
}

func (m *MoqFile_starGenType) ParamsKey_Symbols(params MoqFile_starGenType_Symbols_params, anyParams uint64) MoqFile_starGenType_Symbols_paramsKey {
	m.Scene.T.Helper()
	return MoqFile_starGenType_Symbols_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqFile_starGenType_recorder) Section(name string) *MoqFile_starGenType_Section_fnRecorder {
	return &MoqFile_starGenType_Section_fnRecorder{
		Params: MoqFile_starGenType_Section_params{
			Name: name,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_starGenType_Section_fnRecorder) Any() *MoqFile_starGenType_Section_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Section(r.Params))
		return nil
	}
	return &MoqFile_starGenType_Section_anyParams{Recorder: r}
}

func (a *MoqFile_starGenType_Section_anyParams) Name() *MoqFile_starGenType_Section_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFile_starGenType_Section_fnRecorder) Seq() *MoqFile_starGenType_Section_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Section(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_starGenType_Section_fnRecorder) NoSeq() *MoqFile_starGenType_Section_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Section(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_starGenType_Section_fnRecorder) ReturnResults(result1 *plan9obj.Section) *MoqFile_starGenType_Section_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *plan9obj.Section
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Section_doFn
		DoReturnFn MoqFile_starGenType_Section_doReturnFn
	}{
		Values: &struct {
			Result1 *plan9obj.Section
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFile_starGenType_Section_fnRecorder) AndDo(fn MoqFile_starGenType_Section_doFn) *MoqFile_starGenType_Section_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_starGenType_Section_fnRecorder) DoReturnResults(fn MoqFile_starGenType_Section_doReturnFn) *MoqFile_starGenType_Section_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *plan9obj.Section
		}
		Sequence   uint32
		DoFn       MoqFile_starGenType_Section_doFn
		DoReturnFn MoqFile_starGenType_Section_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_starGenType_Section_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_starGenType_Section_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Section {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFile_starGenType_Section_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_starGenType_Section_paramsKey]*MoqFile_starGenType_Section_results{},
		}
		r.Moq.ResultsByParams_Section = append(r.Moq.ResultsByParams_Section, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Section) {
			copy(r.Moq.ResultsByParams_Section[insertAt+1:], r.Moq.ResultsByParams_Section[insertAt:0])
			r.Moq.ResultsByParams_Section[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Section(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFile_starGenType_Section_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_starGenType_Section_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_starGenType_Section_fnRecorder {
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
					Result1 *plan9obj.Section
				}
				Sequence   uint32
				DoFn       MoqFile_starGenType_Section_doFn
				DoReturnFn MoqFile_starGenType_Section_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile_starGenType) PrettyParams_Section(params MoqFile_starGenType_Section_params) string {
	return fmt.Sprintf("Section(%#v)", params.Name)
}

func (m *MoqFile_starGenType) ParamsKey_Section(params MoqFile_starGenType_Section_params, anyParams uint64) MoqFile_starGenType_Section_paramsKey {
	m.Scene.T.Helper()
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Section.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	return MoqFile_starGenType_Section_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct{ Name hash.Hash }{
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFile_starGenType) Reset() {
	m.ResultsByParams_Close = nil
	m.ResultsByParams_Symbols = nil
	m.ResultsByParams_Section = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFile_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Close {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Close(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Symbols {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Symbols(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Section {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Section(results.Params))
			}
		}
	}
}
