package querier

import (
	"context"
	"time"

	"github.com/grafana/dskit/ring"
	ring_client "github.com/grafana/dskit/ring/client"

	"github.com/grafana/fire/pkg/gen/ingester/v1/ingestv1connect"
)

type responseFromIngesters[T interface{}] struct {
	addr     string
	response T
}

type IngesterFn[T interface{}] func(ingestv1connect.IngesterClient) (T, error)

// IngesterQuerier helps with querying the ingesters.
type IngesterQuerier struct {
	ring            ring.ReadRing
	pool            *ring_client.Pool
	extraQueryDelay time.Duration
}

func NewIngesterQuerier(pool *ring_client.Pool, ring ring.ReadRing, extraQueryDelay time.Duration) *IngesterQuerier {
	return &IngesterQuerier{
		ring:            ring,
		pool:            pool,
		extraQueryDelay: extraQueryDelay,
	}
}

// forAllIngesters runs f, in parallel, for all ingesters
func forAllIngesters[T any](ctx context.Context, q *IngesterQuerier, f IngesterFn[T]) ([]responseFromIngesters[T], error) {
	replicationSet, err := q.ring.GetReplicationSetForOperation(ring.Read)
	if err != nil {
		return nil, err
	}

	return forGivenIngesters(ctx, q, replicationSet, f)
}

// forGivenIngesters runs f, in parallel, for given ingesters
func forGivenIngesters[T any](ctx context.Context, q *IngesterQuerier, replicationSet ring.ReplicationSet, f IngesterFn[T]) ([]responseFromIngesters[T], error) {
	results, err := replicationSet.Do(ctx, q.extraQueryDelay, func(ctx context.Context, ingester *ring.InstanceDesc) (interface{}, error) {
		client, err := q.pool.GetClientFor(ingester.Addr)
		if err != nil {
			return nil, err
		}

		resp, err := f(client.(ingestv1connect.IngesterClient))
		if err != nil {
			return nil, err
		}

		return responseFromIngesters[T]{ingester.Addr, resp}, nil
	})
	if err != nil {
		return nil, err
	}

	responses := make([]responseFromIngesters[T], 0, len(results))
	for _, result := range results {
		responses = append(responses, result.(responseFromIngesters[T]))
	}

	return responses, err
}
