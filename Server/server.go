package main

import (
	"context"
	proto "dscdgrpc/protoc"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
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
	Description   string
	SellerAddress string
	PricePerUnit  int32
	Rating        float32
	RatedBy       []string
	Wishlist      []string
}

var list_of_products []product_struct
var list_of_sellers []seller_struct
var running_id int = 0

func enum_to_category(enumval proto.Category) string {
	if enumval == proto.Category_ELECTRONICS {
		return "ELECTRONICS"
	} else if enumval == proto.Category_FASHION {
		return "FASHION"
	} else if enumval == proto.Category_OTHERS {
		return "OTHERS"
	}
	return ""
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

func NotifySeller(ipport string, item_index int) {
	conn, err := grpc.Dial(ipport, grpc.WithInsecure())
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
	defer conn.Close()
	client := proto.NewMarketSellerClient(conn)
	for _, item := range list_of_products {
		if item.ItemID == int32(item_index) {
			sold_item := proto.Item{ItemId: item.ItemID, ProductName: item.Name, Quantity: item.Qty, PricePerUnit: item.PricePerUnit, Category: category_to_enum(item.Category), Description: item.Description, SellerAddress: item.SellerAddress, Rating: item.Rating}
			message := "Your item " + item.Name + " has been sold"
			notification := &proto.Notification{Message: message, SoldItem: &sold_item}
			_, err := client.NotifySeller(context.Background(), &proto.NotifySellerRequest{Notification: notification})
			if err != nil {
				fmt.Println("❌ Error:", err)
			} else {
				fmt.Println("✅ Notified Seller with Ip:", ipport)
				fmt.Println("---------------------------------")
			}
		}
	}
}

func NotifyBuyer(item_index int) {
	for _, item := range list_of_products {
		if item.ItemID == int32(item_index) {
			for _, ipport := range item.Wishlist {
				conn, err := grpc.Dial(ipport, grpc.WithInsecure())
				if err != nil {
					fmt.Println("❌ Error:", err)
				}
				defer conn.Close()
				client := proto.NewMarketBuyerClient(conn)
				modified_item := proto.Item{ItemId: item.ItemID, ProductName: item.Name, Quantity: item.Qty, PricePerUnit: item.PricePerUnit, Category: category_to_enum(item.Category), Description: item.Description, SellerAddress: item.SellerAddress, Rating: item.Rating}
				message := "Your wishlisted item " + item.Name + " has been modified"
				notification := &proto.Notification{Message: message, SoldItem: &modified_item}
				_, err = client.NotifyBuyer(context.Background(), &proto.NotifyBuyerRequest{Notification: notification})
				if err != nil {
					fmt.Println("❌ Error:", err)
				} else {
					fmt.Println("✅ Notified Buyer with IP:", ipport)
					fmt.Println("---------------------------------")
				}
			}
		}

	}
}

func (s *server) RegisterSeller(ctx context.Context, request *proto.RegisterSellerRequest) (*proto.RegisterSellerResponse, error) {
	if request.Uuid == "" {
		return &proto.RegisterSellerResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ Register Seller Request from:", request.Listening)
		for _,sellers := range list_of_sellers {
			if sellers.Uuid.String() == request.Uuid {
				return &proto.RegisterSellerResponse{Status: "FAIL! (Alreay Registered!)"}, nil
			}
		}

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
	if request.SellerUuid == "" {
		return &proto.SellItemResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ SellItem Request from Seller:", request.Listening)
		newItemId := running_id
		running_id++
		newItemName := request.ProductName
		newItemQty := request.Quantity
		newItemPrice := request.PricePerUnit
		newItemCategory := request.Category
		newItemDesc := request.Description
		newItemSellerUuid, _ := uuid.Parse(request.SellerUuid)
		var newItemSellerAddress string
		for _, seller := range list_of_sellers {
			if seller.Uuid.String() == request.SellerUuid {
				newItemSellerAddress = seller.Ipport
			}
		}
		fmt.Println("New Item ID:", newItemId)
		fmt.Println("New Item Name:", newItemName)
		fmt.Println("New Item Quantity:", newItemQty)
		fmt.Println("New Item Price:", newItemPrice)
		fmt.Println("New Item Category:", enum_to_category(newItemCategory))
		fmt.Println("New Item Description:", newItemDesc)
		fmt.Println("New Item Seller UUID:", newItemSellerUuid)
		fmt.Println("New Item Seller Address:", newItemSellerAddress)
		fmt.Println("---------------------------------")
		newItem := product_struct{ItemID: int32(newItemId), SellerUUID: newItemSellerUuid, Name: newItemName, Price: newItemPrice, Qty: newItemQty, Category: enum_to_category(newItemCategory), SellerAddress: newItemSellerAddress, PricePerUnit: newItemPrice, Rating: 0.0, Description: newItemDesc, RatedBy: nil, Wishlist: nil}
		list_of_products = append(list_of_products, newItem)
		response_message := "SUCCESS! Your Item-ID is " + fmt.Sprint(newItemId)
		return &proto.SellItemResponse{Status: response_message}, nil
	}
}

func (s *server) UpdateItem(ctx context.Context, request *proto.UpdateItemRequest) (*proto.UpdateItemResponse, error) {
	if request.Uuid == "" {
		return &proto.UpdateItemResponse{Status: "FAIL! (Empty UUID)"}, nil
	} else {
		fmt.Println("❗ Update item request from seller:", request.Listening, "for item:", request.ItemId)
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
						NotifyBuyer(j)
						return &proto.UpdateItemResponse{Status: response_message}, nil
					}
				}
			}
		}
		return &proto.UpdateItemResponse{Status: "FAIL! (Not Registered)"}, nil

	}
}

func (s *server) DeleteItem(ctx context.Context, request *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	if request.Uuid == "" {
		return &proto.DeleteItemResponse{Status: "FAIL"}, nil
	} else {
		fmt.Println("❗ Delete item request from seller:", request.Listening, "for item:", request.ItemId)
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
	fmt.Println("❗ Display item request from seller:", request.Listening)
	fmt.Println("---------------------------------")
	for _, sellers := range list_of_sellers {
		println("Seller UUID:", sellers.Uuid.String())
		println("Request UUID:", request.Uuid)
		if sellers.Uuid.String() == request.Uuid {
			for _,products := range list_of_products {
				if products.SellerUUID.String() == request.Uuid {
					item := proto.Item{ItemId: products.ItemID, ProductName: products.Name, Quantity: products.Qty, PricePerUnit: products.PricePerUnit, Category: category_to_enum(products.Category), Description: products.Description, SellerAddress: products.SellerAddress, Rating: products.Rating}
					items := append([]*proto.Item{}, &item)
					status := "SUCCESS"
					return &proto.DisplayItemsResponse{Status: status, Items: items}, nil
				}
			}
		}
	}
	return &proto.DisplayItemsResponse{Status: "FAIL! (No relevant items found)"}, nil
}

func (s *server) SearchItems(ctx context.Context, request *proto.SearchItemRequest) (*proto.SearchItemResponse, error) {
	fmt.Println("❗ Search item request from buyer:", request.Listening, "for:", request.ProductName, request.Category)
	fmt.Println("---------------------------------")
	var items []*proto.Item
	for _, product := range list_of_products {
		if request.ProductName == "" {
			if product.Category == request.Category {
				item := proto.Item{ItemId: product.ItemID, ProductName: product.Name, Quantity: product.Qty, PricePerUnit: product.PricePerUnit, Category: category_to_enum(product.Category), Description: product.Description, SellerAddress: product.SellerAddress, Rating: product.Rating}
				items = append(items, &item)
			} else if request.Category == "ANY" {
				item := proto.Item{ItemId: product.ItemID, ProductName: product.Name, Quantity: product.Qty, PricePerUnit: product.PricePerUnit, Category: category_to_enum(product.Category), Description: product.SellerAddress, SellerAddress: product.SellerAddress, Rating: product.Rating}
				items = append(items, &item)
			}
		} else {
			if product.Name == request.ProductName && product.Category == request.Category {
				item := proto.Item{ItemId: product.ItemID, ProductName: product.Name, Quantity: product.Qty, PricePerUnit: product.PricePerUnit, Category: category_to_enum(product.Category), Description: product.Description, SellerAddress: product.SellerAddress, Rating: product.Rating}
				items = append(items, &item)
			} else if request.Category == "ANY" {
				item := proto.Item{ItemId: product.ItemID, ProductName: product.Name, Quantity: product.Qty, PricePerUnit: product.PricePerUnit, Category: category_to_enum(product.Category), Description: product.Description, SellerAddress: product.SellerAddress, Rating: product.Rating}
				items = append(items, &item)
			}
		}
	}
	status := "SUCCESS"
	return &proto.SearchItemResponse{Status: status, Items: items}, nil
}

func (s *server) BuyItem(ctx context.Context, request *proto.BuyItemRequest) (*proto.BuyItemResponse, error) {
	fmt.Println("❗ Buy item request from buyer:", request.Listening, "for item:", request.ItemId)
	fmt.Println("---------------------------------")
	item_id := request.ItemId
	for j, product := range list_of_products {
		fmt.Println("Quantity:", product.Qty, "Requested:", request.Quantity)
		if product.ItemID == item_id {
			if product.Qty >= request.Quantity {
				list_of_products[j].Qty -= request.Quantity
				NotifySeller(product.SellerAddress, j)
				response_message := "SUCCESS! Item Bought"
				return &proto.BuyItemResponse{Status: response_message}, nil
			} else {
				return &proto.BuyItemResponse{Status: "FAIL! (Sold Out!)"}, nil
			}
		}
	}
	return &proto.BuyItemResponse{Status: "FAIL! (Item-ID Not found)"}, nil
}

func (s *server) AddToWishlist(ctx context.Context, request *proto.AddToWishlistRequest) (*proto.AddToWishlistResponse, error) {
	fmt.Println("❗ Add to wishlist request from buyer:", request.Listening, "for item:", request.ItemId)
	fmt.Println("---------------------------------")
	item_id := request.ItemId
	for j, product := range list_of_products {
		if product.ItemID == item_id {
			list_of_products[j].Wishlist = append(list_of_products[j].Wishlist, request.Listening)
			response_message := "SUCCESS! Item Added to Wishlist"
			return &proto.AddToWishlistResponse{Status: response_message}, nil
		}
	}
	return &proto.AddToWishlistResponse{Status: "FAIL! (Item-ID Not Found)"}, nil
}

func (s *server) RateItem(ctx context.Context, request *proto.RateItemRequest) (*proto.RateItemResponse, error) {
	fmt.Println("❗ Rate item request from buyer:", request.Listening, "for item:", request.ItemId)
	for index, product := range list_of_products {
		if product.ItemID == request.ItemId {
			for _, ratedby := range product.RatedBy {
				if ratedby == request.Listening {
					return &proto.RateItemResponse{Status: "FAIL! (Already Rated)"}, nil
				}
			}
			current_rating := product.Rating
			new_rating := (current_rating + float32(request.Rating)) / (float32(len(product.RatedBy)) + 1)
			product.RatedBy = append(product.RatedBy, request.Listening)
			fmt.Println("New Rating:", new_rating)
			list_of_products[index].Rating = new_rating
			response_message := "SUCCESS! Item Rated"
			return &proto.RateItemResponse{Status: response_message}, nil
		}
	}
	return &proto.RateItemResponse{Status: "FAIL! (Item-ID Not Found)"}, nil
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
