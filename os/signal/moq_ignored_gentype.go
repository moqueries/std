// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package signal

import (
	"fmt"
	"math/bits"
	"os"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Ignored_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Ignored_genType func(sig os.Signal) bool

// MoqIgnored_genType holds the state of a moq of the Ignored_genType type
type MoqIgnored_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIgnored_genType_mock

	ResultsByParams []MoqIgnored_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Sig moq.ParamIndexing
		}
	}
}

// MoqIgnored_genType_mock isolates the mock interface of the Ignored_genType
// type
type MoqIgnored_genType_mock struct {
	Moq *MoqIgnored_genType
}

// MoqIgnored_genType_params holds the params of the Ignored_genType type
type MoqIgnored_genType_params struct{ Sig os.Signal }

// MoqIgnored_genType_paramsKey holds the map key params of the Ignored_genType
// type
type MoqIgnored_genType_paramsKey struct {
	Params struct{ Sig os.Signal }
	Hashes struct{ Sig hash.Hash }
}

// MoqIgnored_genType_resultsByParams contains the results for a given set of
// parameters for the Ignored_genType type
type MoqIgnored_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIgnored_genType_paramsKey]*MoqIgnored_genType_results
}

// MoqIgnored_genType_doFn defines the type of function needed when calling
// AndDo for the Ignored_genType type
type MoqIgnored_genType_doFn func(sig os.Signal)

// MoqIgnored_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Ignored_genType type
type MoqIgnored_genType_doReturnFn func(sig os.Signal) bool

// MoqIgnored_genType_results holds the results of the Ignored_genType type
type MoqIgnored_genType_results struct {
	Params  MoqIgnored_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIgnored_genType_doFn
		DoReturnFn MoqIgnored_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIgnored_genType_fnRecorder routes recorded function calls to the
// MoqIgnored_genType moq
type MoqIgnored_genType_fnRecorder struct {
	Params    MoqIgnored_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIgnored_genType_results
	Moq       *MoqIgnored_genType
}

// MoqIgnored_genType_anyParams isolates the any params functions of the
// Ignored_genType type
type MoqIgnored_genType_anyParams struct {
	Recorder *MoqIgnored_genType_fnRecorder
}

// NewMoqIgnored_genType creates a new moq of the Ignored_genType type
func NewMoqIgnored_genType(scene *moq.Scene, config *moq.Config) *MoqIgnored_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIgnored_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIgnored_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Sig moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Sig moq.ParamIndexing
		}{
			Sig: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Ignored_genType type
func (m *MoqIgnored_genType) Mock() Ignored_genType {
	return func(sig os.Signal) bool {
		m.Scene.T.Helper()
		moq := &MoqIgnored_genType_mock{Moq: m}
		return moq.Fn(sig)
	}
}

func (m *MoqIgnored_genType_mock) Fn(sig os.Signal) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIgnored_genType_params{
		Sig: sig,
	}
	var results *MoqIgnored_genType_results
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
		result.DoFn(sig)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(sig)
	}
	return
}

func (m *MoqIgnored_genType) OnCall(sig os.Signal) *MoqIgnored_genType_fnRecorder {
	return &MoqIgnored_genType_fnRecorder{
		Params: MoqIgnored_genType_params{
			Sig: sig,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIgnored_genType_fnRecorder) Any() *MoqIgnored_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIgnored_genType_anyParams{Recorder: r}
}

func (a *MoqIgnored_genType_anyParams) Sig() *MoqIgnored_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIgnored_genType_fnRecorder) Seq() *MoqIgnored_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIgnored_genType_fnRecorder) NoSeq() *MoqIgnored_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIgnored_genType_fnRecorder) ReturnResults(result1 bool) *MoqIgnored_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIgnored_genType_doFn
		DoReturnFn MoqIgnored_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIgnored_genType_fnRecorder) AndDo(fn MoqIgnored_genType_doFn) *MoqIgnored_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIgnored_genType_fnRecorder) DoReturnResults(fn MoqIgnored_genType_doReturnFn) *MoqIgnored_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIgnored_genType_doFn
		DoReturnFn MoqIgnored_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIgnored_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIgnored_genType_resultsByParams
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
		results = &MoqIgnored_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIgnored_genType_paramsKey]*MoqIgnored_genType_results{},
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
		r.Results = &MoqIgnored_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIgnored_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIgnored_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqIgnored_genType_doFn
				DoReturnFn MoqIgnored_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIgnored_genType) PrettyParams(params MoqIgnored_genType_params) string {
	return fmt.Sprintf("Ignored_genType(%#v)", params.Sig)
}

func (m *MoqIgnored_genType) ParamsKey(params MoqIgnored_genType_params, anyParams uint64) MoqIgnored_genType_paramsKey {
	m.Scene.T.Helper()
	var sigUsed os.Signal
	var sigUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Sig == moq.ParamIndexByValue {
			sigUsed = params.Sig
		} else {
			sigUsedHash = hash.DeepHash(params.Sig)
		}
	}
	return MoqIgnored_genType_paramsKey{
		Params: struct{ Sig os.Signal }{
			Sig: sigUsed,
		},
		Hashes: struct{ Sig hash.Hash }{
			Sig: sigUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIgnored_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIgnored_genType) AssertExpectationsMet() {
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