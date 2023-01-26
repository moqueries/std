// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that net.Interface_starGenType is
// mocked completely
var _ Interface_starGenType = (*MoqInterface_starGenType_mock)(nil)

// Interface_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Interface_starGenType interface {
	Addrs() ([]net.Addr, error)
	MulticastAddrs() ([]net.Addr, error)
}

// MoqInterface_starGenType holds the state of a moq of the
// Interface_starGenType type
type MoqInterface_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqInterface_starGenType_mock

	ResultsByParams_Addrs          []MoqInterface_starGenType_Addrs_resultsByParams
	ResultsByParams_MulticastAddrs []MoqInterface_starGenType_MulticastAddrs_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addrs          struct{}
			MulticastAddrs struct{}
		}
	}
}

// MoqInterface_starGenType_mock isolates the mock interface of the
// Interface_starGenType type
type MoqInterface_starGenType_mock struct {
	Moq *MoqInterface_starGenType
}

// MoqInterface_starGenType_recorder isolates the recorder interface of the
// Interface_starGenType type
type MoqInterface_starGenType_recorder struct {
	Moq *MoqInterface_starGenType
}

// MoqInterface_starGenType_Addrs_params holds the params of the
// Interface_starGenType type
type MoqInterface_starGenType_Addrs_params struct{}

// MoqInterface_starGenType_Addrs_paramsKey holds the map key params of the
// Interface_starGenType type
type MoqInterface_starGenType_Addrs_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqInterface_starGenType_Addrs_resultsByParams contains the results for a
// given set of parameters for the Interface_starGenType type
type MoqInterface_starGenType_Addrs_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqInterface_starGenType_Addrs_paramsKey]*MoqInterface_starGenType_Addrs_results
}

// MoqInterface_starGenType_Addrs_doFn defines the type of function needed when
// calling AndDo for the Interface_starGenType type
type MoqInterface_starGenType_Addrs_doFn func()

// MoqInterface_starGenType_Addrs_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Interface_starGenType type
type MoqInterface_starGenType_Addrs_doReturnFn func() ([]net.Addr, error)

// MoqInterface_starGenType_Addrs_results holds the results of the
// Interface_starGenType type
type MoqInterface_starGenType_Addrs_results struct {
	Params  MoqInterface_starGenType_Addrs_params
	Results []struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_Addrs_doFn
		DoReturnFn MoqInterface_starGenType_Addrs_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqInterface_starGenType_Addrs_fnRecorder routes recorded function calls to
// the MoqInterface_starGenType moq
type MoqInterface_starGenType_Addrs_fnRecorder struct {
	Params    MoqInterface_starGenType_Addrs_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqInterface_starGenType_Addrs_results
	Moq       *MoqInterface_starGenType
}

// MoqInterface_starGenType_Addrs_anyParams isolates the any params functions
// of the Interface_starGenType type
type MoqInterface_starGenType_Addrs_anyParams struct {
	Recorder *MoqInterface_starGenType_Addrs_fnRecorder
}

// MoqInterface_starGenType_MulticastAddrs_params holds the params of the
// Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_params struct{}

// MoqInterface_starGenType_MulticastAddrs_paramsKey holds the map key params
// of the Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqInterface_starGenType_MulticastAddrs_resultsByParams contains the results
// for a given set of parameters for the Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqInterface_starGenType_MulticastAddrs_paramsKey]*MoqInterface_starGenType_MulticastAddrs_results
}

// MoqInterface_starGenType_MulticastAddrs_doFn defines the type of function
// needed when calling AndDo for the Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_doFn func()

// MoqInterface_starGenType_MulticastAddrs_doReturnFn defines the type of
// function needed when calling DoReturnResults for the Interface_starGenType
// type
type MoqInterface_starGenType_MulticastAddrs_doReturnFn func() ([]net.Addr, error)

// MoqInterface_starGenType_MulticastAddrs_results holds the results of the
// Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_results struct {
	Params  MoqInterface_starGenType_MulticastAddrs_params
	Results []struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_MulticastAddrs_doFn
		DoReturnFn MoqInterface_starGenType_MulticastAddrs_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqInterface_starGenType_MulticastAddrs_fnRecorder routes recorded function
// calls to the MoqInterface_starGenType moq
type MoqInterface_starGenType_MulticastAddrs_fnRecorder struct {
	Params    MoqInterface_starGenType_MulticastAddrs_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqInterface_starGenType_MulticastAddrs_results
	Moq       *MoqInterface_starGenType
}

// MoqInterface_starGenType_MulticastAddrs_anyParams isolates the any params
// functions of the Interface_starGenType type
type MoqInterface_starGenType_MulticastAddrs_anyParams struct {
	Recorder *MoqInterface_starGenType_MulticastAddrs_fnRecorder
}

// NewMoqInterface_starGenType creates a new moq of the Interface_starGenType
// type
func NewMoqInterface_starGenType(scene *moq.Scene, config *moq.Config) *MoqInterface_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqInterface_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqInterface_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addrs          struct{}
				MulticastAddrs struct{}
			}
		}{ParameterIndexing: struct {
			Addrs          struct{}
			MulticastAddrs struct{}
		}{
			Addrs:          struct{}{},
			MulticastAddrs: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Interface_starGenType type
func (m *MoqInterface_starGenType) Mock() *MoqInterface_starGenType_mock { return m.Moq }

func (m *MoqInterface_starGenType_mock) Addrs() (result1 []net.Addr, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqInterface_starGenType_Addrs_params{}
	var results *MoqInterface_starGenType_Addrs_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Addrs {
		paramsKey := m.Moq.ParamsKey_Addrs(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Addrs(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Addrs(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Addrs(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

func (m *MoqInterface_starGenType_mock) MulticastAddrs() (result1 []net.Addr, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqInterface_starGenType_MulticastAddrs_params{}
	var results *MoqInterface_starGenType_MulticastAddrs_results
	for _, resultsByParams := range m.Moq.ResultsByParams_MulticastAddrs {
		paramsKey := m.Moq.ParamsKey_MulticastAddrs(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_MulticastAddrs(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_MulticastAddrs(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_MulticastAddrs(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Interface_starGenType type
func (m *MoqInterface_starGenType) OnCall() *MoqInterface_starGenType_recorder {
	return &MoqInterface_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqInterface_starGenType_recorder) Addrs() *MoqInterface_starGenType_Addrs_fnRecorder {
	return &MoqInterface_starGenType_Addrs_fnRecorder{
		Params:   MoqInterface_starGenType_Addrs_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) Any() *MoqInterface_starGenType_Addrs_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Addrs(r.Params))
		return nil
	}
	return &MoqInterface_starGenType_Addrs_anyParams{Recorder: r}
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) Seq() *MoqInterface_starGenType_Addrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Addrs(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) NoSeq() *MoqInterface_starGenType_Addrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Addrs(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) ReturnResults(result1 []net.Addr, result2 error) *MoqInterface_starGenType_Addrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_Addrs_doFn
		DoReturnFn MoqInterface_starGenType_Addrs_doReturnFn
	}{
		Values: &struct {
			Result1 []net.Addr
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) AndDo(fn MoqInterface_starGenType_Addrs_doFn) *MoqInterface_starGenType_Addrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) DoReturnResults(fn MoqInterface_starGenType_Addrs_doReturnFn) *MoqInterface_starGenType_Addrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_Addrs_doFn
		DoReturnFn MoqInterface_starGenType_Addrs_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqInterface_starGenType_Addrs_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Addrs {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqInterface_starGenType_Addrs_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqInterface_starGenType_Addrs_paramsKey]*MoqInterface_starGenType_Addrs_results{},
		}
		r.Moq.ResultsByParams_Addrs = append(r.Moq.ResultsByParams_Addrs, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Addrs) {
			copy(r.Moq.ResultsByParams_Addrs[insertAt+1:], r.Moq.ResultsByParams_Addrs[insertAt:0])
			r.Moq.ResultsByParams_Addrs[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Addrs(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqInterface_starGenType_Addrs_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqInterface_starGenType_Addrs_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqInterface_starGenType_Addrs_fnRecorder {
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
					Result1 []net.Addr
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqInterface_starGenType_Addrs_doFn
				DoReturnFn MoqInterface_starGenType_Addrs_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqInterface_starGenType) PrettyParams_Addrs(params MoqInterface_starGenType_Addrs_params) string {
	return fmt.Sprintf("Addrs()")
}

func (m *MoqInterface_starGenType) ParamsKey_Addrs(params MoqInterface_starGenType_Addrs_params, anyParams uint64) MoqInterface_starGenType_Addrs_paramsKey {
	m.Scene.T.Helper()
	return MoqInterface_starGenType_Addrs_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqInterface_starGenType_recorder) MulticastAddrs() *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	return &MoqInterface_starGenType_MulticastAddrs_fnRecorder{
		Params:   MoqInterface_starGenType_MulticastAddrs_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) Any() *MoqInterface_starGenType_MulticastAddrs_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MulticastAddrs(r.Params))
		return nil
	}
	return &MoqInterface_starGenType_MulticastAddrs_anyParams{Recorder: r}
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) Seq() *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MulticastAddrs(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) NoSeq() *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MulticastAddrs(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) ReturnResults(result1 []net.Addr, result2 error) *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_MulticastAddrs_doFn
		DoReturnFn MoqInterface_starGenType_MulticastAddrs_doReturnFn
	}{
		Values: &struct {
			Result1 []net.Addr
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) AndDo(fn MoqInterface_starGenType_MulticastAddrs_doFn) *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) DoReturnResults(fn MoqInterface_starGenType_MulticastAddrs_doReturnFn) *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []net.Addr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqInterface_starGenType_MulticastAddrs_doFn
		DoReturnFn MoqInterface_starGenType_MulticastAddrs_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqInterface_starGenType_MulticastAddrs_resultsByParams
	for n, res := range r.Moq.ResultsByParams_MulticastAddrs {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqInterface_starGenType_MulticastAddrs_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqInterface_starGenType_MulticastAddrs_paramsKey]*MoqInterface_starGenType_MulticastAddrs_results{},
		}
		r.Moq.ResultsByParams_MulticastAddrs = append(r.Moq.ResultsByParams_MulticastAddrs, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_MulticastAddrs) {
			copy(r.Moq.ResultsByParams_MulticastAddrs[insertAt+1:], r.Moq.ResultsByParams_MulticastAddrs[insertAt:0])
			r.Moq.ResultsByParams_MulticastAddrs[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_MulticastAddrs(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqInterface_starGenType_MulticastAddrs_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqInterface_starGenType_MulticastAddrs_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqInterface_starGenType_MulticastAddrs_fnRecorder {
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
					Result1 []net.Addr
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqInterface_starGenType_MulticastAddrs_doFn
				DoReturnFn MoqInterface_starGenType_MulticastAddrs_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqInterface_starGenType) PrettyParams_MulticastAddrs(params MoqInterface_starGenType_MulticastAddrs_params) string {
	return fmt.Sprintf("MulticastAddrs()")
}

func (m *MoqInterface_starGenType) ParamsKey_MulticastAddrs(params MoqInterface_starGenType_MulticastAddrs_params, anyParams uint64) MoqInterface_starGenType_MulticastAddrs_paramsKey {
	m.Scene.T.Helper()
	return MoqInterface_starGenType_MulticastAddrs_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqInterface_starGenType) Reset() {
	m.ResultsByParams_Addrs = nil
	m.ResultsByParams_MulticastAddrs = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqInterface_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Addrs {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Addrs(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_MulticastAddrs {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_MulticastAddrs(results.Params))
			}
		}
	}
}