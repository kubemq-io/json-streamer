package services

import (
	"context"
	"fmt"
	"github.com/kubemq-io/kubemq-go/queues_stream"
	target "github.com/kubemq-io/kubemq-targets/types"
	"github/kubemq-io/json-streamer/config"
	"github/kubemq-io/json-streamer/pkg/logger"
	"github/kubemq-io/json-streamer/types"
	"time"
)

type Service struct {
	cfg    *config.Config
	logger *logger.Logger
	player *types.Player
	client *queues_stream.QueuesStreamClient
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Init(ctx context.Context, cfg *config.Config) error {
	s.cfg = cfg
	s.logger = logger.NewLogger("player")
	s.player = types.NewPlayer()
	var err error
	s.client, err = queues_stream.NewQueuesStreamClient(ctx,
		queues_stream.WithAddress(s.cfg.Host(), s.cfg.Port()),
		queues_stream.WithClientId("player"),
		queues_stream.WithCheckConnection(true),
		queues_stream.WithAutoReconnect(true),
	)
	if err != nil {
		return err
	}
	s.logger.Info("kubemq connected.")
	return nil
}
func (s *Service) Start(ctx context.Context) error {
	go s.run(ctx)

	return nil
}
func (s *Service) Stop() {
	return
}
func (s *Service) sendToKubeMQ(ctx context.Context, song *types.SongChart) error {
	req := target.NewRequest().
		SetMetadataKeyValue("method", "query").
		SetData([]byte(song.InsertSql(s.cfg.Table)))
	msg := queues_stream.NewQueueMessage().SetChannel(s.cfg.Queue).SetBody(req.MarshalBinary())
	result, err := s.client.Send(ctx, msg)
	if err != nil {
		return err
	}
	if len(result.Results) != 0 {
		res := result.Results[0]
		if res.IsError {
			return fmt.Errorf("%s", res.Error)
		}
	}

	return nil
}
func (s *Service) run(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Duration(s.cfg.Interval) * time.Second):
			song := s.player.PlayRandomSong()
			s.logger.Infof("count: %d,song played: %s", song.Count, song.SongName)
			if err := s.sendToKubeMQ(ctx, song); err != nil {
				s.logger.Errorf("error sending data to kubemq, %s", err.Error())
			}
		case <-ctx.Done():
			return
		}
	}
}
