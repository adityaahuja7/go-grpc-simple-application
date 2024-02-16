package main

import (
	"bufio"
	"context"
	proto "dscdgrpc/protoc"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type client_seller_sever struct {
	proto.UnimplementedMarketSellerServer
}

var notification_ipport string = "localhost:4041"
var uuid_global uuid.UUID = uuid.UUID{}

func enum_to_category(enumval proto.Category) string {
	if enumval == proto.Category_ELECTRONICS {
		return "ELECTRONICS"
	} else if enumval == proto.Category_FASHION {
		return "FASHION"
	} else {
		return "OTHERS"
	}
}

func category_to_enum(category string) proto.Category {
	if category == "ELECTRONICS" {
		return proto.Category_ELECTRONICS
	} else if category == "FASHION" {
		return proto.Category_FASHION
	} else {
		return proto.Category_OTHERS
	}
}

func (s *client_seller_sever) NotifySeller(ctx context.Context, in *proto.NotifySellerRequest) (*empty.Empty, error) {
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

func start_client_seller_server(sigs chan os.Signal) {
	lis, err := net.Listen("tcp", notification_ipport)
	if err != nil {
		fmt.Println("🔴 Client failed to start listening")
		fmt.Println("---------------------------------")
	}
	srv := grpc.NewServer()
	proto.RegisterMarketSellerServer(srv, &client_seller_sever{})
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

func RegisterSeller(client proto.MarketClient) {
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Print("Enter Name:")
	os.Stdin.Read(make([]byte, 1000))
	name, _ = reader.ReadString('\n')
	fmt.Println("❗ Your UUID is:", uuid_global)

	response, err := client.RegisterSeller(context.Background(), &proto.RegisterSellerRequest{Name: name, Listening: notification_ipport, Uuid: uuid_global.String()})
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", response)
	}
}

func SellItem(client proto.MarketClient) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Item Name:")
	os.Stdin.Read(make([]byte, 1000))
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Item Quantity: ")
	qtyStr, _ := reader.ReadString('\n')
	qty, _ := strconv.Atoi(strings.TrimSpace(qtyStr))

	fmt.Print("Enter Price Per Unit: ")
	priceStr, _ := reader.ReadString('\n')
	price_per_unit, _ := strconv.Atoi(strings.TrimSpace(priceStr))

	fmt.Print("Enter Item Category (ELECTRONICS/FASHION/OTHERS): ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)
	var enum_category proto.Category
	if category == "ELECTRONICS" {
		enum_category = proto.Category_ELECTRONICS
	} else if category == "FASHION" {
		enum_category = proto.Category_FASHION
	} else {
		enum_category = proto.Category_OTHERS
	}

	fmt.Print("Enter Description: ")
	Desc, _ := reader.ReadString('\n')
	Desc = strings.TrimSpace(Desc)
	uuid_string := uuid_global.String()
	fmt.Println("❗ Your UUID is:", uuid_global)
	response, err := client.SellItem(context.Background(), &proto.SellItemRequest{SellerUuid: uuid_string, ProductName: name, Category: enum_category, Quantity: int32(qty), Description: Desc, PricePerUnit: int32(price_per_unit), Listening: notification_ipport})
	if err != nil {
		fmt.Println("❌ Error:", err)
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("✅ Response:", response)
		fmt.Println("---------------------------------")
	}
}

func UpdateItem(client proto.MarketClient) {
	var item_id int32
	var new_price int32
	var new_qty int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	fmt.Print("Enter New Price: ")
	fmt.Scan(&new_price)
	fmt.Print("Enter New Quantity: ")
	fmt.Scan(&new_qty)
	uuid_string := uuid_global.String()
	update_request := &proto.UpdateItemRequest{Uuid: uuid_string, ItemId: item_id, NewPrice: new_price, NewQuantity: new_qty}
	response, err := client.UpdateItem(context.Background(), update_request)
	if err != nil {
		fmt.Println("❌ Error:", err)
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("✅ Response:", response)
		fmt.Println("---------------------------------")
	}
}

func DeleteItem(client proto.MarketClient) {
	var item_id int32
	fmt.Print("Enter Item ID: ")
	fmt.Scan(&item_id)
	uuid_string := uuid_global.String()
	delete_request := &proto.DeleteItemRequest{Uuid: uuid_string, ItemId: item_id, Listening: notification_ipport}
	response, err := client.DeleteItem(context.Background(), delete_request)
	if err != nil {
		fmt.Println("❌ Error:", err)
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("✅ Response:", response)
		fmt.Println("---------------------------------")
	}
}

func DisplaySellerItems(client proto.MarketClient) {
	uuid_string := uuid_global.String()
	display_request := &proto.DisplayItemsRequest{Uuid: uuid_string, Listening: notification_ipport}
	response, err := client.DisplayItems(context.Background(), display_request)
	if err != nil {
		fmt.Println("❌ Error:", err)
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("✅ Response:", response.Status)
		for _, product := range response.Items {
			fmt.Println("-----------------")
			fmt.Println("Item ID:", product.ItemId, "\nName:", product.ProductName, "\nPrice:", product.PricePerUnit, "\nQuantity:", product.Quantity, "\nCategory:", product.Category, "\nDescription:", product.Description, "\nSeller Address:", product.SellerAddress, "\nRating:", product.Rating)

		}
	}
	fmt.Println("---------------------------------")
}

func main() {
	uuid_global = uuid.New()
	//Setup connection parameters
	fmt.Println("SELLER CLIENT IS RUNNING...")
	var ip, port string
	fmt.Print("Enter IP: ")
	fmt.Scan(&ip)
	fmt.Print("Enter Port: ")
	fmt.Scan(&port)

	//handle shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("🔴 Client is shutting down...")
		os.Exit(0)
	}()

	//start a seperate routine for client notifications
	go start_client_seller_server(sigs)

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
			RegisterSeller(client)
		case 2:
			SellItem(client)
		case 3:
			UpdateItem(client)
		case 4:
			DeleteItem(client)
		case 5:
			DisplaySellerItems(client)
		case 6:
			return
		}
	}
}
