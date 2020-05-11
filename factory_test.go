package logger

import (
	userclient "github.com/best-expendables/user-service-client"
	"fmt"
	"sync"
	"testing"

	"context"

	"reflect"

	"github.com/best-expendables/trace"
)

// flag -race required
func TestFactory_Logger(t *testing.T) {
	factory := NewLoggerFactory(InfoLevel)

	var wg sync.WaitGroup

	routinesNumber := 10
	wg.Add(routinesNumber)

	for i := 0; i < 10; i++ {
		go func(i int) {
			requestID := fmt.Sprintf("REQUEST_ID_%d", i)
			userID := fmt.Sprintf("USER_ID_%d", i)

			ctx := populatedContext(requestID, userID)

			entry, ok := factory.Logger(ctx).(*entry)
			if !ok {
				t.Fatal("Factory should return type *logger.logger")
			}

			metaAssert := meta{
				UserID:    userID,
				RequestID: requestID,
			}

			if !reflect.DeepEqual(entry.meta, metaAssert) {
				t.Errorf("Meta not equals, expected '%+v', 'actual '%+v", metaAssert, entry.meta)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

// populatedContext return context with trace-id, request-id and user
func populatedContext(requestID, userID string) context.Context {
	ctx := context.TODO()
	ctx = trace.ContextWithRequestID(ctx, requestID)
	ctx = userclient.ContextWithUser(ctx, &userclient.User{Id: userID})
	return ctx
}
