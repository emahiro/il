package a

import "context"

func main() {
	Sample()
	SampleWithContext(context.Background())
}

func Sample() error { // want `change Sample to Example`
	return nil
}

func SampleWithContext(ctx context.Context) error { // want `change SampleWithContext to ExampleWithContext`
	return nil
}
