// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package filepath

import (
	"fmt"
	"math/bits"
	"os"
	"path/filepath"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MoqWalkFunc holds the state of a moq of the WalkFunc type
type MoqWalkFunc struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWalkFunc_mock

	ResultsByParams []MoqWalkFunc_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Info moq.ParamIndexing
			Err  moq.ParamIndexing
		}
	}
}

// MoqWalkFunc_mock isolates the mock interface of the WalkFunc type
type MoqWalkFunc_mock struct {
	Moq *MoqWalkFunc
}

// MoqWalkFunc_params holds the params of the WalkFunc type
type MoqWalkFunc_params struct {
	Path string
	Info os.FileInfo
	Err  error
}

// MoqWalkFunc_paramsKey holds the map key params of the WalkFunc type
type MoqWalkFunc_paramsKey struct {
	Params struct {
		Path string
		Info os.FileInfo
		Err  error
	}
	Hashes struct {
		Path hash.Hash
		Info hash.Hash
		Err  hash.Hash
	}
}

// MoqWalkFunc_resultsByParams contains the results for a given set of
// parameters for the WalkFunc type
type MoqWalkFunc_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWalkFunc_paramsKey]*MoqWalkFunc_results
}

// MoqWalkFunc_doFn defines the type of function needed when calling AndDo for
// the WalkFunc type
type MoqWalkFunc_doFn func(path string, info os.FileInfo, err error)

// MoqWalkFunc_doReturnFn defines the type of function needed when calling
// DoReturnResults for the WalkFunc type
type MoqWalkFunc_doReturnFn func(path string, info os.FileInfo, err error) error

// MoqWalkFunc_results holds the results of the WalkFunc type
type MoqWalkFunc_results struct {
	Params  MoqWalkFunc_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWalkFunc_doFn
		DoReturnFn MoqWalkFunc_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWalkFunc_fnRecorder routes recorded function calls to the MoqWalkFunc moq
type MoqWalkFunc_fnRecorder struct {
	Params    MoqWalkFunc_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWalkFunc_results
	Moq       *MoqWalkFunc
}

// MoqWalkFunc_anyParams isolates the any params functions of the WalkFunc type
type MoqWalkFunc_anyParams struct {
	Recorder *MoqWalkFunc_fnRecorder
}

// NewMoqWalkFunc creates a new moq of the WalkFunc type
func NewMoqWalkFunc(scene *moq.Scene, config *moq.Config) *MoqWalkFunc {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWalkFunc{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWalkFunc_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Info moq.ParamIndexing
				Err  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Info moq.ParamIndexing
			Err  moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Info: moq.ParamIndexByHash,
			Err:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the WalkFunc type
func (m *MoqWalkFunc) Mock() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		m.Scene.T.Helper()
		moq := &MoqWalkFunc_mock{Moq: m}
		return moq.Fn(path, info, err)
	}
}

func (m *MoqWalkFunc_mock) Fn(path string, info os.FileInfo, err error) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqWalkFunc_params{
		Path: path,
		Info: info,
		Err:  err,
	}
	var results *MoqWalkFunc_results
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
		result.DoFn(path, info, err)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(path, info, err)
	}
	return
}

func (m *MoqWalkFunc) OnCall(path string, info os.FileInfo, err error) *MoqWalkFunc_fnRecorder {
	return &MoqWalkFunc_fnRecorder{
		Params: MoqWalkFunc_params{
			Path: path,
			Info: info,
			Err:  err,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqWalkFunc_fnRecorder) Any() *MoqWalkFunc_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqWalkFunc_anyParams{Recorder: r}
}

func (a *MoqWalkFunc_anyParams) Path() *MoqWalkFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqWalkFunc_anyParams) Info() *MoqWalkFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqWalkFunc_anyParams) Err() *MoqWalkFunc_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqWalkFunc_fnRecorder) Seq() *MoqWalkFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWalkFunc_fnRecorder) NoSeq() *MoqWalkFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWalkFunc_fnRecorder) ReturnResults(result1 error) *MoqWalkFunc_fnRecorder {
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
		DoFn       MoqWalkFunc_doFn
		DoReturnFn MoqWalkFunc_doReturnFn
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

func (r *MoqWalkFunc_fnRecorder) AndDo(fn MoqWalkFunc_doFn) *MoqWalkFunc_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWalkFunc_fnRecorder) DoReturnResults(fn MoqWalkFunc_doReturnFn) *MoqWalkFunc_fnRecorder {
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
		DoFn       MoqWalkFunc_doFn
		DoReturnFn MoqWalkFunc_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWalkFunc_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWalkFunc_resultsByParams
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
		results = &MoqWalkFunc_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWalkFunc_paramsKey]*MoqWalkFunc_results{},
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
		r.Results = &MoqWalkFunc_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWalkFunc_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWalkFunc_fnRecorder {
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
				DoFn       MoqWalkFunc_doFn
				DoReturnFn MoqWalkFunc_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWalkFunc) PrettyParams(params MoqWalkFunc_params) string {
	return fmt.Sprintf("WalkFunc(%#v, %#v, %#v)", params.Path, params.Info, params.Err)
}

func (m *MoqWalkFunc) ParamsKey(params MoqWalkFunc_params, anyParams uint64) MoqWalkFunc_paramsKey {
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
	var infoUsed os.FileInfo
	var infoUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Info == moq.ParamIndexByValue {
			infoUsed = params.Info
		} else {
			infoUsedHash = hash.DeepHash(params.Info)
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
	return MoqWalkFunc_paramsKey{
		Params: struct {
			Path string
			Info os.FileInfo
			Err  error
		}{
			Path: pathUsed,
			Info: infoUsed,
			Err:  errUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Info hash.Hash
			Err  hash.Hash
		}{
			Path: pathUsedHash,
			Info: infoUsedHash,
			Err:  errUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWalkFunc) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWalkFunc) AssertExpectationsMet() {
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