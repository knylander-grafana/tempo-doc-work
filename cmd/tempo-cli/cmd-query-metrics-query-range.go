package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/grafana/dskit/user"
	"github.com/grafana/tempo/pkg/api"
	"github.com/grafana/tempo/pkg/tempopb"
)

type metricsQueryCmd struct {
	HostPort string `arg:"" help:"tempo host and port. scheme and path will be provided based on query type. e.g. localhost:3200"`
	TraceQL  string `arg:"" optional:"" help:"traceql query"`
	Start    string `arg:"" optional:"" help:"start time in ISO8601 format"`
	End      string `arg:"" optional:"" help:"end time in ISO8601 format"`

	OrgID      string `help:"optional orgID"`
	UseGRPC    bool   `help:"stream search results over GRPC"`
	Instant    bool   `help:"perform an instant query instead of a range query"`
	PathPrefix string `help:"string to prefix all http paths with"`
}

func (cmd *metricsQueryCmd) Run(_ *globalOptions) error {
	startDate, err := time.Parse(time.RFC3339, cmd.Start)
	if err != nil {
		return err
	}
	start := startDate.UnixNano()

	endDate, err := time.Parse(time.RFC3339, cmd.End)
	if err != nil {
		return err
	}
	end := endDate.UnixNano()

	if cmd.Instant {
		req := &tempopb.QueryInstantRequest{
			Query: cmd.TraceQL,
			Start: uint64(start),
			End:   uint64(end),
		}

		if cmd.UseGRPC {
			return cmd.queryInstantGRPC(req)
		}

		return cmd.queryInstantHTTP(req)
	}

	req := &tempopb.QueryRangeRequest{
		Query: cmd.TraceQL,
		Start: uint64(start),
		End:   uint64(end),
	}

	if cmd.UseGRPC {
		return cmd.queryRangeGRPC(req)
	}

	return cmd.queryRangeHTTP(req)
}

func (cmd *metricsQueryCmd) queryRangeGRPC(req *tempopb.QueryRangeRequest) error {
	ctx := user.InjectOrgID(context.Background(), cmd.OrgID)
	ctx, err := user.InjectIntoGRPCRequest(ctx)
	if err != nil {
		return err
	}

	clientConn, err := grpc.NewClient(cmd.HostPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := tempopb.NewStreamingQuerierClient(clientConn)

	resp, err := client.MetricsQueryRange(ctx, req)
	if err != nil {
		return err
	}

	for {
		searchResp, err := resp.Recv()
		if searchResp != nil {
			err = printAsJSON(searchResp)
			if err != nil {
				return err
			}
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// nolint: goconst // goconst wants us to make http:// a const
func (cmd *metricsQueryCmd) queryRangeHTTP(req *tempopb.QueryRangeRequest) error {
	httpReq, err := http.NewRequest("GET", "http://"+path.Join(cmd.HostPort, cmd.PathPrefix, api.PathMetricsQueryRange), nil)
	if err != nil {
		return err
	}

	httpReq = api.BuildQueryRangeRequest(httpReq, req, "")
	httpReq.Header = http.Header{}
	err = user.InjectOrgIDIntoHTTPRequest(user.InjectOrgID(context.Background(), cmd.OrgID), httpReq)
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("failed to query. body: " + string(body) + " status: " + httpResp.Status)
	}

	resp := &tempopb.QueryRangeResponse{}
	err = jsonpb.Unmarshal(bytes.NewReader(body), resp)
	if err != nil {
		panic("failed to parse resp: " + err.Error())
	}
	err = printAsJSON(resp)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *metricsQueryCmd) queryInstantGRPC(req *tempopb.QueryInstantRequest) error {
	ctx := user.InjectOrgID(context.Background(), cmd.OrgID)
	ctx, err := user.InjectIntoGRPCRequest(ctx)
	if err != nil {
		return err
	}

	clientConn, err := grpc.NewClient(cmd.HostPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := tempopb.NewStreamingQuerierClient(clientConn)

	resp, err := client.MetricsQueryInstant(ctx, req)
	if err != nil {
		return err
	}

	for {
		searchResp, err := resp.Recv()
		if searchResp != nil {
			err = printAsJSON(searchResp)
			if err != nil {
				return err
			}
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

// nolint: goconst // goconst wants us to make http:// a const
func (cmd *metricsQueryCmd) queryInstantHTTP(req *tempopb.QueryInstantRequest) error {
	httpReq, err := http.NewRequest("GET", "http://"+path.Join(cmd.HostPort, cmd.PathPrefix, api.PathMetricsQueryInstant), nil)
	if err != nil {
		return err
	}

	httpReq = api.BuildQueryInstantRequest(httpReq, req)
	httpReq.Header = http.Header{}
	err = user.InjectOrgIDIntoHTTPRequest(user.InjectOrgID(context.Background(), cmd.OrgID), httpReq)
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("failed to query. body: " + string(body) + " status: " + httpResp.Status)
	}

	resp := &tempopb.QueryInstantResponse{}
	err = jsonpb.Unmarshal(bytes.NewReader(body), resp)
	if err != nil {
		panic("failed to parse resp: " + err.Error())
	}
	err = printAsJSON(resp)
	if err != nil {
		return err
	}

	return nil
}
