[TOC]

```
注释:
·go-micro 消息说明
```

<!-- TOC -->

- [1. 说明](#1-%E8%AF%B4%E6%98%8E)
- [2. 代码example](#2-%E4%BB%A3%E7%A0%81example)

<!-- /TOC -->

# 1. 说明
+ 订阅主题，消息的结构体内容都是一一对应的这个例子(PubEvent = api.ServiceName +"_event")就是一个 spservice__event 主题的消息，消息的结构体为 protocol/spservice/pub/proto/Event
+ 代码example(使用nats、并使用zipkin追踪记录)
+ nats 部署参照 https://nats.io/documentation/tutorials/gnatsd-install/

# 2. 代码example
+ Publish
```go
        #github.com/pborman/uuid
		service := micro.NewService(
			micro.Name(constant.Servicename),
			micro.Registry(consul.NewRegistry(registry.Addrs(config.CConsulAddr()))),
			micro.WrapHandler(trace.NewHandlerWrapper(trace.GetZipkinTracer())),
			micro.Broker(nats.NewBroker(broker.Addrs(config.CNatsAddr()))),
			micro.WrapSubscriber(trace.NewSubscriberWrapper(trace.GetZipkinTracer())),
		)

		constant.ServicePublisher = micro.NewPublisher(constant.Servicename, service.Client())
		if err := constant.ServicePublisher.Publish(context.Background(), &natsproto.Event{
					Id:        uuid.NewUUID().String(),
					Timestamp: time.Now().Unix(),
					Message:   fmt.Sprintf("Messaging you all day on %s and index:%s", constant.Servicename,strconv.Itoa(index)),
		}); err != nil {
					fmt.Println("publish message error")
        }
```

+ Sub
 + 非消息队列(主题发布的一个消息会传送到达每个节点)

```go
            #github.com/micro/go-micro/server
            service := micro.NewService(
				micro.Name(constant.Servicename),
				micro.Registry(consul.NewRegistry(registry.Addrs(config.CConsulAddr()))),
				micro.WrapHandler(trace.NewHandlerWrapper(trace.GetZipkinTracer())),
				micro.Broker(nats.NewBroker(broker.Addrs(config.CNatsAddr()))),
				micro.WrapSubscriber(trace.NewSubscriberWrapper(trace.GetZipkinTracer())),
			)
			micro.RegisterSubscriber(constant.Servicename, service.Server(), subEv,server)
			func subEv(ctx context.Context, event *natsproto.Event) error {
				md, _ := metadata.FromContext(ctx)
				fmt.Println("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
				return nil
			}
```


  + 消息队列(主题发布的每个消息只会到达一个节点)

```go
            #github.com/micro/go-micro/server
            service := micro.NewService(
				micro.Name(constant.Servicename),
				micro.Registry(consul.NewRegistry(registry.Addrs(config.CConsulAddr()))),
				micro.WrapHandler(trace.NewHandlerWrapper(trace.GetZipkinTracer())),
				micro.Broker(nats.NewBroker(broker.Addrs(config.CNatsAddr()))),
				micro.WrapSubscriber(trace.NewSubscriberWrapper(trace.GetZipkinTracer())),
			)
			micro.RegisterSubscriber(constant.Servicename, service.Server(), subEv,server.SubscriberQueue("queue.pubsub"))
			func subEv(ctx context.Context, event *natsproto.Event) error {
				md, _ := metadata.FromContext(ctx)
				fmt.Println("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
				return nil
```
