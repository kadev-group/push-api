package manager

import (
	"context"
	"go.uber.org/fx"
)

func Run(
	lc fx.Lifecycle,
	manager *Manager,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			processor := manager.Processor()
			{
				processor.Mailer()
				processor.Cache()
				processor.Queue().Consumers().Mails()
			}
			manager.Repository()
			service := manager.Service()
			{
				service.Mail()
			}

			server := manager.Server()
			{
				if err = server.AMPQ().Handle(); err != nil {
					return err
				}
				server.REST().Run()
			}
			return
		},
		OnStop: func(ctx context.Context) (err error) {
			if err = manager.db.Close(); err != nil {
				return err
			}
			if err = manager.cacheProvider.Close(); err != nil {
				return err
			}
			return nil
		},
	})
}
