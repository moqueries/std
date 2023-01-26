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

// Futimes_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Futimes_genType func(fd int, tv []syscall.Timeval) (err error)

// MoqFutimes_genType holds the state of a moq of the Futimes_genType type
type MoqFutimes_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFutimes_genType_mock

	ResultsByParams []MoqFutimes_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd moq.ParamIndexing
			Tv moq.ParamIndexing
		}
	}
}

// MoqFutimes_genType_mock isolates the mock interface of the Futimes_genType
// type
type MoqFutimes_genType_mock struct {
	Moq *MoqFutimes_genType
}

// MoqFutimes_genType_params holds the params of the Futimes_genType type
type MoqFutimes_genType_params struct {
	Fd int
	Tv []syscall.Timeval
}

// MoqFutimes_genType_paramsKey holds the map key params of the Futimes_genType
// type
type MoqFutimes_genType_paramsKey struct {
	Params struct{ Fd int }
	Hashes struct {
		Fd hash.Hash
		Tv hash.Hash
	}
}

// MoqFutimes_genType_resultsByParams contains the results for a given set of
// parameters for the Futimes_genType type
type MoqFutimes_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFutimes_genType_paramsKey]*MoqFutimes_genType_results
}

// MoqFutimes_genType_doFn defines the type of function needed when calling
// AndDo for the Futimes_genType type
type MoqFutimes_genType_doFn func(fd int, tv []syscall.Timeval)

// MoqFutimes_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Futimes_genType type
type MoqFutimes_genType_doReturnFn func(fd int, tv []syscall.Timeval) (err error)

// MoqFutimes_genType_results holds the results of the Futimes_genType type
type MoqFutimes_genType_results struct {
	Params  MoqFutimes_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFutimes_genType_doFn
		DoReturnFn MoqFutimes_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFutimes_genType_fnRecorder routes recorded function calls to the
// MoqFutimes_genType moq
type MoqFutimes_genType_fnRecorder struct {
	Params    MoqFutimes_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFutimes_genType_results
	Moq       *MoqFutimes_genType
}

// MoqFutimes_genType_anyParams isolates the any params functions of the
// Futimes_genType type
type MoqFutimes_genType_anyParams struct {
	Recorder *MoqFutimes_genType_fnRecorder
}

// NewMoqFutimes_genType creates a new moq of the Futimes_genType type
func NewMoqFutimes_genType(scene *moq.Scene, config *moq.Config) *MoqFutimes_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFutimes_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFutimes_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd moq.ParamIndexing
				Tv moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd moq.ParamIndexing
			Tv moq.ParamIndexing
		}{
			Fd: moq.ParamIndexByValue,
			Tv: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Futimes_genType type
func (m *MoqFutimes_genType) Mock() Futimes_genType {
	return func(fd int, tv []syscall.Timeval) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqFutimes_genType_mock{Moq: m}
		return moq.Fn(fd, tv)
	}
}

func (m *MoqFutimes_genType_mock) Fn(fd int, tv []syscall.Timeval) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFutimes_genType_params{
		Fd: fd,
		Tv: tv,
	}
	var results *MoqFutimes_genType_results
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
		result.DoFn(fd, tv)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(fd, tv)
	}
	return
}

func (m *MoqFutimes_genType) OnCall(fd int, tv []syscall.Timeval) *MoqFutimes_genType_fnRecorder {
	return &MoqFutimes_genType_fnRecorder{
		Params: MoqFutimes_genType_params{
			Fd: fd,
			Tv: tv,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFutimes_genType_fnRecorder) Any() *MoqFutimes_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFutimes_genType_anyParams{Recorder: r}
}

func (a *MoqFutimes_genType_anyParams) Fd() *MoqFutimes_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFutimes_genType_anyParams) Tv() *MoqFutimes_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFutimes_genType_fnRecorder) Seq() *MoqFutimes_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFutimes_genType_fnRecorder) NoSeq() *MoqFutimes_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFutimes_genType_fnRecorder) ReturnResults(err error) *MoqFutimes_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFutimes_genType_doFn
		DoReturnFn MoqFutimes_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFutimes_genType_fnRecorder) AndDo(fn MoqFutimes_genType_doFn) *MoqFutimes_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFutimes_genType_fnRecorder) DoReturnResults(fn MoqFutimes_genType_doReturnFn) *MoqFutimes_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFutimes_genType_doFn
		DoReturnFn MoqFutimes_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFutimes_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFutimes_genType_resultsByParams
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
		results = &MoqFutimes_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFutimes_genType_paramsKey]*MoqFutimes_genType_results{},
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
		r.Results = &MoqFutimes_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFutimes_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFutimes_genType_fnRecorder {
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
				DoFn       MoqFutimes_genType_doFn
				DoReturnFn MoqFutimes_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFutimes_genType) PrettyParams(params MoqFutimes_genType_params) string {
	return fmt.Sprintf("Futimes_genType(%#v, %#v)", params.Fd, params.Tv)
}

func (m *MoqFutimes_genType) ParamsKey(params MoqFutimes_genType_params, anyParams uint64) MoqFutimes_genType_paramsKey {
	m.Scene.T.Helper()
	var fdUsed int
	var fdUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fd == moq.ParamIndexByValue {
			fdUsed = params.Fd
		} else {
			fdUsedHash = hash.DeepHash(params.Fd)
		}
	}
	var tvUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Tv == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The tv parameter can't be indexed by value")
		}
		tvUsedHash = hash.DeepHash(params.Tv)
	}
	return MoqFutimes_genType_paramsKey{
		Params: struct{ Fd int }{
			Fd: fdUsed,
		},
		Hashes: struct {
			Fd hash.Hash
			Tv hash.Hash
		}{
			Fd: fdUsedHash,
			Tv: tvUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFutimes_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFutimes_genType) AssertExpectationsMet() {
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