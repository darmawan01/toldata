// Code generated by github.com/darmawan01/toldata. DO NOT EDIT.
// package: cdl.toldatatest
// source: toldata_test.proto

package test

import (
	"context"
	"errors"
	"github.com/darmawan01/toldata"
	"google.golang.org/protobuf/proto"
	"io"

	nats "github.com/nats-io/nats.go"
)

// Workaround for template problem
func _eof() error {
	return io.EOF
}

type TestServiceToldataInterface interface {
	ToldataHealthCheck(ctx context.Context, req *toldata.Empty) (*toldata.ToldataHealthCheckInfo, error)

	GetTestA(ctx context.Context, req *TestARequest) (*TestAResponse, error)

	GetTestAB(ctx context.Context, req *TestARequest) (*TestAResponse, error)

	GetTestGetIP(ctx context.Context, req *Empty) (*TestGetIPResponse, error)

	FeedData(stream TestService_FeedDataToldataServer)

	StreamData(req *StreamDataRequest, stream TestService_StreamDataToldataServer) error

	StreamDataAlt1(req *StreamDataRequest, stream TestService_StreamDataAlt1ToldataServer) error

	TestEmpty(ctx context.Context, req *Empty) (*Empty, error)
}

type TestServiceToldataClient struct {
	Bus *toldata.Bus
}

type TestServiceToldataServer struct {
	Bus     *toldata.Bus
	Service TestServiceToldataInterface
}

func NewTestServiceToldataClient(bus *toldata.Bus) *TestServiceToldataClient {
	s := &TestServiceToldataClient{Bus: bus}
	return s
}

func NewTestServiceToldataServer(bus *toldata.Bus, service TestServiceToldataInterface) *TestServiceToldataServer {
	s := &TestServiceToldataServer{Bus: bus, Service: service}
	return s
}

func (service *TestServiceToldataClient) ToldataHealthCheck(ctx context.Context, req *toldata.Empty) (*toldata.ToldataHealthCheckInfo, error) {
	functionName := "cdl.toldatatest/TestService/ToldataHealthCheck"

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &toldata.ToldataHealthCheckInfo{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (service *TestServiceToldataClient) GetTestA(ctx context.Context, req *TestARequest) (*TestAResponse, error) {
	functionName := "cdl.toldatatest/TestService/GetTestA"

	if req == nil {
		return nil, errors.New("empty-request")
	}

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &TestAResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (service *TestServiceToldataClient) GetTestAB(ctx context.Context, req *TestARequest) (*TestAResponse, error) {
	functionName := "cdl.toldatatest/TestService/GetTestAB"

	if req == nil {
		return nil, errors.New("empty-request")
	}

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &TestAResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (service *TestServiceToldataClient) GetTestGetIP(ctx context.Context, req *Empty) (*TestGetIPResponse, error) {
	functionName := "cdl.toldatatest/TestService/GetTestGetIP"

	if req == nil {
		return nil, errors.New("empty-request")
	}

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &TestGetIPResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

type TestService_FeedDataToldataServer interface {
	Receive() (*FeedDataRequest, error)
	OnData(*FeedDataRequest) error
	Done(resp *FeedDataResponse) error

	GetResponse() (*FeedDataResponse, error)

	TriggerEOF()
	Error(err error)
	OnExit(func())
	Exit()
}

type TestService_FeedDataToldataServerImpl struct {
	request         chan *FeedDataRequest
	isRequestClosed bool

	response chan *FeedDataResponse

	cancel chan struct{}
	eof    chan struct{}
	err    chan error
	done   chan struct{}

	isEOF      bool
	isCanceled bool

	streamErr error

	Context context.Context
}

func CreateTestService_FeedDataToldataServerImpl(ctx context.Context) *TestService_FeedDataToldataServerImpl {
	t := &TestService_FeedDataToldataServerImpl{}

	t.Context = ctx
	t.request = make(chan *FeedDataRequest)
	t.response = make(chan *FeedDataResponse)
	t.cancel = make(chan struct{})
	t.eof = make(chan struct{})
	t.done = make(chan struct{})
	t.err = make(chan error)
	return t
}

func (impl *TestService_FeedDataToldataServerImpl) Exit() {
	close(impl.done)
}

func (impl *TestService_FeedDataToldataServerImpl) OnExit(fn func()) {
	go func() {
		select {
		case <-impl.done:
			fn()
		}
	}()
}

func (impl *TestService_FeedDataToldataServerImpl) TriggerEOF() {
	if impl.streamErr != nil {
		return
	}
	if impl.isEOF == false {
		close(impl.eof)
		impl.isEOF = true
	}
}

func (impl *TestService_FeedDataToldataServerImpl) Receive() (*FeedDataRequest, error) {

	if impl.streamErr != nil {
		return nil, impl.streamErr
	}
	if impl.isEOF {
		return nil, io.EOF
	}

	select {
	case data := <-impl.request:
		return data, impl.streamErr
	case <-impl.cancel:
		return nil, impl.streamErr
	case <-impl.eof:
		return nil, io.EOF
	case err := <-impl.err:

		return nil, err

	}
}

func (impl *TestService_FeedDataToldataServerImpl) OnData(req *FeedDataRequest) error {
	if impl.streamErr != nil {
		return impl.streamErr
	}

	select {
	case err := <-impl.err:
		return err
	case impl.request <- req:
		return nil
	}
}

func (impl *TestService_FeedDataToldataServerImpl) Done(resp *FeedDataResponse) error {
	if impl.streamErr != nil {
		return impl.streamErr
	}

	select {
	case impl.response <- resp:
		return nil
	case err := <-impl.err:
		return err

	}
}

func (impl *TestService_FeedDataToldataServerImpl) GetResponse() (*FeedDataResponse, error) {
	if impl.streamErr != nil {
		return nil, impl.streamErr
	}

	select {
	case err := <-impl.err:

		return nil, err

	case <-impl.cancel:

		return nil, errors.New("canceled")

	case response := <-impl.response:
		return response, nil

	}
}

func (impl *TestService_FeedDataToldataServerImpl) Cancel() {
	if impl.isCanceled == false {
		close(impl.cancel)
		impl.isCanceled = true
	}
}

func (impl *TestService_FeedDataToldataServerImpl) Error(err error) {
	impl.err <- err
	impl.streamErr = err
}

type TestServiceToldataClient_FeedData struct {
	Context context.Context
	Service *TestServiceToldataClient
	ID      string
}

func (client *TestServiceToldataClient_FeedData) Send(req *FeedDataRequest) error {
	functionName := "cdl.toldatatest/TestService/FeedData_Send_" + client.ID
	if req == nil {
		return errors.New("empty-request")
	}

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return errors.New(functionName + ":" + err.Error())
	}

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, reqRaw)
	if err != nil {
		return errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		return nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return errors.New(pErr.ErrorMessage)
		} else {
			return err
		}
	}
}

func (client *TestServiceToldataClient_FeedData) Done() (*FeedDataResponse, error) {
	functionName := "cdl.toldatatest/TestService/FeedData_Done_" + client.ID

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, nil)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &FeedDataResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (impl *TestService_FeedDataToldataServerImpl) Subscribe(service *TestServiceToldataServer, id string) error {
	bus := service.Bus
	var sub *nats.Subscription
	var subscriptions []*nats.Subscription
	var err error

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/FeedData_Send_"+id, "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input FeedDataRequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}

		err = impl.OnData(&input)

		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, zero)
		}

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/FeedData_Done_"+id, "cdl.toldatatest/TestService", func(m *nats.Msg) {

		defer impl.Exit()
		impl.TriggerEOF()
		result, err := impl.GetResponse()

		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		raw, err := proto.Marshal(result)
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

	})

	subscriptions = append(subscriptions, sub)

	impl.OnExit(func() {
		for i := range subscriptions {
			subscriptions[i].Unsubscribe()
		}
	})

	return err
}

func (service *TestServiceToldataClient) FeedData(ctx context.Context) (*TestServiceToldataClient_FeedData, error) {
	functionName := "cdl.toldatatest/TestService/FeedData"

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, nil)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error

		p := &toldata.StreamInfo{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return &TestServiceToldataClient_FeedData{
			ID:      p.ID,
			Context: ctx,
			Service: service,
		}, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

type TestService_StreamDataToldataServer interface {
	GetResponse() (*StreamDataResponse, error)

	Send(*StreamDataResponse) error

	TriggerEOF()
	Error(err error)
	OnExit(func())
	Exit()
}

type TestService_StreamDataToldataServerImpl struct {
	request         chan *StreamDataRequest
	isRequestClosed bool

	response chan *StreamDataResponse

	cancel chan struct{}
	eof    chan struct{}
	err    chan error
	done   chan struct{}

	isEOF      bool
	isCanceled bool

	streamErr error

	Context context.Context
}

func CreateTestService_StreamDataToldataServerImpl(ctx context.Context) *TestService_StreamDataToldataServerImpl {
	t := &TestService_StreamDataToldataServerImpl{}

	t.Context = ctx
	t.request = make(chan *StreamDataRequest)
	t.response = make(chan *StreamDataResponse)
	t.cancel = make(chan struct{})
	t.eof = make(chan struct{})
	t.done = make(chan struct{})
	t.err = make(chan error)
	return t
}

func (impl *TestService_StreamDataToldataServerImpl) Exit() {
	close(impl.done)
}

func (impl *TestService_StreamDataToldataServerImpl) OnExit(fn func()) {
	go func() {
		select {
		case <-impl.done:
			fn()
		}
	}()
}

func (impl *TestService_StreamDataToldataServerImpl) TriggerEOF() {
	if impl.streamErr != nil {
		return
	}
	if impl.isEOF == false {
		close(impl.eof)
		impl.isEOF = true
	}
}

func (impl *TestService_StreamDataToldataServerImpl) GetResponse() (*StreamDataResponse, error) {
	if impl.streamErr != nil {
		return nil, impl.streamErr
	}

	select {
	case err := <-impl.err:

		impl.Exit()

		return nil, err

	case <-impl.cancel:

		impl.Exit()

		return nil, errors.New("canceled")

	case response := <-impl.response:
		return response, nil

	case <-impl.eof:
		impl.Exit()
		return nil, io.EOF

	}
}

func (impl *TestService_StreamDataToldataServerImpl) Send(req *StreamDataResponse) error {

	if impl.isEOF {
		return io.EOF
	}

	if impl.streamErr != nil {
		return impl.streamErr
	}
	select {
	case impl.response <- req:
		return impl.streamErr

	case <-impl.cancel:
		return impl.streamErr
	case <-impl.eof:
		return io.EOF
	case err := <-impl.err:
		return err

	}
}

func (impl *TestService_StreamDataToldataServerImpl) Cancel() {
	if impl.isCanceled == false {
		close(impl.cancel)
		impl.isCanceled = true
	}
}

func (impl *TestService_StreamDataToldataServerImpl) Error(err error) {
	impl.err <- err
	impl.streamErr = err
}

type TestServiceToldataClient_StreamData struct {
	Context context.Context
	Service *TestServiceToldataClient
	ID      string
}

func (client *TestServiceToldataClient_StreamData) Receive() (*StreamDataResponse, error) {
	functionName := "cdl.toldatatest/TestService/StreamData_Receive_" + client.ID

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, nil)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &StreamDataResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (client *TestServiceToldataClient_StreamData) Done() (*StreamDataResponse, error) {
	functionName := "cdl.toldatatest/TestService/StreamData_Done_" + client.ID

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, nil)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &StreamDataResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (impl *TestService_StreamDataToldataServerImpl) Subscribe(service *TestServiceToldataServer, id string) error {
	bus := service.Bus
	var sub *nats.Subscription
	var subscriptions []*nats.Subscription
	var err error

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/StreamData_Receive_"+id, "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input StreamDataRequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}

		response, err := impl.GetResponse()
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}

		raw, err := proto.Marshal(response)
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

	})

	subscriptions = append(subscriptions, sub)

	impl.OnExit(func() {
		for i := range subscriptions {
			subscriptions[i].Unsubscribe()
		}
	})

	return err
}

func (service *TestServiceToldataClient) StreamData(ctx context.Context, req *StreamDataRequest) (*TestServiceToldataClient_StreamData, error) {
	functionName := "cdl.toldatatest/TestService/StreamData"
	if req == nil {
		return nil, errors.New("empty-request")
	}
	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error

		p := &toldata.StreamInfo{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return &TestServiceToldataClient_StreamData{
			ID:      p.ID,
			Context: ctx,
			Service: service,
		}, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

type TestService_StreamDataAlt1ToldataServer interface {
	GetResponse() (*StreamDataResponse, error)

	Send(*StreamDataResponse) error

	TriggerEOF()
	Error(err error)
	OnExit(func())
	Exit()
}

type TestService_StreamDataAlt1ToldataServerImpl struct {
	request         chan *StreamDataRequest
	isRequestClosed bool

	response chan *StreamDataResponse

	cancel chan struct{}
	eof    chan struct{}
	err    chan error
	done   chan struct{}

	isEOF      bool
	isCanceled bool

	streamErr error

	Context context.Context
}

func CreateTestService_StreamDataAlt1ToldataServerImpl(ctx context.Context) *TestService_StreamDataAlt1ToldataServerImpl {
	t := &TestService_StreamDataAlt1ToldataServerImpl{}

	t.Context = ctx
	t.request = make(chan *StreamDataRequest)
	t.response = make(chan *StreamDataResponse)
	t.cancel = make(chan struct{})
	t.eof = make(chan struct{})
	t.done = make(chan struct{})
	t.err = make(chan error)
	return t
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) Exit() {
	close(impl.done)
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) OnExit(fn func()) {
	go func() {
		select {
		case <-impl.done:
			fn()
		}
	}()
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) TriggerEOF() {
	if impl.streamErr != nil {
		return
	}
	if impl.isEOF == false {
		close(impl.eof)
		impl.isEOF = true
	}
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) GetResponse() (*StreamDataResponse, error) {
	if impl.streamErr != nil {
		return nil, impl.streamErr
	}

	select {
	case err := <-impl.err:

		impl.Exit()

		return nil, err

	case <-impl.cancel:

		impl.Exit()

		return nil, errors.New("canceled")

	case response := <-impl.response:
		return response, nil

	case <-impl.eof:
		impl.Exit()
		return nil, io.EOF

	}
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) Send(req *StreamDataResponse) error {

	if impl.isEOF {
		return io.EOF
	}

	if impl.streamErr != nil {
		return impl.streamErr
	}
	select {
	case impl.response <- req:
		return impl.streamErr

	case <-impl.cancel:
		return impl.streamErr
	case <-impl.eof:
		return io.EOF
	case err := <-impl.err:
		return err

	}
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) Cancel() {
	if impl.isCanceled == false {
		close(impl.cancel)
		impl.isCanceled = true
	}
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) Error(err error) {
	impl.err <- err
	impl.streamErr = err
}

type TestServiceToldataClient_StreamDataAlt1 struct {
	Context context.Context
	Service *TestServiceToldataClient
	ID      string
}

func (client *TestServiceToldataClient_StreamDataAlt1) Receive() (*StreamDataResponse, error) {
	functionName := "cdl.toldatatest/TestService/StreamDataAlt1_Receive_" + client.ID

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, nil)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &StreamDataResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (client *TestServiceToldataClient_StreamDataAlt1) Done() (*StreamDataResponse, error) {
	functionName := "cdl.toldatatest/TestService/StreamDataAlt1_Done_" + client.ID

	result, err := client.Service.Bus.Connection.RequestWithContext(client.Context, functionName, nil)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &StreamDataResponse{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (impl *TestService_StreamDataAlt1ToldataServerImpl) Subscribe(service *TestServiceToldataServer, id string) error {
	bus := service.Bus
	var sub *nats.Subscription
	var subscriptions []*nats.Subscription
	var err error

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/StreamDataAlt1_Receive_"+id, "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input StreamDataRequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}

		response, err := impl.GetResponse()
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}

		raw, err := proto.Marshal(response)
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

	})

	subscriptions = append(subscriptions, sub)

	impl.OnExit(func() {
		for i := range subscriptions {
			subscriptions[i].Unsubscribe()
		}
	})

	return err
}

func (service *TestServiceToldataClient) StreamDataAlt1(ctx context.Context, req *StreamDataRequest) (*TestServiceToldataClient_StreamDataAlt1, error) {
	functionName := "cdl.toldatatest/TestService/StreamDataAlt1"
	if req == nil {
		return nil, errors.New("empty-request")
	}
	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)

	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error

		p := &toldata.StreamInfo{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return &TestServiceToldataClient_StreamDataAlt1{
			ID:      p.ID,
			Context: ctx,
			Service: service,
		}, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (service *TestServiceToldataClient) TestEmpty(ctx context.Context, req *Empty) (*Empty, error) {
	functionName := "cdl.toldatatest/TestService/TestEmpty"

	if req == nil {
		return nil, errors.New("empty-request")
	}

	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, errors.New(functionName + ":" + err.Error())
	}

	if result.Data[0] == 0 {
		// 0 means no error
		p := &Empty{}
		err = proto.Unmarshal(result.Data[1:], p)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else {
		var pErr toldata.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}

func (service *TestServiceToldataServer) SubscribeTestService() (<-chan struct{}, error) {
	bus := service.Bus

	var err error
	var sub *nats.Subscription
	var subscriptions []*nats.Subscription

	done := make(chan struct{})

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/GetTestA", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input TestARequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.GetTestA(bus.Context, &input)

		if m.Reply != "" {
			if err != nil {
				bus.HandleError(m.Reply, err)
			} else {
				raw, err := proto.Marshal(result)
				if err != nil {
					bus.HandleError(m.Reply, err)
				} else {
					zero := []byte{0}
					bus.Connection.Publish(m.Reply, append(zero, raw...))
				}
			}
		}

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/GetTestAB", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input TestARequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.GetTestAB(bus.Context, &input)

		if m.Reply != "" {
			if err != nil {
				bus.HandleError(m.Reply, err)
			} else {
				raw, err := proto.Marshal(result)
				if err != nil {
					bus.HandleError(m.Reply, err)
				} else {
					zero := []byte{0}
					bus.Connection.Publish(m.Reply, append(zero, raw...))
				}
			}
		}

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/GetTestGetIP", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input Empty
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.GetTestGetIP(bus.Context, &input)

		if m.Reply != "" {
			if err != nil {
				bus.HandleError(m.Reply, err)
			} else {
				raw, err := proto.Marshal(result)
				if err != nil {
					bus.HandleError(m.Reply, err)
				} else {
					zero := []byte{0}
					bus.Connection.Publish(m.Reply, append(zero, raw...))
				}
			}
		}

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/FeedData", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		stream := CreateTestService_FeedDataToldataServerImpl(bus.Context)

		stream.Subscribe(service, m.Reply)

		raw, err := proto.Marshal(&toldata.StreamInfo{
			ID: m.Reply,
		})
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

		service.Service.FeedData(stream)

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/StreamData", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		stream := CreateTestService_StreamDataToldataServerImpl(bus.Context)

		stream.Subscribe(service, m.Reply)

		raw, err := proto.Marshal(&toldata.StreamInfo{
			ID: m.Reply,
		})
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

		var input StreamDataRequest
		err = proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		err = service.Service.StreamData(&input, stream)
		if err != nil {
			stream.Error(err)
			bus.HandleError(m.Reply, err)
			return
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, zero)
		}
		stream.TriggerEOF()

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/StreamDataAlt1", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		stream := CreateTestService_StreamDataAlt1ToldataServerImpl(bus.Context)

		stream.Subscribe(service, m.Reply)

		raw, err := proto.Marshal(&toldata.StreamInfo{
			ID: m.Reply,
		})
		if err != nil {
			bus.HandleError(m.Reply, err)
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, append(zero, raw...))
		}

		var input StreamDataRequest
		err = proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		err = service.Service.StreamDataAlt1(&input, stream)
		if err != nil {
			stream.Error(err)
			bus.HandleError(m.Reply, err)
			return
		} else {
			zero := []byte{0}
			bus.Connection.Publish(m.Reply, zero)
		}
		stream.TriggerEOF()

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/TestEmpty", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input Empty
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.TestEmpty(bus.Context, &input)

		if m.Reply != "" {
			if err != nil {
				bus.HandleError(m.Reply, err)
			} else {
				raw, err := proto.Marshal(result)
				if err != nil {
					bus.HandleError(m.Reply, err)
				} else {
					zero := []byte{0}
					bus.Connection.Publish(m.Reply, append(zero, raw...))
				}
			}
		}

	})

	subscriptions = append(subscriptions, sub)

	sub, err = bus.Connection.QueueSubscribe("cdl.toldatatest/TestService/ToldataHealthCheck", "cdl.toldatatest/TestService", func(m *nats.Msg) {
		var input toldata.Empty
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.ToldataHealthCheck(bus.Context, &input)

		if m.Reply != "" {
			if err != nil {
				bus.HandleError(m.Reply, err)
			} else {
				raw, err := proto.Marshal(result)
				if err != nil {
					bus.HandleError(m.Reply, err)
				} else {
					zero := []byte{0}
					bus.Connection.Publish(m.Reply, append(zero, raw...))
				}
			}
		}

	})

	subscriptions = append(subscriptions, sub)

	go func() {
		defer close(done)

		select {
		case <-bus.Context.Done():
			for i := range subscriptions {
				subscriptions[i].Unsubscribe()
			}
		}
	}()

	return done, err
}
