// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Umask_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Umask_genType func(mask int) (oldmask int)

// MoqUmask_genType holds the state of a moq of the Umask_genType type
type MoqUmask_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUmask_genType_mock

	ResultsByParams []MoqUmask_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Mask moq.ParamIndexing
		}
	}
}

// MoqUmask_genType_mock isolates the mock interface of the Umask_genType type
type MoqUmask_genType_mock struct {
	Moq *MoqUmask_genType
}

// MoqUmask_genType_params holds the params of the Umask_genType type
type MoqUmask_genType_params struct{ Mask int }

// MoqUmask_genType_paramsKey holds the map key params of the Umask_genType
// type
type MoqUmask_genType_paramsKey struct {
	Params struct{ Mask int }
	Hashes struct{ Mask hash.Hash }
}

// MoqUmask_genType_resultsByParams contains the results for a given set of
// parameters for the Umask_genType type
type MoqUmask_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUmask_genType_paramsKey]*MoqUmask_genType_results
}

// MoqUmask_genType_doFn defines the type of function needed when calling AndDo
// for the Umask_genType type
type MoqUmask_genType_doFn func(mask int)

// MoqUmask_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Umask_genType type
type MoqUmask_genType_doReturnFn func(mask int) (oldmask int)

// MoqUmask_genType_results holds the results of the Umask_genType type
type MoqUmask_genType_results struct {
	Params  MoqUmask_genType_params
	Results []struct {
		Values     *struct{ Oldmask int }
		Sequence   uint32
		DoFn       MoqUmask_genType_doFn
		DoReturnFn MoqUmask_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUmask_genType_fnRecorder routes recorded function calls to the
// MoqUmask_genType moq
type MoqUmask_genType_fnRecorder struct {
	Params    MoqUmask_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUmask_genType_results
	Moq       *MoqUmask_genType
}

// MoqUmask_genType_anyParams isolates the any params functions of the
// Umask_genType type
type MoqUmask_genType_anyParams struct {
	Recorder *MoqUmask_genType_fnRecorder
}

// NewMoqUmask_genType creates a new moq of the Umask_genType type
func NewMoqUmask_genType(scene *moq.Scene, config *moq.Config) *MoqUmask_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUmask_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUmask_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Mask moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Mask moq.ParamIndexing
		}{
			Mask: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Umask_genType type
func (m *MoqUmask_genType) Mock() Umask_genType {
	return func(mask int) (_ int) { m.Scene.T.Helper(); moq := &MoqUmask_genType_mock{Moq: m}; return moq.Fn(mask) }
}

func (m *MoqUmask_genType_mock) Fn(mask int) (oldmask int) {
	m.Moq.Scene.T.Helper()
	params := MoqUmask_genType_params{
		Mask: mask,
	}
	var results *MoqUmask_genType_results
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
		result.DoFn(mask)
	}

	if result.Values != nil {
		oldmask = result.Values.Oldmask
	}
	if result.DoReturnFn != nil {
		oldmask = result.DoReturnFn(mask)
	}
	return
}

func (m *MoqUmask_genType) OnCall(mask int) *MoqUmask_genType_fnRecorder {
	return &MoqUmask_genType_fnRecorder{
		Params: MoqUmask_genType_params{
			Mask: mask,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUmask_genType_fnRecorder) Any() *MoqUmask_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUmask_genType_anyParams{Recorder: r}
}

func (a *MoqUmask_genType_anyParams) Mask() *MoqUmask_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqUmask_genType_fnRecorder) Seq() *MoqUmask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUmask_genType_fnRecorder) NoSeq() *MoqUmask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUmask_genType_fnRecorder) ReturnResults(oldmask int) *MoqUmask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Oldmask int }
		Sequence   uint32
		DoFn       MoqUmask_genType_doFn
		DoReturnFn MoqUmask_genType_doReturnFn
	}{
		Values: &struct{ Oldmask int }{
			Oldmask: oldmask,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUmask_genType_fnRecorder) AndDo(fn MoqUmask_genType_doFn) *MoqUmask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUmask_genType_fnRecorder) DoReturnResults(fn MoqUmask_genType_doReturnFn) *MoqUmask_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Oldmask int }
		Sequence   uint32
		DoFn       MoqUmask_genType_doFn
		DoReturnFn MoqUmask_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUmask_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUmask_genType_resultsByParams
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
		results = &MoqUmask_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUmask_genType_paramsKey]*MoqUmask_genType_results{},
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
		r.Results = &MoqUmask_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUmask_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUmask_genType_fnRecorder {
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
				Values     *struct{ Oldmask int }
				Sequence   uint32
				DoFn       MoqUmask_genType_doFn
				DoReturnFn MoqUmask_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUmask_genType) PrettyParams(params MoqUmask_genType_params) string {
	return fmt.Sprintf("Umask_genType(%#v)", params.Mask)
}

func (m *MoqUmask_genType) ParamsKey(params MoqUmask_genType_params, anyParams uint64) MoqUmask_genType_paramsKey {
	m.Scene.T.Helper()
	var maskUsed int
	var maskUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Mask == moq.ParamIndexByValue {
			maskUsed = params.Mask
		} else {
			maskUsedHash = hash.DeepHash(params.Mask)
		}
	}
	return MoqUmask_genType_paramsKey{
		Params: struct{ Mask int }{
			Mask: maskUsed,
		},
		Hashes: struct{ Mask hash.Hash }{
			Mask: maskUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqUmask_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUmask_genType) AssertExpectationsMet() {
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