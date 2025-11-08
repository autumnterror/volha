package products

import (
	"context"
	"fmt"
	"gateway/config"
	"gateway/internal/utils/format"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

type Client struct {
	api productsRPC.ProductsClient
}

func New(
	cfg *config.Config,
) (*Client, error) {
	const op = "grpc.products.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(cfg.RetriesCount)),
		grpcretry.WithPerRetryTimeout(cfg.Timeout),
		grpcretry.WithBackoff(grpcretry.BackoffExponential(cfg.Backoff)),
	}

	//logOpts := []grpclog.Option{
	//	grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	//}

	cc, err := grpc.NewClient(
		cfg.AddrProducts,
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...),
			//	grpclog.UnaryClientInterceptor(interceptorLogger(), logOpts...),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5 * time.Minute,
			Timeout:             11 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return nil, format.Error(op, err)
	}

	log.Println("start connection...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cc.Connect()
	for {
		state := cc.GetState()
		if state == connectivity.Ready {
			log.Println("CONNECT!!")
			break
		}

		if !cc.WaitForStateChange(ctx, state) {
			return nil, fmt.Errorf("connection timeout")
		}
	}

	return &Client{
		api: productsRPC.NewProductsClient(cc),
	}, nil
}

//func interceptorLogger() grpclog.Logger {
//	const op = "balance.log"
//	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
//		log.Println(format.Params(op, fields))
//	})
//}
