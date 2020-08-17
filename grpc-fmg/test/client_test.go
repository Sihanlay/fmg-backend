package test

//import (
//	"context"
//	"fmt"
//	"google.golang.org/grpc"
//	"testing"
//)
//
//func TestClient(t *testing.T) {
//	conn, _ := grpc.Dial("127.0.0.1:6000", grpc.WithInsecure())
//
//	v := Address.NewAddressInClient(conn)
//	fmt.Println(v.CreatAddress(context.TODO(), &Address.CreatAddressRequest{
//		Detail: "guanwf",
//		CityId: 1,
//		CountryId: 1,
//		ProvinceId: 0,
//		DistrictId: 11,
//	}))

//}