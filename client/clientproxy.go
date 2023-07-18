/**
* Generated by go-doudou v2.1.5.
* You can edit it as your need.
 */
package client

import (
	"context"
	"time"

	"github.com/VINDA-98/doudou-test/dto"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/slok/goresilience"
	"github.com/slok/goresilience/circuitbreaker"
	rerrors "github.com/slok/goresilience/errors"
	"github.com/slok/goresilience/metrics"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
)

type DoudouTestClientProxy struct {
	client *DoudouTestClient
	logger zerolog.Logger
	runner goresilience.Runner
}

func (receiver *DoudouTestClientProxy) PostUser(ctx context.Context, _headers map[string]string, user dto.GddUser, options Options) (_resp *resty.Response, data int32, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.PostUser(
			ctx,
			_headers,
			user,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call PostUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call PostUser fail")
	}
	return
}
func (receiver *DoudouTestClientProxy) GetUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.GddUser, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetUser_Id fail")
	}
	return
}
func (receiver *DoudouTestClientProxy) PutUser(ctx context.Context, _headers map[string]string, user dto.GddUser, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.PutUser(
			ctx,
			_headers,
			user,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call PutUser fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call PutUser fail")
	}
	return
}
func (receiver *DoudouTestClientProxy) DeleteUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, re = receiver.client.DeleteUser_Id(
			ctx,
			_headers,
			id,
			options,
		)
		if re != nil {
			return errors.Wrap(re, "call DeleteUser_Id fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		re = errors.Wrap(_err, "call DeleteUser_Id fail")
	}
	return
}
func (receiver *DoudouTestClientProxy) GetUsers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error) {
	if _err := receiver.runner.Run(ctx, func(ctx context.Context) error {
		_resp, data, err = receiver.client.GetUsers(
			ctx,
			_headers,
			parameter,
			options,
		)
		if err != nil {
			return errors.Wrap(err, "call GetUsers fail")
		}
		return nil
	}); _err != nil {
		// you can implement your fallback logic here
		if errors.Is(_err, rerrors.ErrCircuitOpen) {
			receiver.logger.Error().Err(_err).Msg("")
		}
		err = errors.Wrap(_err, "call GetUsers fail")
	}
	return
}

type ProxyOption func(*DoudouTestClientProxy)

func WithRunner(runner goresilience.Runner) ProxyOption {
	return func(proxy *DoudouTestClientProxy) {
		proxy.runner = runner
	}
}

func WithLogger(logger zerolog.Logger) ProxyOption {
	return func(proxy *DoudouTestClientProxy) {
		proxy.logger = logger
	}
}

func NewDoudouTestClientProxy(client *DoudouTestClient, opts ...ProxyOption) *DoudouTestClientProxy {
	cp := &DoudouTestClientProxy{
		client: client,
		logger: zlogger.Logger,
	}

	for _, opt := range opts {
		opt(cp)
	}

	if cp.runner == nil {
		var mid []goresilience.Middleware
		mid = append(mid, metrics.NewMiddleware("github.com/VINDA-98/doudou-test_client", metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)))
		mid = append(mid, circuitbreaker.NewMiddleware(circuitbreaker.Config{
			ErrorPercentThresholdToOpen:        50,
			MinimumRequestToOpen:               6,
			SuccessfulRequiredOnHalfOpen:       1,
			WaitDurationInOpenState:            5 * time.Second,
			MetricsSlidingWindowBucketQuantity: 10,
			MetricsBucketDuration:              1 * time.Second,
		}),
			timeout.NewMiddleware(timeout.Config{
				Timeout: 3 * time.Minute,
			}),
			retry.NewMiddleware(retry.Config{
				Times: 3,
			}))

		cp.runner = goresilience.RunnerChain(mid...)
	}

	return cp
}
