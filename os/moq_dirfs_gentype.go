// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"io/fs"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// DirFS_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type DirFS_genType func(dir string) fs.FS

// MoqDirFS_genType holds the state of a moq of the DirFS_genType type
type MoqDirFS_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDirFS_genType_mock

	ResultsByParams []MoqDirFS_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dir moq.ParamIndexing
		}
	}
}

// MoqDirFS_genType_mock isolates the mock interface of the DirFS_genType type
type MoqDirFS_genType_mock struct {
	Moq *MoqDirFS_genType
}

// MoqDirFS_genType_params holds the params of the DirFS_genType type
type MoqDirFS_genType_params struct{ Dir string }

// MoqDirFS_genType_paramsKey holds the map key params of the DirFS_genType
// type
type MoqDirFS_genType_paramsKey struct {
	Params struct{ Dir string }
	Hashes struct{ Dir hash.Hash }
}

// MoqDirFS_genType_resultsByParams contains the results for a given set of
// parameters for the DirFS_genType type
type MoqDirFS_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDirFS_genType_paramsKey]*MoqDirFS_genType_results
}

// MoqDirFS_genType_doFn defines the type of function needed when calling AndDo
// for the DirFS_genType type
type MoqDirFS_genType_doFn func(dir string)

// MoqDirFS_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the DirFS_genType type
type MoqDirFS_genType_doReturnFn func(dir string) fs.FS

// MoqDirFS_genType_results holds the results of the DirFS_genType type
type MoqDirFS_genType_results struct {
	Params  MoqDirFS_genType_params
	Results []struct {
		Values *struct {
			Result1 fs.FS
		}
		Sequence   uint32
		DoFn       MoqDirFS_genType_doFn
		DoReturnFn MoqDirFS_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDirFS_genType_fnRecorder routes recorded function calls to the
// MoqDirFS_genType moq
type MoqDirFS_genType_fnRecorder struct {
	Params    MoqDirFS_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDirFS_genType_results
	Moq       *MoqDirFS_genType
}

// MoqDirFS_genType_anyParams isolates the any params functions of the
// DirFS_genType type
type MoqDirFS_genType_anyParams struct {
	Recorder *MoqDirFS_genType_fnRecorder
}

// NewMoqDirFS_genType creates a new moq of the DirFS_genType type
func NewMoqDirFS_genType(scene *moq.Scene, config *moq.Config) *MoqDirFS_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDirFS_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDirFS_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dir moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dir moq.ParamIndexing
		}{
			Dir: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DirFS_genType type
func (m *MoqDirFS_genType) Mock() DirFS_genType {
	return func(dir string) fs.FS { m.Scene.T.Helper(); moq := &MoqDirFS_genType_mock{Moq: m}; return moq.Fn(dir) }
}

func (m *MoqDirFS_genType_mock) Fn(dir string) (result1 fs.FS) {
	m.Moq.Scene.T.Helper()
	params := MoqDirFS_genType_params{
		Dir: dir,
	}
	var results *MoqDirFS_genType_results
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
		result.DoFn(dir)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(dir)
	}
	return
}

func (m *MoqDirFS_genType) OnCall(dir string) *MoqDirFS_genType_fnRecorder {
	return &MoqDirFS_genType_fnRecorder{
		Params: MoqDirFS_genType_params{
			Dir: dir,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDirFS_genType_fnRecorder) Any() *MoqDirFS_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDirFS_genType_anyParams{Recorder: r}
}

func (a *MoqDirFS_genType_anyParams) Dir() *MoqDirFS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqDirFS_genType_fnRecorder) Seq() *MoqDirFS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDirFS_genType_fnRecorder) NoSeq() *MoqDirFS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDirFS_genType_fnRecorder) ReturnResults(result1 fs.FS) *MoqDirFS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FS
		}
		Sequence   uint32
		DoFn       MoqDirFS_genType_doFn
		DoReturnFn MoqDirFS_genType_doReturnFn
	}{
		Values: &struct {
			Result1 fs.FS
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDirFS_genType_fnRecorder) AndDo(fn MoqDirFS_genType_doFn) *MoqDirFS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDirFS_genType_fnRecorder) DoReturnResults(fn MoqDirFS_genType_doReturnFn) *MoqDirFS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FS
		}
		Sequence   uint32
		DoFn       MoqDirFS_genType_doFn
		DoReturnFn MoqDirFS_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDirFS_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDirFS_genType_resultsByParams
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
		results = &MoqDirFS_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDirFS_genType_paramsKey]*MoqDirFS_genType_results{},
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
		r.Results = &MoqDirFS_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDirFS_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDirFS_genType_fnRecorder {
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
					Result1 fs.FS
				}
				Sequence   uint32
				DoFn       MoqDirFS_genType_doFn
				DoReturnFn MoqDirFS_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDirFS_genType) PrettyParams(params MoqDirFS_genType_params) string {
	return fmt.Sprintf("DirFS_genType(%#v)", params.Dir)
}

func (m *MoqDirFS_genType) ParamsKey(params MoqDirFS_genType_params, anyParams uint64) MoqDirFS_genType_paramsKey {
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
	return MoqDirFS_genType_paramsKey{
		Params: struct{ Dir string }{
			Dir: dirUsed,
		},
		Hashes: struct{ Dir hash.Hash }{
			Dir: dirUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDirFS_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDirFS_genType) AssertExpectationsMet() {
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
