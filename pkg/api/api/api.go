package api

import (
	"context"
	"gin_session_auth/pkg/api/api/metric"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/client-go/kubernetes"
)

type ClientManager struct {
	KubeClient *kubernetes.Clientset
}

func NewClientManager() *ClientManager {
	result := &ClientManager{}

	return result
}

func GetMultiMetric(ip string) (*metric.MultiMetric, error) {
	host := ip + ":9323"
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	grpcClient := metric.NewMetricCollectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	res, err := grpcClient.GetMultiMetric(ctx, &metric.Request{})
	if err != nil {
		cancel()
		return nil, err
	}

	cancel()

	return res, nil
}
