// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Mprotect_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Mprotect_genType func(b []byte, prot int) (err error)

// MoqMprotect_genType holds the state of a moq of the Mprotect_genType type
type MoqMprotect_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMprotect_genType_mock

	ResultsByParams []MoqMprotect_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			B    moq.ParamIndexing
			Prot moq.ParamIndexing
		}
	}
}

// MoqMprotect_genType_mock isolates the mock interface of the Mprotect_genType
// type
type MoqMprotect_genType_mock struct {
	Moq *MoqMprotect_genType
}

// MoqMprotect_genType_params holds the params of the Mprotect_genType type
type MoqMprotect_genType_params struct {
	B    []byte
	Prot int
}

// MoqMprotect_genType_paramsKey holds the map key params of the
// Mprotect_genType type
type MoqMprotect_genType_paramsKey struct {
	Params struct{ Prot int }
	Hashes struct {
		B    hash.Hash
		Prot hash.Hash
	}
}

// MoqMprotect_genType_resultsByParams contains the results for a given set of
// parameters for the Mprotect_genType type
type MoqMprotect_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMprotect_genType_paramsKey]*MoqMprotect_genType_results
}

// MoqMprotect_genType_doFn defines the type of function needed when calling
// AndDo for the Mprotect_genType type
type MoqMprotect_genType_doFn func(b []byte, prot int)

// MoqMprotect_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Mprotect_genType type
type MoqMprotect_genType_doReturnFn func(b []byte, prot int) (err error)

// MoqMprotect_genType_results holds the results of the Mprotect_genType type
type MoqMprotect_genType_results struct {
	Params  MoqMprotect_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMprotect_genType_doFn
		DoReturnFn MoqMprotect_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMprotect_genType_fnRecorder routes recorded function calls to the
// MoqMprotect_genType moq
type MoqMprotect_genType_fnRecorder struct {
	Params    MoqMprotect_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMprotect_genType_results
	Moq       *MoqMprotect_genType
}

// MoqMprotect_genType_anyParams isolates the any params functions of the
// Mprotect_genType type
type MoqMprotect_genType_anyParams struct {
	Recorder *MoqMprotect_genType_fnRecorder
}

// NewMoqMprotect_genType creates a new moq of the Mprotect_genType type
func NewMoqMprotect_genType(scene *moq.Scene, config *moq.Config) *MoqMprotect_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMprotect_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMprotect_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				B    moq.ParamIndexing
				Prot moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			B    moq.ParamIndexing
			Prot moq.ParamIndexing
		}{
			B:    moq.ParamIndexByHash,
			Prot: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Mprotect_genType type
func (m *MoqMprotect_genType) Mock() Mprotect_genType {
	return func(b []byte, prot int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqMprotect_genType_mock{Moq: m}
		return moq.Fn(b, prot)
	}
}

func (m *MoqMprotect_genType_mock) Fn(b []byte, prot int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqMprotect_genType_params{
		B:    b,
		Prot: prot,
	}
	var results *MoqMprotect_genType_results
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
		result.DoFn(b, prot)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(b, prot)
	}
	return
}

func (m *MoqMprotect_genType) OnCall(b []byte, prot int) *MoqMprotect_genType_fnRecorder {
	return &MoqMprotect_genType_fnRecorder{
		Params: MoqMprotect_genType_params{
			B:    b,
			Prot: prot,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMprotect_genType_fnRecorder) Any() *MoqMprotect_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMprotect_genType_anyParams{Recorder: r}
}

func (a *MoqMprotect_genType_anyParams) B() *MoqMprotect_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqMprotect_genType_anyParams) Prot() *MoqMprotect_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqMprotect_genType_fnRecorder) Seq() *MoqMprotect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMprotect_genType_fnRecorder) NoSeq() *MoqMprotect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMprotect_genType_fnRecorder) ReturnResults(err error) *MoqMprotect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMprotect_genType_doFn
		DoReturnFn MoqMprotect_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMprotect_genType_fnRecorder) AndDo(fn MoqMprotect_genType_doFn) *MoqMprotect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMprotect_genType_fnRecorder) DoReturnResults(fn MoqMprotect_genType_doReturnFn) *MoqMprotect_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqMprotect_genType_doFn
		DoReturnFn MoqMprotect_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMprotect_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMprotect_genType_resultsByParams
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
		results = &MoqMprotect_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMprotect_genType_paramsKey]*MoqMprotect_genType_results{},
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
		r.Results = &MoqMprotect_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMprotect_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMprotect_genType_fnRecorder {
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
				DoFn       MoqMprotect_genType_doFn
				DoReturnFn MoqMprotect_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMprotect_genType) PrettyParams(params MoqMprotect_genType_params) string {
	return fmt.Sprintf("Mprotect_genType(%#v, %#v)", params.B, params.Prot)
}

func (m *MoqMprotect_genType) ParamsKey(params MoqMprotect_genType_params, anyParams uint64) MoqMprotect_genType_paramsKey {
	m.Scene.T.Helper()
	var bUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.B == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The b parameter can't be indexed by value")
		}
		bUsedHash = hash.DeepHash(params.B)
	}
	var protUsed int
	var protUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Prot == moq.ParamIndexByValue {
			protUsed = params.Prot
		} else {
			protUsedHash = hash.DeepHash(params.Prot)
		}
	}
	return MoqMprotect_genType_paramsKey{
		Params: struct{ Prot int }{
			Prot: protUsed,
		},
		Hashes: struct {
			B    hash.Hash
			Prot hash.Hash
		}{
			B:    bUsedHash,
			Prot: protUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMprotect_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMprotect_genType) AssertExpectationsMet() {
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
