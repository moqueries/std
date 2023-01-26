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

// Uname_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Uname_genType func(buf *syscall.Utsname) (err error)

// MoqUname_genType holds the state of a moq of the Uname_genType type
type MoqUname_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUname_genType_mock

	ResultsByParams []MoqUname_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Buf moq.ParamIndexing
		}
	}
}

// MoqUname_genType_mock isolates the mock interface of the Uname_genType type
type MoqUname_genType_mock struct {
	Moq *MoqUname_genType
}

// MoqUname_genType_params holds the params of the Uname_genType type
type MoqUname_genType_params struct{ Buf *syscall.Utsname }

// MoqUname_genType_paramsKey holds the map key params of the Uname_genType
// type
type MoqUname_genType_paramsKey struct {
	Params struct{ Buf *syscall.Utsname }
	Hashes struct{ Buf hash.Hash }
}

// MoqUname_genType_resultsByParams contains the results for a given set of
// parameters for the Uname_genType type
type MoqUname_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUname_genType_paramsKey]*MoqUname_genType_results
}

// MoqUname_genType_doFn defines the type of function needed when calling AndDo
// for the Uname_genType type
type MoqUname_genType_doFn func(buf *syscall.Utsname)

// MoqUname_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Uname_genType type
type MoqUname_genType_doReturnFn func(buf *syscall.Utsname) (err error)

// MoqUname_genType_results holds the results of the Uname_genType type
type MoqUname_genType_results struct {
	Params  MoqUname_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUname_genType_doFn
		DoReturnFn MoqUname_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUname_genType_fnRecorder routes recorded function calls to the
// MoqUname_genType moq
type MoqUname_genType_fnRecorder struct {
	Params    MoqUname_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUname_genType_results
	Moq       *MoqUname_genType
}

// MoqUname_genType_anyParams isolates the any params functions of the
// Uname_genType type
type MoqUname_genType_anyParams struct {
	Recorder *MoqUname_genType_fnRecorder
}

// NewMoqUname_genType creates a new moq of the Uname_genType type
func NewMoqUname_genType(scene *moq.Scene, config *moq.Config) *MoqUname_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUname_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUname_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Buf moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Buf moq.ParamIndexing
		}{
			Buf: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Uname_genType type
func (m *MoqUname_genType) Mock() Uname_genType {
	return func(buf *syscall.Utsname) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqUname_genType_mock{Moq: m}
		return moq.Fn(buf)
	}
}

func (m *MoqUname_genType_mock) Fn(buf *syscall.Utsname) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqUname_genType_params{
		Buf: buf,
	}
	var results *MoqUname_genType_results
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
		result.DoFn(buf)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(buf)
	}
	return
}

func (m *MoqUname_genType) OnCall(buf *syscall.Utsname) *MoqUname_genType_fnRecorder {
	return &MoqUname_genType_fnRecorder{
		Params: MoqUname_genType_params{
			Buf: buf,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUname_genType_fnRecorder) Any() *MoqUname_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUname_genType_anyParams{Recorder: r}
}

func (a *MoqUname_genType_anyParams) Buf() *MoqUname_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqUname_genType_fnRecorder) Seq() *MoqUname_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUname_genType_fnRecorder) NoSeq() *MoqUname_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUname_genType_fnRecorder) ReturnResults(err error) *MoqUname_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUname_genType_doFn
		DoReturnFn MoqUname_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUname_genType_fnRecorder) AndDo(fn MoqUname_genType_doFn) *MoqUname_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUname_genType_fnRecorder) DoReturnResults(fn MoqUname_genType_doReturnFn) *MoqUname_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqUname_genType_doFn
		DoReturnFn MoqUname_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUname_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUname_genType_resultsByParams
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
		results = &MoqUname_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUname_genType_paramsKey]*MoqUname_genType_results{},
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
		r.Results = &MoqUname_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUname_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUname_genType_fnRecorder {
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
				DoFn       MoqUname_genType_doFn
				DoReturnFn MoqUname_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUname_genType) PrettyParams(params MoqUname_genType_params) string {
	return fmt.Sprintf("Uname_genType(%#v)", params.Buf)
}

func (m *MoqUname_genType) ParamsKey(params MoqUname_genType_params, anyParams uint64) MoqUname_genType_paramsKey {
	m.Scene.T.Helper()
	var bufUsed *syscall.Utsname
	var bufUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Buf == moq.ParamIndexByValue {
			bufUsed = params.Buf
		} else {
			bufUsedHash = hash.DeepHash(params.Buf)
		}
	}
	return MoqUname_genType_paramsKey{
		Params: struct{ Buf *syscall.Utsname }{
			Buf: bufUsed,
		},
		Hashes: struct{ Buf hash.Hash }{
			Buf: bufUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqUname_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUname_genType) AssertExpectationsMet() {
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