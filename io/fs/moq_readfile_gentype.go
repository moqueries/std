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

// ReadFile_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type ReadFile_genType func(fsys fs.FS, name string) ([]byte, error)

// MoqReadFile_genType holds the state of a moq of the ReadFile_genType type
type MoqReadFile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReadFile_genType_mock

	ResultsByParams []MoqReadFile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fsys moq.ParamIndexing
			Name moq.ParamIndexing
		}
	}
}

// MoqReadFile_genType_mock isolates the mock interface of the ReadFile_genType
// type
type MoqReadFile_genType_mock struct {
	Moq *MoqReadFile_genType
}

// MoqReadFile_genType_params holds the params of the ReadFile_genType type
type MoqReadFile_genType_params struct {
	Fsys fs.FS
	Name string
}

// MoqReadFile_genType_paramsKey holds the map key params of the
// ReadFile_genType type
type MoqReadFile_genType_paramsKey struct {
	Params struct {
		Fsys fs.FS
		Name string
	}
	Hashes struct {
		Fsys hash.Hash
		Name hash.Hash
	}
}

// MoqReadFile_genType_resultsByParams contains the results for a given set of
// parameters for the ReadFile_genType type
type MoqReadFile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadFile_genType_paramsKey]*MoqReadFile_genType_results
}

// MoqReadFile_genType_doFn defines the type of function needed when calling
// AndDo for the ReadFile_genType type
type MoqReadFile_genType_doFn func(fsys fs.FS, name string)

// MoqReadFile_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadFile_genType type
type MoqReadFile_genType_doReturnFn func(fsys fs.FS, name string) ([]byte, error)

// MoqReadFile_genType_results holds the results of the ReadFile_genType type
type MoqReadFile_genType_results struct {
	Params  MoqReadFile_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadFile_genType_doFn
		DoReturnFn MoqReadFile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadFile_genType_fnRecorder routes recorded function calls to the
// MoqReadFile_genType moq
type MoqReadFile_genType_fnRecorder struct {
	Params    MoqReadFile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadFile_genType_results
	Moq       *MoqReadFile_genType
}

// MoqReadFile_genType_anyParams isolates the any params functions of the
// ReadFile_genType type
type MoqReadFile_genType_anyParams struct {
	Recorder *MoqReadFile_genType_fnRecorder
}

// NewMoqReadFile_genType creates a new moq of the ReadFile_genType type
func NewMoqReadFile_genType(scene *moq.Scene, config *moq.Config) *MoqReadFile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReadFile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReadFile_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fsys moq.ParamIndexing
				Name moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fsys moq.ParamIndexing
			Name moq.ParamIndexing
		}{
			Fsys: moq.ParamIndexByHash,
			Name: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ReadFile_genType type
func (m *MoqReadFile_genType) Mock() ReadFile_genType {
	return func(fsys fs.FS, name string) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqReadFile_genType_mock{Moq: m}
		return moq.Fn(fsys, name)
	}
}

func (m *MoqReadFile_genType_mock) Fn(fsys fs.FS, name string) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadFile_genType_params{
		Fsys: fsys,
		Name: name,
	}
	var results *MoqReadFile_genType_results
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
		result.DoFn(fsys, name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(fsys, name)
	}
	return
}

func (m *MoqReadFile_genType) OnCall(fsys fs.FS, name string) *MoqReadFile_genType_fnRecorder {
	return &MoqReadFile_genType_fnRecorder{
		Params: MoqReadFile_genType_params{
			Fsys: fsys,
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReadFile_genType_fnRecorder) Any() *MoqReadFile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReadFile_genType_anyParams{Recorder: r}
}

func (a *MoqReadFile_genType_anyParams) Fsys() *MoqReadFile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReadFile_genType_anyParams) Name() *MoqReadFile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqReadFile_genType_fnRecorder) Seq() *MoqReadFile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadFile_genType_fnRecorder) NoSeq() *MoqReadFile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadFile_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqReadFile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadFile_genType_doFn
		DoReturnFn MoqReadFile_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReadFile_genType_fnRecorder) AndDo(fn MoqReadFile_genType_doFn) *MoqReadFile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadFile_genType_fnRecorder) DoReturnResults(fn MoqReadFile_genType_doReturnFn) *MoqReadFile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadFile_genType_doFn
		DoReturnFn MoqReadFile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadFile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadFile_genType_resultsByParams
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
		results = &MoqReadFile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadFile_genType_paramsKey]*MoqReadFile_genType_results{},
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
		r.Results = &MoqReadFile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadFile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadFile_genType_fnRecorder {
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
					Result1 []byte
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReadFile_genType_doFn
				DoReturnFn MoqReadFile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadFile_genType) PrettyParams(params MoqReadFile_genType_params) string {
	return fmt.Sprintf("ReadFile_genType(%#v, %#v)", params.Fsys, params.Name)
}

func (m *MoqReadFile_genType) ParamsKey(params MoqReadFile_genType_params, anyParams uint64) MoqReadFile_genType_paramsKey {
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
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	return MoqReadFile_genType_paramsKey{
		Params: struct {
			Fsys fs.FS
			Name string
		}{
			Fsys: fsysUsed,
			Name: nameUsed,
		},
		Hashes: struct {
			Fsys hash.Hash
			Name hash.Hash
		}{
			Fsys: fsysUsedHash,
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReadFile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReadFile_genType) AssertExpectationsMet() {
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
