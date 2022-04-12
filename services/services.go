package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Shahboz4131/api-gateway-bot/config"
	pb "github.com/Shahboz4131/api-gateway-bot/genproto"
)

type IServiceManager interface {
	BotService() pb.BotServiceClient
}

type serviceManager struct {
	botService pb.BotServiceClient
}

func (s *serviceManager) BotService() pb.BotServiceClient {
	return s.botService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BotServiceHost, conf.BotServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		botService: pb.NewBotServiceClient(connTask),
	}

	return serviceManager, nil
}
