// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// GetsockoptInet4Addr_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type GetsockoptInet4Addr_genType func(fd, level, opt int) (value [4]byte, err error)

// MoqGetsockoptInet4Addr_genType holds the state of a moq of the
// GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetsockoptInet4Addr_genType_mock

	ResultsByParams []MoqGetsockoptInet4Addr_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd    moq.ParamIndexing
			Level moq.ParamIndexing
			Opt   moq.ParamIndexing
		}
	}
}

// MoqGetsockoptInet4Addr_genType_mock isolates the mock interface of the
// GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_mock struct {
	Moq *MoqGetsockoptInet4Addr_genType
}

// MoqGetsockoptInet4Addr_genType_params holds the params of the
// GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_params struct{ Fd, Level, Opt int }

// MoqGetsockoptInet4Addr_genType_paramsKey holds the map key params of the
// GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_paramsKey struct {
	Params struct{ Fd, Level, Opt int }
	Hashes struct{ Fd, Level, Opt hash.Hash }
}

// MoqGetsockoptInet4Addr_genType_resultsByParams contains the results for a
// given set of parameters for the GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetsockoptInet4Addr_genType_paramsKey]*MoqGetsockoptInet4Addr_genType_results
}

// MoqGetsockoptInet4Addr_genType_doFn defines the type of function needed when
// calling AndDo for the GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_doFn func(fd, level, opt int)

// MoqGetsockoptInet4Addr_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_doReturnFn func(fd, level, opt int) (value [4]byte, err error)

// MoqGetsockoptInet4Addr_genType_results holds the results of the
// GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_results struct {
	Params  MoqGetsockoptInet4Addr_genType_params
	Results []struct {
		Values *struct {
			Value [4]byte
			Err   error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptInet4Addr_genType_doFn
		DoReturnFn MoqGetsockoptInet4Addr_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetsockoptInet4Addr_genType_fnRecorder routes recorded function calls to
// the MoqGetsockoptInet4Addr_genType moq
type MoqGetsockoptInet4Addr_genType_fnRecorder struct {
	Params    MoqGetsockoptInet4Addr_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetsockoptInet4Addr_genType_results
	Moq       *MoqGetsockoptInet4Addr_genType
}

// MoqGetsockoptInet4Addr_genType_anyParams isolates the any params functions
// of the GetsockoptInet4Addr_genType type
type MoqGetsockoptInet4Addr_genType_anyParams struct {
	Recorder *MoqGetsockoptInet4Addr_genType_fnRecorder
}

// NewMoqGetsockoptInet4Addr_genType creates a new moq of the
// GetsockoptInet4Addr_genType type
func NewMoqGetsockoptInet4Addr_genType(scene *moq.Scene, config *moq.Config) *MoqGetsockoptInet4Addr_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetsockoptInet4Addr_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetsockoptInet4Addr_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd    moq.ParamIndexing
				Level moq.ParamIndexing
				Opt   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd    moq.ParamIndexing
			Level moq.ParamIndexing
			Opt   moq.ParamIndexing
		}{
			Fd:    moq.ParamIndexByValue,
			Level: moq.ParamIndexByValue,
			Opt:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the GetsockoptInet4Addr_genType type
func (m *MoqGetsockoptInet4Addr_genType) Mock() GetsockoptInet4Addr_genType {
	return func(fd, level, opt int) (_ [4]byte, _ error) {
		m.Scene.T.Helper()
		moq := &MoqGetsockoptInet4Addr_genType_mock{Moq: m}
		return moq.Fn(fd, level, opt)
	}
}

func (m *MoqGetsockoptInet4Addr_genType_mock) Fn(fd, level, opt int) (value [4]byte, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqGetsockoptInet4Addr_genType_params{
		Fd:    fd,
		Level: level,
		Opt:   opt,
	}
	var results *MoqGetsockoptInet4Addr_genType_results
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
		result.DoFn(fd, level, opt)
	}

	if result.Values != nil {
		value = result.Values.Value
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		value, err = result.DoReturnFn(fd, level, opt)
	}
	return
}

func (m *MoqGetsockoptInet4Addr_genType) OnCall(fd, level, opt int) *MoqGetsockoptInet4Addr_genType_fnRecorder {
	return &MoqGetsockoptInet4Addr_genType_fnRecorder{
		Params: MoqGetsockoptInet4Addr_genType_params{
			Fd:    fd,
			Level: level,
			Opt:   opt,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) Any() *MoqGetsockoptInet4Addr_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetsockoptInet4Addr_genType_anyParams{Recorder: r}
}

func (a *MoqGetsockoptInet4Addr_genType_anyParams) Fd() *MoqGetsockoptInet4Addr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGetsockoptInet4Addr_genType_anyParams) Level() *MoqGetsockoptInet4Addr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqGetsockoptInet4Addr_genType_anyParams) Opt() *MoqGetsockoptInet4Addr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) Seq() *MoqGetsockoptInet4Addr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) NoSeq() *MoqGetsockoptInet4Addr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) ReturnResults(value [4]byte, err error) *MoqGetsockoptInet4Addr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Value [4]byte
			Err   error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptInet4Addr_genType_doFn
		DoReturnFn MoqGetsockoptInet4Addr_genType_doReturnFn
	}{
		Values: &struct {
			Value [4]byte
			Err   error
		}{
			Value: value,
			Err:   err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) AndDo(fn MoqGetsockoptInet4Addr_genType_doFn) *MoqGetsockoptInet4Addr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) DoReturnResults(fn MoqGetsockoptInet4Addr_genType_doReturnFn) *MoqGetsockoptInet4Addr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Value [4]byte
			Err   error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptInet4Addr_genType_doFn
		DoReturnFn MoqGetsockoptInet4Addr_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetsockoptInet4Addr_genType_resultsByParams
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
		results = &MoqGetsockoptInet4Addr_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetsockoptInet4Addr_genType_paramsKey]*MoqGetsockoptInet4Addr_genType_results{},
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
		r.Results = &MoqGetsockoptInet4Addr_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetsockoptInet4Addr_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetsockoptInet4Addr_genType_fnRecorder {
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
					Value [4]byte
					Err   error
				}
				Sequence   uint32
				DoFn       MoqGetsockoptInet4Addr_genType_doFn
				DoReturnFn MoqGetsockoptInet4Addr_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetsockoptInet4Addr_genType) PrettyParams(params MoqGetsockoptInet4Addr_genType_params) string {
	return fmt.Sprintf("GetsockoptInet4Addr_genType(%#v, %#v, %#v)", params.Fd, params.Level, params.Opt)
}

func (m *MoqGetsockoptInet4Addr_genType) ParamsKey(params MoqGetsockoptInet4Addr_genType_params, anyParams uint64) MoqGetsockoptInet4Addr_genType_paramsKey {
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
	var levelUsed int
	var levelUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Level == moq.ParamIndexByValue {
			levelUsed = params.Level
		} else {
			levelUsedHash = hash.DeepHash(params.Level)
		}
	}
	var optUsed int
	var optUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Opt == moq.ParamIndexByValue {
			optUsed = params.Opt
		} else {
			optUsedHash = hash.DeepHash(params.Opt)
		}
	}
	return MoqGetsockoptInet4Addr_genType_paramsKey{
		Params: struct{ Fd, Level, Opt int }{
			Fd:    fdUsed,
			Level: levelUsed,
			Opt:   optUsed,
		},
		Hashes: struct{ Fd, Level, Opt hash.Hash }{
			Fd:    fdUsedHash,
			Level: levelUsedHash,
			Opt:   optUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGetsockoptInet4Addr_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetsockoptInet4Addr_genType) AssertExpectationsMet() {
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