package main

import (
	"context"
	proto "dscdgrpc/protoc"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
	proto.UnimplementedMarketServer
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
	Rating        float32
	RatedBy       []string
	Wishlist      []string
}

var list_of_products []product_struct
var list_of_sellers []seller_struct
var running_id int = 0
var ctx = context.Background()


func NotifySeller(ipport string, item_index int) {
	conn, err := grpc.Dial(ipport, grpc.WithInsecure())
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
	defer conn.Close()
	item := proto.Item{ItemId: 0, ProductName: "Test", Quantity: 0, PricePerUnit: "0", Category: "Test", Description: "Test", SellerAddress: "Test"}
	notification_request := proto.NotifySellerRequest{Notification: &item}
	client := proto.NewMarketSellerClient(conn)
	client.NotifySeller(ctx, &notification_request)
	fmt.Println("✅ Notified Seller at", ipport)
}

func (s *server) RegisterSeller(ctx context.Context, request *proto.RegisterSellerRequest) (*proto.RegisterSellerResponse, error) {
	if request.Uuid == "" {
		return &proto.RegisterSellerResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ Register Seller Request from:", request.Name)
		newSellerName := request.Name
		newSellerUuid, _ := uuid.Parse(request.Uuid)
		newSellerListening := request.Listening
		fmt.Println("New Seller:", newSellerName, newSellerUuid, newSellerListening)
		fmt.Println("---------------------------------")
		newSeller := seller_struct{Uuid: newSellerUuid, Name: newSellerName, Ipport: newSellerListening}
		list_of_sellers = append(list_of_sellers, newSeller)

	}
	return &proto.RegisterSellerResponse{Status: "SUCCESS"}, nil
}

func (s *server) SellItem(ctx context.Context, request *proto.SellItemRequest) (*proto.SellItemResponse, error) {
	client_ip, _ := peer.FromContext(ctx)
	if request.SellerUuid == "" {
		return &proto.SellItemResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ SellItem Request from Seller:", client_ip.Addr.String())
		newItemId := running_id
		running_id++
		newItemName := request.ProductName
		newItemQty := request.Quantity
		newItemPrice := request.PricePerUnit
		newItemCategory := request.Category
		newItemDesc := request.Description
		newItemSellerUuid, _ := uuid.Parse(request.SellerUuid)
		newItemSellerAddress := client_ip.Addr.String()
		fmt.Println("New Item:", newItemId, newItemName, newItemQty, newItemPrice, newItemCategory, newItemDesc, newItemSellerAddress)
		fmt.Println("---------------------------------")
		newItem := product_struct{ItemID: int32(newItemId), SellerUUID: newItemSellerUuid, Name: newItemName, Price: newItemPrice, Qty: newItemQty, Category: newItemCategory, SellerAddress: newItemSellerAddress, PricePerUnit: newItemPrice, Rating: 0.0}
		list_of_products = append(list_of_products, newItem)
		response_message := "SUCCESS! Your Item-ID is " + fmt.Sprint(newItemId)
		return &proto.SellItemResponse{Status: response_message}, nil
	}
}

func (s *server) UpdateItem(ctx context.Context, request *proto.UpdateItemRequest) (*proto.UpdateItemResponse, error) {
	client_ip, _ := peer.FromContext(ctx)
	if request.Uuid == "" {
		return &proto.UpdateItemResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ Update item request from seller:", client_ip.Addr.String())
		item_id := request.ItemId
		seller_uuid := request.Uuid
		for _, seller := range list_of_sellers {
			if seller.Uuid.String() == seller_uuid {
				for j, product := range list_of_products {
					if product.ItemID == item_id && product.SellerUUID.String() == seller_uuid {
						list_of_products[j].PricePerUnit = request.NewPrice
						list_of_products[j].Qty = request.NewQuantity
						response_message := "SUCCESS! Item Updated"
						fmt.Println("---------------------------------")
						return &proto.UpdateItemResponse{Status: response_message}, nil
					}
				}
			}
		}
		return &proto.UpdateItemResponse{Status: "FAIL"}, nil

	}
}

func (s *server) DeleteItem(ctx context.Context, request *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	client_ip, _ := peer.FromContext(ctx)
	if request.Uuid == "" {
		return &proto.DeleteItemResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ Delete item request from seller:", client_ip.Addr.String())
		item_id := request.ItemId
		seller_uuid := request.Uuid
		for _, seller := range list_of_sellers {
			if seller.Uuid.String() == seller_uuid {
				for j, product := range list_of_products {
					if product.ItemID == item_id && product.SellerUUID.String() == seller_uuid {
						list_of_products = append(list_of_products[:j], list_of_products[j+1:]...)
						response_message := "SUCCESS! Item Deleted"
						fmt.Println("---------------------------------")
						return &proto.DeleteItemResponse{Status: response_message}, nil
					}
				}
			}
		}
		return &proto.DeleteItemResponse{Status: "FAIL"}, nil
	}
}

func (s *server) DisplayItems(ctx context.Context, request *proto.DisplayItemsRequest) (*proto.DisplayItemsResponse, error) {
	client_ip, _ := peer.FromContext(ctx)
	fmt.Println("❗ Display item request from seller:", client_ip.Addr.String())
	for _, sellers := range list_of_sellers {
		if sellers.Uuid.String() == request.Uuid {
			var items []*proto.Item
			for _, product := range list_of_products {
				item := proto.Item{ItemId: product.ItemID, ProductName: product.Name, Quantity: product.Qty, PricePerUnit: fmt.Sprint(product.PricePerUnit), Category: product.Category, Description: product.SellerAddress, SellerAddress: product.SellerAddress, Rating: product.Rating}
				items = append(items, &item)
			}
			status := "SUCCESS"
			return &proto.DisplayItemsResponse{Status: status, Items: items}, nil
		}
	}
	return &proto.DisplayItemsResponse{Status: "FAIL!"}, nil
}

func main() {
	list_of_products = make([]product_struct, 0, 10)
	list_of_sellers = make([]seller_struct, 0, 10)
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
	proto.RegisterMarketServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		fmt.Println("🔴 Server failed to start")
		panic(e)
	}
}
