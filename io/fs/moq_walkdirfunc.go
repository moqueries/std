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

// MoqWalkDirFunc holds the state of a moq of the WalkDirFunc type
type MoqWalkDirFunc struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWalkDirFunc_mock

	ResultsByParams []MoqWalkDirFunc_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			D    moq.ParamIndexing
			Err  moq.ParamIndexing
		}
	}
}

// MoqWalkDirFunc_mock isolates the mock interface of the WalkDirFunc type
type MoqWalkDirFunc_mock struct {
	Moq *MoqWalkDirFunc
}

// MoqWalkDirFunc_params holds the params of the WalkDirFunc type
type MoqWalkDirFunc_params struct {
	Path string
	D    fs.DirEntry
	Err  error
}

// MoqWalkDirFunc_paramsKey holds the map key params of the WalkDirFunc type
type MoqWalkDirFunc_paramsKey struct {
	Params struct {
		Path string
		D    fs.DirEntry
		Err  error
	}
	Hashes struct {
		Path hash.Hash
		D    hash.Hash
		Err  hash.Hash
	}
}

// MoqWalkDirFunc_resultsByParams contains the results for a given set of
// parameters for the WalkDirFunc type
type MoqWalkDirFunc_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWalkDirFunc_paramsKey]*MoqWalkDirFunc_results
}

// MoqWalkDirFunc_doFn defines the type of function needed when calling AndDo
// for the WalkDirFunc type
type MoqWalkDirFunc_doFn func(path string, d fs.DirEntry, err error)

// MoqWalkDirFunc_doReturnFn defines the type of function needed when calling
// DoReturnResults for the WalkDirFunc type
type MoqWalkDirFunc_doReturnFn func(path string, d fs.DirEntry, err error) error

// MoqWalkDirFunc_results holds the results of the WalkDirFunc type
type MoqWalkDirFunc_results struct {
	Params  MoqWalkDirFunc_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWalkDirFunc_doFn
		DoReturnFn MoqWalkDirFunc_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWalkDirFunc_fnRecorder routes recorded function calls to the
// MoqWalkDirFunc moq
type MoqWalkDirFunc_fnRecorder struct {
	Params    MoqWalkDirFunc_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWalkDirFunc_results
	Moq       *MoqWalkDirFunc
}

// MoqWalkDirFunc_anyParams isolates the any params functions of the
// WalkDirFunc type
type MoqWalkDirFunc_anyParams struct {
	Recorder *MoqWalkDirFunc_fnRecorder
}

// NewMoqWalkDirFunc creates a new moq of the WalkDirFunc type
func NewMoqWalkDirFunc(scene *moq.Scene, config *moq.Config) *MoqWalkDirFunc {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWalkDirFunc{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWalkDirFunc_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				D    moq.ParamIndexing
				Err  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			D    moq.ParamIndexing
			Err  moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			D:    moq.ParamIndexByHash,
			Err:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the WalkDirFunc type
func (m *MoqWalkDirFunc) Mock() fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		m.Scene.T.Helper()
		moq := &MoqWalkDirFunc_mock{Moq: m}
		return moq.Fn(path, d, err)
	}
}

func (m *MoqWalkDirFunc_mock) Fn(path string, d fs.DirEntry, err error) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqWalkDirFunc_params{
		Path: path,
		D:    d,
		Err:  err,
	}
	var results *MoqWalkDirFunc_results
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
		result.DoFn(path, d, err)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(path, d, err)
	}
	return
}

func (m *MoqWalkDirFunc) OnCall(path string, d fs.DirEntry, err error) *MoqWalkDirFunc_fnRecorder {
	return &MoqWalkDirFunc_fnRecorder{
		Params: MoqWalkDirFunc_params{
			Path: path,
			D:    d,
			Err:  err,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqWalkDirFunc_fnRecorder) Any() *MoqWalkDirFunc_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqWalkDirFunc_anyParams{Recorder: r}
}

func (a *MoqWalkDirFunc_anyParams) Path() *MoqWalkDirFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqWalkDirFunc_anyParams) D() *MoqWalkDirFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqWalkDirFunc_anyParams) Err() *MoqWalkDirFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqWalkDirFunc_fnRecorder) Seq() *MoqWalkDirFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWalkDirFunc_fnRecorder) NoSeq() *MoqWalkDirFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWalkDirFunc_fnRecorder) ReturnResults(result1 error) *MoqWalkDirFunc_fnRecorder {
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
		DoFn       MoqWalkDirFunc_doFn
		DoReturnFn MoqWalkDirFunc_doReturnFn
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

func (r *MoqWalkDirFunc_fnRecorder) AndDo(fn MoqWalkDirFunc_doFn) *MoqWalkDirFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWalkDirFunc_fnRecorder) DoReturnResults(fn MoqWalkDirFunc_doReturnFn) *MoqWalkDirFunc_fnRecorder {
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
		DoFn       MoqWalkDirFunc_doFn
		DoReturnFn MoqWalkDirFunc_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWalkDirFunc_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWalkDirFunc_resultsByParams
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
		results = &MoqWalkDirFunc_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWalkDirFunc_paramsKey]*MoqWalkDirFunc_results{},
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
		r.Results = &MoqWalkDirFunc_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWalkDirFunc_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWalkDirFunc_fnRecorder {
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
				DoFn       MoqWalkDirFunc_doFn
				DoReturnFn MoqWalkDirFunc_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWalkDirFunc) PrettyParams(params MoqWalkDirFunc_params) string {
	return fmt.Sprintf("WalkDirFunc(%#v, %#v, %#v)", params.Path, params.D, params.Err)
}

func (m *MoqWalkDirFunc) ParamsKey(params MoqWalkDirFunc_params, anyParams uint64) MoqWalkDirFunc_paramsKey {
	m.Scene.T.Helper()
	var pathUsed string
	var pathUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Path == moq.ParamIndexByValue {
			pathUsed = params.Path
		} else {
			pathUsedHash = hash.DeepHash(params.Path)
		}
	}
	var dUsed fs.DirEntry
	var dUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.D == moq.ParamIndexByValue {
			dUsed = params.D
		} else {
			dUsedHash = hash.DeepHash(params.D)
		}
	}
	var errUsed error
	var errUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Err == moq.ParamIndexByValue {
			errUsed = params.Err
		} else {
			errUsedHash = hash.DeepHash(params.Err)
		}
	}
	return MoqWalkDirFunc_paramsKey{
		Params: struct {
			Path string
			D    fs.DirEntry
			Err  error
		}{
			Path: pathUsed,
			D:    dUsed,
			Err:  errUsed,
		},
		Hashes: struct {
			Path hash.Hash
			D    hash.Hash
			Err  hash.Hash
		}{
			Path: pathUsedHash,
			D:    dUsedHash,
			Err:  errUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWalkDirFunc) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWalkDirFunc) AssertExpectationsMet() {
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
