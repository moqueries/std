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

// Accept_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Accept_genType func(fd int) (nfd int, sa syscall.Sockaddr, err error)

// MoqAccept_genType holds the state of a moq of the Accept_genType type
type MoqAccept_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAccept_genType_mock

	ResultsByParams []MoqAccept_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd moq.ParamIndexing
		}
	}
}

// MoqAccept_genType_mock isolates the mock interface of the Accept_genType
// type
type MoqAccept_genType_mock struct {
	Moq *MoqAccept_genType
}

// MoqAccept_genType_params holds the params of the Accept_genType type
type MoqAccept_genType_params struct{ Fd int }

// MoqAccept_genType_paramsKey holds the map key params of the Accept_genType
// type
type MoqAccept_genType_paramsKey struct {
	Params struct{ Fd int }
	Hashes struct{ Fd hash.Hash }
}

// MoqAccept_genType_resultsByParams contains the results for a given set of
// parameters for the Accept_genType type
type MoqAccept_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAccept_genType_paramsKey]*MoqAccept_genType_results
}

// MoqAccept_genType_doFn defines the type of function needed when calling
// AndDo for the Accept_genType type
type MoqAccept_genType_doFn func(fd int)

// MoqAccept_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Accept_genType type
type MoqAccept_genType_doReturnFn func(fd int) (nfd int, sa syscall.Sockaddr, err error)

// MoqAccept_genType_results holds the results of the Accept_genType type
type MoqAccept_genType_results struct {
	Params  MoqAccept_genType_params
	Results []struct {
		Values *struct {
			Nfd int
			Sa  syscall.Sockaddr
			Err error
		}
		Sequence   uint32
		DoFn       MoqAccept_genType_doFn
		DoReturnFn MoqAccept_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAccept_genType_fnRecorder routes recorded function calls to the
// MoqAccept_genType moq
type MoqAccept_genType_fnRecorder struct {
	Params    MoqAccept_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAccept_genType_results
	Moq       *MoqAccept_genType
}

// MoqAccept_genType_anyParams isolates the any params functions of the
// Accept_genType type
type MoqAccept_genType_anyParams struct {
	Recorder *MoqAccept_genType_fnRecorder
}

// NewMoqAccept_genType creates a new moq of the Accept_genType type
func NewMoqAccept_genType(scene *moq.Scene, config *moq.Config) *MoqAccept_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAccept_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAccept_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd moq.ParamIndexing
		}{
			Fd: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Accept_genType type
func (m *MoqAccept_genType) Mock() Accept_genType {
	return func(fd int) (_ int, _ syscall.Sockaddr, _ error) {
		m.Scene.T.Helper()
		moq := &MoqAccept_genType_mock{Moq: m}
		return moq.Fn(fd)
	}
}

func (m *MoqAccept_genType_mock) Fn(fd int) (nfd int, sa syscall.Sockaddr, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqAccept_genType_params{
		Fd: fd,
	}
	var results *MoqAccept_genType_results
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
		result.DoFn(fd)
	}

	if result.Values != nil {
		nfd = result.Values.Nfd
		sa = result.Values.Sa
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		nfd, sa, err = result.DoReturnFn(fd)
	}
	return
}

func (m *MoqAccept_genType) OnCall(fd int) *MoqAccept_genType_fnRecorder {
	return &MoqAccept_genType_fnRecorder{
		Params: MoqAccept_genType_params{
			Fd: fd,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAccept_genType_fnRecorder) Any() *MoqAccept_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAccept_genType_anyParams{Recorder: r}
}

func (a *MoqAccept_genType_anyParams) Fd() *MoqAccept_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqAccept_genType_fnRecorder) Seq() *MoqAccept_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAccept_genType_fnRecorder) NoSeq() *MoqAccept_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAccept_genType_fnRecorder) ReturnResults(nfd int, sa syscall.Sockaddr, err error) *MoqAccept_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Nfd int
			Sa  syscall.Sockaddr
			Err error
		}
		Sequence   uint32
		DoFn       MoqAccept_genType_doFn
		DoReturnFn MoqAccept_genType_doReturnFn
	}{
		Values: &struct {
			Nfd int
			Sa  syscall.Sockaddr
			Err error
		}{
			Nfd: nfd,
			Sa:  sa,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAccept_genType_fnRecorder) AndDo(fn MoqAccept_genType_doFn) *MoqAccept_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAccept_genType_fnRecorder) DoReturnResults(fn MoqAccept_genType_doReturnFn) *MoqAccept_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Nfd int
			Sa  syscall.Sockaddr
			Err error
		}
		Sequence   uint32
		DoFn       MoqAccept_genType_doFn
		DoReturnFn MoqAccept_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAccept_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAccept_genType_resultsByParams
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
		results = &MoqAccept_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAccept_genType_paramsKey]*MoqAccept_genType_results{},
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
		r.Results = &MoqAccept_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAccept_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAccept_genType_fnRecorder {
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
					Nfd int
					Sa  syscall.Sockaddr
					Err error
				}
				Sequence   uint32
				DoFn       MoqAccept_genType_doFn
				DoReturnFn MoqAccept_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAccept_genType) PrettyParams(params MoqAccept_genType_params) string {
	return fmt.Sprintf("Accept_genType(%#v)", params.Fd)
}

func (m *MoqAccept_genType) ParamsKey(params MoqAccept_genType_params, anyParams uint64) MoqAccept_genType_paramsKey {
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
	return MoqAccept_genType_paramsKey{
		Params: struct{ Fd int }{
			Fd: fdUsed,
		},
		Hashes: struct{ Fd hash.Hash }{
			Fd: fdUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAccept_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAccept_genType) AssertExpectationsMet() {
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
