// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package netip

import (
	"fmt"
	"math/bits"
	"net/netip"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// PrefixFrom_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type PrefixFrom_genType func(ip netip.Addr, bits int) netip.Prefix

// MoqPrefixFrom_genType holds the state of a moq of the PrefixFrom_genType
// type
type MoqPrefixFrom_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPrefixFrom_genType_mock

	ResultsByParams []MoqPrefixFrom_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ip   moq.ParamIndexing
			Bits moq.ParamIndexing
		}
	}
}

// MoqPrefixFrom_genType_mock isolates the mock interface of the
// PrefixFrom_genType type
type MoqPrefixFrom_genType_mock struct {
	Moq *MoqPrefixFrom_genType
}

// MoqPrefixFrom_genType_params holds the params of the PrefixFrom_genType type
type MoqPrefixFrom_genType_params struct {
	Ip   netip.Addr
	Bits int
}

// MoqPrefixFrom_genType_paramsKey holds the map key params of the
// PrefixFrom_genType type
type MoqPrefixFrom_genType_paramsKey struct {
	Params struct {
		Ip   netip.Addr
		Bits int
	}
	Hashes struct {
		Ip   hash.Hash
		Bits hash.Hash
	}
}

// MoqPrefixFrom_genType_resultsByParams contains the results for a given set
// of parameters for the PrefixFrom_genType type
type MoqPrefixFrom_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPrefixFrom_genType_paramsKey]*MoqPrefixFrom_genType_results
}

// MoqPrefixFrom_genType_doFn defines the type of function needed when calling
// AndDo for the PrefixFrom_genType type
type MoqPrefixFrom_genType_doFn func(ip netip.Addr, bits int)

// MoqPrefixFrom_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the PrefixFrom_genType type
type MoqPrefixFrom_genType_doReturnFn func(ip netip.Addr, bits int) netip.Prefix

// MoqPrefixFrom_genType_results holds the results of the PrefixFrom_genType
// type
type MoqPrefixFrom_genType_results struct {
	Params  MoqPrefixFrom_genType_params
	Results []struct {
		Values *struct {
			Result1 netip.Prefix
		}
		Sequence   uint32
		DoFn       MoqPrefixFrom_genType_doFn
		DoReturnFn MoqPrefixFrom_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPrefixFrom_genType_fnRecorder routes recorded function calls to the
// MoqPrefixFrom_genType moq
type MoqPrefixFrom_genType_fnRecorder struct {
	Params    MoqPrefixFrom_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPrefixFrom_genType_results
	Moq       *MoqPrefixFrom_genType
}

// MoqPrefixFrom_genType_anyParams isolates the any params functions of the
// PrefixFrom_genType type
type MoqPrefixFrom_genType_anyParams struct {
	Recorder *MoqPrefixFrom_genType_fnRecorder
}

// NewMoqPrefixFrom_genType creates a new moq of the PrefixFrom_genType type
func NewMoqPrefixFrom_genType(scene *moq.Scene, config *moq.Config) *MoqPrefixFrom_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPrefixFrom_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPrefixFrom_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ip   moq.ParamIndexing
				Bits moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ip   moq.ParamIndexing
			Bits moq.ParamIndexing
		}{
			Ip:   moq.ParamIndexByHash,
			Bits: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PrefixFrom_genType type
func (m *MoqPrefixFrom_genType) Mock() PrefixFrom_genType {
	return func(ip netip.Addr, bits int) netip.Prefix {
		m.Scene.T.Helper()
		moq := &MoqPrefixFrom_genType_mock{Moq: m}
		return moq.Fn(ip, bits)
	}
}

func (m *MoqPrefixFrom_genType_mock) Fn(ip netip.Addr, bits int) (result1 netip.Prefix) {
	m.Moq.Scene.T.Helper()
	params := MoqPrefixFrom_genType_params{
		Ip:   ip,
		Bits: bits,
	}
	var results *MoqPrefixFrom_genType_results
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
		result.DoFn(ip, bits)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ip, bits)
	}
	return
}

func (m *MoqPrefixFrom_genType) OnCall(ip netip.Addr, bits int) *MoqPrefixFrom_genType_fnRecorder {
	return &MoqPrefixFrom_genType_fnRecorder{
		Params: MoqPrefixFrom_genType_params{
			Ip:   ip,
			Bits: bits,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPrefixFrom_genType_fnRecorder) Any() *MoqPrefixFrom_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPrefixFrom_genType_anyParams{Recorder: r}
}

func (a *MoqPrefixFrom_genType_anyParams) Ip() *MoqPrefixFrom_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqPrefixFrom_genType_anyParams) Bits() *MoqPrefixFrom_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqPrefixFrom_genType_fnRecorder) Seq() *MoqPrefixFrom_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPrefixFrom_genType_fnRecorder) NoSeq() *MoqPrefixFrom_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPrefixFrom_genType_fnRecorder) ReturnResults(result1 netip.Prefix) *MoqPrefixFrom_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 netip.Prefix
		}
		Sequence   uint32
		DoFn       MoqPrefixFrom_genType_doFn
		DoReturnFn MoqPrefixFrom_genType_doReturnFn
	}{
		Values: &struct {
			Result1 netip.Prefix
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPrefixFrom_genType_fnRecorder) AndDo(fn MoqPrefixFrom_genType_doFn) *MoqPrefixFrom_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPrefixFrom_genType_fnRecorder) DoReturnResults(fn MoqPrefixFrom_genType_doReturnFn) *MoqPrefixFrom_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 netip.Prefix
		}
		Sequence   uint32
		DoFn       MoqPrefixFrom_genType_doFn
		DoReturnFn MoqPrefixFrom_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPrefixFrom_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPrefixFrom_genType_resultsByParams
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
		results = &MoqPrefixFrom_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPrefixFrom_genType_paramsKey]*MoqPrefixFrom_genType_results{},
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
		r.Results = &MoqPrefixFrom_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPrefixFrom_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPrefixFrom_genType_fnRecorder {
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
					Result1 netip.Prefix
				}
				Sequence   uint32
				DoFn       MoqPrefixFrom_genType_doFn
				DoReturnFn MoqPrefixFrom_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPrefixFrom_genType) PrettyParams(params MoqPrefixFrom_genType_params) string {
	return fmt.Sprintf("PrefixFrom_genType(%#v, %#v)", params.Ip, params.Bits)
}

func (m *MoqPrefixFrom_genType) ParamsKey(params MoqPrefixFrom_genType_params, anyParams uint64) MoqPrefixFrom_genType_paramsKey {
	m.Scene.T.Helper()
	var ipUsed netip.Addr
	var ipUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Ip == moq.ParamIndexByValue {
			ipUsed = params.Ip
		} else {
			ipUsedHash = hash.DeepHash(params.Ip)
		}
	}
	var bitsUsed int
	var bitsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Bits == moq.ParamIndexByValue {
			bitsUsed = params.Bits
		} else {
			bitsUsedHash = hash.DeepHash(params.Bits)
		}
	}
	return MoqPrefixFrom_genType_paramsKey{
		Params: struct {
			Ip   netip.Addr
			Bits int
		}{
			Ip:   ipUsed,
			Bits: bitsUsed,
		},
		Hashes: struct {
			Ip   hash.Hash
			Bits hash.Hash
		}{
			Ip:   ipUsedHash,
			Bits: bitsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPrefixFrom_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPrefixFrom_genType) AssertExpectationsMet() {
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
