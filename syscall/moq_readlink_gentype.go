// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Readlink_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Readlink_genType func(path string, buf []byte) (n int, err error)

// MoqReadlink_genType holds the state of a moq of the Readlink_genType type
type MoqReadlink_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReadlink_genType_mock

	ResultsByParams []MoqReadlink_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Buf  moq.ParamIndexing
		}
	}
}

// MoqReadlink_genType_mock isolates the mock interface of the Readlink_genType
// type
type MoqReadlink_genType_mock struct {
	Moq *MoqReadlink_genType
}

// MoqReadlink_genType_params holds the params of the Readlink_genType type
type MoqReadlink_genType_params struct {
	Path string
	Buf  []byte
}

// MoqReadlink_genType_paramsKey holds the map key params of the
// Readlink_genType type
type MoqReadlink_genType_paramsKey struct {
	Params struct{ Path string }
	Hashes struct {
		Path hash.Hash
		Buf  hash.Hash
	}
}

// MoqReadlink_genType_resultsByParams contains the results for a given set of
// parameters for the Readlink_genType type
type MoqReadlink_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadlink_genType_paramsKey]*MoqReadlink_genType_results
}

// MoqReadlink_genType_doFn defines the type of function needed when calling
// AndDo for the Readlink_genType type
type MoqReadlink_genType_doFn func(path string, buf []byte)

// MoqReadlink_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Readlink_genType type
type MoqReadlink_genType_doReturnFn func(path string, buf []byte) (n int, err error)

// MoqReadlink_genType_results holds the results of the Readlink_genType type
type MoqReadlink_genType_results struct {
	Params  MoqReadlink_genType_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReadlink_genType_doFn
		DoReturnFn MoqReadlink_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadlink_genType_fnRecorder routes recorded function calls to the
// MoqReadlink_genType moq
type MoqReadlink_genType_fnRecorder struct {
	Params    MoqReadlink_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadlink_genType_results
	Moq       *MoqReadlink_genType
}

// MoqReadlink_genType_anyParams isolates the any params functions of the
// Readlink_genType type
type MoqReadlink_genType_anyParams struct {
	Recorder *MoqReadlink_genType_fnRecorder
}

// NewMoqReadlink_genType creates a new moq of the Readlink_genType type
func NewMoqReadlink_genType(scene *moq.Scene, config *moq.Config) *MoqReadlink_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReadlink_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReadlink_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Buf  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Buf  moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Buf:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Readlink_genType type
func (m *MoqReadlink_genType) Mock() Readlink_genType {
	return func(path string, buf []byte) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqReadlink_genType_mock{Moq: m}
		return moq.Fn(path, buf)
	}
}

func (m *MoqReadlink_genType_mock) Fn(path string, buf []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadlink_genType_params{
		Path: path,
		Buf:  buf,
	}
	var results *MoqReadlink_genType_results
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
		result.DoFn(path, buf)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(path, buf)
	}
	return
}

func (m *MoqReadlink_genType) OnCall(path string, buf []byte) *MoqReadlink_genType_fnRecorder {
	return &MoqReadlink_genType_fnRecorder{
		Params: MoqReadlink_genType_params{
			Path: path,
			Buf:  buf,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReadlink_genType_fnRecorder) Any() *MoqReadlink_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReadlink_genType_anyParams{Recorder: r}
}

func (a *MoqReadlink_genType_anyParams) Path() *MoqReadlink_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReadlink_genType_anyParams) Buf() *MoqReadlink_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqReadlink_genType_fnRecorder) Seq() *MoqReadlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadlink_genType_fnRecorder) NoSeq() *MoqReadlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadlink_genType_fnRecorder) ReturnResults(n int, err error) *MoqReadlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReadlink_genType_doFn
		DoReturnFn MoqReadlink_genType_doReturnFn
	}{
		Values: &struct {
			N   int
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReadlink_genType_fnRecorder) AndDo(fn MoqReadlink_genType_doFn) *MoqReadlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadlink_genType_fnRecorder) DoReturnResults(fn MoqReadlink_genType_doReturnFn) *MoqReadlink_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReadlink_genType_doFn
		DoReturnFn MoqReadlink_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadlink_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadlink_genType_resultsByParams
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
		results = &MoqReadlink_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadlink_genType_paramsKey]*MoqReadlink_genType_results{},
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
		r.Results = &MoqReadlink_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadlink_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadlink_genType_fnRecorder {
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
					N   int
					Err error
				}
				Sequence   uint32
				DoFn       MoqReadlink_genType_doFn
				DoReturnFn MoqReadlink_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadlink_genType) PrettyParams(params MoqReadlink_genType_params) string {
	return fmt.Sprintf("Readlink_genType(%#v, %#v)", params.Path, params.Buf)
}

func (m *MoqReadlink_genType) ParamsKey(params MoqReadlink_genType_params, anyParams uint64) MoqReadlink_genType_paramsKey {
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
	var bufUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Buf == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The buf parameter can't be indexed by value")
		}
		bufUsedHash = hash.DeepHash(params.Buf)
	}
	return MoqReadlink_genType_paramsKey{
		Params: struct{ Path string }{
			Path: pathUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Buf  hash.Hash
		}{
			Path: pathUsedHash,
			Buf:  bufUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReadlink_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReadlink_genType) AssertExpectationsMet() {
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
