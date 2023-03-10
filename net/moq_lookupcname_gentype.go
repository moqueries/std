// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LookupCNAME_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LookupCNAME_genType func(host string) (cname string, err error)

// MoqLookupCNAME_genType holds the state of a moq of the LookupCNAME_genType
// type
type MoqLookupCNAME_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLookupCNAME_genType_mock

	ResultsByParams []MoqLookupCNAME_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Host moq.ParamIndexing
		}
	}
}

// MoqLookupCNAME_genType_mock isolates the mock interface of the
// LookupCNAME_genType type
type MoqLookupCNAME_genType_mock struct {
	Moq *MoqLookupCNAME_genType
}

// MoqLookupCNAME_genType_params holds the params of the LookupCNAME_genType
// type
type MoqLookupCNAME_genType_params struct{ Host string }

// MoqLookupCNAME_genType_paramsKey holds the map key params of the
// LookupCNAME_genType type
type MoqLookupCNAME_genType_paramsKey struct {
	Params struct{ Host string }
	Hashes struct{ Host hash.Hash }
}

// MoqLookupCNAME_genType_resultsByParams contains the results for a given set
// of parameters for the LookupCNAME_genType type
type MoqLookupCNAME_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLookupCNAME_genType_paramsKey]*MoqLookupCNAME_genType_results
}

// MoqLookupCNAME_genType_doFn defines the type of function needed when calling
// AndDo for the LookupCNAME_genType type
type MoqLookupCNAME_genType_doFn func(host string)

// MoqLookupCNAME_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LookupCNAME_genType type
type MoqLookupCNAME_genType_doReturnFn func(host string) (cname string, err error)

// MoqLookupCNAME_genType_results holds the results of the LookupCNAME_genType
// type
type MoqLookupCNAME_genType_results struct {
	Params  MoqLookupCNAME_genType_params
	Results []struct {
		Values *struct {
			Cname string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupCNAME_genType_doFn
		DoReturnFn MoqLookupCNAME_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLookupCNAME_genType_fnRecorder routes recorded function calls to the
// MoqLookupCNAME_genType moq
type MoqLookupCNAME_genType_fnRecorder struct {
	Params    MoqLookupCNAME_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLookupCNAME_genType_results
	Moq       *MoqLookupCNAME_genType
}

// MoqLookupCNAME_genType_anyParams isolates the any params functions of the
// LookupCNAME_genType type
type MoqLookupCNAME_genType_anyParams struct {
	Recorder *MoqLookupCNAME_genType_fnRecorder
}

// NewMoqLookupCNAME_genType creates a new moq of the LookupCNAME_genType type
func NewMoqLookupCNAME_genType(scene *moq.Scene, config *moq.Config) *MoqLookupCNAME_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLookupCNAME_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLookupCNAME_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Host moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Host moq.ParamIndexing
		}{
			Host: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the LookupCNAME_genType type
func (m *MoqLookupCNAME_genType) Mock() LookupCNAME_genType {
	return func(host string) (_ string, _ error) {
		m.Scene.T.Helper()
		moq := &MoqLookupCNAME_genType_mock{Moq: m}
		return moq.Fn(host)
	}
}

func (m *MoqLookupCNAME_genType_mock) Fn(host string) (cname string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqLookupCNAME_genType_params{
		Host: host,
	}
	var results *MoqLookupCNAME_genType_results
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
		result.DoFn(host)
	}

	if result.Values != nil {
		cname = result.Values.Cname
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		cname, err = result.DoReturnFn(host)
	}
	return
}

func (m *MoqLookupCNAME_genType) OnCall(host string) *MoqLookupCNAME_genType_fnRecorder {
	return &MoqLookupCNAME_genType_fnRecorder{
		Params: MoqLookupCNAME_genType_params{
			Host: host,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLookupCNAME_genType_fnRecorder) Any() *MoqLookupCNAME_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLookupCNAME_genType_anyParams{Recorder: r}
}

func (a *MoqLookupCNAME_genType_anyParams) Host() *MoqLookupCNAME_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLookupCNAME_genType_fnRecorder) Seq() *MoqLookupCNAME_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLookupCNAME_genType_fnRecorder) NoSeq() *MoqLookupCNAME_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLookupCNAME_genType_fnRecorder) ReturnResults(cname string, err error) *MoqLookupCNAME_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Cname string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupCNAME_genType_doFn
		DoReturnFn MoqLookupCNAME_genType_doReturnFn
	}{
		Values: &struct {
			Cname string
			Err   error
		}{
			Cname: cname,
			Err:   err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLookupCNAME_genType_fnRecorder) AndDo(fn MoqLookupCNAME_genType_doFn) *MoqLookupCNAME_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLookupCNAME_genType_fnRecorder) DoReturnResults(fn MoqLookupCNAME_genType_doReturnFn) *MoqLookupCNAME_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Cname string
			Err   error
		}
		Sequence   uint32
		DoFn       MoqLookupCNAME_genType_doFn
		DoReturnFn MoqLookupCNAME_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLookupCNAME_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLookupCNAME_genType_resultsByParams
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
		results = &MoqLookupCNAME_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLookupCNAME_genType_paramsKey]*MoqLookupCNAME_genType_results{},
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
		r.Results = &MoqLookupCNAME_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLookupCNAME_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLookupCNAME_genType_fnRecorder {
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
					Cname string
					Err   error
				}
				Sequence   uint32
				DoFn       MoqLookupCNAME_genType_doFn
				DoReturnFn MoqLookupCNAME_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLookupCNAME_genType) PrettyParams(params MoqLookupCNAME_genType_params) string {
	return fmt.Sprintf("LookupCNAME_genType(%#v)", params.Host)
}

func (m *MoqLookupCNAME_genType) ParamsKey(params MoqLookupCNAME_genType_params, anyParams uint64) MoqLookupCNAME_genType_paramsKey {
	m.Scene.T.Helper()
	var hostUsed string
	var hostUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Host == moq.ParamIndexByValue {
			hostUsed = params.Host
		} else {
			hostUsedHash = hash.DeepHash(params.Host)
		}
	}
	return MoqLookupCNAME_genType_paramsKey{
		Params: struct{ Host string }{
			Host: hostUsed,
		},
		Hashes: struct{ Host hash.Hash }{
			Host: hostUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLookupCNAME_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLookupCNAME_genType) AssertExpectationsMet() {
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
