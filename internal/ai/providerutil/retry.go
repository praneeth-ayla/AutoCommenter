package providerutil

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

const (
	DefaultRetryDelay = 5 * time.Second
	MaxRetryAttempts  = 3
	PerRequestTimeout = 60 * time.Second
)

// DoWithRetry runs fn with retry, timeout and rate-limit handling.
// fn itself is synchronous; this helper runs it in a goroutine and enforces timeout.
func DoWithRetry[T any](maxAttempts int, timeout time.Duration, fn func() (T, error)) (T, error) {
	var zero T
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), timeout) // Create a context with a timeout for this attempt.
		done := make(chan struct{})
		var result T

		go func() {
			result, lastErr = fn()
			close(done) // Signal that the function has completed.
		}()

		select {
		case <-ctx.Done():
			lastErr = fmt.Errorf("request timed out after %s", timeout) // Handle context cancellation due to timeout.
		case <-done:
			// fn completed, lastErr is already set
		}
		cancel() // Release the context resources.

		if lastErr == nil {
			return result, nil // Success, return the result.
		}

		if delay, isRateLimit := CheckRateLimitError(lastErr); isRateLimit {
			if attempt < maxAttempts {
				SleepWithJitter(delay) // Wait with jitter before the next retry if it's a rate limit error.
				continue
			}
			return zero, fmt.Errorf("rate limit after %d attempts: %w", maxAttempts, lastErr)
		}

		// Non-rate-limit error: fail fast
		return zero, fmt.Errorf("operation failed: %w", lastErr) // Return immediately for non-rate-limit errors.
	}

	return zero, fmt.Errorf("retries exhausted after %d attempts: %w", maxAttempts, lastErr) // All retries failed.
}

func SleepWithJitter(base time.Duration) {
	if base <= 0 {
		base = DefaultRetryDelay // Use default delay if provided delay is non-positive.
	}
	j := time.Duration(rand.Int63n(int64(base / 2))) // Calculate a random jitter up to half of the base delay.
	time.Sleep(base + j)                             // Sleep for base delay plus jitter.
}

func CheckRateLimitError(err error) (time.Duration, bool) {
	if err == nil {
		return 0, false
	}

	errStr := err.Error()
	// Check for common string patterns indicating rate limiting.
	if strings.Contains(errStr, "RESOURCE_EXHAUSTED") ||
		strings.Contains(errStr, "429") ||
		strings.Contains(errStr, "Quota exceeded") {
		return extractRetryDelay(errStr), true
	}

	var apiErr *apierror.APIError
	if errors.As(err, &apiErr) {
		if status := apiErr.GRPCStatus(); status != nil && status.Code() == codes.ResourceExhausted {
			// Check for gRPC status code ResourceExhausted.
			for _, detail := range status.Details() {
				if retryInfo, ok := detail.(*errdetails.RetryInfo); ok {
					if retryInfo.RetryDelay != nil {
						if d := retryInfo.RetryDelay.AsDuration(); d > 0 {
							return d, true // Use retry delay from RetryInfo if available and positive.
						}
					}
				}
			}
			return DefaultRetryDelay, true // Fallback to default delay if RetryInfo is not available or zero.
		}
	}

	return 0, false
}

func extractRetryDelay(errStr string) time.Duration {
	re := regexp.MustCompile(`retry in ([0-9.]+)s`) // Regex to find "retry in X.Ys" pattern.
	if matches := re.FindStringSubmatch(errStr); len(matches) > 1 {
		if seconds, err := strconv.ParseFloat(matches[1], 64); err == nil {
			return time.Duration(seconds * float64(time.Second)) // Parse seconds and convert to duration.
		}
	}
	return DefaultRetryDelay // Return default delay if parsing fails or pattern not found.
}