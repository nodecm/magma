/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers_test

import (
	"errors"
	"flag"
	"strconv"
	"testing"
	"time"

	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/magmad"
	magmad_protos "magma/orc8r/cloud/go/services/magmad/protos"
	magmad_test_init "magma/orc8r/cloud/go/services/magmad/test_init"
	"magma/orc8r/cloud/go/services/metricsd/exporters"
	"magma/orc8r/cloud/go/services/metricsd/servicers"

	"github.com/golang/glog"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type testMetricExporter struct {
	queue []exporters.Sample

	// error to return
	retErr error
}

const (
	MetricName = protos.MetricName_process_virtual_memory_bytes
	LabelName  = protos.MetricLabelName_result
	LabelValue = "success"
)

// Set verbosity so we can capture exporter error logging
var _ = flag.Set("vmodule", "*=2")

func (e *testMetricExporter) Submit(metrics []exporters.MetricAndContext) error {
	for _, metricAndContext := range metrics {
		family, ctx := metricAndContext.Family, metricAndContext.Context
		for _, metric := range family.GetMetric() {
			e.queue = append(
				e.queue,
				exporters.GetSamplesForMetrics(ctx.DecodedName, family.GetType(), metric, ctx.OriginatingEntity)...,
			)
		}
	}

	return e.retErr
}

func (e *testMetricExporter) Start() {}

func TestCollect(t *testing.T) {
	magmad_test_init.StartTestService(t)

	e := &testMetricExporter{}
	ctx := context.Background()
	srv := servicers.NewMetricsControllerServer()
	srv.RegisterExporter(e)

	// Create test network
	testNetworkId, err := magmad.RegisterNetwork(
		&magmad_protos.MagmadNetworkRecord{Name: "Test Network Name"},
		"metricsd_servicer_test_network")
	if err != nil {
		t.Fatalf("Magmad Register Network '%s' Error: %s", testNetworkId, err)
	}

	// Register a fake gateway
	gatewayId := "2876171d-bf38-4254-b4da-71a713952904"
	hwId := protos.AccessGatewayID{Id: gatewayId}
	logicalId, err := magmad.RegisterGateway(testNetworkId,
		&magmad_protos.AccessGatewayRecord{HwId: &hwId, Name: "bla"})
	if err != nil || logicalId == "" {
		t.Fatalf("Magmad Register Error: %s, logical ID: %#v", err, logicalId)
	}

	name := strconv.Itoa(int(MetricName))
	key := strconv.Itoa(int(LabelName))
	value := LabelValue
	float := 1.0
	int_val := uint64(1)
	counter_type := dto.MetricType_COUNTER
	counters := protos.MetricsContainer{
		GatewayId: gatewayId,
		Family: []*dto.MetricFamily{{
			Type: &counter_type,
			Name: &name,
			Metric: []*dto.Metric{
				{
					Label:   []*dto.LabelPair{{Name: &key, Value: &value}},
					Counter: &dto.Counter{Value: &float}}}}}}

	gauge_type := dto.MetricType_GAUGE
	gauges := protos.MetricsContainer{
		GatewayId: gatewayId,
		Family: []*dto.MetricFamily{{
			Type: &gauge_type,
			Name: &name,
			Metric: []*dto.Metric{
				{
					Label: []*dto.LabelPair{{Name: &key, Value: &value}},
					Gauge: &dto.Gauge{Value: &float}}}}}}

	summary_type := dto.MetricType_SUMMARY
	summaries := protos.MetricsContainer{
		GatewayId: gatewayId,
		Family: []*dto.MetricFamily{{
			Type: &summary_type,
			Name: &name,
			Metric: []*dto.Metric{
				{
					Label:   []*dto.LabelPair{{Name: &key, Value: &value}},
					Summary: &dto.Summary{SampleSum: &float, SampleCount: &int_val}}}}}}

	histogram_type := dto.MetricType_HISTOGRAM
	histograms := protos.MetricsContainer{
		GatewayId: gatewayId,
		Family: []*dto.MetricFamily{{
			Type: &histogram_type,
			Name: &name,
			Metric: []*dto.Metric{
				{
					Label: []*dto.LabelPair{{Name: &key, Value: &value}},
					Histogram: &dto.Histogram{SampleSum: &float, SampleCount: &int_val,
						Bucket: []*dto.Bucket{
							{CumulativeCount: &int_val,
								UpperBound: &float}}}}}}}}

	// Collect counters
	_, err = srv.Collect(ctx, &counters)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(e.queue))
	assert.Equal(t, strconv.FormatFloat(float, 'f', -1, 64), e.queue[0].Value())
	// clear queue
	e.queue = e.queue[:0]

	// Collect gauges
	_, err = srv.Collect(ctx, &gauges)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(e.queue))
	assert.Equal(t, strconv.FormatFloat(float, 'f', -1, 64), e.queue[0].Value())
	e.queue = e.queue[:0]

	// Collect summaries
	_, err = srv.Collect(ctx, &summaries)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(e.queue))
	assert.Equal(t, strconv.FormatUint(int_val, 10), e.queue[0].Value())
	assert.Equal(t, strconv.FormatFloat(float, 'f', -1, 64), e.queue[0].Value())
	e.queue = e.queue[:0]

	// Collect histograms
	_, err = srv.Collect(ctx, &histograms)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(e.queue))
	assert.Equal(t, strconv.FormatUint(int_val, 10), e.queue[0].Value())
	assert.Equal(t, strconv.FormatFloat(float, 'E', -1, 64), e.queue[1].Value())
	assert.Equal(t, strconv.FormatFloat(float, 'E', -1, 64), e.queue[2].Value())
	assert.Equal(t, strconv.FormatUint(int_val, 10), e.queue[3].Value())
	e.queue = e.queue[:0]

	// Test Collect with empty collection
	_, err = srv.Collect(ctx, &protos.MetricsContainer{})
	assert.NoError(t, err)

	// Exporter error should not result in error returned from Collect
	// But with verbosity set to 2 at the top of the test, we should log
	prevInfoLines := glog.Stats.Info.Lines()
	e.retErr = errors.New("mock exporter error")
	_, err = srv.Collect(ctx, &gauges)
	assert.NoError(t, err)
	assert.Equal(t, prevInfoLines+1, glog.Stats.Info.Lines())
}

func TestConsume(t *testing.T) {
	metricsChan := make(chan *dto.MetricFamily)
	e := &testMetricExporter{}

	srv := servicers.NewMetricsControllerServer()
	srv.RegisterExporter(e)

	go srv.ConsumeCloudMetrics(metricsChan, "Host_name_place_holder")
	fam1 := "test1"
	fam2 := "test2"
	go func() {
		metricsChan <- &dto.MetricFamily{Name: &fam1, Metric: []*dto.Metric{{}}}
		metricsChan <- &dto.MetricFamily{Name: &fam2, Metric: []*dto.Metric{{}}}
	}()
	time.Sleep(time.Second)
	assert.Equal(t, 2, len(e.queue))
}
