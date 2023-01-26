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

// Getrusage_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Getrusage_genType func(who int, rusage *syscall.Rusage) (err error)

// MoqGetrusage_genType holds the state of a moq of the Getrusage_genType type
type MoqGetrusage_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetrusage_genType_mock

	ResultsByParams []MoqGetrusage_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Who    moq.ParamIndexing
			Rusage moq.ParamIndexing
		}
	}
}

// MoqGetrusage_genType_mock isolates the mock interface of the
// Getrusage_genType type
type MoqGetrusage_genType_mock struct {
	Moq *MoqGetrusage_genType
}

// MoqGetrusage_genType_params holds the params of the Getrusage_genType type
type MoqGetrusage_genType_params struct {
	Who    int
	Rusage *syscall.Rusage
}

// MoqGetrusage_genType_paramsKey holds the map key params of the
// Getrusage_genType type
type MoqGetrusage_genType_paramsKey struct {
	Params struct {
		Who    int
		Rusage *syscall.Rusage
	}
	Hashes struct {
		Who    hash.Hash
		Rusage hash.Hash
	}
}

// MoqGetrusage_genType_resultsByParams contains the results for a given set of
// parameters for the Getrusage_genType type
type MoqGetrusage_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetrusage_genType_paramsKey]*MoqGetrusage_genType_results
}

// MoqGetrusage_genType_doFn defines the type of function needed when calling
// AndDo for the Getrusage_genType type
type MoqGetrusage_genType_doFn func(who int, rusage *syscall.Rusage)

// MoqGetrusage_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Getrusage_genType type
type MoqGetrusage_genType_doReturnFn func(who int, rusage *syscall.Rusage) (err error)

// MoqGetrusage_genType_results holds the results of the Getrusage_genType type
type MoqGetrusage_genType_results struct {
	Params  MoqGetrusage_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqGetrusage_genType_doFn
		DoReturnFn MoqGetrusage_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetrusage_genType_fnRecorder routes recorded function calls to the
// MoqGetrusage_genType moq
type MoqGetrusage_genType_fnRecorder struct {
	Params    MoqGetrusage_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetrusage_genType_results
	Moq       *MoqGetrusage_genType
}

// MoqGetrusage_genType_anyParams isolates the any params functions of the
// Getrusage_genType type
type MoqGetrusage_genType_anyParams struct {
	Recorder *MoqGetrusage_genType_fnRecorder
}

// NewMoqGetrusage_genType creates a new moq of the Getrusage_genType type
func NewMoqGetrusage_genType(scene *moq.Scene, config *moq.Config) *MoqGetrusage_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetrusage_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetrusage_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Who    moq.ParamIndexing
				Rusage moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Who    moq.ParamIndexing
			Rusage moq.ParamIndexing
		}{
			Who:    moq.ParamIndexByValue,
			Rusage: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Getrusage_genType type
func (m *MoqGetrusage_genType) Mock() Getrusage_genType {
	return func(who int, rusage *syscall.Rusage) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqGetrusage_genType_mock{Moq: m}
		return moq.Fn(who, rusage)
	}
}

func (m *MoqGetrusage_genType_mock) Fn(who int, rusage *syscall.Rusage) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqGetrusage_genType_params{
		Who:    who,
		Rusage: rusage,
	}
	var results *MoqGetrusage_genType_results
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
		result.DoFn(who, rusage)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(who, rusage)
	}
	return
}

func (m *MoqGetrusage_genType) OnCall(who int, rusage *syscall.Rusage) *MoqGetrusage_genType_fnRecorder {
	return &MoqGetrusage_genType_fnRecorder{
		Params: MoqGetrusage_genType_params{
			Who:    who,
			Rusage: rusage,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetrusage_genType_fnRecorder) Any() *MoqGetrusage_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetrusage_genType_anyParams{Recorder: r}
}

func (a *MoqGetrusage_genType_anyParams) Who() *MoqGetrusage_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGetrusage_genType_anyParams) Rusage() *MoqGetrusage_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqGetrusage_genType_fnRecorder) Seq() *MoqGetrusage_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetrusage_genType_fnRecorder) NoSeq() *MoqGetrusage_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetrusage_genType_fnRecorder) ReturnResults(err error) *MoqGetrusage_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqGetrusage_genType_doFn
		DoReturnFn MoqGetrusage_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetrusage_genType_fnRecorder) AndDo(fn MoqGetrusage_genType_doFn) *MoqGetrusage_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetrusage_genType_fnRecorder) DoReturnResults(fn MoqGetrusage_genType_doReturnFn) *MoqGetrusage_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqGetrusage_genType_doFn
		DoReturnFn MoqGetrusage_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetrusage_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetrusage_genType_resultsByParams
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
		results = &MoqGetrusage_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetrusage_genType_paramsKey]*MoqGetrusage_genType_results{},
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
		r.Results = &MoqGetrusage_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetrusage_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetrusage_genType_fnRecorder {
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
				DoFn       MoqGetrusage_genType_doFn
				DoReturnFn MoqGetrusage_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetrusage_genType) PrettyParams(params MoqGetrusage_genType_params) string {
	return fmt.Sprintf("Getrusage_genType(%#v, %#v)", params.Who, params.Rusage)
}

func (m *MoqGetrusage_genType) ParamsKey(params MoqGetrusage_genType_params, anyParams uint64) MoqGetrusage_genType_paramsKey {
	m.Scene.T.Helper()
	var whoUsed int
	var whoUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Who == moq.ParamIndexByValue {
			whoUsed = params.Who
		} else {
			whoUsedHash = hash.DeepHash(params.Who)
		}
	}
	var rusageUsed *syscall.Rusage
	var rusageUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Rusage == moq.ParamIndexByValue {
			rusageUsed = params.Rusage
		} else {
			rusageUsedHash = hash.DeepHash(params.Rusage)
		}
	}
	return MoqGetrusage_genType_paramsKey{
		Params: struct {
			Who    int
			Rusage *syscall.Rusage
		}{
			Who:    whoUsed,
			Rusage: rusageUsed,
		},
		Hashes: struct {
			Who    hash.Hash
			Rusage hash.Hash
		}{
			Who:    whoUsedHash,
			Rusage: rusageUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGetrusage_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetrusage_genType) AssertExpectationsMet() {
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