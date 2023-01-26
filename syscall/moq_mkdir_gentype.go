// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Mkdir_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Mkdir_genType func(path string, mode uint32) (err error)

// MoqMkdir_genType holds the state of a moq of the Mkdir_genType type
type MoqMkdir_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMkdir_genType_mock

	ResultsByParams []MoqMkdir_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Mode moq.ParamIndexing
		}
	}
}

// MoqMkdir_genType_mock isolates the mock interface of the Mkdir_genType type
type MoqMkdir_genType_mock struct {
	Moq *MoqMkdir_genType
}

// MoqMkdir_genType_params holds the params of the Mkdir_genType type
type MoqMkdir_genType_params struct {
	Path string
	Mode uint32
}

// MoqMkdir_genType_paramsKey holds the map key params of the Mkdir_genType
// type
type MoqMkdir_genType_paramsKey struct {
	Params struct {
		Path string
		Mode uint32
	}
	Hashes struct {
		Path hash.Hash
		Mode hash.Hash
	}
}

// MoqMkdir_genType_resultsByParams contains the results for a given set of
// parameters for the Mkdir_genType type
type MoqMkdir_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMkdir_genType_paramsKey]*MoqMkdir_genType_results
}

// MoqMkdir_genType_doFn defines the type of function needed when calling AndDo
// for the Mkdir_genType type
type MoqMkdir_genType_doFn func(path string, mode uint32)

// MoqMkdir_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Mkdir_genType type
type MoqMkdir_genType_doReturnFn func(path string, mode uint32) (err error)

// MoqMkdir_genType_results holds the results of the Mkdir_genType type
type MoqMkdir_genType_results struct {
	Params  MoqMkdir_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMkdir_genType_doFn
		DoReturnFn MoqMkdir_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMkdir_genType_fnRecorder routes recorded function calls to the
// MoqMkdir_genType moq
type MoqMkdir_genType_fnRecorder struct {
	Params    MoqMkdir_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMkdir_genType_results
	Moq       *MoqMkdir_genType
}

// MoqMkdir_genType_anyParams isolates the any params functions of the
// Mkdir_genType type
type MoqMkdir_genType_anyParams struct {
	Recorder *MoqMkdir_genType_fnRecorder
}

// NewMoqMkdir_genType creates a new moq of the Mkdir_genType type
func NewMoqMkdir_genType(scene *moq.Scene, config *moq.Config) *MoqMkdir_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMkdir_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMkdir_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Mode moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Mode moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Mode: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Mkdir_genType type
func (m *MoqMkdir_genType) Mock() Mkdir_genType {
	return func(path string, mode uint32) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqMkdir_genType_mock{Moq: m}
		return moq.Fn(path, mode)
	}
}

func (m *MoqMkdir_genType_mock) Fn(path string, mode uint32) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqMkdir_genType_params{
		Path: path,
		Mode: mode,
	}
	var results *MoqMkdir_genType_results
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
		result.DoFn(path, mode)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(path, mode)
	}
	return
}

func (m *MoqMkdir_genType) OnCall(path string, mode uint32) *MoqMkdir_genType_fnRecorder {
	return &MoqMkdir_genType_fnRecorder{
		Params: MoqMkdir_genType_params{
			Path: path,
			Mode: mode,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMkdir_genType_fnRecorder) Any() *MoqMkdir_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMkdir_genType_anyParams{Recorder: r}
}

func (a *MoqMkdir_genType_anyParams) Path() *MoqMkdir_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqMkdir_genType_anyParams) Mode() *MoqMkdir_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqMkdir_genType_fnRecorder) Seq() *MoqMkdir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMkdir_genType_fnRecorder) NoSeq() *MoqMkdir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMkdir_genType_fnRecorder) ReturnResults(err error) *MoqMkdir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMkdir_genType_doFn
		DoReturnFn MoqMkdir_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMkdir_genType_fnRecorder) AndDo(fn MoqMkdir_genType_doFn) *MoqMkdir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMkdir_genType_fnRecorder) DoReturnResults(fn MoqMkdir_genType_doReturnFn) *MoqMkdir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMkdir_genType_doFn
		DoReturnFn MoqMkdir_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMkdir_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMkdir_genType_resultsByParams
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
		results = &MoqMkdir_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMkdir_genType_paramsKey]*MoqMkdir_genType_results{},
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
		r.Results = &MoqMkdir_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMkdir_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMkdir_genType_fnRecorder {
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
				Values     *struct{ Err error }
				Sequence   uint32
				DoFn       MoqMkdir_genType_doFn
				DoReturnFn MoqMkdir_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMkdir_genType) PrettyParams(params MoqMkdir_genType_params) string {
	return fmt.Sprintf("Mkdir_genType(%#v, %#v)", params.Path, params.Mode)
}

func (m *MoqMkdir_genType) ParamsKey(params MoqMkdir_genType_params, anyParams uint64) MoqMkdir_genType_paramsKey {
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
	var modeUsed uint32
	var modeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Mode == moq.ParamIndexByValue {
			modeUsed = params.Mode
		} else {
			modeUsedHash = hash.DeepHash(params.Mode)
		}
	}
	return MoqMkdir_genType_paramsKey{
		Params: struct {
			Path string
			Mode uint32
		}{
			Path: pathUsed,
			Mode: modeUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Mode hash.Hash
		}{
			Path: pathUsedHash,
			Mode: modeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMkdir_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMkdir_genType) AssertExpectationsMet() {
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