package main

//
// func RunGW(srv apiv1.ClustersServiceServer) {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()
//
// 	// Register gRPC server endpoint
// 	// Note: Make sure the gRPC server is running properly and accessible
// 	mux := runtime.NewServeMux()
// 	// opts := []grpc.DialOption{grpc.WithInsecure()}
// 	err := apiv1.RegisterClustersServiceHandlerServer(ctx, mux, srv)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	go func() {
// 		log.Println("Listening HTTP on :8081")
//
// 		http.ListenAndServe(":8081", mux)
// 	}()
// 	// Start HTTP server (and proxy calls to gRPC server endpoint)
//
// }
