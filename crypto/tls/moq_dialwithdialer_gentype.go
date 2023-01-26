// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package tls

import (
	"crypto/tls"
	"fmt"
	"math/bits"
	"net"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// DialWithDialer_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type DialWithDialer_genType func(dialer *net.Dialer, network, addr string, config *tls.Config) (*tls.Conn, error)

// MoqDialWithDialer_genType holds the state of a moq of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDialWithDialer_genType_mock

	ResultsByParams []MoqDialWithDialer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dialer  moq.ParamIndexing
			Network moq.ParamIndexing
			Addr    moq.ParamIndexing
			Config  moq.ParamIndexing
		}
	}
}

// MoqDialWithDialer_genType_mock isolates the mock interface of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType_mock struct {
	Moq *MoqDialWithDialer_genType
}

// MoqDialWithDialer_genType_params holds the params of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType_params struct {
	Dialer        *net.Dialer
	Network, Addr string
	Config        *tls.Config
}

// MoqDialWithDialer_genType_paramsKey holds the map key params of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType_paramsKey struct {
	Params struct {
		Dialer        *net.Dialer
		Network, Addr string
		Config        *tls.Config
	}
	Hashes struct {
		Dialer        hash.Hash
		Network, Addr hash.Hash
		Config        hash.Hash
	}
}

// MoqDialWithDialer_genType_resultsByParams contains the results for a given
// set of parameters for the DialWithDialer_genType type
type MoqDialWithDialer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDialWithDialer_genType_paramsKey]*MoqDialWithDialer_genType_results
}

// MoqDialWithDialer_genType_doFn defines the type of function needed when
// calling AndDo for the DialWithDialer_genType type
type MoqDialWithDialer_genType_doFn func(dialer *net.Dialer, network, addr string, config *tls.Config)

// MoqDialWithDialer_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the DialWithDialer_genType type
type MoqDialWithDialer_genType_doReturnFn func(dialer *net.Dialer, network, addr string, config *tls.Config) (*tls.Conn, error)

// MoqDialWithDialer_genType_results holds the results of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType_results struct {
	Params  MoqDialWithDialer_genType_params
	Results []struct {
		Values *struct {
			Result1 *tls.Conn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialWithDialer_genType_doFn
		DoReturnFn MoqDialWithDialer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDialWithDialer_genType_fnRecorder routes recorded function calls to the
// MoqDialWithDialer_genType moq
type MoqDialWithDialer_genType_fnRecorder struct {
	Params    MoqDialWithDialer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDialWithDialer_genType_results
	Moq       *MoqDialWithDialer_genType
}

// MoqDialWithDialer_genType_anyParams isolates the any params functions of the
// DialWithDialer_genType type
type MoqDialWithDialer_genType_anyParams struct {
	Recorder *MoqDialWithDialer_genType_fnRecorder
}

// NewMoqDialWithDialer_genType creates a new moq of the DialWithDialer_genType
// type
func NewMoqDialWithDialer_genType(scene *moq.Scene, config *moq.Config) *MoqDialWithDialer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDialWithDialer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDialWithDialer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dialer  moq.ParamIndexing
				Network moq.ParamIndexing
				Addr    moq.ParamIndexing
				Config  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dialer  moq.ParamIndexing
			Network moq.ParamIndexing
			Addr    moq.ParamIndexing
			Config  moq.ParamIndexing
		}{
			Dialer:  moq.ParamIndexByHash,
			Network: moq.ParamIndexByValue,
			Addr:    moq.ParamIndexByValue,
			Config:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DialWithDialer_genType type
func (m *MoqDialWithDialer_genType) Mock() DialWithDialer_genType {
	return func(dialer *net.Dialer, network, addr string, config *tls.Config) (*tls.Conn, error) {
		m.Scene.T.Helper()
		moq := &MoqDialWithDialer_genType_mock{Moq: m}
		return moq.Fn(dialer, network, addr, config)
	}
}

func (m *MoqDialWithDialer_genType_mock) Fn(dialer *net.Dialer, network, addr string, config *tls.Config) (result1 *tls.Conn, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDialWithDialer_genType_params{
		Dialer:  dialer,
		Network: network,
		Addr:    addr,
		Config:  config,
	}
	var results *MoqDialWithDialer_genType_results
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
		result.DoFn(dialer, network, addr, config)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(dialer, network, addr, config)
	}
	return
}

func (m *MoqDialWithDialer_genType) OnCall(dialer *net.Dialer, network, addr string, config *tls.Config) *MoqDialWithDialer_genType_fnRecorder {
	return &MoqDialWithDialer_genType_fnRecorder{
		Params: MoqDialWithDialer_genType_params{
			Dialer:  dialer,
			Network: network,
			Addr:    addr,
			Config:  config,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDialWithDialer_genType_fnRecorder) Any() *MoqDialWithDialer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDialWithDialer_genType_anyParams{Recorder: r}
}

func (a *MoqDialWithDialer_genType_anyParams) Dialer() *MoqDialWithDialer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqDialWithDialer_genType_anyParams) Network() *MoqDialWithDialer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqDialWithDialer_genType_anyParams) Addr() *MoqDialWithDialer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqDialWithDialer_genType_anyParams) Config() *MoqDialWithDialer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqDialWithDialer_genType_fnRecorder) Seq() *MoqDialWithDialer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDialWithDialer_genType_fnRecorder) NoSeq() *MoqDialWithDialer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDialWithDialer_genType_fnRecorder) ReturnResults(result1 *tls.Conn, result2 error) *MoqDialWithDialer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *tls.Conn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialWithDialer_genType_doFn
		DoReturnFn MoqDialWithDialer_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *tls.Conn
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDialWithDialer_genType_fnRecorder) AndDo(fn MoqDialWithDialer_genType_doFn) *MoqDialWithDialer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDialWithDialer_genType_fnRecorder) DoReturnResults(fn MoqDialWithDialer_genType_doReturnFn) *MoqDialWithDialer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *tls.Conn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialWithDialer_genType_doFn
		DoReturnFn MoqDialWithDialer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDialWithDialer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDialWithDialer_genType_resultsByParams
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
		results = &MoqDialWithDialer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDialWithDialer_genType_paramsKey]*MoqDialWithDialer_genType_results{},
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
		r.Results = &MoqDialWithDialer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDialWithDialer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDialWithDialer_genType_fnRecorder {
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
					Result1 *tls.Conn
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqDialWithDialer_genType_doFn
				DoReturnFn MoqDialWithDialer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDialWithDialer_genType) PrettyParams(params MoqDialWithDialer_genType_params) string {
	return fmt.Sprintf("DialWithDialer_genType(%#v, %#v, %#v, %#v)", params.Dialer, params.Network, params.Addr, params.Config)
}

func (m *MoqDialWithDialer_genType) ParamsKey(params MoqDialWithDialer_genType_params, anyParams uint64) MoqDialWithDialer_genType_paramsKey {
	m.Scene.T.Helper()
	var dialerUsed *net.Dialer
	var dialerUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dialer == moq.ParamIndexByValue {
			dialerUsed = params.Dialer
		} else {
			dialerUsedHash = hash.DeepHash(params.Dialer)
		}
	}
	var networkUsed string
	var networkUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Network == moq.ParamIndexByValue {
			networkUsed = params.Network
		} else {
			networkUsedHash = hash.DeepHash(params.Network)
		}
	}
	var addrUsed string
	var addrUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	var configUsed *tls.Config
	var configUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Config == moq.ParamIndexByValue {
			configUsed = params.Config
		} else {
			configUsedHash = hash.DeepHash(params.Config)
		}
	}
	return MoqDialWithDialer_genType_paramsKey{
		Params: struct {
			Dialer        *net.Dialer
			Network, Addr string
			Config        *tls.Config
		}{
			Dialer:  dialerUsed,
			Network: networkUsed,
			Addr:    addrUsed,
			Config:  configUsed,
		},
		Hashes: struct {
			Dialer        hash.Hash
			Network, Addr hash.Hash
			Config        hash.Hash
		}{
			Dialer:  dialerUsedHash,
			Network: networkUsedHash,
			Addr:    addrUsedHash,
			Config:  configUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDialWithDialer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDialWithDialer_genType) AssertExpectationsMet() {
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
