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

// Server_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Server_genType func(conn net.Conn, config *tls.Config) *tls.Conn

// MoqServer_genType holds the state of a moq of the Server_genType type
type MoqServer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqServer_genType_mock

	ResultsByParams []MoqServer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Conn   moq.ParamIndexing
			Config moq.ParamIndexing
		}
	}
}

// MoqServer_genType_mock isolates the mock interface of the Server_genType
// type
type MoqServer_genType_mock struct {
	Moq *MoqServer_genType
}

// MoqServer_genType_params holds the params of the Server_genType type
type MoqServer_genType_params struct {
	Conn   net.Conn
	Config *tls.Config
}

// MoqServer_genType_paramsKey holds the map key params of the Server_genType
// type
type MoqServer_genType_paramsKey struct {
	Params struct {
		Conn   net.Conn
		Config *tls.Config
	}
	Hashes struct {
		Conn   hash.Hash
		Config hash.Hash
	}
}

// MoqServer_genType_resultsByParams contains the results for a given set of
// parameters for the Server_genType type
type MoqServer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqServer_genType_paramsKey]*MoqServer_genType_results
}

// MoqServer_genType_doFn defines the type of function needed when calling
// AndDo for the Server_genType type
type MoqServer_genType_doFn func(conn net.Conn, config *tls.Config)

// MoqServer_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Server_genType type
type MoqServer_genType_doReturnFn func(conn net.Conn, config *tls.Config) *tls.Conn

// MoqServer_genType_results holds the results of the Server_genType type
type MoqServer_genType_results struct {
	Params  MoqServer_genType_params
	Results []struct {
		Values *struct {
			Result1 *tls.Conn
		}
		Sequence   uint32
		DoFn       MoqServer_genType_doFn
		DoReturnFn MoqServer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqServer_genType_fnRecorder routes recorded function calls to the
// MoqServer_genType moq
type MoqServer_genType_fnRecorder struct {
	Params    MoqServer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqServer_genType_results
	Moq       *MoqServer_genType
}

// MoqServer_genType_anyParams isolates the any params functions of the
// Server_genType type
type MoqServer_genType_anyParams struct {
	Recorder *MoqServer_genType_fnRecorder
}

// NewMoqServer_genType creates a new moq of the Server_genType type
func NewMoqServer_genType(scene *moq.Scene, config *moq.Config) *MoqServer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqServer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqServer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Conn   moq.ParamIndexing
				Config moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Conn   moq.ParamIndexing
			Config moq.ParamIndexing
		}{
			Conn:   moq.ParamIndexByHash,
			Config: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Server_genType type
func (m *MoqServer_genType) Mock() Server_genType {
	return func(conn net.Conn, config *tls.Config) *tls.Conn {
		m.Scene.T.Helper()
		moq := &MoqServer_genType_mock{Moq: m}
		return moq.Fn(conn, config)
	}
}

func (m *MoqServer_genType_mock) Fn(conn net.Conn, config *tls.Config) (result1 *tls.Conn) {
	m.Moq.Scene.T.Helper()
	params := MoqServer_genType_params{
		Conn:   conn,
		Config: config,
	}
	var results *MoqServer_genType_results
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
		result.DoFn(conn, config)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(conn, config)
	}
	return
}

func (m *MoqServer_genType) OnCall(conn net.Conn, config *tls.Config) *MoqServer_genType_fnRecorder {
	return &MoqServer_genType_fnRecorder{
		Params: MoqServer_genType_params{
			Conn:   conn,
			Config: config,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqServer_genType_fnRecorder) Any() *MoqServer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqServer_genType_anyParams{Recorder: r}
}

func (a *MoqServer_genType_anyParams) Conn() *MoqServer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqServer_genType_anyParams) Config() *MoqServer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqServer_genType_fnRecorder) Seq() *MoqServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqServer_genType_fnRecorder) NoSeq() *MoqServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqServer_genType_fnRecorder) ReturnResults(result1 *tls.Conn) *MoqServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *tls.Conn
		}
		Sequence   uint32
		DoFn       MoqServer_genType_doFn
		DoReturnFn MoqServer_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *tls.Conn
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqServer_genType_fnRecorder) AndDo(fn MoqServer_genType_doFn) *MoqServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqServer_genType_fnRecorder) DoReturnResults(fn MoqServer_genType_doReturnFn) *MoqServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *tls.Conn
		}
		Sequence   uint32
		DoFn       MoqServer_genType_doFn
		DoReturnFn MoqServer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqServer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqServer_genType_resultsByParams
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
		results = &MoqServer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqServer_genType_paramsKey]*MoqServer_genType_results{},
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
		r.Results = &MoqServer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqServer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqServer_genType_fnRecorder {
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
				}
				Sequence   uint32
				DoFn       MoqServer_genType_doFn
				DoReturnFn MoqServer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqServer_genType) PrettyParams(params MoqServer_genType_params) string {
	return fmt.Sprintf("Server_genType(%#v, %#v)", params.Conn, params.Config)
}

func (m *MoqServer_genType) ParamsKey(params MoqServer_genType_params, anyParams uint64) MoqServer_genType_paramsKey {
	m.Scene.T.Helper()
	var connUsed net.Conn
	var connUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Conn == moq.ParamIndexByValue {
			connUsed = params.Conn
		} else {
			connUsedHash = hash.DeepHash(params.Conn)
		}
	}
	var configUsed *tls.Config
	var configUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Config == moq.ParamIndexByValue {
			configUsed = params.Config
		} else {
			configUsedHash = hash.DeepHash(params.Config)
		}
	}
	return MoqServer_genType_paramsKey{
		Params: struct {
			Conn   net.Conn
			Config *tls.Config
		}{
			Conn:   connUsed,
			Config: configUsed,
		},
		Hashes: struct {
			Conn   hash.Hash
			Config hash.Hash
		}{
			Conn:   connUsedHash,
			Config: configUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqServer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqServer_genType) AssertExpectationsMet() {
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
