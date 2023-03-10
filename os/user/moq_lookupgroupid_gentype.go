// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package user

import (
	"fmt"
	"math/bits"
	"os/user"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LookupGroupId_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LookupGroupId_genType func(gid string) (*user.Group, error)

// MoqLookupGroupId_genType holds the state of a moq of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLookupGroupId_genType_mock

	ResultsByParams []MoqLookupGroupId_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Gid moq.ParamIndexing
		}
	}
}

// MoqLookupGroupId_genType_mock isolates the mock interface of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType_mock struct {
	Moq *MoqLookupGroupId_genType
}

// MoqLookupGroupId_genType_params holds the params of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType_params struct{ Gid string }

// MoqLookupGroupId_genType_paramsKey holds the map key params of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType_paramsKey struct {
	Params struct{ Gid string }
	Hashes struct{ Gid hash.Hash }
}

// MoqLookupGroupId_genType_resultsByParams contains the results for a given
// set of parameters for the LookupGroupId_genType type
type MoqLookupGroupId_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLookupGroupId_genType_paramsKey]*MoqLookupGroupId_genType_results
}

// MoqLookupGroupId_genType_doFn defines the type of function needed when
// calling AndDo for the LookupGroupId_genType type
type MoqLookupGroupId_genType_doFn func(gid string)

// MoqLookupGroupId_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LookupGroupId_genType type
type MoqLookupGroupId_genType_doReturnFn func(gid string) (*user.Group, error)

// MoqLookupGroupId_genType_results holds the results of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType_results struct {
	Params  MoqLookupGroupId_genType_params
	Results []struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroupId_genType_doFn
		DoReturnFn MoqLookupGroupId_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLookupGroupId_genType_fnRecorder routes recorded function calls to the
// MoqLookupGroupId_genType moq
type MoqLookupGroupId_genType_fnRecorder struct {
	Params    MoqLookupGroupId_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLookupGroupId_genType_results
	Moq       *MoqLookupGroupId_genType
}

// MoqLookupGroupId_genType_anyParams isolates the any params functions of the
// LookupGroupId_genType type
type MoqLookupGroupId_genType_anyParams struct {
	Recorder *MoqLookupGroupId_genType_fnRecorder
}

// NewMoqLookupGroupId_genType creates a new moq of the LookupGroupId_genType
// type
func NewMoqLookupGroupId_genType(scene *moq.Scene, config *moq.Config) *MoqLookupGroupId_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLookupGroupId_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLookupGroupId_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Gid moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Gid moq.ParamIndexing
		}{
			Gid: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the LookupGroupId_genType type
func (m *MoqLookupGroupId_genType) Mock() LookupGroupId_genType {
	return func(gid string) (*user.Group, error) {
		m.Scene.T.Helper()
		moq := &MoqLookupGroupId_genType_mock{Moq: m}
		return moq.Fn(gid)
	}
}

func (m *MoqLookupGroupId_genType_mock) Fn(gid string) (result1 *user.Group, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqLookupGroupId_genType_params{
		Gid: gid,
	}
	var results *MoqLookupGroupId_genType_results
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
		result.DoFn(gid)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(gid)
	}
	return
}

func (m *MoqLookupGroupId_genType) OnCall(gid string) *MoqLookupGroupId_genType_fnRecorder {
	return &MoqLookupGroupId_genType_fnRecorder{
		Params: MoqLookupGroupId_genType_params{
			Gid: gid,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLookupGroupId_genType_fnRecorder) Any() *MoqLookupGroupId_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLookupGroupId_genType_anyParams{Recorder: r}
}

func (a *MoqLookupGroupId_genType_anyParams) Gid() *MoqLookupGroupId_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLookupGroupId_genType_fnRecorder) Seq() *MoqLookupGroupId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLookupGroupId_genType_fnRecorder) NoSeq() *MoqLookupGroupId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLookupGroupId_genType_fnRecorder) ReturnResults(result1 *user.Group, result2 error) *MoqLookupGroupId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroupId_genType_doFn
		DoReturnFn MoqLookupGroupId_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *user.Group
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLookupGroupId_genType_fnRecorder) AndDo(fn MoqLookupGroupId_genType_doFn) *MoqLookupGroupId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLookupGroupId_genType_fnRecorder) DoReturnResults(fn MoqLookupGroupId_genType_doReturnFn) *MoqLookupGroupId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *user.Group
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLookupGroupId_genType_doFn
		DoReturnFn MoqLookupGroupId_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLookupGroupId_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLookupGroupId_genType_resultsByParams
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
		results = &MoqLookupGroupId_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLookupGroupId_genType_paramsKey]*MoqLookupGroupId_genType_results{},
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
		r.Results = &MoqLookupGroupId_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLookupGroupId_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLookupGroupId_genType_fnRecorder {
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
					Result1 *user.Group
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqLookupGroupId_genType_doFn
				DoReturnFn MoqLookupGroupId_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLookupGroupId_genType) PrettyParams(params MoqLookupGroupId_genType_params) string {
	return fmt.Sprintf("LookupGroupId_genType(%#v)", params.Gid)
}

func (m *MoqLookupGroupId_genType) ParamsKey(params MoqLookupGroupId_genType_params, anyParams uint64) MoqLookupGroupId_genType_paramsKey {
	m.Scene.T.Helper()
	var gidUsed string
	var gidUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Gid == moq.ParamIndexByValue {
			gidUsed = params.Gid
		} else {
			gidUsedHash = hash.DeepHash(params.Gid)
		}
	}
	return MoqLookupGroupId_genType_paramsKey{
		Params: struct{ Gid string }{
			Gid: gidUsed,
		},
		Hashes: struct{ Gid hash.Hash }{
			Gid: gidUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLookupGroupId_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLookupGroupId_genType) AssertExpectationsMet() {
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
