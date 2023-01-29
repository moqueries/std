// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Setresgid_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Setresgid_genType func(rgid, egid, sgid int) (err error)

// MoqSetresgid_genType holds the state of a moq of the Setresgid_genType type
type MoqSetresgid_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetresgid_genType_mock

	ResultsByParams []MoqSetresgid_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rgid moq.ParamIndexing
			Egid moq.ParamIndexing
			Sgid moq.ParamIndexing
		}
	}
}

// MoqSetresgid_genType_mock isolates the mock interface of the
// Setresgid_genType type
type MoqSetresgid_genType_mock struct {
	Moq *MoqSetresgid_genType
}

// MoqSetresgid_genType_params holds the params of the Setresgid_genType type
type MoqSetresgid_genType_params struct{ Rgid, Egid, Sgid int }

// MoqSetresgid_genType_paramsKey holds the map key params of the
// Setresgid_genType type
type MoqSetresgid_genType_paramsKey struct {
	Params struct{ Rgid, Egid, Sgid int }
	Hashes struct{ Rgid, Egid, Sgid hash.Hash }
}

// MoqSetresgid_genType_resultsByParams contains the results for a given set of
// parameters for the Setresgid_genType type
type MoqSetresgid_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetresgid_genType_paramsKey]*MoqSetresgid_genType_results
}

// MoqSetresgid_genType_doFn defines the type of function needed when calling
// AndDo for the Setresgid_genType type
type MoqSetresgid_genType_doFn func(rgid, egid, sgid int)

// MoqSetresgid_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Setresgid_genType type
type MoqSetresgid_genType_doReturnFn func(rgid, egid, sgid int) (err error)

// MoqSetresgid_genType_results holds the results of the Setresgid_genType type
type MoqSetresgid_genType_results struct {
	Params  MoqSetresgid_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetresgid_genType_doFn
		DoReturnFn MoqSetresgid_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetresgid_genType_fnRecorder routes recorded function calls to the
// MoqSetresgid_genType moq
type MoqSetresgid_genType_fnRecorder struct {
	Params    MoqSetresgid_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetresgid_genType_results
	Moq       *MoqSetresgid_genType
}

// MoqSetresgid_genType_anyParams isolates the any params functions of the
// Setresgid_genType type
type MoqSetresgid_genType_anyParams struct {
	Recorder *MoqSetresgid_genType_fnRecorder
}

// NewMoqSetresgid_genType creates a new moq of the Setresgid_genType type
func NewMoqSetresgid_genType(scene *moq.Scene, config *moq.Config) *MoqSetresgid_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetresgid_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetresgid_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rgid moq.ParamIndexing
				Egid moq.ParamIndexing
				Sgid moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rgid moq.ParamIndexing
			Egid moq.ParamIndexing
			Sgid moq.ParamIndexing
		}{
			Rgid: moq.ParamIndexByValue,
			Egid: moq.ParamIndexByValue,
			Sgid: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Setresgid_genType type
func (m *MoqSetresgid_genType) Mock() Setresgid_genType {
	return func(rgid, egid, sgid int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqSetresgid_genType_mock{Moq: m}
		return moq.Fn(rgid, egid, sgid)
	}
}

func (m *MoqSetresgid_genType_mock) Fn(rgid, egid, sgid int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqSetresgid_genType_params{
		Rgid: rgid,
		Egid: egid,
		Sgid: sgid,
	}
	var results *MoqSetresgid_genType_results
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
		result.DoFn(rgid, egid, sgid)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(rgid, egid, sgid)
	}
	return
}

func (m *MoqSetresgid_genType) OnCall(rgid, egid, sgid int) *MoqSetresgid_genType_fnRecorder {
	return &MoqSetresgid_genType_fnRecorder{
		Params: MoqSetresgid_genType_params{
			Rgid: rgid,
			Egid: egid,
			Sgid: sgid,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetresgid_genType_fnRecorder) Any() *MoqSetresgid_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetresgid_genType_anyParams{Recorder: r}
}

func (a *MoqSetresgid_genType_anyParams) Rgid() *MoqSetresgid_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSetresgid_genType_anyParams) Egid() *MoqSetresgid_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqSetresgid_genType_anyParams) Sgid() *MoqSetresgid_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqSetresgid_genType_fnRecorder) Seq() *MoqSetresgid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetresgid_genType_fnRecorder) NoSeq() *MoqSetresgid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetresgid_genType_fnRecorder) ReturnResults(err error) *MoqSetresgid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetresgid_genType_doFn
		DoReturnFn MoqSetresgid_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetresgid_genType_fnRecorder) AndDo(fn MoqSetresgid_genType_doFn) *MoqSetresgid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetresgid_genType_fnRecorder) DoReturnResults(fn MoqSetresgid_genType_doReturnFn) *MoqSetresgid_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqSetresgid_genType_doFn
		DoReturnFn MoqSetresgid_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetresgid_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetresgid_genType_resultsByParams
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
		results = &MoqSetresgid_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetresgid_genType_paramsKey]*MoqSetresgid_genType_results{},
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
		r.Results = &MoqSetresgid_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetresgid_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetresgid_genType_fnRecorder {
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
				DoFn       MoqSetresgid_genType_doFn
				DoReturnFn MoqSetresgid_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetresgid_genType) PrettyParams(params MoqSetresgid_genType_params) string {
	return fmt.Sprintf("Setresgid_genType(%#v, %#v, %#v)", params.Rgid, params.Egid, params.Sgid)
}

func (m *MoqSetresgid_genType) ParamsKey(params MoqSetresgid_genType_params, anyParams uint64) MoqSetresgid_genType_paramsKey {
	m.Scene.T.Helper()
	var rgidUsed int
	var rgidUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Rgid == moq.ParamIndexByValue {
			rgidUsed = params.Rgid
		} else {
			rgidUsedHash = hash.DeepHash(params.Rgid)
		}
	}
	var egidUsed int
	var egidUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Egid == moq.ParamIndexByValue {
			egidUsed = params.Egid
		} else {
			egidUsedHash = hash.DeepHash(params.Egid)
		}
	}
	var sgidUsed int
	var sgidUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Sgid == moq.ParamIndexByValue {
			sgidUsed = params.Sgid
		} else {
			sgidUsedHash = hash.DeepHash(params.Sgid)
		}
	}
	return MoqSetresgid_genType_paramsKey{
		Params: struct{ Rgid, Egid, Sgid int }{
			Rgid: rgidUsed,
			Egid: egidUsed,
			Sgid: sgidUsed,
		},
		Hashes: struct{ Rgid, Egid, Sgid hash.Hash }{
			Rgid: rgidUsedHash,
			Egid: egidUsedHash,
			Sgid: sgidUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetresgid_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetresgid_genType) AssertExpectationsMet() {
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
