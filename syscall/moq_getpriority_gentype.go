// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Getpriority_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Getpriority_genType func(which int, who int) (prio int, err error)

// MoqGetpriority_genType holds the state of a moq of the Getpriority_genType
// type
type MoqGetpriority_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetpriority_genType_mock

	ResultsByParams []MoqGetpriority_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Which moq.ParamIndexing
			Who   moq.ParamIndexing
		}
	}
}

// MoqGetpriority_genType_mock isolates the mock interface of the
// Getpriority_genType type
type MoqGetpriority_genType_mock struct {
	Moq *MoqGetpriority_genType
}

// MoqGetpriority_genType_params holds the params of the Getpriority_genType
// type
type MoqGetpriority_genType_params struct {
	Which int
	Who   int
}

// MoqGetpriority_genType_paramsKey holds the map key params of the
// Getpriority_genType type
type MoqGetpriority_genType_paramsKey struct {
	Params struct {
		Which int
		Who   int
	}
	Hashes struct {
		Which hash.Hash
		Who   hash.Hash
	}
}

// MoqGetpriority_genType_resultsByParams contains the results for a given set
// of parameters for the Getpriority_genType type
type MoqGetpriority_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetpriority_genType_paramsKey]*MoqGetpriority_genType_results
}

// MoqGetpriority_genType_doFn defines the type of function needed when calling
// AndDo for the Getpriority_genType type
type MoqGetpriority_genType_doFn func(which int, who int)

// MoqGetpriority_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Getpriority_genType type
type MoqGetpriority_genType_doReturnFn func(which int, who int) (prio int, err error)

// MoqGetpriority_genType_results holds the results of the Getpriority_genType
// type
type MoqGetpriority_genType_results struct {
	Params  MoqGetpriority_genType_params
	Results []struct {
		Values *struct {
			Prio int
			Err  error
		}
		Sequence   uint32
		DoFn       MoqGetpriority_genType_doFn
		DoReturnFn MoqGetpriority_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetpriority_genType_fnRecorder routes recorded function calls to the
// MoqGetpriority_genType moq
type MoqGetpriority_genType_fnRecorder struct {
	Params    MoqGetpriority_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetpriority_genType_results
	Moq       *MoqGetpriority_genType
}

// MoqGetpriority_genType_anyParams isolates the any params functions of the
// Getpriority_genType type
type MoqGetpriority_genType_anyParams struct {
	Recorder *MoqGetpriority_genType_fnRecorder
}

// NewMoqGetpriority_genType creates a new moq of the Getpriority_genType type
func NewMoqGetpriority_genType(scene *moq.Scene, config *moq.Config) *MoqGetpriority_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetpriority_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetpriority_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Which moq.ParamIndexing
				Who   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Which moq.ParamIndexing
			Who   moq.ParamIndexing
		}{
			Which: moq.ParamIndexByValue,
			Who:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Getpriority_genType type
func (m *MoqGetpriority_genType) Mock() Getpriority_genType {
	return func(which int, who int) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqGetpriority_genType_mock{Moq: m}
		return moq.Fn(which, who)
	}
}

func (m *MoqGetpriority_genType_mock) Fn(which int, who int) (prio int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqGetpriority_genType_params{
		Which: which,
		Who:   who,
	}
	var results *MoqGetpriority_genType_results
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
		result.DoFn(which, who)
	}

	if result.Values != nil {
		prio = result.Values.Prio
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		prio, err = result.DoReturnFn(which, who)
	}
	return
}

func (m *MoqGetpriority_genType) OnCall(which int, who int) *MoqGetpriority_genType_fnRecorder {
	return &MoqGetpriority_genType_fnRecorder{
		Params: MoqGetpriority_genType_params{
			Which: which,
			Who:   who,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetpriority_genType_fnRecorder) Any() *MoqGetpriority_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetpriority_genType_anyParams{Recorder: r}
}

func (a *MoqGetpriority_genType_anyParams) Which() *MoqGetpriority_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGetpriority_genType_anyParams) Who() *MoqGetpriority_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqGetpriority_genType_fnRecorder) Seq() *MoqGetpriority_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetpriority_genType_fnRecorder) NoSeq() *MoqGetpriority_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetpriority_genType_fnRecorder) ReturnResults(prio int, err error) *MoqGetpriority_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Prio int
			Err  error
		}
		Sequence   uint32
		DoFn       MoqGetpriority_genType_doFn
		DoReturnFn MoqGetpriority_genType_doReturnFn
	}{
		Values: &struct {
			Prio int
			Err  error
		}{
			Prio: prio,
			Err:  err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetpriority_genType_fnRecorder) AndDo(fn MoqGetpriority_genType_doFn) *MoqGetpriority_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetpriority_genType_fnRecorder) DoReturnResults(fn MoqGetpriority_genType_doReturnFn) *MoqGetpriority_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Prio int
			Err  error
		}
		Sequence   uint32
		DoFn       MoqGetpriority_genType_doFn
		DoReturnFn MoqGetpriority_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetpriority_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetpriority_genType_resultsByParams
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
		results = &MoqGetpriority_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetpriority_genType_paramsKey]*MoqGetpriority_genType_results{},
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
		r.Results = &MoqGetpriority_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetpriority_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetpriority_genType_fnRecorder {
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
					Prio int
					Err  error
				}
				Sequence   uint32
				DoFn       MoqGetpriority_genType_doFn
				DoReturnFn MoqGetpriority_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetpriority_genType) PrettyParams(params MoqGetpriority_genType_params) string {
	return fmt.Sprintf("Getpriority_genType(%#v, %#v)", params.Which, params.Who)
}

func (m *MoqGetpriority_genType) ParamsKey(params MoqGetpriority_genType_params, anyParams uint64) MoqGetpriority_genType_paramsKey {
	m.Scene.T.Helper()
	var whichUsed int
	var whichUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Which == moq.ParamIndexByValue {
			whichUsed = params.Which
		} else {
			whichUsedHash = hash.DeepHash(params.Which)
		}
	}
	var whoUsed int
	var whoUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Who == moq.ParamIndexByValue {
			whoUsed = params.Who
		} else {
			whoUsedHash = hash.DeepHash(params.Who)
		}
	}
	return MoqGetpriority_genType_paramsKey{
		Params: struct {
			Which int
			Who   int
		}{
			Which: whichUsed,
			Who:   whoUsed,
		},
		Hashes: struct {
			Which hash.Hash
			Who   hash.Hash
		}{
			Which: whichUsedHash,
			Who:   whoUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGetpriority_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetpriority_genType) AssertExpectationsMet() {
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
