// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: transactions.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_TransactionsAPI_ListTransactions_0(ctx context.Context, marshaler runtime.Marshaler, client TransactionsAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListTransactionsRequest
	var metadata runtime.ServerMetadata

	msg, err := client.ListTransactions(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TransactionsAPI_ListTransactions_0(ctx context.Context, marshaler runtime.Marshaler, server TransactionsAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListTransactionsRequest
	var metadata runtime.ServerMetadata

	msg, err := server.ListTransactions(ctx, &protoReq)
	return msg, metadata, err

}

func request_TransactionsAPI_GetTransaction_0(ctx context.Context, marshaler runtime.Marshaler, client TransactionsAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTransactionRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.GetTransaction(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TransactionsAPI_GetTransaction_0(ctx context.Context, marshaler runtime.Marshaler, server TransactionsAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTransactionRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.GetTransaction(ctx, &protoReq)
	return msg, metadata, err

}

func request_TransactionsAPI_AddTransaction_0(ctx context.Context, marshaler runtime.Marshaler, client TransactionsAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddTransactionRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AddTransaction(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TransactionsAPI_AddTransaction_0(ctx context.Context, marshaler runtime.Marshaler, server TransactionsAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddTransactionRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AddTransaction(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterTransactionsAPIHandlerServer registers the http handlers for service TransactionsAPI to "mux".
// UnaryRPC     :call TransactionsAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTransactionsAPIHandlerFromEndpoint instead.
func RegisterTransactionsAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TransactionsAPIServer) error {

	mux.Handle("GET", pattern_TransactionsAPI_ListTransactions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/pb.TransactionsAPI/ListTransactions", runtime.WithHTTPPathPattern("/transactions/list_transactions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TransactionsAPI_ListTransactions_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_ListTransactions_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_TransactionsAPI_GetTransaction_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/pb.TransactionsAPI/GetTransaction", runtime.WithHTTPPathPattern("/transactions/get_transaction/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TransactionsAPI_GetTransaction_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_GetTransaction_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_TransactionsAPI_AddTransaction_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/pb.TransactionsAPI/AddTransaction", runtime.WithHTTPPathPattern("/transactions/add_transaction"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TransactionsAPI_AddTransaction_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_AddTransaction_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTransactionsAPIHandlerFromEndpoint is same as RegisterTransactionsAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTransactionsAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTransactionsAPIHandler(ctx, mux, conn)
}

// RegisterTransactionsAPIHandler registers the http handlers for service TransactionsAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTransactionsAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTransactionsAPIHandlerClient(ctx, mux, NewTransactionsAPIClient(conn))
}

// RegisterTransactionsAPIHandlerClient registers the http handlers for service TransactionsAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TransactionsAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TransactionsAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TransactionsAPIClient" to call the correct interceptors.
func RegisterTransactionsAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TransactionsAPIClient) error {

	mux.Handle("GET", pattern_TransactionsAPI_ListTransactions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/pb.TransactionsAPI/ListTransactions", runtime.WithHTTPPathPattern("/transactions/list_transactions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TransactionsAPI_ListTransactions_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_ListTransactions_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_TransactionsAPI_GetTransaction_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/pb.TransactionsAPI/GetTransaction", runtime.WithHTTPPathPattern("/transactions/get_transaction/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TransactionsAPI_GetTransaction_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_GetTransaction_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_TransactionsAPI_AddTransaction_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/pb.TransactionsAPI/AddTransaction", runtime.WithHTTPPathPattern("/transactions/add_transaction"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TransactionsAPI_AddTransaction_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TransactionsAPI_AddTransaction_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TransactionsAPI_ListTransactions_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"transactions", "list_transactions"}, ""))

	pattern_TransactionsAPI_GetTransaction_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"transactions", "get_transaction", "id"}, ""))

	pattern_TransactionsAPI_AddTransaction_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"transactions", "add_transaction"}, ""))
)

var (
	forward_TransactionsAPI_ListTransactions_0 = runtime.ForwardResponseMessage

	forward_TransactionsAPI_GetTransaction_0 = runtime.ForwardResponseMessage

	forward_TransactionsAPI_AddTransaction_0 = runtime.ForwardResponseMessage
)
