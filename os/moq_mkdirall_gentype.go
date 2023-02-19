// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"math/bits"
	"os"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MkdirAll_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type MkdirAll_genType func(path string, perm os.FileMode) error

// MoqMkdirAll_genType holds the state of a moq of the MkdirAll_genType type
type MoqMkdirAll_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMkdirAll_genType_mock

	ResultsByParams []MoqMkdirAll_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Perm moq.ParamIndexing
		}
	}
}

// MoqMkdirAll_genType_mock isolates the mock interface of the MkdirAll_genType
// type
type MoqMkdirAll_genType_mock struct {
	Moq *MoqMkdirAll_genType
}

// MoqMkdirAll_genType_params holds the params of the MkdirAll_genType type
type MoqMkdirAll_genType_params struct {
	Path string
	Perm os.FileMode
}

// MoqMkdirAll_genType_paramsKey holds the map key params of the
// MkdirAll_genType type
type MoqMkdirAll_genType_paramsKey struct {
	Params struct {
		Path string
		Perm os.FileMode
	}
	Hashes struct {
		Path hash.Hash
		Perm hash.Hash
	}
}

// MoqMkdirAll_genType_resultsByParams contains the results for a given set of
// parameters for the MkdirAll_genType type
type MoqMkdirAll_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMkdirAll_genType_paramsKey]*MoqMkdirAll_genType_results
}

// MoqMkdirAll_genType_doFn defines the type of function needed when calling
// AndDo for the MkdirAll_genType type
type MoqMkdirAll_genType_doFn func(path string, perm os.FileMode)

// MoqMkdirAll_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the MkdirAll_genType type
type MoqMkdirAll_genType_doReturnFn func(path string, perm os.FileMode) error

// MoqMkdirAll_genType_results holds the results of the MkdirAll_genType type
type MoqMkdirAll_genType_results struct {
	Params  MoqMkdirAll_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqMkdirAll_genType_doFn
		DoReturnFn MoqMkdirAll_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMkdirAll_genType_fnRecorder routes recorded function calls to the
// MoqMkdirAll_genType moq
type MoqMkdirAll_genType_fnRecorder struct {
	Params    MoqMkdirAll_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMkdirAll_genType_results
	Moq       *MoqMkdirAll_genType
}

// MoqMkdirAll_genType_anyParams isolates the any params functions of the
// MkdirAll_genType type
type MoqMkdirAll_genType_anyParams struct {
	Recorder *MoqMkdirAll_genType_fnRecorder
}

// NewMoqMkdirAll_genType creates a new moq of the MkdirAll_genType type
func NewMoqMkdirAll_genType(scene *moq.Scene, config *moq.Config) *MoqMkdirAll_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMkdirAll_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMkdirAll_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Perm moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Perm moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Perm: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the MkdirAll_genType type
func (m *MoqMkdirAll_genType) Mock() MkdirAll_genType {
	return func(path string, perm os.FileMode) error {
		m.Scene.T.Helper()
		moq := &MoqMkdirAll_genType_mock{Moq: m}
		return moq.Fn(path, perm)
	}
}

func (m *MoqMkdirAll_genType_mock) Fn(path string, perm os.FileMode) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqMkdirAll_genType_params{
		Path: path,
		Perm: perm,
	}
	var results *MoqMkdirAll_genType_results
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
		result.DoFn(path, perm)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(path, perm)
	}
	return
}

func (m *MoqMkdirAll_genType) OnCall(path string, perm os.FileMode) *MoqMkdirAll_genType_fnRecorder {
	return &MoqMkdirAll_genType_fnRecorder{
		Params: MoqMkdirAll_genType_params{
			Path: path,
			Perm: perm,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMkdirAll_genType_fnRecorder) Any() *MoqMkdirAll_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMkdirAll_genType_anyParams{Recorder: r}
}

func (a *MoqMkdirAll_genType_anyParams) Path() *MoqMkdirAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqMkdirAll_genType_anyParams) Perm() *MoqMkdirAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqMkdirAll_genType_fnRecorder) Seq() *MoqMkdirAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMkdirAll_genType_fnRecorder) NoSeq() *MoqMkdirAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMkdirAll_genType_fnRecorder) ReturnResults(result1 error) *MoqMkdirAll_genType_fnRecorder {
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
		DoFn       MoqMkdirAll_genType_doFn
		DoReturnFn MoqMkdirAll_genType_doReturnFn
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

func (r *MoqMkdirAll_genType_fnRecorder) AndDo(fn MoqMkdirAll_genType_doFn) *MoqMkdirAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMkdirAll_genType_fnRecorder) DoReturnResults(fn MoqMkdirAll_genType_doReturnFn) *MoqMkdirAll_genType_fnRecorder {
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
		DoFn       MoqMkdirAll_genType_doFn
		DoReturnFn MoqMkdirAll_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMkdirAll_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMkdirAll_genType_resultsByParams
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
		results = &MoqMkdirAll_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMkdirAll_genType_paramsKey]*MoqMkdirAll_genType_results{},
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
		r.Results = &MoqMkdirAll_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMkdirAll_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMkdirAll_genType_fnRecorder {
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
				DoFn       MoqMkdirAll_genType_doFn
				DoReturnFn MoqMkdirAll_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMkdirAll_genType) PrettyParams(params MoqMkdirAll_genType_params) string {
	return fmt.Sprintf("MkdirAll_genType(%#v, %#v)", params.Path, params.Perm)
}

func (m *MoqMkdirAll_genType) ParamsKey(params MoqMkdirAll_genType_params, anyParams uint64) MoqMkdirAll_genType_paramsKey {
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
	var permUsed os.FileMode
	var permUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Perm == moq.ParamIndexByValue {
			permUsed = params.Perm
		} else {
			permUsedHash = hash.DeepHash(params.Perm)
		}
	}
	return MoqMkdirAll_genType_paramsKey{
		Params: struct {
			Path string
			Perm os.FileMode
		}{
			Path: pathUsed,
			Perm: permUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Perm hash.Hash
		}{
			Path: pathUsedHash,
			Perm: permUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMkdirAll_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMkdirAll_genType) AssertExpectationsMet() {
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
