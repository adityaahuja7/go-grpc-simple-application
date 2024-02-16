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

var notification_ipport string

func enum_to_category(enumval proto.Category) string {
	if enumval == proto.Category_ELECTRONICS {
		return "ELECTRONICS"
	} else if enumval == proto.Category_FASHION{
		return "FASHION"
	} else {
		return "OTHERS"
	}
}

func category_to_enum (category string) proto.Category {
	if category == "ELECTRONICS" {
		return proto.Category_ELECTRONICS
	} else if category == "FASHION" {
		return proto.Category_FASHION
	} else {
		return proto.Category_OTHERS
	}
}

type client_buyer_sever struct {
	proto.UnimplementedMarketBuyerServer
}

func (s *client_buyer_sever) NotifyBuyer(ctx context.Context, in *proto.NotifyBuyerRequest) (*empty.Empty, error) {
	fmt.Println("📩 Notification from Market:", in.Notification.Message)
	fmt.Println("Item ID:", in.Notification.SoldItem.ItemId)
	fmt.Println("Item Name:", in.Notification.SoldItem.ProductName)
	fmt.Println("Description:", in.Notification.SoldItem.Description)
	fmt.Println("Remaining Quantity:", in.Notification.SoldItem.Quantity)
	fmt.Println("Price:", in.Notification.SoldItem.PricePerUnit)
	fmt.Println("Rating:", in.Notification.SoldItem.Rating)
	fmt.Println("Seller Address:", in.Notification.SoldItem.SellerAddress)
	fmt.Println("---------------------------------")
	return &empty.Empty{}, nil
}

func start_client_buyer_server(sigs chan os.Signal) {
	lis, err := net.Listen("tcp", notification_ipport)
	if err != nil {
		fmt.Println("🔴 Client failed to start listening")
		fmt.Println("---------------------------------")
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
	response, err := client.SearchItems(context.Background(), &proto.SearchItemRequest{ProductName: product_name, Category: product_category})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response.Status)
		if len(response.Items) == 0 {
			fmt.Println("No items found!")
		}
		for _, item := range response.Items {
			fmt.Println("------")
			fmt.Println("Item ID:", item.ItemId, "\nName:", item.ProductName, "\nCategory:", item.Category, "\nPrice:", item.PricePerUnit, "\nQuantity:", item.Quantity, "\nDescription", item.Description, "\nRating:", item.Rating)
			fmt.Println("------")
		}
	}
}

func buyItem(client proto.MarketClient) {
	var item_id, quantity int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	fmt.Print("Enter Quantity: ")
	fmt.Scan(&quantity)
	response, err := client.BuyItem(context.Background(), &proto.BuyItemRequest{ItemId: item_id, Quantity: quantity, Listening: notification_ipport})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
		fmt.Println("---------------------------------")
	}
}

func addToWishlist(client proto.MarketClient) {
	var item_id int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	response, err := client.AddToWishlist(context.Background(), &proto.AddToWishlistRequest{ItemId: item_id, Listening: notification_ipport})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
		fmt.Println("---------------------------------")
	}
}

func rateItem(client proto.MarketClient) {
	var item_id, rating int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	fmt.Print("Enter Rating (1-5):")
	fmt.Scan(&rating)
	response, err := client.RateItem(context.Background(), &proto.RateItemRequest{ItemId: item_id, Rating: rating, Listening: notification_ipport})
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
	fmt.Print("Enter Notification IP:PORT: ")
	fmt.Scan(&notification_ipport)

	//handle shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("🔴 Client is shutting down...")
		os.Exit(0)
	}()

	//start a seperate routine for client notifications
	go start_client_buyer_server(sigs)

	//connect to server
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	} else {
		defer conn.Close()
		fmt.Println("🟢 Connected to server on:", ip+":"+port)
	}
	client := proto.NewMarketClient(conn)

	//main menu
	fmt.Println("🟢 Welcome to the Market")
	fmt.Println("1. Search Item")
	fmt.Println("2. Buy Item")
	fmt.Println("3. Add to Wishlist")
	fmt.Println("4. Rate Item")
	fmt.Println("5. Exit")
	fmt.Println("---------------------------------")
	for {
		var choice int
		fmt.Print("Enter Choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			searchItem(client)
		case 2:
			buyItem(client)
		case 3:
			addToWishlist(client)
		case 4:
			rateItem(client)
		case 5:
			os.Exit(0)
		}
	}
}
