// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SetBlockProfileRate_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type SetBlockProfileRate_genType func(rate int)

// MoqSetBlockProfileRate_genType holds the state of a moq of the
// SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetBlockProfileRate_genType_mock

	ResultsByParams []MoqSetBlockProfileRate_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rate moq.ParamIndexing
		}
	}
}

// MoqSetBlockProfileRate_genType_mock isolates the mock interface of the
// SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_mock struct {
	Moq *MoqSetBlockProfileRate_genType
}

// MoqSetBlockProfileRate_genType_params holds the params of the
// SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_params struct{ Rate int }

// MoqSetBlockProfileRate_genType_paramsKey holds the map key params of the
// SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_paramsKey struct {
	Params struct{ Rate int }
	Hashes struct{ Rate hash.Hash }
}

// MoqSetBlockProfileRate_genType_resultsByParams contains the results for a
// given set of parameters for the SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetBlockProfileRate_genType_paramsKey]*MoqSetBlockProfileRate_genType_results
}

// MoqSetBlockProfileRate_genType_doFn defines the type of function needed when
// calling AndDo for the SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_doFn func(rate int)

// MoqSetBlockProfileRate_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_doReturnFn func(rate int)

// MoqSetBlockProfileRate_genType_results holds the results of the
// SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_results struct {
	Params  MoqSetBlockProfileRate_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetBlockProfileRate_genType_doFn
		DoReturnFn MoqSetBlockProfileRate_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetBlockProfileRate_genType_fnRecorder routes recorded function calls to
// the MoqSetBlockProfileRate_genType moq
type MoqSetBlockProfileRate_genType_fnRecorder struct {
	Params    MoqSetBlockProfileRate_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetBlockProfileRate_genType_results
	Moq       *MoqSetBlockProfileRate_genType
}

// MoqSetBlockProfileRate_genType_anyParams isolates the any params functions
// of the SetBlockProfileRate_genType type
type MoqSetBlockProfileRate_genType_anyParams struct {
	Recorder *MoqSetBlockProfileRate_genType_fnRecorder
}

// NewMoqSetBlockProfileRate_genType creates a new moq of the
// SetBlockProfileRate_genType type
func NewMoqSetBlockProfileRate_genType(scene *moq.Scene, config *moq.Config) *MoqSetBlockProfileRate_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetBlockProfileRate_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetBlockProfileRate_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rate moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rate moq.ParamIndexing
		}{
			Rate: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SetBlockProfileRate_genType type
func (m *MoqSetBlockProfileRate_genType) Mock() SetBlockProfileRate_genType {
	return func(rate int) { m.Scene.T.Helper(); moq := &MoqSetBlockProfileRate_genType_mock{Moq: m}; moq.Fn(rate) }
}

func (m *MoqSetBlockProfileRate_genType_mock) Fn(rate int) {
	m.Moq.Scene.T.Helper()
	params := MoqSetBlockProfileRate_genType_params{
		Rate: rate,
	}
	var results *MoqSetBlockProfileRate_genType_results
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
		result.DoFn(rate)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(rate)
	}
	return
}

func (m *MoqSetBlockProfileRate_genType) OnCall(rate int) *MoqSetBlockProfileRate_genType_fnRecorder {
	return &MoqSetBlockProfileRate_genType_fnRecorder{
		Params: MoqSetBlockProfileRate_genType_params{
			Rate: rate,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) Any() *MoqSetBlockProfileRate_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetBlockProfileRate_genType_anyParams{Recorder: r}
}

func (a *MoqSetBlockProfileRate_genType_anyParams) Rate() *MoqSetBlockProfileRate_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) Seq() *MoqSetBlockProfileRate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) NoSeq() *MoqSetBlockProfileRate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) ReturnResults() *MoqSetBlockProfileRate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetBlockProfileRate_genType_doFn
		DoReturnFn MoqSetBlockProfileRate_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) AndDo(fn MoqSetBlockProfileRate_genType_doFn) *MoqSetBlockProfileRate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) DoReturnResults(fn MoqSetBlockProfileRate_genType_doReturnFn) *MoqSetBlockProfileRate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetBlockProfileRate_genType_doFn
		DoReturnFn MoqSetBlockProfileRate_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetBlockProfileRate_genType_resultsByParams
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
		results = &MoqSetBlockProfileRate_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetBlockProfileRate_genType_paramsKey]*MoqSetBlockProfileRate_genType_results{},
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
		r.Results = &MoqSetBlockProfileRate_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetBlockProfileRate_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetBlockProfileRate_genType_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqSetBlockProfileRate_genType_doFn
				DoReturnFn MoqSetBlockProfileRate_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetBlockProfileRate_genType) PrettyParams(params MoqSetBlockProfileRate_genType_params) string {
	return fmt.Sprintf("SetBlockProfileRate_genType(%#v)", params.Rate)
}

func (m *MoqSetBlockProfileRate_genType) ParamsKey(params MoqSetBlockProfileRate_genType_params, anyParams uint64) MoqSetBlockProfileRate_genType_paramsKey {
	m.Scene.T.Helper()
	var rateUsed int
	var rateUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Rate == moq.ParamIndexByValue {
			rateUsed = params.Rate
		} else {
			rateUsedHash = hash.DeepHash(params.Rate)
		}
	}
	return MoqSetBlockProfileRate_genType_paramsKey{
		Params: struct{ Rate int }{
			Rate: rateUsed,
		},
		Hashes: struct{ Rate hash.Hash }{
			Rate: rateUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetBlockProfileRate_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetBlockProfileRate_genType) AssertExpectationsMet() {
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