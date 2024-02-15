package main

import (
	"context"
	proto "dscdgrpc/protoc"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type client_sever struct {
	proto.UnimplementedMarketSellerServer
}


func start_client_server(portChan chan string, sigs chan os.Signal) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		fmt.Println("🔴 Client failed to start listening")
	}else{
		portChan <- lis.Addr().String()
	}
	srv := grpc.NewServer()
	proto.RegisterMarketSellerServer(srv, &client_sever{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		fmt.Println("🔴 Server failed to start")
		panic(e)
	}
	go func(){
		<-sigs
		srv.Stop()
	}()
}

func (s *client_sever) NotifyClient(ctx context.Context, in *proto.NotifyClientRequest) (*proto.NotifyClientResponse, error) {
	fmt.Println("📩 Notification from server:", in.Notification)
	return &proto.NotifyClientResponse{Status: "OK!"}, nil
}


func RegisterSeller(client proto.SellerMarketClient, portChan chan string) uuid.UUID {
	var name_first, name_last string
	uuid := uuid.New()
	fmt.Print("Enter Name (FirstName LastName):")
	fmt.Scan(&name_first, &name_last)
	fmt.Println("❗ Your UUID is:", uuid)
	portchan := <-portChan
	response, err := client.RegisterSeller(context.Background(), &proto.RegisterSellerRequest{Name: name_first + " " + name_last, Listening: portchan, Uuid: uuid.String()})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
	return uuid
}

func SellItem(client proto.SellerMarketClient) {
	var name string
	var qty, price_per_unit int
	var category string
	var Desc, Addr string
	var uid string
	fmt.Print("Enter Item Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Item Quantity: ")
	fmt.Scan(&qty)
	fmt.Print("Enter Price Per Unit: ")
	fmt.Scan(&price_per_unit)
	fmt.Print("Enter Item Category: ")
	fmt.Scan(&category)
	fmt.Print("Enter Description: ")
	fmt.Scan(&Desc)
	fmt.Print("Enter Seller address: ")
	fmt.Scan(&Addr)
	fmt.Print("Enter Seller UUID: ")
	fmt.Scan(&uid)
	seller_uuid, _ := uuid.Parse(uid)
	seller_uuid_string := seller_uuid.String()
	request := &proto.SellItemRequest{SellerUuid: seller_uuid_string, ProductName: name, Category: category, Quantity: int32(qty), Description: Desc, SellerAddress: Addr, PricePerUnit: int32(price_per_unit)}
	fmt.Println(request)
	response, err := client.SellItem(context.Background(), &proto.SellItemRequest{SellerUuid: uid, ProductName: name, Category: category, Quantity: int32(qty), Description: Desc, SellerAddress: Addr, PricePerUnit: int32(price_per_unit)})
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
	go func(){
		<-sigs
		fmt.Println("🔴 Client is shutting down...")
		os.Exit(0)
	}()

	//start a seperate routine for client notifications
	go start_client_server(portChan, sigs)

	//connect to server
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	} else {
		defer conn.Close()
		fmt.Println("🟢 Connected to server on:", ip+":"+port)
	}
	client := proto.NewSellerMarketClient(conn)


	fmt.Println("1. Register Seller")
	fmt.Println("2. Sell Item")
	fmt.Println("3. Update Item")
	fmt.Println("4. Delete Item")
	fmt.Println("5. Display All Items")
	fmt.Println("6. Exit")
	for {
		var choice int
		fmt.Print("Enter Choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			RegisterSeller(client, portChan)
		case 2:
			SellItem(client)
		case 6:
			return
		}
	}
}
