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

// Chmod_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Chmod_genType func(name string, mode os.FileMode) error

// MoqChmod_genType holds the state of a moq of the Chmod_genType type
type MoqChmod_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqChmod_genType_mock

	ResultsByParams []MoqChmod_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name moq.ParamIndexing
			Mode moq.ParamIndexing
		}
	}
}

// MoqChmod_genType_mock isolates the mock interface of the Chmod_genType type
type MoqChmod_genType_mock struct {
	Moq *MoqChmod_genType
}

// MoqChmod_genType_params holds the params of the Chmod_genType type
type MoqChmod_genType_params struct {
	Name string
	Mode os.FileMode
}

// MoqChmod_genType_paramsKey holds the map key params of the Chmod_genType
// type
type MoqChmod_genType_paramsKey struct {
	Params struct {
		Name string
		Mode os.FileMode
	}
	Hashes struct {
		Name hash.Hash
		Mode hash.Hash
	}
}

// MoqChmod_genType_resultsByParams contains the results for a given set of
// parameters for the Chmod_genType type
type MoqChmod_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqChmod_genType_paramsKey]*MoqChmod_genType_results
}

// MoqChmod_genType_doFn defines the type of function needed when calling AndDo
// for the Chmod_genType type
type MoqChmod_genType_doFn func(name string, mode os.FileMode)

// MoqChmod_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Chmod_genType type
type MoqChmod_genType_doReturnFn func(name string, mode os.FileMode) error

// MoqChmod_genType_results holds the results of the Chmod_genType type
type MoqChmod_genType_results struct {
	Params  MoqChmod_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqChmod_genType_doFn
		DoReturnFn MoqChmod_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqChmod_genType_fnRecorder routes recorded function calls to the
// MoqChmod_genType moq
type MoqChmod_genType_fnRecorder struct {
	Params    MoqChmod_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqChmod_genType_results
	Moq       *MoqChmod_genType
}

// MoqChmod_genType_anyParams isolates the any params functions of the
// Chmod_genType type
type MoqChmod_genType_anyParams struct {
	Recorder *MoqChmod_genType_fnRecorder
}

// NewMoqChmod_genType creates a new moq of the Chmod_genType type
func NewMoqChmod_genType(scene *moq.Scene, config *moq.Config) *MoqChmod_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqChmod_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqChmod_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Name moq.ParamIndexing
				Mode moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Name moq.ParamIndexing
			Mode moq.ParamIndexing
		}{
			Name: moq.ParamIndexByValue,
			Mode: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Chmod_genType type
func (m *MoqChmod_genType) Mock() Chmod_genType {
	return func(name string, mode os.FileMode) error {
		m.Scene.T.Helper()
		moq := &MoqChmod_genType_mock{Moq: m}
		return moq.Fn(name, mode)
	}
}

func (m *MoqChmod_genType_mock) Fn(name string, mode os.FileMode) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqChmod_genType_params{
		Name: name,
		Mode: mode,
	}
	var results *MoqChmod_genType_results
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
		result.DoFn(name, mode)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(name, mode)
	}
	return
}

func (m *MoqChmod_genType) OnCall(name string, mode os.FileMode) *MoqChmod_genType_fnRecorder {
	return &MoqChmod_genType_fnRecorder{
		Params: MoqChmod_genType_params{
			Name: name,
			Mode: mode,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqChmod_genType_fnRecorder) Any() *MoqChmod_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqChmod_genType_anyParams{Recorder: r}
}

func (a *MoqChmod_genType_anyParams) Name() *MoqChmod_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqChmod_genType_anyParams) Mode() *MoqChmod_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqChmod_genType_fnRecorder) Seq() *MoqChmod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqChmod_genType_fnRecorder) NoSeq() *MoqChmod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqChmod_genType_fnRecorder) ReturnResults(result1 error) *MoqChmod_genType_fnRecorder {
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
		DoFn       MoqChmod_genType_doFn
		DoReturnFn MoqChmod_genType_doReturnFn
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

func (r *MoqChmod_genType_fnRecorder) AndDo(fn MoqChmod_genType_doFn) *MoqChmod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqChmod_genType_fnRecorder) DoReturnResults(fn MoqChmod_genType_doReturnFn) *MoqChmod_genType_fnRecorder {
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
		DoFn       MoqChmod_genType_doFn
		DoReturnFn MoqChmod_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqChmod_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqChmod_genType_resultsByParams
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
		results = &MoqChmod_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqChmod_genType_paramsKey]*MoqChmod_genType_results{},
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
		r.Results = &MoqChmod_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqChmod_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqChmod_genType_fnRecorder {
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
				DoFn       MoqChmod_genType_doFn
				DoReturnFn MoqChmod_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqChmod_genType) PrettyParams(params MoqChmod_genType_params) string {
	return fmt.Sprintf("Chmod_genType(%#v, %#v)", params.Name, params.Mode)
}

func (m *MoqChmod_genType) ParamsKey(params MoqChmod_genType_params, anyParams uint64) MoqChmod_genType_paramsKey {
	m.Scene.T.Helper()
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	var modeUsed os.FileMode
	var modeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Mode == moq.ParamIndexByValue {
			modeUsed = params.Mode
		} else {
			modeUsedHash = hash.DeepHash(params.Mode)
		}
	}
	return MoqChmod_genType_paramsKey{
		Params: struct {
			Name string
			Mode os.FileMode
		}{
			Name: nameUsed,
			Mode: modeUsed,
		},
		Hashes: struct {
			Name hash.Hash
			Mode hash.Hash
		}{
			Name: nameUsedHash,
			Mode: modeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqChmod_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqChmod_genType) AssertExpectationsMet() {
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
