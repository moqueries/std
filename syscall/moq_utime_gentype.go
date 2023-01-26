// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"syscall"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Utime_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Utime_genType func(path string, buf *syscall.Utimbuf) (err error)

// MoqUtime_genType holds the state of a moq of the Utime_genType type
type MoqUtime_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUtime_genType_mock

	ResultsByParams []MoqUtime_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Buf  moq.ParamIndexing
		}
	}
}

// MoqUtime_genType_mock isolates the mock interface of the Utime_genType type
type MoqUtime_genType_mock struct {
	Moq *MoqUtime_genType
}

// MoqUtime_genType_params holds the params of the Utime_genType type
type MoqUtime_genType_params struct {
	Path string
	Buf  *syscall.Utimbuf
}

// MoqUtime_genType_paramsKey holds the map key params of the Utime_genType
// type
type MoqUtime_genType_paramsKey struct {
	Params struct {
		Path string
		Buf  *syscall.Utimbuf
	}
	Hashes struct {
		Path hash.Hash
		Buf  hash.Hash
	}
}

// MoqUtime_genType_resultsByParams contains the results for a given set of
// parameters for the Utime_genType type
type MoqUtime_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUtime_genType_paramsKey]*MoqUtime_genType_results
}

// MoqUtime_genType_doFn defines the type of function needed when calling AndDo
// for the Utime_genType type
type MoqUtime_genType_doFn func(path string, buf *syscall.Utimbuf)

// MoqUtime_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Utime_genType type
type MoqUtime_genType_doReturnFn func(path string, buf *syscall.Utimbuf) (err error)

// MoqUtime_genType_results holds the results of the Utime_genType type
type MoqUtime_genType_results struct {
	Params  MoqUtime_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUtime_genType_doFn
		DoReturnFn MoqUtime_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUtime_genType_fnRecorder routes recorded function calls to the
// MoqUtime_genType moq
type MoqUtime_genType_fnRecorder struct {
	Params    MoqUtime_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUtime_genType_results
	Moq       *MoqUtime_genType
}

// MoqUtime_genType_anyParams isolates the any params functions of the
// Utime_genType type
type MoqUtime_genType_anyParams struct {
	Recorder *MoqUtime_genType_fnRecorder
}

// NewMoqUtime_genType creates a new moq of the Utime_genType type
func NewMoqUtime_genType(scene *moq.Scene, config *moq.Config) *MoqUtime_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUtime_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUtime_genType_mock{},

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

// Mock returns the moq implementation of the Utime_genType type
func (m *MoqUtime_genType) Mock() Utime_genType {
	return func(path string, buf *syscall.Utimbuf) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqUtime_genType_mock{Moq: m}
		return moq.Fn(path, buf)
	}
}

func (m *MoqUtime_genType_mock) Fn(path string, buf *syscall.Utimbuf) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqUtime_genType_params{
		Path: path,
		Buf:  buf,
	}
	var results *MoqUtime_genType_results
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
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(path, buf)
	}
	return
}

func (m *MoqUtime_genType) OnCall(path string, buf *syscall.Utimbuf) *MoqUtime_genType_fnRecorder {
	return &MoqUtime_genType_fnRecorder{
		Params: MoqUtime_genType_params{
			Path: path,
			Buf:  buf,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUtime_genType_fnRecorder) Any() *MoqUtime_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUtime_genType_anyParams{Recorder: r}
}

func (a *MoqUtime_genType_anyParams) Path() *MoqUtime_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqUtime_genType_anyParams) Buf() *MoqUtime_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqUtime_genType_fnRecorder) Seq() *MoqUtime_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUtime_genType_fnRecorder) NoSeq() *MoqUtime_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUtime_genType_fnRecorder) ReturnResults(err error) *MoqUtime_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUtime_genType_doFn
		DoReturnFn MoqUtime_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUtime_genType_fnRecorder) AndDo(fn MoqUtime_genType_doFn) *MoqUtime_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUtime_genType_fnRecorder) DoReturnResults(fn MoqUtime_genType_doReturnFn) *MoqUtime_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUtime_genType_doFn
		DoReturnFn MoqUtime_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUtime_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUtime_genType_resultsByParams
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
		results = &MoqUtime_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUtime_genType_paramsKey]*MoqUtime_genType_results{},
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
		r.Results = &MoqUtime_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUtime_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUtime_genType_fnRecorder {
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
				DoFn       MoqUtime_genType_doFn
				DoReturnFn MoqUtime_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUtime_genType) PrettyParams(params MoqUtime_genType_params) string {
	return fmt.Sprintf("Utime_genType(%#v, %#v)", params.Path, params.Buf)
}

func (m *MoqUtime_genType) ParamsKey(params MoqUtime_genType_params, anyParams uint64) MoqUtime_genType_paramsKey {
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
	var bufUsed *syscall.Utimbuf
	var bufUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Buf == moq.ParamIndexByValue {
			bufUsed = params.Buf
		} else {
			bufUsedHash = hash.DeepHash(params.Buf)
		}
	}
	return MoqUtime_genType_paramsKey{
		Params: struct {
			Path string
			Buf  *syscall.Utimbuf
		}{
			Path: pathUsed,
			Buf:  bufUsed,
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
func (m *MoqUtime_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUtime_genType) AssertExpectationsMet() {
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