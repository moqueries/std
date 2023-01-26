// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SetsockoptByte_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SetsockoptByte_genType func(fd, level, opt int, value byte) (err error)

// MoqSetsockoptByte_genType holds the state of a moq of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetsockoptByte_genType_mock

	ResultsByParams []MoqSetsockoptByte_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd    moq.ParamIndexing
			Level moq.ParamIndexing
			Opt   moq.ParamIndexing
			Value moq.ParamIndexing
		}
	}
}

// MoqSetsockoptByte_genType_mock isolates the mock interface of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType_mock struct {
	Moq *MoqSetsockoptByte_genType
}

// MoqSetsockoptByte_genType_params holds the params of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType_params struct {
	Fd, Level, Opt int
	Value          byte
}

// MoqSetsockoptByte_genType_paramsKey holds the map key params of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType_paramsKey struct {
	Params struct {
		Fd, Level, Opt int
		Value          byte
	}
	Hashes struct {
		Fd, Level, Opt hash.Hash
		Value          hash.Hash
	}
}

// MoqSetsockoptByte_genType_resultsByParams contains the results for a given
// set of parameters for the SetsockoptByte_genType type
type MoqSetsockoptByte_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetsockoptByte_genType_paramsKey]*MoqSetsockoptByte_genType_results
}

// MoqSetsockoptByte_genType_doFn defines the type of function needed when
// calling AndDo for the SetsockoptByte_genType type
type MoqSetsockoptByte_genType_doFn func(fd, level, opt int, value byte)

// MoqSetsockoptByte_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the SetsockoptByte_genType type
type MoqSetsockoptByte_genType_doReturnFn func(fd, level, opt int, value byte) (err error)

// MoqSetsockoptByte_genType_results holds the results of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType_results struct {
	Params  MoqSetsockoptByte_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetsockoptByte_genType_doFn
		DoReturnFn MoqSetsockoptByte_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetsockoptByte_genType_fnRecorder routes recorded function calls to the
// MoqSetsockoptByte_genType moq
type MoqSetsockoptByte_genType_fnRecorder struct {
	Params    MoqSetsockoptByte_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetsockoptByte_genType_results
	Moq       *MoqSetsockoptByte_genType
}

// MoqSetsockoptByte_genType_anyParams isolates the any params functions of the
// SetsockoptByte_genType type
type MoqSetsockoptByte_genType_anyParams struct {
	Recorder *MoqSetsockoptByte_genType_fnRecorder
}

// NewMoqSetsockoptByte_genType creates a new moq of the SetsockoptByte_genType
// type
func NewMoqSetsockoptByte_genType(scene *moq.Scene, config *moq.Config) *MoqSetsockoptByte_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetsockoptByte_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetsockoptByte_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd    moq.ParamIndexing
				Level moq.ParamIndexing
				Opt   moq.ParamIndexing
				Value moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd    moq.ParamIndexing
			Level moq.ParamIndexing
			Opt   moq.ParamIndexing
			Value moq.ParamIndexing
		}{
			Fd:    moq.ParamIndexByValue,
			Level: moq.ParamIndexByValue,
			Opt:   moq.ParamIndexByValue,
			Value: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SetsockoptByte_genType type
func (m *MoqSetsockoptByte_genType) Mock() SetsockoptByte_genType {
	return func(fd, level, opt int, value byte) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqSetsockoptByte_genType_mock{Moq: m}
		return moq.Fn(fd, level, opt, value)
	}
}

func (m *MoqSetsockoptByte_genType_mock) Fn(fd, level, opt int, value byte) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqSetsockoptByte_genType_params{
		Fd:    fd,
		Level: level,
		Opt:   opt,
		Value: value,
	}
	var results *MoqSetsockoptByte_genType_results
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
		result.DoFn(fd, level, opt, value)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(fd, level, opt, value)
	}
	return
}

func (m *MoqSetsockoptByte_genType) OnCall(fd, level, opt int, value byte) *MoqSetsockoptByte_genType_fnRecorder {
	return &MoqSetsockoptByte_genType_fnRecorder{
		Params: MoqSetsockoptByte_genType_params{
			Fd:    fd,
			Level: level,
			Opt:   opt,
			Value: value,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetsockoptByte_genType_fnRecorder) Any() *MoqSetsockoptByte_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetsockoptByte_genType_anyParams{Recorder: r}
}

func (a *MoqSetsockoptByte_genType_anyParams) Fd() *MoqSetsockoptByte_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSetsockoptByte_genType_anyParams) Level() *MoqSetsockoptByte_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqSetsockoptByte_genType_anyParams) Opt() *MoqSetsockoptByte_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqSetsockoptByte_genType_anyParams) Value() *MoqSetsockoptByte_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqSetsockoptByte_genType_fnRecorder) Seq() *MoqSetsockoptByte_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetsockoptByte_genType_fnRecorder) NoSeq() *MoqSetsockoptByte_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetsockoptByte_genType_fnRecorder) ReturnResults(err error) *MoqSetsockoptByte_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetsockoptByte_genType_doFn
		DoReturnFn MoqSetsockoptByte_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetsockoptByte_genType_fnRecorder) AndDo(fn MoqSetsockoptByte_genType_doFn) *MoqSetsockoptByte_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetsockoptByte_genType_fnRecorder) DoReturnResults(fn MoqSetsockoptByte_genType_doReturnFn) *MoqSetsockoptByte_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetsockoptByte_genType_doFn
		DoReturnFn MoqSetsockoptByte_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetsockoptByte_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetsockoptByte_genType_resultsByParams
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
		results = &MoqSetsockoptByte_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetsockoptByte_genType_paramsKey]*MoqSetsockoptByte_genType_results{},
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
		r.Results = &MoqSetsockoptByte_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetsockoptByte_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetsockoptByte_genType_fnRecorder {
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
				DoFn       MoqSetsockoptByte_genType_doFn
				DoReturnFn MoqSetsockoptByte_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetsockoptByte_genType) PrettyParams(params MoqSetsockoptByte_genType_params) string {
	return fmt.Sprintf("SetsockoptByte_genType(%#v, %#v, %#v, %#v)", params.Fd, params.Level, params.Opt, params.Value)
}

func (m *MoqSetsockoptByte_genType) ParamsKey(params MoqSetsockoptByte_genType_params, anyParams uint64) MoqSetsockoptByte_genType_paramsKey {
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
	var valueUsed byte
	var valueUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	return MoqSetsockoptByte_genType_paramsKey{
		Params: struct {
			Fd, Level, Opt int
			Value          byte
		}{
			Fd:    fdUsed,
			Level: levelUsed,
			Opt:   optUsed,
			Value: valueUsed,
		},
		Hashes: struct {
			Fd, Level, Opt hash.Hash
			Value          hash.Hash
		}{
			Fd:    fdUsedHash,
			Level: levelUsedHash,
			Opt:   optUsedHash,
			Value: valueUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetsockoptByte_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetsockoptByte_genType) AssertExpectationsMet() {
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