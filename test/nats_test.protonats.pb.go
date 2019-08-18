// Code generated by github.com/citradigital/protonats. DO NOT EDIT.
// package: cdl.protonats
// source: nats_test.proto
package test
import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/citradigital/protonats"
	nats "github.com/nats-io/go-nats"
)



type TestServiceInterface interface {
	GetTestA(ctx context.Context, req *TestARequest) (*TestAResponse, error)

}



type TestServiceProtonatsClient struct {
	Bus *protonats.Bus
}

type TestServiceProtonatsServer struct {
	Bus *protonats.Bus
	Service TestServiceInterface
}

func NewTestServiceProtonatsClient(bus *protonats.Bus) * TestServiceProtonatsClient {
	s := &TestServiceProtonatsClient{ Bus: bus }
	return s
}

func NewTestServiceProtonatsServer(bus *protonats.Bus, service TestServiceInterface) * TestServiceProtonatsServer {
	s := &TestServiceProtonatsServer{ Bus: bus, Service: service }
	return s
}


	
func (service *TestServiceProtonatsClient) GetTestA(ctx context.Context, req *TestARequest) (*TestAResponse, error) {
	functionName := "cdl.protonats/TestService/GetTestA"
	
	reqRaw, err := proto.Marshal(req)

	result, err := service.Bus.Connection.RequestWithContext(ctx, functionName, reqRaw)
	if err != nil {
		return nil, err
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
		var pErr protonats.ErrorMessage
		err = proto.Unmarshal(result.Data[1:], &pErr)
		if err == nil {
			return nil, errors.New(pErr.ErrorMessage)
		} else {
			return nil, err
		}
	}
}







func (service *TestServiceProtonatsServer) SubscribeTestService() (<-chan struct{}, error) {
	bus := service.Bus
	
	var err error
	var sub *nats.Subscription
	var subscriptions []*nats.Subscription
	
	done := make(chan struct{})
	
		

	sub, err = bus.Connection.QueueSubscribe("cdl.protonats/TestService/GetTestA", "TestService", func(m *nats.Msg) {
		var input TestARequest
		err := proto.Unmarshal(m.Data, &input)
		if err != nil {
			bus.HandleError(m.Reply, err)
			return
		}
		result, err := service.Service.GetTestA(bus.Context, &input)

		if m.Reply != ""  {
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


