package es

// func TestReadStream(t *testing.T) {
// 	es, err := NewESClient(config.DBConfig{
// 		Scheme: "esdb",
// 		Host:   "localhost",
// 		Port:   "2113",
// 		User:   "admin",
// 		Pass:   "changeit",
// 	})
// 	assert.Nil(t, err)
// 	r, err := es.Db.(context.Background(), "$$stream_by_category", esdb.ReadStreamOptions{
// 		Direction: esdb.Forwards,
// 		From:      esdb.Start{},
// 	}, 12)
// 	assert.Nil(t, err)

// 	defer r.Close()
// 	for {
// 		event, err := r.Recv()
// 		if errors.Is(err, io.EOF) {
// 			break
// 		}
// 		assert.Nil(t, err)
// 		if event == nil {
// 			break
// 		}
// 		fmt.Printf("event: %+v\n", string(event.Event.Data))
// 	}

// }
