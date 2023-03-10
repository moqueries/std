// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package build

import (
	"fmt"
	"go/build"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ImportDir_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ImportDir_genType func(dir string, mode build.ImportMode) (*build.Package, error)

// MoqImportDir_genType holds the state of a moq of the ImportDir_genType type
type MoqImportDir_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqImportDir_genType_mock

	ResultsByParams []MoqImportDir_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dir  moq.ParamIndexing
			Mode moq.ParamIndexing
		}
	}
}

// MoqImportDir_genType_mock isolates the mock interface of the
// ImportDir_genType type
type MoqImportDir_genType_mock struct {
	Moq *MoqImportDir_genType
}

// MoqImportDir_genType_params holds the params of the ImportDir_genType type
type MoqImportDir_genType_params struct {
	Dir  string
	Mode build.ImportMode
}

// MoqImportDir_genType_paramsKey holds the map key params of the
// ImportDir_genType type
type MoqImportDir_genType_paramsKey struct {
	Params struct {
		Dir  string
		Mode build.ImportMode
	}
	Hashes struct {
		Dir  hash.Hash
		Mode hash.Hash
	}
}

// MoqImportDir_genType_resultsByParams contains the results for a given set of
// parameters for the ImportDir_genType type
type MoqImportDir_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqImportDir_genType_paramsKey]*MoqImportDir_genType_results
}

// MoqImportDir_genType_doFn defines the type of function needed when calling
// AndDo for the ImportDir_genType type
type MoqImportDir_genType_doFn func(dir string, mode build.ImportMode)

// MoqImportDir_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ImportDir_genType type
type MoqImportDir_genType_doReturnFn func(dir string, mode build.ImportMode) (*build.Package, error)

// MoqImportDir_genType_results holds the results of the ImportDir_genType type
type MoqImportDir_genType_results struct {
	Params  MoqImportDir_genType_params
	Results []struct {
		Values *struct {
			Result1 *build.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImportDir_genType_doFn
		DoReturnFn MoqImportDir_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqImportDir_genType_fnRecorder routes recorded function calls to the
// MoqImportDir_genType moq
type MoqImportDir_genType_fnRecorder struct {
	Params    MoqImportDir_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqImportDir_genType_results
	Moq       *MoqImportDir_genType
}

// MoqImportDir_genType_anyParams isolates the any params functions of the
// ImportDir_genType type
type MoqImportDir_genType_anyParams struct {
	Recorder *MoqImportDir_genType_fnRecorder
}

// NewMoqImportDir_genType creates a new moq of the ImportDir_genType type
func NewMoqImportDir_genType(scene *moq.Scene, config *moq.Config) *MoqImportDir_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqImportDir_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqImportDir_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dir  moq.ParamIndexing
				Mode moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dir  moq.ParamIndexing
			Mode moq.ParamIndexing
		}{
			Dir:  moq.ParamIndexByValue,
			Mode: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ImportDir_genType type
func (m *MoqImportDir_genType) Mock() ImportDir_genType {
	return func(dir string, mode build.ImportMode) (*build.Package, error) {
		m.Scene.T.Helper()
		moq := &MoqImportDir_genType_mock{Moq: m}
		return moq.Fn(dir, mode)
	}
}

func (m *MoqImportDir_genType_mock) Fn(dir string, mode build.ImportMode) (result1 *build.Package, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqImportDir_genType_params{
		Dir:  dir,
		Mode: mode,
	}
	var results *MoqImportDir_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(dir, mode)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(dir, mode)
	}
	return
}

func (m *MoqImportDir_genType) OnCall(dir string, mode build.ImportMode) *MoqImportDir_genType_fnRecorder {
	return &MoqImportDir_genType_fnRecorder{
		Params: MoqImportDir_genType_params{
			Dir:  dir,
			Mode: mode,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqImportDir_genType_fnRecorder) Any() *MoqImportDir_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqImportDir_genType_anyParams{Recorder: r}
}

func (a *MoqImportDir_genType_anyParams) Dir() *MoqImportDir_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqImportDir_genType_anyParams) Mode() *MoqImportDir_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqImportDir_genType_fnRecorder) Seq() *MoqImportDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqImportDir_genType_fnRecorder) NoSeq() *MoqImportDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqImportDir_genType_fnRecorder) ReturnResults(result1 *build.Package, result2 error) *MoqImportDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *build.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImportDir_genType_doFn
		DoReturnFn MoqImportDir_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *build.Package
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqImportDir_genType_fnRecorder) AndDo(fn MoqImportDir_genType_doFn) *MoqImportDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqImportDir_genType_fnRecorder) DoReturnResults(fn MoqImportDir_genType_doReturnFn) *MoqImportDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *build.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImportDir_genType_doFn
		DoReturnFn MoqImportDir_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqImportDir_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqImportDir_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqImportDir_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqImportDir_genType_paramsKey]*MoqImportDir_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqImportDir_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqImportDir_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqImportDir_genType_fnRecorder {
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
					Result1 *build.Package
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqImportDir_genType_doFn
				DoReturnFn MoqImportDir_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqImportDir_genType) PrettyParams(params MoqImportDir_genType_params) string {
	return fmt.Sprintf("ImportDir_genType(%#v, %#v)", params.Dir, params.Mode)
}

func (m *MoqImportDir_genType) ParamsKey(params MoqImportDir_genType_params, anyParams uint64) MoqImportDir_genType_paramsKey {
	m.Scene.T.Helper()
	var dirUsed string
	var dirUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dir == moq.ParamIndexByValue {
			dirUsed = params.Dir
		} else {
			dirUsedHash = hash.DeepHash(params.Dir)
		}
	}
	var modeUsed build.ImportMode
	var modeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Mode == moq.ParamIndexByValue {
			modeUsed = params.Mode
		} else {
			modeUsedHash = hash.DeepHash(params.Mode)
		}
	}
	return MoqImportDir_genType_paramsKey{
		Params: struct {
			Dir  string
			Mode build.ImportMode
		}{
			Dir:  dirUsed,
			Mode: modeUsed,
		},
		Hashes: struct {
			Dir  hash.Hash
			Mode hash.Hash
		}{
			Dir:  dirUsedHash,
			Mode: modeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqImportDir_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqImportDir_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}
