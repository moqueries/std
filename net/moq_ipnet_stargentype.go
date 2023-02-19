// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that net.IPNet_starGenType is mocked
// completely
var _ IPNet_starGenType = (*MoqIPNet_starGenType_mock)(nil)

// IPNet_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type IPNet_starGenType interface {
	Contains(ip net.IP) bool
	Network() string
	String() string
}

// MoqIPNet_starGenType holds the state of a moq of the IPNet_starGenType type
type MoqIPNet_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIPNet_starGenType_mock

	ResultsByParams_Contains []MoqIPNet_starGenType_Contains_resultsByParams
	ResultsByParams_Network  []MoqIPNet_starGenType_Network_resultsByParams
	ResultsByParams_String   []MoqIPNet_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Contains struct {
				Ip moq.ParamIndexing
			}
			Network struct{}
			String  struct{}
		}
	}
}

// MoqIPNet_starGenType_mock isolates the mock interface of the
// IPNet_starGenType type
type MoqIPNet_starGenType_mock struct {
	Moq *MoqIPNet_starGenType
}

// MoqIPNet_starGenType_recorder isolates the recorder interface of the
// IPNet_starGenType type
type MoqIPNet_starGenType_recorder struct {
	Moq *MoqIPNet_starGenType
}

// MoqIPNet_starGenType_Contains_params holds the params of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Contains_params struct{ Ip net.IP }

// MoqIPNet_starGenType_Contains_paramsKey holds the map key params of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Contains_paramsKey struct {
	Params struct{}
	Hashes struct{ Ip hash.Hash }
}

// MoqIPNet_starGenType_Contains_resultsByParams contains the results for a
// given set of parameters for the IPNet_starGenType type
type MoqIPNet_starGenType_Contains_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPNet_starGenType_Contains_paramsKey]*MoqIPNet_starGenType_Contains_results
}

// MoqIPNet_starGenType_Contains_doFn defines the type of function needed when
// calling AndDo for the IPNet_starGenType type
type MoqIPNet_starGenType_Contains_doFn func(ip net.IP)

// MoqIPNet_starGenType_Contains_doReturnFn defines the type of function needed
// when calling DoReturnResults for the IPNet_starGenType type
type MoqIPNet_starGenType_Contains_doReturnFn func(ip net.IP) bool

// MoqIPNet_starGenType_Contains_results holds the results of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Contains_results struct {
	Params  MoqIPNet_starGenType_Contains_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_Contains_doFn
		DoReturnFn MoqIPNet_starGenType_Contains_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPNet_starGenType_Contains_fnRecorder routes recorded function calls to
// the MoqIPNet_starGenType moq
type MoqIPNet_starGenType_Contains_fnRecorder struct {
	Params    MoqIPNet_starGenType_Contains_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPNet_starGenType_Contains_results
	Moq       *MoqIPNet_starGenType
}

// MoqIPNet_starGenType_Contains_anyParams isolates the any params functions of
// the IPNet_starGenType type
type MoqIPNet_starGenType_Contains_anyParams struct {
	Recorder *MoqIPNet_starGenType_Contains_fnRecorder
}

// MoqIPNet_starGenType_Network_params holds the params of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Network_params struct{}

// MoqIPNet_starGenType_Network_paramsKey holds the map key params of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Network_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPNet_starGenType_Network_resultsByParams contains the results for a
// given set of parameters for the IPNet_starGenType type
type MoqIPNet_starGenType_Network_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPNet_starGenType_Network_paramsKey]*MoqIPNet_starGenType_Network_results
}

// MoqIPNet_starGenType_Network_doFn defines the type of function needed when
// calling AndDo for the IPNet_starGenType type
type MoqIPNet_starGenType_Network_doFn func()

// MoqIPNet_starGenType_Network_doReturnFn defines the type of function needed
// when calling DoReturnResults for the IPNet_starGenType type
type MoqIPNet_starGenType_Network_doReturnFn func() string

// MoqIPNet_starGenType_Network_results holds the results of the
// IPNet_starGenType type
type MoqIPNet_starGenType_Network_results struct {
	Params  MoqIPNet_starGenType_Network_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_Network_doFn
		DoReturnFn MoqIPNet_starGenType_Network_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPNet_starGenType_Network_fnRecorder routes recorded function calls to
// the MoqIPNet_starGenType moq
type MoqIPNet_starGenType_Network_fnRecorder struct {
	Params    MoqIPNet_starGenType_Network_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPNet_starGenType_Network_results
	Moq       *MoqIPNet_starGenType
}

// MoqIPNet_starGenType_Network_anyParams isolates the any params functions of
// the IPNet_starGenType type
type MoqIPNet_starGenType_Network_anyParams struct {
	Recorder *MoqIPNet_starGenType_Network_fnRecorder
}

// MoqIPNet_starGenType_String_params holds the params of the IPNet_starGenType
// type
type MoqIPNet_starGenType_String_params struct{}

// MoqIPNet_starGenType_String_paramsKey holds the map key params of the
// IPNet_starGenType type
type MoqIPNet_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPNet_starGenType_String_resultsByParams contains the results for a given
// set of parameters for the IPNet_starGenType type
type MoqIPNet_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPNet_starGenType_String_paramsKey]*MoqIPNet_starGenType_String_results
}

// MoqIPNet_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the IPNet_starGenType type
type MoqIPNet_starGenType_String_doFn func()

// MoqIPNet_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the IPNet_starGenType type
type MoqIPNet_starGenType_String_doReturnFn func() string

// MoqIPNet_starGenType_String_results holds the results of the
// IPNet_starGenType type
type MoqIPNet_starGenType_String_results struct {
	Params  MoqIPNet_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_String_doFn
		DoReturnFn MoqIPNet_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPNet_starGenType_String_fnRecorder routes recorded function calls to the
// MoqIPNet_starGenType moq
type MoqIPNet_starGenType_String_fnRecorder struct {
	Params    MoqIPNet_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPNet_starGenType_String_results
	Moq       *MoqIPNet_starGenType
}

// MoqIPNet_starGenType_String_anyParams isolates the any params functions of
// the IPNet_starGenType type
type MoqIPNet_starGenType_String_anyParams struct {
	Recorder *MoqIPNet_starGenType_String_fnRecorder
}

// NewMoqIPNet_starGenType creates a new moq of the IPNet_starGenType type
func NewMoqIPNet_starGenType(scene *moq.Scene, config *moq.Config) *MoqIPNet_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIPNet_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIPNet_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Contains struct {
					Ip moq.ParamIndexing
				}
				Network struct{}
				String  struct{}
			}
		}{ParameterIndexing: struct {
			Contains struct {
				Ip moq.ParamIndexing
			}
			Network struct{}
			String  struct{}
		}{
			Contains: struct {
				Ip moq.ParamIndexing
			}{
				Ip: moq.ParamIndexByHash,
			},
			Network: struct{}{},
			String:  struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the IPNet_starGenType type
func (m *MoqIPNet_starGenType) Mock() *MoqIPNet_starGenType_mock { return m.Moq }

func (m *MoqIPNet_starGenType_mock) Contains(ip net.IP) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIPNet_starGenType_Contains_params{
		Ip: ip,
	}
	var results *MoqIPNet_starGenType_Contains_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Contains {
		paramsKey := m.Moq.ParamsKey_Contains(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Contains(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Contains(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Contains(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(ip)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ip)
	}
	return
}

func (m *MoqIPNet_starGenType_mock) Network() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqIPNet_starGenType_Network_params{}
	var results *MoqIPNet_starGenType_Network_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Network {
		paramsKey := m.Moq.ParamsKey_Network(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Network(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Network(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Network(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqIPNet_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqIPNet_starGenType_String_params{}
	var results *MoqIPNet_starGenType_String_results
	for _, resultsByParams := range m.Moq.ResultsByParams_String {
		paramsKey := m.Moq.ParamsKey_String(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_String(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_String(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_String(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the IPNet_starGenType type
func (m *MoqIPNet_starGenType) OnCall() *MoqIPNet_starGenType_recorder {
	return &MoqIPNet_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqIPNet_starGenType_recorder) Contains(ip net.IP) *MoqIPNet_starGenType_Contains_fnRecorder {
	return &MoqIPNet_starGenType_Contains_fnRecorder{
		Params: MoqIPNet_starGenType_Contains_params{
			Ip: ip,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) Any() *MoqIPNet_starGenType_Contains_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Contains(r.Params))
		return nil
	}
	return &MoqIPNet_starGenType_Contains_anyParams{Recorder: r}
}

func (a *MoqIPNet_starGenType_Contains_anyParams) Ip() *MoqIPNet_starGenType_Contains_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) Seq() *MoqIPNet_starGenType_Contains_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Contains(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) NoSeq() *MoqIPNet_starGenType_Contains_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Contains(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) ReturnResults(result1 bool) *MoqIPNet_starGenType_Contains_fnRecorder {
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
		DoFn       MoqIPNet_starGenType_Contains_doFn
		DoReturnFn MoqIPNet_starGenType_Contains_doReturnFn
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

func (r *MoqIPNet_starGenType_Contains_fnRecorder) AndDo(fn MoqIPNet_starGenType_Contains_doFn) *MoqIPNet_starGenType_Contains_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) DoReturnResults(fn MoqIPNet_starGenType_Contains_doReturnFn) *MoqIPNet_starGenType_Contains_fnRecorder {
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
		DoFn       MoqIPNet_starGenType_Contains_doFn
		DoReturnFn MoqIPNet_starGenType_Contains_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPNet_starGenType_Contains_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Contains {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIPNet_starGenType_Contains_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPNet_starGenType_Contains_paramsKey]*MoqIPNet_starGenType_Contains_results{},
		}
		r.Moq.ResultsByParams_Contains = append(r.Moq.ResultsByParams_Contains, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Contains) {
			copy(r.Moq.ResultsByParams_Contains[insertAt+1:], r.Moq.ResultsByParams_Contains[insertAt:0])
			r.Moq.ResultsByParams_Contains[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Contains(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIPNet_starGenType_Contains_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPNet_starGenType_Contains_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPNet_starGenType_Contains_fnRecorder {
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
				DoFn       MoqIPNet_starGenType_Contains_doFn
				DoReturnFn MoqIPNet_starGenType_Contains_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPNet_starGenType) PrettyParams_Contains(params MoqIPNet_starGenType_Contains_params) string {
	return fmt.Sprintf("Contains(%#v)", params.Ip)
}

func (m *MoqIPNet_starGenType) ParamsKey_Contains(params MoqIPNet_starGenType_Contains_params, anyParams uint64) MoqIPNet_starGenType_Contains_paramsKey {
	m.Scene.T.Helper()
	var ipUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Contains.Ip == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The ip parameter of the Contains function can't be indexed by value")
		}
		ipUsedHash = hash.DeepHash(params.Ip)
	}
	return MoqIPNet_starGenType_Contains_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Ip hash.Hash }{
			Ip: ipUsedHash,
		},
	}
}

func (m *MoqIPNet_starGenType_recorder) Network() *MoqIPNet_starGenType_Network_fnRecorder {
	return &MoqIPNet_starGenType_Network_fnRecorder{
		Params:   MoqIPNet_starGenType_Network_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) Any() *MoqIPNet_starGenType_Network_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	return &MoqIPNet_starGenType_Network_anyParams{Recorder: r}
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) Seq() *MoqIPNet_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) NoSeq() *MoqIPNet_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) ReturnResults(result1 string) *MoqIPNet_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_Network_doFn
		DoReturnFn MoqIPNet_starGenType_Network_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) AndDo(fn MoqIPNet_starGenType_Network_doFn) *MoqIPNet_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) DoReturnResults(fn MoqIPNet_starGenType_Network_doReturnFn) *MoqIPNet_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_Network_doFn
		DoReturnFn MoqIPNet_starGenType_Network_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPNet_starGenType_Network_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Network {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIPNet_starGenType_Network_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPNet_starGenType_Network_paramsKey]*MoqIPNet_starGenType_Network_results{},
		}
		r.Moq.ResultsByParams_Network = append(r.Moq.ResultsByParams_Network, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Network) {
			copy(r.Moq.ResultsByParams_Network[insertAt+1:], r.Moq.ResultsByParams_Network[insertAt:0])
			r.Moq.ResultsByParams_Network[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Network(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIPNet_starGenType_Network_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPNet_starGenType_Network_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPNet_starGenType_Network_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqIPNet_starGenType_Network_doFn
				DoReturnFn MoqIPNet_starGenType_Network_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPNet_starGenType) PrettyParams_Network(params MoqIPNet_starGenType_Network_params) string {
	return fmt.Sprintf("Network()")
}

func (m *MoqIPNet_starGenType) ParamsKey_Network(params MoqIPNet_starGenType_Network_params, anyParams uint64) MoqIPNet_starGenType_Network_paramsKey {
	m.Scene.T.Helper()
	return MoqIPNet_starGenType_Network_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqIPNet_starGenType_recorder) String() *MoqIPNet_starGenType_String_fnRecorder {
	return &MoqIPNet_starGenType_String_fnRecorder{
		Params:   MoqIPNet_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPNet_starGenType_String_fnRecorder) Any() *MoqIPNet_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqIPNet_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqIPNet_starGenType_String_fnRecorder) Seq() *MoqIPNet_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPNet_starGenType_String_fnRecorder) NoSeq() *MoqIPNet_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPNet_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqIPNet_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_String_doFn
		DoReturnFn MoqIPNet_starGenType_String_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIPNet_starGenType_String_fnRecorder) AndDo(fn MoqIPNet_starGenType_String_doFn) *MoqIPNet_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPNet_starGenType_String_fnRecorder) DoReturnResults(fn MoqIPNet_starGenType_String_doReturnFn) *MoqIPNet_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPNet_starGenType_String_doFn
		DoReturnFn MoqIPNet_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPNet_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPNet_starGenType_String_resultsByParams
	for n, res := range r.Moq.ResultsByParams_String {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIPNet_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPNet_starGenType_String_paramsKey]*MoqIPNet_starGenType_String_results{},
		}
		r.Moq.ResultsByParams_String = append(r.Moq.ResultsByParams_String, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_String) {
			copy(r.Moq.ResultsByParams_String[insertAt+1:], r.Moq.ResultsByParams_String[insertAt:0])
			r.Moq.ResultsByParams_String[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_String(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIPNet_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPNet_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPNet_starGenType_String_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqIPNet_starGenType_String_doFn
				DoReturnFn MoqIPNet_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPNet_starGenType) PrettyParams_String(params MoqIPNet_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqIPNet_starGenType) ParamsKey_String(params MoqIPNet_starGenType_String_params, anyParams uint64) MoqIPNet_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqIPNet_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqIPNet_starGenType) Reset() {
	m.ResultsByParams_Contains = nil
	m.ResultsByParams_Network = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIPNet_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Contains {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Contains(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Network {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Network(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
}
