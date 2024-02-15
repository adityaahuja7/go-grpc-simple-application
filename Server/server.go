package main

import (
	"context"
	proto "dscdgrpc/protoc"
	"errors"
	"fmt"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)


type server struct {
	proto.UnimplementedSellerMarketServer
}

type seller_struct struct {
	Uuid   uuid.UUID 
	Name   string    
	Ipport string    
}

type product_struct struct {
	ItemID        int32     
	SellerUUID    uuid.UUID 
	Name          string    
	Price         int32     
	Qty           int32    
	Category      string    
	SellerAddress string    
	PricePerUnit  int32     
	Rating        []string  
}

var list_of_products []product_struct
var list_of_sellers []seller_struct

var ctx = context.Background()

func main() {
	list_of_products = make([]product_struct, 0,10)
	list_of_sellers = make([]seller_struct, 0,10)
	fmt.Println("MARKET SERVER IS RUNNING...")
	var port string
	fmt.Print("Enter Port: ")
	fmt.Scanln(&port)
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("🔴 Server failed to start")
		panic(err)
	} else {
		fmt.Println("🟢 Server is running on port " + port)
	}
	srv := grpc.NewServer()
	proto.RegisterSellerMarketServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		fmt.Println("🔴 Server failed to start")
		panic(e)
	} 
}

func NotifySeller(ipport string, item_index int) {
	conn, err := grpc.Dial(ipport, grpc.WithInsecure())
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
	defer conn.Close()
	item := proto.Item{ItemId: 0, ProductName: "Test", Quantity: 0, PricePerUnit: "0", Category: "Test", Description: "Test", SellerAddress: "Test"}
	notification_request := proto.NotifyClientRequest{Notification: &item}
	client := proto.NewMarketSellerClient(conn)
	client.NotifyClient(ctx, &notification_request)
	fmt.Println("✅ Notified Seller at", ipport)
}

func (s *server) RegisterSeller(ctx context.Context, request *proto.RegisterSellerRequest) (*proto.RegisterSellerResponse, error) {
	if request.Uuid == "" {
		return &proto.RegisterSellerResponse{Status: "FAIL"}, nil
	} else {
		newSellerName := request.Name
		newSellerUuid, _ := uuid.Parse(request.Uuid)
		fmt.Println("New Seller:", newSellerName, newSellerUuid)
		NotifySeller(request.Listening, 0)
	}
	return &proto.RegisterSellerResponse{Status: "SUCCESS"}, nil
}

func (s *server) SellItem(ctx context.Context, request *proto.SellItemRequest) (*proto.SellItemResponse, error) {
	client_ip, _ := peer.FromContext(ctx)
	if request.SellerUuid == "" {
		return nil, errors.New("FAIL!")
	} else {
		fmt.Println("❗ SellItem Request from Seller:", client_ip.Addr.String())
		var newItemId int32 = 0
		newItemName := request.ProductName
		newItemQty := request.Quantity
		newItemPrice := request.PricePerUnit
		newItemCategory := request.Category
		newItemDesc := request.Description
		newItemAddr := request.SellerAddress
		newItemSellerUuid, _ := uuid.Parse(request.SellerUuid)
		fmt.Println("New Item:", newItemId, newItemName, newItemQty, newItemPrice, newItemCategory, newItemDesc, newItemAddr, newItemSellerUuid)
	}
	return &proto.SellItemResponse{Status: "SUCCESS"}, nil
}
