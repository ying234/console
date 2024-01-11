// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha1/kafka_connect.proto

package dataplanev1alpha1connect

import (
	context "context"
	fmt "fmt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha1"
)

// KafkaConnectServiceGatewayServer implements the gRPC server API for the KafkaConnectService
// service.
type KafkaConnectServiceGatewayServer struct {
	v1alpha1.UnimplementedKafkaConnectServiceServer
	listConnectClusters connect_gateway.UnaryHandler[v1alpha1.ListConnectClustersRequest, v1alpha1.ListConnectClustersResponse]
	getConnectCluster   connect_gateway.UnaryHandler[v1alpha1.GetConnectClusterRequest, v1alpha1.GetConnectClusterResponse]
	listConnectors      connect_gateway.UnaryHandler[v1alpha1.ListConnectorsRequest, v1alpha1.ListConnectorsResponse]
	createConnector     connect_gateway.UnaryHandler[v1alpha1.CreateConnectorRequest, v1alpha1.CreateConnectorResponse]
	restartConnector    connect_gateway.UnaryHandler[v1alpha1.RestartConnectorRequest, emptypb.Empty]
	getConnector        connect_gateway.UnaryHandler[v1alpha1.GetConnectorRequest, v1alpha1.GetConnectorResponse]
	pauseConnector      connect_gateway.UnaryHandler[v1alpha1.PauseConnectorRequest, emptypb.Empty]
	resumeConnector     connect_gateway.UnaryHandler[v1alpha1.ResumeConnectorRequest, emptypb.Empty]
	deleteConnector     connect_gateway.UnaryHandler[v1alpha1.DeleteConnectorRequest, emptypb.Empty]
	upsertConnector     connect_gateway.UnaryHandler[v1alpha1.UpsertConnectorRequest, v1alpha1.UpsertConnectorResponse]
	getConnectorConfig  connect_gateway.UnaryHandler[v1alpha1.GetConnectorConfigRequest, v1alpha1.GetConnectorConfigResponse]
}

// NewKafkaConnectServiceGatewayServer constructs a Connect-Gateway gRPC server for the
// KafkaConnectService service.
func NewKafkaConnectServiceGatewayServer(svc KafkaConnectServiceHandler, opts ...connect_gateway.HandlerOption) *KafkaConnectServiceGatewayServer {
	return &KafkaConnectServiceGatewayServer{
		listConnectClusters: connect_gateway.NewUnaryHandler(KafkaConnectServiceListConnectClustersProcedure, svc.ListConnectClusters, opts...),
		getConnectCluster:   connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectClusterProcedure, svc.GetConnectCluster, opts...),
		listConnectors:      connect_gateway.NewUnaryHandler(KafkaConnectServiceListConnectorsProcedure, svc.ListConnectors, opts...),
		createConnector:     connect_gateway.NewUnaryHandler(KafkaConnectServiceCreateConnectorProcedure, svc.CreateConnector, opts...),
		restartConnector:    connect_gateway.NewUnaryHandler(KafkaConnectServiceRestartConnectorProcedure, svc.RestartConnector, opts...),
		getConnector:        connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectorProcedure, svc.GetConnector, opts...),
		pauseConnector:      connect_gateway.NewUnaryHandler(KafkaConnectServicePauseConnectorProcedure, svc.PauseConnector, opts...),
		resumeConnector:     connect_gateway.NewUnaryHandler(KafkaConnectServiceResumeConnectorProcedure, svc.ResumeConnector, opts...),
		deleteConnector:     connect_gateway.NewUnaryHandler(KafkaConnectServiceDeleteConnectorProcedure, svc.DeleteConnector, opts...),
		upsertConnector:     connect_gateway.NewUnaryHandler(KafkaConnectServiceUpsertConnectorProcedure, svc.UpsertConnector, opts...),
		getConnectorConfig:  connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectorConfigProcedure, svc.GetConnectorConfig, opts...),
	}
}

func (s *KafkaConnectServiceGatewayServer) ListConnectClusters(ctx context.Context, req *v1alpha1.ListConnectClustersRequest) (*v1alpha1.ListConnectClustersResponse, error) {
	return s.listConnectClusters(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnectCluster(ctx context.Context, req *v1alpha1.GetConnectClusterRequest) (*v1alpha1.GetConnectClusterResponse, error) {
	return s.getConnectCluster(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ListConnectors(ctx context.Context, req *v1alpha1.ListConnectorsRequest) (*v1alpha1.ListConnectorsResponse, error) {
	return s.listConnectors(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) CreateConnector(ctx context.Context, req *v1alpha1.CreateConnectorRequest) (*v1alpha1.CreateConnectorResponse, error) {
	return s.createConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) RestartConnector(ctx context.Context, req *v1alpha1.RestartConnectorRequest) (*emptypb.Empty, error) {
	return s.restartConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnector(ctx context.Context, req *v1alpha1.GetConnectorRequest) (*v1alpha1.GetConnectorResponse, error) {
	return s.getConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) PauseConnector(ctx context.Context, req *v1alpha1.PauseConnectorRequest) (*emptypb.Empty, error) {
	return s.pauseConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ResumeConnector(ctx context.Context, req *v1alpha1.ResumeConnectorRequest) (*emptypb.Empty, error) {
	return s.resumeConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) DeleteConnector(ctx context.Context, req *v1alpha1.DeleteConnectorRequest) (*emptypb.Empty, error) {
	return s.deleteConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) UpsertConnector(ctx context.Context, req *v1alpha1.UpsertConnectorRequest) (*v1alpha1.UpsertConnectorResponse, error) {
	return s.upsertConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnectorConfig(ctx context.Context, req *v1alpha1.GetConnectorConfigRequest) (*v1alpha1.GetConnectorConfigResponse, error) {
	return s.getConnectorConfig(ctx, req)
}

// RegisterKafkaConnectServiceHandlerGatewayServer registers the Connect handlers for the
// KafkaConnectService "svc" to "mux".
func RegisterKafkaConnectServiceHandlerGatewayServer(mux *runtime.ServeMux, svc KafkaConnectServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha1.RegisterKafkaConnectServiceHandlerServer(context.TODO(), mux, NewKafkaConnectServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}