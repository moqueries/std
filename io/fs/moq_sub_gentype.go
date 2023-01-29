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

// Sub_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Sub_genType func(fsys fs.FS, dir string) (fs.FS, error)

// MoqSub_genType holds the state of a moq of the Sub_genType type
type MoqSub_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSub_genType_mock

	ResultsByParams []MoqSub_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fsys moq.ParamIndexing
			Dir  moq.ParamIndexing
		}
	}
}

// MoqSub_genType_mock isolates the mock interface of the Sub_genType type
type MoqSub_genType_mock struct {
	Moq *MoqSub_genType
}

// MoqSub_genType_params holds the params of the Sub_genType type
type MoqSub_genType_params struct {
	Fsys fs.FS
	Dir  string
}

// MoqSub_genType_paramsKey holds the map key params of the Sub_genType type
type MoqSub_genType_paramsKey struct {
	Params struct {
		Fsys fs.FS
		Dir  string
	}
	Hashes struct {
		Fsys hash.Hash
		Dir  hash.Hash
	}
}

// MoqSub_genType_resultsByParams contains the results for a given set of
// parameters for the Sub_genType type
type MoqSub_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSub_genType_paramsKey]*MoqSub_genType_results
}

// MoqSub_genType_doFn defines the type of function needed when calling AndDo
// for the Sub_genType type
type MoqSub_genType_doFn func(fsys fs.FS, dir string)

// MoqSub_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Sub_genType type
type MoqSub_genType_doReturnFn func(fsys fs.FS, dir string) (fs.FS, error)

// MoqSub_genType_results holds the results of the Sub_genType type
type MoqSub_genType_results struct {
	Params  MoqSub_genType_params
	Results []struct {
		Values *struct {
			Result1 fs.FS
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSub_genType_doFn
		DoReturnFn MoqSub_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSub_genType_fnRecorder routes recorded function calls to the
// MoqSub_genType moq
type MoqSub_genType_fnRecorder struct {
	Params    MoqSub_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSub_genType_results
	Moq       *MoqSub_genType
}

// MoqSub_genType_anyParams isolates the any params functions of the
// Sub_genType type
type MoqSub_genType_anyParams struct {
	Recorder *MoqSub_genType_fnRecorder
}

// NewMoqSub_genType creates a new moq of the Sub_genType type
func NewMoqSub_genType(scene *moq.Scene, config *moq.Config) *MoqSub_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSub_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSub_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fsys moq.ParamIndexing
				Dir  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fsys moq.ParamIndexing
			Dir  moq.ParamIndexing
		}{
			Fsys: moq.ParamIndexByHash,
			Dir:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Sub_genType type
func (m *MoqSub_genType) Mock() Sub_genType {
	return func(fsys fs.FS, dir string) (fs.FS, error) {
		m.Scene.T.Helper()
		moq := &MoqSub_genType_mock{Moq: m}
		return moq.Fn(fsys, dir)
	}
}

func (m *MoqSub_genType_mock) Fn(fsys fs.FS, dir string) (result1 fs.FS, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSub_genType_params{
		Fsys: fsys,
		Dir:  dir,
	}
	var results *MoqSub_genType_results
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
		result.DoFn(fsys, dir)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(fsys, dir)
	}
	return
}

func (m *MoqSub_genType) OnCall(fsys fs.FS, dir string) *MoqSub_genType_fnRecorder {
	return &MoqSub_genType_fnRecorder{
		Params: MoqSub_genType_params{
			Fsys: fsys,
			Dir:  dir,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSub_genType_fnRecorder) Any() *MoqSub_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSub_genType_anyParams{Recorder: r}
}

func (a *MoqSub_genType_anyParams) Fsys() *MoqSub_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSub_genType_anyParams) Dir() *MoqSub_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSub_genType_fnRecorder) Seq() *MoqSub_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSub_genType_fnRecorder) NoSeq() *MoqSub_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSub_genType_fnRecorder) ReturnResults(result1 fs.FS, result2 error) *MoqSub_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FS
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSub_genType_doFn
		DoReturnFn MoqSub_genType_doReturnFn
	}{
		Values: &struct {
			Result1 fs.FS
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSub_genType_fnRecorder) AndDo(fn MoqSub_genType_doFn) *MoqSub_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSub_genType_fnRecorder) DoReturnResults(fn MoqSub_genType_doReturnFn) *MoqSub_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 fs.FS
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSub_genType_doFn
		DoReturnFn MoqSub_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSub_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSub_genType_resultsByParams
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
		results = &MoqSub_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSub_genType_paramsKey]*MoqSub_genType_results{},
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
		r.Results = &MoqSub_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSub_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSub_genType_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqSub_genType_doFn
				DoReturnFn MoqSub_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSub_genType) PrettyParams(params MoqSub_genType_params) string {
	return fmt.Sprintf("Sub_genType(%#v, %#v)", params.Fsys, params.Dir)
}

func (m *MoqSub_genType) ParamsKey(params MoqSub_genType_params, anyParams uint64) MoqSub_genType_paramsKey {
	m.Scene.T.Helper()
	var fsysUsed fs.FS
	var fsysUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fsys == moq.ParamIndexByValue {
			fsysUsed = params.Fsys
		} else {
			fsysUsedHash = hash.DeepHash(params.Fsys)
		}
	}
	var dirUsed string
	var dirUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Dir == moq.ParamIndexByValue {
			dirUsed = params.Dir
		} else {
			dirUsedHash = hash.DeepHash(params.Dir)
		}
	}
	return MoqSub_genType_paramsKey{
		Params: struct {
			Fsys fs.FS
			Dir  string
		}{
			Fsys: fsysUsed,
			Dir:  dirUsed,
		},
		Hashes: struct {
			Fsys hash.Hash
			Dir  hash.Hash
		}{
			Fsys: fsysUsedHash,
			Dir:  dirUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSub_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSub_genType) AssertExpectationsMet() {
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