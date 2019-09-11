package mock_proto_gen

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch07/grpc-docker/go/proto-gen"
	"testing"
	"time"
)


// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}



func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moclProdInfoClient := NewMockProductInfoClient(ctrl)

	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
	price := float32(700.0)
	req := &pb.Product{Name: name, Description: description, Price: price}


	moclProdInfoClient.
		EXPECT().AddProduct(gomock.Any(), &rpcMsg{msg: req},).
		Return(&wrapper.StringValue{Value: "Product:" + name}, nil)

	testAddProduct(t, moclProdInfoClient)


}

func testAddProduct(t *testing.T, client pb.ProductInfoClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
	price := float32(700.0)

	r, err := client.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})

	if err != nil || r.GetValue() != "Product:Sumsung dS10" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.GetValue())


}