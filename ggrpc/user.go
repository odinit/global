package ggrpc

import (
	"context"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var confMap map[string]interface{}

func init() {
	file, err := os.Open("../config/config.yaml")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	confMap = make(map[string]interface{}, 10)
	err = yaml.Unmarshal(bytes, &confMap)
	if err != nil {
		panic(err)
	}
}

func CheckLogin(token string) (*AuthResponse, error) {
	conn, err := grpc.Dial(confMap["grpc_url"].(string), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用user.pb.go中的NewAuthServiceClient方法
	userServiceClient := NewUserServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetAuthInfo方法
	resp, err := userServiceClient.GetAuthInfo(context.Background(), &AuthRequest{Token: token})
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetNetSpaceByUserUUID(uuid string) (*NetSpaceResponse, error) {
	conn, err := grpc.Dial(confMap["grpc_url"].(string), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	// 退出时关闭链接
	defer conn.Close()

	NetSpaceServiceClient := NewNetSpaceServiceClient(conn)

	resp, err := NetSpaceServiceClient.GetNetSpaceByUserUUID(context.Background(), &NetSpaceRequest{UserUUID: uuid})
	if err != nil {
		return nil, err
	}
	return resp, err
}
