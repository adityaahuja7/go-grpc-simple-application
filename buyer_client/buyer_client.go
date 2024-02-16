package main

import (
	"context"
	proto "dscdgrpc/protoc"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type client_buyer_sever struct {
	proto.UnimplementedMarketBuyerServer
}

func (s *client_buyer_sever) NotifyClient(ctx context.Context, in *proto.NotifyBuyerRequest) (*empty.Empty, error) {
	fmt.Println("📩 Notification from Market:", in.Notification)
	return &empty.Empty{}, nil
}

func start_client_buyer_server(portChan chan string, sigs chan os.Signal) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		fmt.Println("🔴 Client failed to start listening")
		fmt.Println("---------------------------------")
	} else {
		portChan <- lis.Addr().String()
	}
	srv := grpc.NewServer()
	proto.RegisterMarketBuyerServer(srv, &client_buyer_sever{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		fmt.Println("🔴 Server failed to start")
		fmt.Println("---------------------------------")
		panic(e)
	}
	go func() {
		<-sigs
		srv.Stop()
	}()
}

func searchItem(client proto.MarketClient) {
	var product_name, product_category string
	fmt.Print("Enter Product Name: ")
	fmt.Scan(&product_name)
	fmt.Print("Enter Product Category: ")
	fmt.Scan(&product_category)
	response, err := client.SearchItems(context.Background(), &proto.SearchItemRequest{ProductName: product_name, ProductCategory: product_category})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
}

func buyItem(client proto.MarketClient, portChan chan string) {
	var item_id, quantity int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	fmt.Print("Enter Quantity: ")
	fmt.Scan(&quantity)
	buyer_address := <-portChan
	response, err := client.BuyItem(context.Background(), &proto.BuyItemRequest{ItemId: item_id, Quantity: quantity, Listening: buyer_address})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
}

func addToWishlist(client proto.MarketClient, portChan chan string) {
	var item_id int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	buyer_address := <-portChan
	response, err := client.AddToWishlist(context.Background(), &proto.AddToWishlistRequest{ItemId: item_id, Listening: buyer_address})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
}

func rateItem(client proto.MarketClient) {
	var item_id, rating int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	fmt.Println("Enter Rating (1-5):")
	fmt.Scan(&rating)
	response, err := client.RateItem(context.Background(), &proto.RateItemRequest{ItemId: item_id, Rating: rating})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
}

func main() {

	//Setup connection parameters
	fmt.Println("SELLER CLIENT IS RUNNING...")
	var ip, port string
	fmt.Print("Enter IP: ")
	fmt.Scan(&ip)
	fmt.Print("Enter Port: ")
	fmt.Scan(&port)
	portChan := make(chan string)

	//handle shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("🔴 Client is shutting down...")
		os.Exit(0)
	}()

	//start a seperate routine for client notifications
	go start_client_buyer_server(portChan, sigs)

	//connect to server
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	} else {
		defer conn.Close()
		fmt.Println("🟢 Connected to server on:", ip+":"+port)
	}
	client := proto.NewMarketClient(conn)

	fmt.Println("1. Register Seller")
	fmt.Println("2. Sell Item")
	fmt.Println("3. Update Item")
	fmt.Println("4. Delete Item")
	fmt.Println("5. Display All Items")
	fmt.Println("6. Exit")
	fmt.Println("---------------------------------")
	for {
		var choice int
		fmt.Print("Enter Choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			searchItem(client)
		case 2:
			buyItem(client, portChan)
		case 3:
			addToWishlist(client, portChan)
		case 4:
			rateItem(client)
		}
	}
}
