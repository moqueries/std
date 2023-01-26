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

// EpollWait_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type EpollWait_genType func(epfd int, events []syscall.EpollEvent, msec int) (n int, err error)

// MoqEpollWait_genType holds the state of a moq of the EpollWait_genType type
type MoqEpollWait_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEpollWait_genType_mock

	ResultsByParams []MoqEpollWait_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Epfd   moq.ParamIndexing
			Events moq.ParamIndexing
			Msec   moq.ParamIndexing
		}
	}
}

// MoqEpollWait_genType_mock isolates the mock interface of the
// EpollWait_genType type
type MoqEpollWait_genType_mock struct {
	Moq *MoqEpollWait_genType
}

// MoqEpollWait_genType_params holds the params of the EpollWait_genType type
type MoqEpollWait_genType_params struct {
	Epfd   int
	Events []syscall.EpollEvent
	Msec   int
}

// MoqEpollWait_genType_paramsKey holds the map key params of the
// EpollWait_genType type
type MoqEpollWait_genType_paramsKey struct {
	Params struct {
		Epfd int
		Msec int
	}
	Hashes struct {
		Epfd   hash.Hash
		Events hash.Hash
		Msec   hash.Hash
	}
}

// MoqEpollWait_genType_resultsByParams contains the results for a given set of
// parameters for the EpollWait_genType type
type MoqEpollWait_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEpollWait_genType_paramsKey]*MoqEpollWait_genType_results
}

// MoqEpollWait_genType_doFn defines the type of function needed when calling
// AndDo for the EpollWait_genType type
type MoqEpollWait_genType_doFn func(epfd int, events []syscall.EpollEvent, msec int)

// MoqEpollWait_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the EpollWait_genType type
type MoqEpollWait_genType_doReturnFn func(epfd int, events []syscall.EpollEvent, msec int) (n int, err error)

// MoqEpollWait_genType_results holds the results of the EpollWait_genType type
type MoqEpollWait_genType_results struct {
	Params  MoqEpollWait_genType_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqEpollWait_genType_doFn
		DoReturnFn MoqEpollWait_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEpollWait_genType_fnRecorder routes recorded function calls to the
// MoqEpollWait_genType moq
type MoqEpollWait_genType_fnRecorder struct {
	Params    MoqEpollWait_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEpollWait_genType_results
	Moq       *MoqEpollWait_genType
}

// MoqEpollWait_genType_anyParams isolates the any params functions of the
// EpollWait_genType type
type MoqEpollWait_genType_anyParams struct {
	Recorder *MoqEpollWait_genType_fnRecorder
}

// NewMoqEpollWait_genType creates a new moq of the EpollWait_genType type
func NewMoqEpollWait_genType(scene *moq.Scene, config *moq.Config) *MoqEpollWait_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEpollWait_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEpollWait_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Epfd   moq.ParamIndexing
				Events moq.ParamIndexing
				Msec   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Epfd   moq.ParamIndexing
			Events moq.ParamIndexing
			Msec   moq.ParamIndexing
		}{
			Epfd:   moq.ParamIndexByValue,
			Events: moq.ParamIndexByHash,
			Msec:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the EpollWait_genType type
func (m *MoqEpollWait_genType) Mock() EpollWait_genType {
	return func(epfd int, events []syscall.EpollEvent, msec int) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqEpollWait_genType_mock{Moq: m}
		return moq.Fn(epfd, events, msec)
	}
}

func (m *MoqEpollWait_genType_mock) Fn(epfd int, events []syscall.EpollEvent, msec int) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqEpollWait_genType_params{
		Epfd:   epfd,
		Events: events,
		Msec:   msec,
	}
	var results *MoqEpollWait_genType_results
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
		result.DoFn(epfd, events, msec)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(epfd, events, msec)
	}
	return
}

func (m *MoqEpollWait_genType) OnCall(epfd int, events []syscall.EpollEvent, msec int) *MoqEpollWait_genType_fnRecorder {
	return &MoqEpollWait_genType_fnRecorder{
		Params: MoqEpollWait_genType_params{
			Epfd:   epfd,
			Events: events,
			Msec:   msec,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqEpollWait_genType_fnRecorder) Any() *MoqEpollWait_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqEpollWait_genType_anyParams{Recorder: r}
}

func (a *MoqEpollWait_genType_anyParams) Epfd() *MoqEpollWait_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEpollWait_genType_anyParams) Events() *MoqEpollWait_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqEpollWait_genType_anyParams) Msec() *MoqEpollWait_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqEpollWait_genType_fnRecorder) Seq() *MoqEpollWait_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEpollWait_genType_fnRecorder) NoSeq() *MoqEpollWait_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEpollWait_genType_fnRecorder) ReturnResults(n int, err error) *MoqEpollWait_genType_fnRecorder {
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
		DoFn       MoqEpollWait_genType_doFn
		DoReturnFn MoqEpollWait_genType_doReturnFn
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

func (r *MoqEpollWait_genType_fnRecorder) AndDo(fn MoqEpollWait_genType_doFn) *MoqEpollWait_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEpollWait_genType_fnRecorder) DoReturnResults(fn MoqEpollWait_genType_doReturnFn) *MoqEpollWait_genType_fnRecorder {
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
		DoFn       MoqEpollWait_genType_doFn
		DoReturnFn MoqEpollWait_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEpollWait_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEpollWait_genType_resultsByParams
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
		results = &MoqEpollWait_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEpollWait_genType_paramsKey]*MoqEpollWait_genType_results{},
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
		r.Results = &MoqEpollWait_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEpollWait_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEpollWait_genType_fnRecorder {
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
				DoFn       MoqEpollWait_genType_doFn
				DoReturnFn MoqEpollWait_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEpollWait_genType) PrettyParams(params MoqEpollWait_genType_params) string {
	return fmt.Sprintf("EpollWait_genType(%#v, %#v, %#v)", params.Epfd, params.Events, params.Msec)
}

func (m *MoqEpollWait_genType) ParamsKey(params MoqEpollWait_genType_params, anyParams uint64) MoqEpollWait_genType_paramsKey {
	m.Scene.T.Helper()
	var epfdUsed int
	var epfdUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Epfd == moq.ParamIndexByValue {
			epfdUsed = params.Epfd
		} else {
			epfdUsedHash = hash.DeepHash(params.Epfd)
		}
	}
	var eventsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Events == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The events parameter can't be indexed by value")
		}
		eventsUsedHash = hash.DeepHash(params.Events)
	}
	var msecUsed int
	var msecUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Msec == moq.ParamIndexByValue {
			msecUsed = params.Msec
		} else {
			msecUsedHash = hash.DeepHash(params.Msec)
		}
	}
	return MoqEpollWait_genType_paramsKey{
		Params: struct {
			Epfd int
			Msec int
		}{
			Epfd: epfdUsed,
			Msec: msecUsed,
		},
		Hashes: struct {
			Epfd   hash.Hash
			Events hash.Hash
			Msec   hash.Hash
		}{
			Epfd:   epfdUsedHash,
			Events: eventsUsedHash,
			Msec:   msecUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEpollWait_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEpollWait_genType) AssertExpectationsMet() {
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
