// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Symlink_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Symlink_genType func(oldpath string, newpath string) (err error)

// MoqSymlink_genType holds the state of a moq of the Symlink_genType type
type MoqSymlink_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSymlink_genType_mock

	ResultsByParams []MoqSymlink_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Oldpath moq.ParamIndexing
			Newpath moq.ParamIndexing
		}
	}
}

// MoqSymlink_genType_mock isolates the mock interface of the Symlink_genType
// type
type MoqSymlink_genType_mock struct {
	Moq *MoqSymlink_genType
}

// MoqSymlink_genType_params holds the params of the Symlink_genType type
type MoqSymlink_genType_params struct {
	Oldpath string
	Newpath string
}

// MoqSymlink_genType_paramsKey holds the map key params of the Symlink_genType
// type
type MoqSymlink_genType_paramsKey struct {
	Params struct {
		Oldpath string
		Newpath string
	}
	Hashes struct {
		Oldpath hash.Hash
		Newpath hash.Hash
	}
}

// MoqSymlink_genType_resultsByParams contains the results for a given set of
// parameters for the Symlink_genType type
type MoqSymlink_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSymlink_genType_paramsKey]*MoqSymlink_genType_results
}

// MoqSymlink_genType_doFn defines the type of function needed when calling
// AndDo for the Symlink_genType type
type MoqSymlink_genType_doFn func(oldpath string, newpath string)

// MoqSymlink_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Symlink_genType type
type MoqSymlink_genType_doReturnFn func(oldpath string, newpath string) (err error)

// MoqSymlink_genType_results holds the results of the Symlink_genType type
type MoqSymlink_genType_results struct {
	Params  MoqSymlink_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSymlink_genType_doFn
		DoReturnFn MoqSymlink_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSymlink_genType_fnRecorder routes recorded function calls to the
// MoqSymlink_genType moq
type MoqSymlink_genType_fnRecorder struct {
	Params    MoqSymlink_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSymlink_genType_results
	Moq       *MoqSymlink_genType
}

// MoqSymlink_genType_anyParams isolates the any params functions of the
// Symlink_genType type
type MoqSymlink_genType_anyParams struct {
	Recorder *MoqSymlink_genType_fnRecorder
}

// NewMoqSymlink_genType creates a new moq of the Symlink_genType type
func NewMoqSymlink_genType(scene *moq.Scene, config *moq.Config) *MoqSymlink_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSymlink_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSymlink_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Oldpath moq.ParamIndexing
				Newpath moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Oldpath moq.ParamIndexing
			Newpath moq.ParamIndexing
		}{
			Oldpath: moq.ParamIndexByValue,
			Newpath: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Symlink_genType type
func (m *MoqSymlink_genType) Mock() Symlink_genType {
	return func(oldpath string, newpath string) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqSymlink_genType_mock{Moq: m}
		return moq.Fn(oldpath, newpath)
	}
}

func (m *MoqSymlink_genType_mock) Fn(oldpath string, newpath string) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqSymlink_genType_params{
		Oldpath: oldpath,
		Newpath: newpath,
	}
	var results *MoqSymlink_genType_results
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
		result.DoFn(oldpath, newpath)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(oldpath, newpath)
	}
	return
}

func (m *MoqSymlink_genType) OnCall(oldpath string, newpath string) *MoqSymlink_genType_fnRecorder {
	return &MoqSymlink_genType_fnRecorder{
		Params: MoqSymlink_genType_params{
			Oldpath: oldpath,
			Newpath: newpath,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSymlink_genType_fnRecorder) Any() *MoqSymlink_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSymlink_genType_anyParams{Recorder: r}
}

func (a *MoqSymlink_genType_anyParams) Oldpath() *MoqSymlink_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSymlink_genType_anyParams) Newpath() *MoqSymlink_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSymlink_genType_fnRecorder) Seq() *MoqSymlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSymlink_genType_fnRecorder) NoSeq() *MoqSymlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSymlink_genType_fnRecorder) ReturnResults(err error) *MoqSymlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSymlink_genType_doFn
		DoReturnFn MoqSymlink_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSymlink_genType_fnRecorder) AndDo(fn MoqSymlink_genType_doFn) *MoqSymlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSymlink_genType_fnRecorder) DoReturnResults(fn MoqSymlink_genType_doReturnFn) *MoqSymlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSymlink_genType_doFn
		DoReturnFn MoqSymlink_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSymlink_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSymlink_genType_resultsByParams
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
		results = &MoqSymlink_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSymlink_genType_paramsKey]*MoqSymlink_genType_results{},
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
		r.Results = &MoqSymlink_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSymlink_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSymlink_genType_fnRecorder {
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
				DoFn       MoqSymlink_genType_doFn
				DoReturnFn MoqSymlink_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSymlink_genType) PrettyParams(params MoqSymlink_genType_params) string {
	return fmt.Sprintf("Symlink_genType(%#v, %#v)", params.Oldpath, params.Newpath)
}

func (m *MoqSymlink_genType) ParamsKey(params MoqSymlink_genType_params, anyParams uint64) MoqSymlink_genType_paramsKey {
	m.Scene.T.Helper()
	var oldpathUsed string
	var oldpathUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Oldpath == moq.ParamIndexByValue {
			oldpathUsed = params.Oldpath
		} else {
			oldpathUsedHash = hash.DeepHash(params.Oldpath)
		}
	}
	var newpathUsed string
	var newpathUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Newpath == moq.ParamIndexByValue {
			newpathUsed = params.Newpath
		} else {
			newpathUsedHash = hash.DeepHash(params.Newpath)
		}
	}
	return MoqSymlink_genType_paramsKey{
		Params: struct {
			Oldpath string
			Newpath string
		}{
			Oldpath: oldpathUsed,
			Newpath: newpathUsed,
		},
		Hashes: struct {
			Oldpath hash.Hash
			Newpath hash.Hash
		}{
			Oldpath: oldpathUsedHash,
			Newpath: newpathUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSymlink_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSymlink_genType) AssertExpectationsMet() {
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