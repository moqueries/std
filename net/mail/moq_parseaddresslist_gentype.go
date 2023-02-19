// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package mail

import (
	"fmt"
	"math/bits"
	"net/mail"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseAddressList_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseAddressList_genType func(list string) ([]*mail.Address, error)

// MoqParseAddressList_genType holds the state of a moq of the
// ParseAddressList_genType type
type MoqParseAddressList_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseAddressList_genType_mock

	ResultsByParams []MoqParseAddressList_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			List moq.ParamIndexing
		}
	}
}

// MoqParseAddressList_genType_mock isolates the mock interface of the
// ParseAddressList_genType type
type MoqParseAddressList_genType_mock struct {
	Moq *MoqParseAddressList_genType
}

// MoqParseAddressList_genType_params holds the params of the
// ParseAddressList_genType type
type MoqParseAddressList_genType_params struct{ List string }

// MoqParseAddressList_genType_paramsKey holds the map key params of the
// ParseAddressList_genType type
type MoqParseAddressList_genType_paramsKey struct {
	Params struct{ List string }
	Hashes struct{ List hash.Hash }
}

// MoqParseAddressList_genType_resultsByParams contains the results for a given
// set of parameters for the ParseAddressList_genType type
type MoqParseAddressList_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseAddressList_genType_paramsKey]*MoqParseAddressList_genType_results
}

// MoqParseAddressList_genType_doFn defines the type of function needed when
// calling AndDo for the ParseAddressList_genType type
type MoqParseAddressList_genType_doFn func(list string)

// MoqParseAddressList_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ParseAddressList_genType type
type MoqParseAddressList_genType_doReturnFn func(list string) ([]*mail.Address, error)

// MoqParseAddressList_genType_results holds the results of the
// ParseAddressList_genType type
type MoqParseAddressList_genType_results struct {
	Params  MoqParseAddressList_genType_params
	Results []struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseAddressList_genType_doFn
		DoReturnFn MoqParseAddressList_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseAddressList_genType_fnRecorder routes recorded function calls to the
// MoqParseAddressList_genType moq
type MoqParseAddressList_genType_fnRecorder struct {
	Params    MoqParseAddressList_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseAddressList_genType_results
	Moq       *MoqParseAddressList_genType
}

// MoqParseAddressList_genType_anyParams isolates the any params functions of
// the ParseAddressList_genType type
type MoqParseAddressList_genType_anyParams struct {
	Recorder *MoqParseAddressList_genType_fnRecorder
}

// NewMoqParseAddressList_genType creates a new moq of the
// ParseAddressList_genType type
func NewMoqParseAddressList_genType(scene *moq.Scene, config *moq.Config) *MoqParseAddressList_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseAddressList_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseAddressList_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				List moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			List moq.ParamIndexing
		}{
			List: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseAddressList_genType type
func (m *MoqParseAddressList_genType) Mock() ParseAddressList_genType {
	return func(list string) ([]*mail.Address, error) {
		m.Scene.T.Helper()
		moq := &MoqParseAddressList_genType_mock{Moq: m}
		return moq.Fn(list)
	}
}

func (m *MoqParseAddressList_genType_mock) Fn(list string) (result1 []*mail.Address, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseAddressList_genType_params{
		List: list,
	}
	var results *MoqParseAddressList_genType_results
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
		result.DoFn(list)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(list)
	}
	return
}

func (m *MoqParseAddressList_genType) OnCall(list string) *MoqParseAddressList_genType_fnRecorder {
	return &MoqParseAddressList_genType_fnRecorder{
		Params: MoqParseAddressList_genType_params{
			List: list,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseAddressList_genType_fnRecorder) Any() *MoqParseAddressList_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseAddressList_genType_anyParams{Recorder: r}
}

func (a *MoqParseAddressList_genType_anyParams) List() *MoqParseAddressList_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParseAddressList_genType_fnRecorder) Seq() *MoqParseAddressList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseAddressList_genType_fnRecorder) NoSeq() *MoqParseAddressList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseAddressList_genType_fnRecorder) ReturnResults(result1 []*mail.Address, result2 error) *MoqParseAddressList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseAddressList_genType_doFn
		DoReturnFn MoqParseAddressList_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []*mail.Address
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseAddressList_genType_fnRecorder) AndDo(fn MoqParseAddressList_genType_doFn) *MoqParseAddressList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseAddressList_genType_fnRecorder) DoReturnResults(fn MoqParseAddressList_genType_doReturnFn) *MoqParseAddressList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseAddressList_genType_doFn
		DoReturnFn MoqParseAddressList_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseAddressList_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseAddressList_genType_resultsByParams
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
		results = &MoqParseAddressList_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseAddressList_genType_paramsKey]*MoqParseAddressList_genType_results{},
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
		r.Results = &MoqParseAddressList_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseAddressList_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseAddressList_genType_fnRecorder {
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
					Result1 []*mail.Address
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqParseAddressList_genType_doFn
				DoReturnFn MoqParseAddressList_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseAddressList_genType) PrettyParams(params MoqParseAddressList_genType_params) string {
	return fmt.Sprintf("ParseAddressList_genType(%#v)", params.List)
}

func (m *MoqParseAddressList_genType) ParamsKey(params MoqParseAddressList_genType_params, anyParams uint64) MoqParseAddressList_genType_paramsKey {
	m.Scene.T.Helper()
	var listUsed string
	var listUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.List == moq.ParamIndexByValue {
			listUsed = params.List
		} else {
			listUsedHash = hash.DeepHash(params.List)
		}
	}
	return MoqParseAddressList_genType_paramsKey{
		Params: struct{ List string }{
			List: listUsed,
		},
		Hashes: struct{ List hash.Hash }{
			List: listUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseAddressList_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseAddressList_genType) AssertExpectationsMet() {
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
