package a

import "context"

func main() {
	Sample()
	SampleWithContext(context.Background())
}

func Sample() error {
	return nil
}

func SampleWithContext(ctx context.Context) error {
	return nil
}
