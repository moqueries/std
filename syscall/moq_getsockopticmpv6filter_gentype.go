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

// GetsockoptICMPv6Filter_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type GetsockoptICMPv6Filter_genType func(fd, level, opt int) (*syscall.ICMPv6Filter, error)

// MoqGetsockoptICMPv6Filter_genType holds the state of a moq of the
// GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetsockoptICMPv6Filter_genType_mock

	ResultsByParams []MoqGetsockoptICMPv6Filter_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd    moq.ParamIndexing
			Level moq.ParamIndexing
			Opt   moq.ParamIndexing
		}
	}
}

// MoqGetsockoptICMPv6Filter_genType_mock isolates the mock interface of the
// GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_mock struct {
	Moq *MoqGetsockoptICMPv6Filter_genType
}

// MoqGetsockoptICMPv6Filter_genType_params holds the params of the
// GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_params struct{ Fd, Level, Opt int }

// MoqGetsockoptICMPv6Filter_genType_paramsKey holds the map key params of the
// GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_paramsKey struct {
	Params struct{ Fd, Level, Opt int }
	Hashes struct{ Fd, Level, Opt hash.Hash }
}

// MoqGetsockoptICMPv6Filter_genType_resultsByParams contains the results for a
// given set of parameters for the GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetsockoptICMPv6Filter_genType_paramsKey]*MoqGetsockoptICMPv6Filter_genType_results
}

// MoqGetsockoptICMPv6Filter_genType_doFn defines the type of function needed
// when calling AndDo for the GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_doFn func(fd, level, opt int)

// MoqGetsockoptICMPv6Filter_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the GetsockoptICMPv6Filter_genType
// type
type MoqGetsockoptICMPv6Filter_genType_doReturnFn func(fd, level, opt int) (*syscall.ICMPv6Filter, error)

// MoqGetsockoptICMPv6Filter_genType_results holds the results of the
// GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_results struct {
	Params  MoqGetsockoptICMPv6Filter_genType_params
	Results []struct {
		Values *struct {
			Result1 *syscall.ICMPv6Filter
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptICMPv6Filter_genType_doFn
		DoReturnFn MoqGetsockoptICMPv6Filter_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetsockoptICMPv6Filter_genType_fnRecorder routes recorded function calls
// to the MoqGetsockoptICMPv6Filter_genType moq
type MoqGetsockoptICMPv6Filter_genType_fnRecorder struct {
	Params    MoqGetsockoptICMPv6Filter_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetsockoptICMPv6Filter_genType_results
	Moq       *MoqGetsockoptICMPv6Filter_genType
}

// MoqGetsockoptICMPv6Filter_genType_anyParams isolates the any params
// functions of the GetsockoptICMPv6Filter_genType type
type MoqGetsockoptICMPv6Filter_genType_anyParams struct {
	Recorder *MoqGetsockoptICMPv6Filter_genType_fnRecorder
}

// NewMoqGetsockoptICMPv6Filter_genType creates a new moq of the
// GetsockoptICMPv6Filter_genType type
func NewMoqGetsockoptICMPv6Filter_genType(scene *moq.Scene, config *moq.Config) *MoqGetsockoptICMPv6Filter_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetsockoptICMPv6Filter_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetsockoptICMPv6Filter_genType_mock{},

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

// Mock returns the moq implementation of the GetsockoptICMPv6Filter_genType
// type
func (m *MoqGetsockoptICMPv6Filter_genType) Mock() GetsockoptICMPv6Filter_genType {
	return func(fd, level, opt int) (*syscall.ICMPv6Filter, error) {
		m.Scene.T.Helper()
		moq := &MoqGetsockoptICMPv6Filter_genType_mock{Moq: m}
		return moq.Fn(fd, level, opt)
	}
}

func (m *MoqGetsockoptICMPv6Filter_genType_mock) Fn(fd, level, opt int) (result1 *syscall.ICMPv6Filter, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqGetsockoptICMPv6Filter_genType_params{
		Fd:    fd,
		Level: level,
		Opt:   opt,
	}
	var results *MoqGetsockoptICMPv6Filter_genType_results
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
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(fd, level, opt)
	}
	return
}

func (m *MoqGetsockoptICMPv6Filter_genType) OnCall(fd, level, opt int) *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	return &MoqGetsockoptICMPv6Filter_genType_fnRecorder{
		Params: MoqGetsockoptICMPv6Filter_genType_params{
			Fd:    fd,
			Level: level,
			Opt:   opt,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) Any() *MoqGetsockoptICMPv6Filter_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetsockoptICMPv6Filter_genType_anyParams{Recorder: r}
}

func (a *MoqGetsockoptICMPv6Filter_genType_anyParams) Fd() *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGetsockoptICMPv6Filter_genType_anyParams) Level() *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqGetsockoptICMPv6Filter_genType_anyParams) Opt() *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) Seq() *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) NoSeq() *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) ReturnResults(result1 *syscall.ICMPv6Filter, result2 error) *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *syscall.ICMPv6Filter
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptICMPv6Filter_genType_doFn
		DoReturnFn MoqGetsockoptICMPv6Filter_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *syscall.ICMPv6Filter
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) AndDo(fn MoqGetsockoptICMPv6Filter_genType_doFn) *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) DoReturnResults(fn MoqGetsockoptICMPv6Filter_genType_doReturnFn) *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *syscall.ICMPv6Filter
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqGetsockoptICMPv6Filter_genType_doFn
		DoReturnFn MoqGetsockoptICMPv6Filter_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetsockoptICMPv6Filter_genType_resultsByParams
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
		results = &MoqGetsockoptICMPv6Filter_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetsockoptICMPv6Filter_genType_paramsKey]*MoqGetsockoptICMPv6Filter_genType_results{},
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
		r.Results = &MoqGetsockoptICMPv6Filter_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetsockoptICMPv6Filter_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetsockoptICMPv6Filter_genType_fnRecorder {
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
					Result1 *syscall.ICMPv6Filter
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqGetsockoptICMPv6Filter_genType_doFn
				DoReturnFn MoqGetsockoptICMPv6Filter_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetsockoptICMPv6Filter_genType) PrettyParams(params MoqGetsockoptICMPv6Filter_genType_params) string {
	return fmt.Sprintf("GetsockoptICMPv6Filter_genType(%#v, %#v, %#v)", params.Fd, params.Level, params.Opt)
}

func (m *MoqGetsockoptICMPv6Filter_genType) ParamsKey(params MoqGetsockoptICMPv6Filter_genType_params, anyParams uint64) MoqGetsockoptICMPv6Filter_genType_paramsKey {
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
	return MoqGetsockoptICMPv6Filter_genType_paramsKey{
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
func (m *MoqGetsockoptICMPv6Filter_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetsockoptICMPv6Filter_genType) AssertExpectationsMet() {
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
