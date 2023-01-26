// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Setuid_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Setuid_genType func(uid int) (err error)

// MoqSetuid_genType holds the state of a moq of the Setuid_genType type
type MoqSetuid_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetuid_genType_mock

	ResultsByParams []MoqSetuid_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Uid moq.ParamIndexing
		}
	}
}

// MoqSetuid_genType_mock isolates the mock interface of the Setuid_genType
// type
type MoqSetuid_genType_mock struct {
	Moq *MoqSetuid_genType
}

// MoqSetuid_genType_params holds the params of the Setuid_genType type
type MoqSetuid_genType_params struct{ Uid int }

// MoqSetuid_genType_paramsKey holds the map key params of the Setuid_genType
// type
type MoqSetuid_genType_paramsKey struct {
	Params struct{ Uid int }
	Hashes struct{ Uid hash.Hash }
}

// MoqSetuid_genType_resultsByParams contains the results for a given set of
// parameters for the Setuid_genType type
type MoqSetuid_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetuid_genType_paramsKey]*MoqSetuid_genType_results
}

// MoqSetuid_genType_doFn defines the type of function needed when calling
// AndDo for the Setuid_genType type
type MoqSetuid_genType_doFn func(uid int)

// MoqSetuid_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Setuid_genType type
type MoqSetuid_genType_doReturnFn func(uid int) (err error)

// MoqSetuid_genType_results holds the results of the Setuid_genType type
type MoqSetuid_genType_results struct {
	Params  MoqSetuid_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetuid_genType_doFn
		DoReturnFn MoqSetuid_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetuid_genType_fnRecorder routes recorded function calls to the
// MoqSetuid_genType moq
type MoqSetuid_genType_fnRecorder struct {
	Params    MoqSetuid_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetuid_genType_results
	Moq       *MoqSetuid_genType
}

// MoqSetuid_genType_anyParams isolates the any params functions of the
// Setuid_genType type
type MoqSetuid_genType_anyParams struct {
	Recorder *MoqSetuid_genType_fnRecorder
}

// NewMoqSetuid_genType creates a new moq of the Setuid_genType type
func NewMoqSetuid_genType(scene *moq.Scene, config *moq.Config) *MoqSetuid_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetuid_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetuid_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Uid moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Uid moq.ParamIndexing
		}{
			Uid: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Setuid_genType type
func (m *MoqSetuid_genType) Mock() Setuid_genType {
	return func(uid int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqSetuid_genType_mock{Moq: m}
		return moq.Fn(uid)
	}
}

func (m *MoqSetuid_genType_mock) Fn(uid int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqSetuid_genType_params{
		Uid: uid,
	}
	var results *MoqSetuid_genType_results
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
		result.DoFn(uid)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(uid)
	}
	return
}

func (m *MoqSetuid_genType) OnCall(uid int) *MoqSetuid_genType_fnRecorder {
	return &MoqSetuid_genType_fnRecorder{
		Params: MoqSetuid_genType_params{
			Uid: uid,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetuid_genType_fnRecorder) Any() *MoqSetuid_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetuid_genType_anyParams{Recorder: r}
}

func (a *MoqSetuid_genType_anyParams) Uid() *MoqSetuid_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSetuid_genType_fnRecorder) Seq() *MoqSetuid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetuid_genType_fnRecorder) NoSeq() *MoqSetuid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetuid_genType_fnRecorder) ReturnResults(err error) *MoqSetuid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetuid_genType_doFn
		DoReturnFn MoqSetuid_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetuid_genType_fnRecorder) AndDo(fn MoqSetuid_genType_doFn) *MoqSetuid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetuid_genType_fnRecorder) DoReturnResults(fn MoqSetuid_genType_doReturnFn) *MoqSetuid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetuid_genType_doFn
		DoReturnFn MoqSetuid_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetuid_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetuid_genType_resultsByParams
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
		results = &MoqSetuid_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetuid_genType_paramsKey]*MoqSetuid_genType_results{},
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
		r.Results = &MoqSetuid_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetuid_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetuid_genType_fnRecorder {
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
				DoFn       MoqSetuid_genType_doFn
				DoReturnFn MoqSetuid_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetuid_genType) PrettyParams(params MoqSetuid_genType_params) string {
	return fmt.Sprintf("Setuid_genType(%#v)", params.Uid)
}

func (m *MoqSetuid_genType) ParamsKey(params MoqSetuid_genType_params, anyParams uint64) MoqSetuid_genType_paramsKey {
	m.Scene.T.Helper()
	var uidUsed int
	var uidUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Uid == moq.ParamIndexByValue {
			uidUsed = params.Uid
		} else {
			uidUsedHash = hash.DeepHash(params.Uid)
		}
	}
	return MoqSetuid_genType_paramsKey{
		Params: struct{ Uid int }{
			Uid: uidUsed,
		},
		Hashes: struct{ Uid hash.Hash }{
			Uid: uidUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetuid_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetuid_genType) AssertExpectationsMet() {
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
