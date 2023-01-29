// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package fs

import (
	"fmt"
	"io/fs"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that fs.GlobFS is mocked completely
var _ fs.GlobFS = (*MoqGlobFS_mock)(nil)

// MoqGlobFS holds the state of a moq of the GlobFS type
type MoqGlobFS struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGlobFS_mock

	ResultsByParams_Open []MoqGlobFS_Open_resultsByParams
	ResultsByParams_Glob []MoqGlobFS_Glob_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Open struct {
				Name moq.ParamIndexing
			}
			Glob struct {
				Pattern moq.ParamIndexing
			}
		}
	}
	// MoqGlobFS_mock isolates the mock interface of the GlobFS type
}

type MoqGlobFS_mock struct {
	Moq *MoqGlobFS
}

// MoqGlobFS_recorder isolates the recorder interface of the GlobFS type
type MoqGlobFS_recorder struct {
	Moq *MoqGlobFS
}

// MoqGlobFS_Open_params holds the params of the GlobFS type
type MoqGlobFS_Open_params struct{ Name string }

// MoqGlobFS_Open_paramsKey holds the map key params of the GlobFS type
type MoqGlobFS_Open_paramsKey struct {
	Params struct{ Name string }
	Hashes struct{ Name hash.Hash }
}

// MoqGlobFS_Open_resultsByParams contains the results for a given set of
// parameters for the GlobFS type
type MoqGlobFS_Open_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGlobFS_Open_paramsKey]*MoqGlobFS_Open_results
}

// MoqGlobFS_Open_doFn defines the type of function needed when calling AndDo
// for the GlobFS type
type MoqGlobFS_Open_doFn func(name string)

// MoqGlobFS_Open_doReturnFn defines the type of function needed when calling
// DoReturnResults for the GlobFS type
type MoqGlobFS_Open_doReturnFn func(name string) (fs.File, error)

// MoqGlobFS_Open_results holds the results of the GlobFS type
type MoqGlobFS_Open_results struct {
	Params  MoqGlobFS_Open_params
	Results []struct {
		Values *struct {
			Result1 fs.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Open_doFn
		DoReturnFn MoqGlobFS_Open_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGlobFS_Open_fnRecorder routes recorded function calls to the MoqGlobFS
// moq
type MoqGlobFS_Open_fnRecorder struct {
	Params    MoqGlobFS_Open_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGlobFS_Open_results
	Moq       *MoqGlobFS
}

// MoqGlobFS_Open_anyParams isolates the any params functions of the GlobFS
// type
type MoqGlobFS_Open_anyParams struct {
	Recorder *MoqGlobFS_Open_fnRecorder
}

// MoqGlobFS_Glob_params holds the params of the GlobFS type
type MoqGlobFS_Glob_params struct{ Pattern string }

// MoqGlobFS_Glob_paramsKey holds the map key params of the GlobFS type
type MoqGlobFS_Glob_paramsKey struct {
	Params struct{ Pattern string }
	Hashes struct{ Pattern hash.Hash }
}

// MoqGlobFS_Glob_resultsByParams contains the results for a given set of
// parameters for the GlobFS type
type MoqGlobFS_Glob_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGlobFS_Glob_paramsKey]*MoqGlobFS_Glob_results
}

// MoqGlobFS_Glob_doFn defines the type of function needed when calling AndDo
// for the GlobFS type
type MoqGlobFS_Glob_doFn func(pattern string)

// MoqGlobFS_Glob_doReturnFn defines the type of function needed when calling
// DoReturnResults for the GlobFS type
type MoqGlobFS_Glob_doReturnFn func(pattern string) ([]string, error)

// MoqGlobFS_Glob_results holds the results of the GlobFS type
type MoqGlobFS_Glob_results struct {
	Params  MoqGlobFS_Glob_params
	Results []struct {
		Values *struct {
			Result1 []string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Glob_doFn
		DoReturnFn MoqGlobFS_Glob_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGlobFS_Glob_fnRecorder routes recorded function calls to the MoqGlobFS
// moq
type MoqGlobFS_Glob_fnRecorder struct {
	Params    MoqGlobFS_Glob_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGlobFS_Glob_results
	Moq       *MoqGlobFS
}

// MoqGlobFS_Glob_anyParams isolates the any params functions of the GlobFS
// type
type MoqGlobFS_Glob_anyParams struct {
	Recorder *MoqGlobFS_Glob_fnRecorder
}

// NewMoqGlobFS creates a new moq of the GlobFS type
func NewMoqGlobFS(scene *moq.Scene, config *moq.Config) *MoqGlobFS {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGlobFS{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGlobFS_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Open struct {
					Name moq.ParamIndexing
				}
				Glob struct {
					Pattern moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Open struct {
				Name moq.ParamIndexing
			}
			Glob struct {
				Pattern moq.ParamIndexing
			}
		}{
			Open: struct {
				Name moq.ParamIndexing
			}{
				Name: moq.ParamIndexByValue,
			},
			Glob: struct {
				Pattern moq.ParamIndexing
			}{
				Pattern: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the GlobFS type
func (m *MoqGlobFS) Mock() *MoqGlobFS_mock { return m.Moq }

func (m *MoqGlobFS_mock) Open(name string) (result1 fs.File, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqGlobFS_Open_params{
		Name: name,
	}
	var results *MoqGlobFS_Open_results
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
		result.DoFn(name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(name)
	}
	return
}

func (m *MoqGlobFS_mock) Glob(pattern string) (result1 []string, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqGlobFS_Glob_params{
		Pattern: pattern,
	}
	var results *MoqGlobFS_Glob_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Glob {
		paramsKey := m.Moq.ParamsKey_Glob(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Glob(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Glob(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Glob(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(pattern)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(pattern)
	}
	return
}

// OnCall returns the recorder implementation of the GlobFS type
func (m *MoqGlobFS) OnCall() *MoqGlobFS_recorder {
	return &MoqGlobFS_recorder{
		Moq: m,
	}
}

func (m *MoqGlobFS_recorder) Open(name string) *MoqGlobFS_Open_fnRecorder {
	return &MoqGlobFS_Open_fnRecorder{
		Params: MoqGlobFS_Open_params{
			Name: name,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqGlobFS_Open_fnRecorder) Any() *MoqGlobFS_Open_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	return &MoqGlobFS_Open_anyParams{Recorder: r}
}

func (a *MoqGlobFS_Open_anyParams) Name() *MoqGlobFS_Open_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqGlobFS_Open_fnRecorder) Seq() *MoqGlobFS_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGlobFS_Open_fnRecorder) NoSeq() *MoqGlobFS_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Open(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGlobFS_Open_fnRecorder) ReturnResults(result1 fs.File, result2 error) *MoqGlobFS_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Open_doFn
		DoReturnFn MoqGlobFS_Open_doReturnFn
	}{
		Values: &struct {
			Result1 fs.File
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGlobFS_Open_fnRecorder) AndDo(fn MoqGlobFS_Open_doFn) *MoqGlobFS_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGlobFS_Open_fnRecorder) DoReturnResults(fn MoqGlobFS_Open_doReturnFn) *MoqGlobFS_Open_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.File
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Open_doFn
		DoReturnFn MoqGlobFS_Open_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGlobFS_Open_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGlobFS_Open_resultsByParams
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
		results = &MoqGlobFS_Open_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGlobFS_Open_paramsKey]*MoqGlobFS_Open_results{},
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
		r.Results = &MoqGlobFS_Open_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGlobFS_Open_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGlobFS_Open_fnRecorder {
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
					Result1 fs.File
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqGlobFS_Open_doFn
				DoReturnFn MoqGlobFS_Open_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGlobFS) PrettyParams_Open(params MoqGlobFS_Open_params) string {
	return fmt.Sprintf("Open(%#v)", params.Name)
}

func (m *MoqGlobFS) ParamsKey_Open(params MoqGlobFS_Open_params, anyParams uint64) MoqGlobFS_Open_paramsKey {
	m.Scene.T.Helper()
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Open.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	return MoqGlobFS_Open_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct{ Name hash.Hash }{
			Name: nameUsedHash,
		},
	}
}

func (m *MoqGlobFS_recorder) Glob(pattern string) *MoqGlobFS_Glob_fnRecorder {
	return &MoqGlobFS_Glob_fnRecorder{
		Params: MoqGlobFS_Glob_params{
			Pattern: pattern,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqGlobFS_Glob_fnRecorder) Any() *MoqGlobFS_Glob_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Glob(r.Params))
		return nil
	}
	return &MoqGlobFS_Glob_anyParams{Recorder: r}
}

func (a *MoqGlobFS_Glob_anyParams) Pattern() *MoqGlobFS_Glob_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqGlobFS_Glob_fnRecorder) Seq() *MoqGlobFS_Glob_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Glob(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGlobFS_Glob_fnRecorder) NoSeq() *MoqGlobFS_Glob_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Glob(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGlobFS_Glob_fnRecorder) ReturnResults(result1 []string, result2 error) *MoqGlobFS_Glob_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Glob_doFn
		DoReturnFn MoqGlobFS_Glob_doReturnFn
	}{
		Values: &struct {
			Result1 []string
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGlobFS_Glob_fnRecorder) AndDo(fn MoqGlobFS_Glob_doFn) *MoqGlobFS_Glob_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGlobFS_Glob_fnRecorder) DoReturnResults(fn MoqGlobFS_Glob_doReturnFn) *MoqGlobFS_Glob_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGlobFS_Glob_doFn
		DoReturnFn MoqGlobFS_Glob_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGlobFS_Glob_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGlobFS_Glob_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Glob {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqGlobFS_Glob_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGlobFS_Glob_paramsKey]*MoqGlobFS_Glob_results{},
		}
		r.Moq.ResultsByParams_Glob = append(r.Moq.ResultsByParams_Glob, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Glob) {
			copy(r.Moq.ResultsByParams_Glob[insertAt+1:], r.Moq.ResultsByParams_Glob[insertAt:0])
			r.Moq.ResultsByParams_Glob[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Glob(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqGlobFS_Glob_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGlobFS_Glob_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGlobFS_Glob_fnRecorder {
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
					Result1 []string
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqGlobFS_Glob_doFn
				DoReturnFn MoqGlobFS_Glob_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGlobFS) PrettyParams_Glob(params MoqGlobFS_Glob_params) string {
	return fmt.Sprintf("Glob(%#v)", params.Pattern)
}

func (m *MoqGlobFS) ParamsKey_Glob(params MoqGlobFS_Glob_params, anyParams uint64) MoqGlobFS_Glob_paramsKey {
	m.Scene.T.Helper()
	var patternUsed string
	var patternUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Glob.Pattern == moq.ParamIndexByValue {
			patternUsed = params.Pattern
		} else {
			patternUsedHash = hash.DeepHash(params.Pattern)
		}
	}
	return MoqGlobFS_Glob_paramsKey{
		Params: struct{ Pattern string }{
			Pattern: patternUsed,
		},
		Hashes: struct{ Pattern hash.Hash }{
			Pattern: patternUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGlobFS) Reset() { m.ResultsByParams_Open = nil; m.ResultsByParams_Glob = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGlobFS) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Open {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Open(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Glob {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Glob(results.Params))
			}
		}
	}
}
