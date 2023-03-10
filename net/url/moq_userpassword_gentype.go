// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package url

import (
	"fmt"
	"math/bits"
	"net/url"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// UserPassword_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type UserPassword_genType func(username, password string) *url.Userinfo

// MoqUserPassword_genType holds the state of a moq of the UserPassword_genType
// type
type MoqUserPassword_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUserPassword_genType_mock

	ResultsByParams []MoqUserPassword_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Username moq.ParamIndexing
			Password moq.ParamIndexing
		}
	}
}

// MoqUserPassword_genType_mock isolates the mock interface of the
// UserPassword_genType type
type MoqUserPassword_genType_mock struct {
	Moq *MoqUserPassword_genType
}

// MoqUserPassword_genType_params holds the params of the UserPassword_genType
// type
type MoqUserPassword_genType_params struct{ Username, Password string }

// MoqUserPassword_genType_paramsKey holds the map key params of the
// UserPassword_genType type
type MoqUserPassword_genType_paramsKey struct {
	Params struct{ Username, Password string }
	Hashes struct{ Username, Password hash.Hash }
}

// MoqUserPassword_genType_resultsByParams contains the results for a given set
// of parameters for the UserPassword_genType type
type MoqUserPassword_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUserPassword_genType_paramsKey]*MoqUserPassword_genType_results
}

// MoqUserPassword_genType_doFn defines the type of function needed when
// calling AndDo for the UserPassword_genType type
type MoqUserPassword_genType_doFn func(username, password string)

// MoqUserPassword_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the UserPassword_genType type
type MoqUserPassword_genType_doReturnFn func(username, password string) *url.Userinfo

// MoqUserPassword_genType_results holds the results of the
// UserPassword_genType type
type MoqUserPassword_genType_results struct {
	Params  MoqUserPassword_genType_params
	Results []struct {
		Values *struct {
			Result1 *url.Userinfo
		}
		Sequence   uint32
		DoFn       MoqUserPassword_genType_doFn
		DoReturnFn MoqUserPassword_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUserPassword_genType_fnRecorder routes recorded function calls to the
// MoqUserPassword_genType moq
type MoqUserPassword_genType_fnRecorder struct {
	Params    MoqUserPassword_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUserPassword_genType_results
	Moq       *MoqUserPassword_genType
}

// MoqUserPassword_genType_anyParams isolates the any params functions of the
// UserPassword_genType type
type MoqUserPassword_genType_anyParams struct {
	Recorder *MoqUserPassword_genType_fnRecorder
}

// NewMoqUserPassword_genType creates a new moq of the UserPassword_genType
// type
func NewMoqUserPassword_genType(scene *moq.Scene, config *moq.Config) *MoqUserPassword_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUserPassword_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUserPassword_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Username moq.ParamIndexing
				Password moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Username moq.ParamIndexing
			Password moq.ParamIndexing
		}{
			Username: moq.ParamIndexByValue,
			Password: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the UserPassword_genType type
func (m *MoqUserPassword_genType) Mock() UserPassword_genType {
	return func(username, password string) *url.Userinfo {
		m.Scene.T.Helper()
		moq := &MoqUserPassword_genType_mock{Moq: m}
		return moq.Fn(username, password)
	}
}

func (m *MoqUserPassword_genType_mock) Fn(username, password string) (result1 *url.Userinfo) {
	m.Moq.Scene.T.Helper()
	params := MoqUserPassword_genType_params{
		Username: username,
		Password: password,
	}
	var results *MoqUserPassword_genType_results
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
		result.DoFn(username, password)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(username, password)
	}
	return
}

func (m *MoqUserPassword_genType) OnCall(username, password string) *MoqUserPassword_genType_fnRecorder {
	return &MoqUserPassword_genType_fnRecorder{
		Params: MoqUserPassword_genType_params{
			Username: username,
			Password: password,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUserPassword_genType_fnRecorder) Any() *MoqUserPassword_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUserPassword_genType_anyParams{Recorder: r}
}

func (a *MoqUserPassword_genType_anyParams) Username() *MoqUserPassword_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqUserPassword_genType_anyParams) Password() *MoqUserPassword_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqUserPassword_genType_fnRecorder) Seq() *MoqUserPassword_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUserPassword_genType_fnRecorder) NoSeq() *MoqUserPassword_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUserPassword_genType_fnRecorder) ReturnResults(result1 *url.Userinfo) *MoqUserPassword_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *url.Userinfo
		}
		Sequence   uint32
		DoFn       MoqUserPassword_genType_doFn
		DoReturnFn MoqUserPassword_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *url.Userinfo
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUserPassword_genType_fnRecorder) AndDo(fn MoqUserPassword_genType_doFn) *MoqUserPassword_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUserPassword_genType_fnRecorder) DoReturnResults(fn MoqUserPassword_genType_doReturnFn) *MoqUserPassword_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *url.Userinfo
		}
		Sequence   uint32
		DoFn       MoqUserPassword_genType_doFn
		DoReturnFn MoqUserPassword_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUserPassword_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUserPassword_genType_resultsByParams
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
		results = &MoqUserPassword_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUserPassword_genType_paramsKey]*MoqUserPassword_genType_results{},
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
		r.Results = &MoqUserPassword_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUserPassword_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUserPassword_genType_fnRecorder {
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
					Result1 *url.Userinfo
				}
				Sequence   uint32
				DoFn       MoqUserPassword_genType_doFn
				DoReturnFn MoqUserPassword_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUserPassword_genType) PrettyParams(params MoqUserPassword_genType_params) string {
	return fmt.Sprintf("UserPassword_genType(%#v, %#v)", params.Username, params.Password)
}

func (m *MoqUserPassword_genType) ParamsKey(params MoqUserPassword_genType_params, anyParams uint64) MoqUserPassword_genType_paramsKey {
	m.Scene.T.Helper()
	var usernameUsed string
	var usernameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Username == moq.ParamIndexByValue {
			usernameUsed = params.Username
		} else {
			usernameUsedHash = hash.DeepHash(params.Username)
		}
	}
	var passwordUsed string
	var passwordUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Password == moq.ParamIndexByValue {
			passwordUsed = params.Password
		} else {
			passwordUsedHash = hash.DeepHash(params.Password)
		}
	}
	return MoqUserPassword_genType_paramsKey{
		Params: struct{ Username, Password string }{
			Username: usernameUsed,
			Password: passwordUsed,
		},
		Hashes: struct{ Username, Password hash.Hash }{
			Username: usernameUsedHash,
			Password: passwordUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqUserPassword_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUserPassword_genType) AssertExpectationsMet() {
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
