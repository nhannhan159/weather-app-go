package common

import (
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
	zipkinOpentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func InitializeTracer(config GlobalConfig) {
	reporter := zipkinHttp.NewReporter(config.Tracing.ZipkinUrl)
	defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint(config.AppName, fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	tracer := zipkinOpentracing.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(tracer)
}
