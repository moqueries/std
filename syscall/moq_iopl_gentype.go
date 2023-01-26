// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Iopl_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Iopl_genType func(level int) (err error)

// MoqIopl_genType holds the state of a moq of the Iopl_genType type
type MoqIopl_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIopl_genType_mock

	ResultsByParams []MoqIopl_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Level moq.ParamIndexing
		}
	}
}

// MoqIopl_genType_mock isolates the mock interface of the Iopl_genType type
type MoqIopl_genType_mock struct {
	Moq *MoqIopl_genType
}

// MoqIopl_genType_params holds the params of the Iopl_genType type
type MoqIopl_genType_params struct{ Level int }

// MoqIopl_genType_paramsKey holds the map key params of the Iopl_genType type
type MoqIopl_genType_paramsKey struct {
	Params struct{ Level int }
	Hashes struct{ Level hash.Hash }
}

// MoqIopl_genType_resultsByParams contains the results for a given set of
// parameters for the Iopl_genType type
type MoqIopl_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIopl_genType_paramsKey]*MoqIopl_genType_results
}

// MoqIopl_genType_doFn defines the type of function needed when calling AndDo
// for the Iopl_genType type
type MoqIopl_genType_doFn func(level int)

// MoqIopl_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Iopl_genType type
type MoqIopl_genType_doReturnFn func(level int) (err error)

// MoqIopl_genType_results holds the results of the Iopl_genType type
type MoqIopl_genType_results struct {
	Params  MoqIopl_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqIopl_genType_doFn
		DoReturnFn MoqIopl_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIopl_genType_fnRecorder routes recorded function calls to the
// MoqIopl_genType moq
type MoqIopl_genType_fnRecorder struct {
	Params    MoqIopl_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIopl_genType_results
	Moq       *MoqIopl_genType
}

// MoqIopl_genType_anyParams isolates the any params functions of the
// Iopl_genType type
type MoqIopl_genType_anyParams struct {
	Recorder *MoqIopl_genType_fnRecorder
}

// NewMoqIopl_genType creates a new moq of the Iopl_genType type
func NewMoqIopl_genType(scene *moq.Scene, config *moq.Config) *MoqIopl_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIopl_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIopl_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Level moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Level moq.ParamIndexing
		}{
			Level: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Iopl_genType type
func (m *MoqIopl_genType) Mock() Iopl_genType {
	return func(level int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqIopl_genType_mock{Moq: m}
		return moq.Fn(level)
	}
}

func (m *MoqIopl_genType_mock) Fn(level int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqIopl_genType_params{
		Level: level,
	}
	var results *MoqIopl_genType_results
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
		result.DoFn(level)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(level)
	}
	return
}

func (m *MoqIopl_genType) OnCall(level int) *MoqIopl_genType_fnRecorder {
	return &MoqIopl_genType_fnRecorder{
		Params: MoqIopl_genType_params{
			Level: level,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIopl_genType_fnRecorder) Any() *MoqIopl_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIopl_genType_anyParams{Recorder: r}
}

func (a *MoqIopl_genType_anyParams) Level() *MoqIopl_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIopl_genType_fnRecorder) Seq() *MoqIopl_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIopl_genType_fnRecorder) NoSeq() *MoqIopl_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIopl_genType_fnRecorder) ReturnResults(err error) *MoqIopl_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqIopl_genType_doFn
		DoReturnFn MoqIopl_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIopl_genType_fnRecorder) AndDo(fn MoqIopl_genType_doFn) *MoqIopl_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIopl_genType_fnRecorder) DoReturnResults(fn MoqIopl_genType_doReturnFn) *MoqIopl_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqIopl_genType_doFn
		DoReturnFn MoqIopl_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIopl_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIopl_genType_resultsByParams
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
		results = &MoqIopl_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIopl_genType_paramsKey]*MoqIopl_genType_results{},
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
		r.Results = &MoqIopl_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIopl_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIopl_genType_fnRecorder {
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
				DoFn       MoqIopl_genType_doFn
				DoReturnFn MoqIopl_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIopl_genType) PrettyParams(params MoqIopl_genType_params) string {
	return fmt.Sprintf("Iopl_genType(%#v)", params.Level)
}

func (m *MoqIopl_genType) ParamsKey(params MoqIopl_genType_params, anyParams uint64) MoqIopl_genType_paramsKey {
	m.Scene.T.Helper()
	var levelUsed int
	var levelUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Level == moq.ParamIndexByValue {
			levelUsed = params.Level
		} else {
			levelUsedHash = hash.DeepHash(params.Level)
		}
	}
	return MoqIopl_genType_paramsKey{
		Params: struct{ Level int }{
			Level: levelUsed,
		},
		Hashes: struct{ Level hash.Hash }{
			Level: levelUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIopl_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIopl_genType) AssertExpectationsMet() {
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