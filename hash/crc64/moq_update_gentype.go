// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package crc64

import (
	"fmt"
	"hash/crc64"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Update_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Update_genType func(crc uint64, tab *crc64.Table, p []byte) uint64

// MoqUpdate_genType holds the state of a moq of the Update_genType type
type MoqUpdate_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUpdate_genType_mock

	ResultsByParams []MoqUpdate_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Crc moq.ParamIndexing
			Tab moq.ParamIndexing
			P   moq.ParamIndexing
		}
	}
}

// MoqUpdate_genType_mock isolates the mock interface of the Update_genType
// type
type MoqUpdate_genType_mock struct {
	Moq *MoqUpdate_genType
}

// MoqUpdate_genType_params holds the params of the Update_genType type
type MoqUpdate_genType_params struct {
	Crc uint64
	Tab *crc64.Table
	P   []byte
}

// MoqUpdate_genType_paramsKey holds the map key params of the Update_genType
// type
type MoqUpdate_genType_paramsKey struct {
	Params struct {
		Crc uint64
		Tab *crc64.Table
	}
	Hashes struct {
		Crc hash.Hash
		Tab hash.Hash
		P   hash.Hash
	}
}

// MoqUpdate_genType_resultsByParams contains the results for a given set of
// parameters for the Update_genType type
type MoqUpdate_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUpdate_genType_paramsKey]*MoqUpdate_genType_results
}

// MoqUpdate_genType_doFn defines the type of function needed when calling
// AndDo for the Update_genType type
type MoqUpdate_genType_doFn func(crc uint64, tab *crc64.Table, p []byte)

// MoqUpdate_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Update_genType type
type MoqUpdate_genType_doReturnFn func(crc uint64, tab *crc64.Table, p []byte) uint64

// MoqUpdate_genType_results holds the results of the Update_genType type
type MoqUpdate_genType_results struct {
	Params  MoqUpdate_genType_params
	Results []struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqUpdate_genType_doFn
		DoReturnFn MoqUpdate_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUpdate_genType_fnRecorder routes recorded function calls to the
// MoqUpdate_genType moq
type MoqUpdate_genType_fnRecorder struct {
	Params    MoqUpdate_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUpdate_genType_results
	Moq       *MoqUpdate_genType
}

// MoqUpdate_genType_anyParams isolates the any params functions of the
// Update_genType type
type MoqUpdate_genType_anyParams struct {
	Recorder *MoqUpdate_genType_fnRecorder
}

// NewMoqUpdate_genType creates a new moq of the Update_genType type
func NewMoqUpdate_genType(scene *moq.Scene, config *moq.Config) *MoqUpdate_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUpdate_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUpdate_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Crc moq.ParamIndexing
				Tab moq.ParamIndexing
				P   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Crc moq.ParamIndexing
			Tab moq.ParamIndexing
			P   moq.ParamIndexing
		}{
			Crc: moq.ParamIndexByValue,
			Tab: moq.ParamIndexByHash,
			P:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Update_genType type
func (m *MoqUpdate_genType) Mock() Update_genType {
	return func(crc uint64, tab *crc64.Table, p []byte) uint64 {
		m.Scene.T.Helper()
		moq := &MoqUpdate_genType_mock{Moq: m}
		return moq.Fn(crc, tab, p)
	}
}

func (m *MoqUpdate_genType_mock) Fn(crc uint64, tab *crc64.Table, p []byte) (result1 uint64) {
	m.Moq.Scene.T.Helper()
	params := MoqUpdate_genType_params{
		Crc: crc,
		Tab: tab,
		P:   p,
	}
	var results *MoqUpdate_genType_results
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
		result.DoFn(crc, tab, p)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(crc, tab, p)
	}
	return
}

func (m *MoqUpdate_genType) OnCall(crc uint64, tab *crc64.Table, p []byte) *MoqUpdate_genType_fnRecorder {
	return &MoqUpdate_genType_fnRecorder{
		Params: MoqUpdate_genType_params{
			Crc: crc,
			Tab: tab,
			P:   p,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUpdate_genType_fnRecorder) Any() *MoqUpdate_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUpdate_genType_anyParams{Recorder: r}
}

func (a *MoqUpdate_genType_anyParams) Crc() *MoqUpdate_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqUpdate_genType_anyParams) Tab() *MoqUpdate_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqUpdate_genType_anyParams) P() *MoqUpdate_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqUpdate_genType_fnRecorder) Seq() *MoqUpdate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUpdate_genType_fnRecorder) NoSeq() *MoqUpdate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUpdate_genType_fnRecorder) ReturnResults(result1 uint64) *MoqUpdate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqUpdate_genType_doFn
		DoReturnFn MoqUpdate_genType_doReturnFn
	}{
		Values: &struct {
			Result1 uint64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUpdate_genType_fnRecorder) AndDo(fn MoqUpdate_genType_doFn) *MoqUpdate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUpdate_genType_fnRecorder) DoReturnResults(fn MoqUpdate_genType_doReturnFn) *MoqUpdate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqUpdate_genType_doFn
		DoReturnFn MoqUpdate_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUpdate_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUpdate_genType_resultsByParams
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
		results = &MoqUpdate_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUpdate_genType_paramsKey]*MoqUpdate_genType_results{},
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
		r.Results = &MoqUpdate_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUpdate_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUpdate_genType_fnRecorder {
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
					Result1 uint64
				}
				Sequence   uint32
				DoFn       MoqUpdate_genType_doFn
				DoReturnFn MoqUpdate_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUpdate_genType) PrettyParams(params MoqUpdate_genType_params) string {
	return fmt.Sprintf("Update_genType(%#v, %#v, %#v)", params.Crc, params.Tab, params.P)
}

func (m *MoqUpdate_genType) ParamsKey(params MoqUpdate_genType_params, anyParams uint64) MoqUpdate_genType_paramsKey {
	m.Scene.T.Helper()
	var crcUsed uint64
	var crcUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Crc == moq.ParamIndexByValue {
			crcUsed = params.Crc
		} else {
			crcUsedHash = hash.DeepHash(params.Crc)
		}
	}
	var tabUsed *crc64.Table
	var tabUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Tab == moq.ParamIndexByValue {
			tabUsed = params.Tab
		} else {
			tabUsedHash = hash.DeepHash(params.Tab)
		}
	}
	var pUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqUpdate_genType_paramsKey{
		Params: struct {
			Crc uint64
			Tab *crc64.Table
		}{
			Crc: crcUsed,
			Tab: tabUsed,
		},
		Hashes: struct {
			Crc hash.Hash
			Tab hash.Hash
			P   hash.Hash
		}{
			Crc: crcUsedHash,
			Tab: tabUsedHash,
			P:   pUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqUpdate_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUpdate_genType) AssertExpectationsMet() {
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
