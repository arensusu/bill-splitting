// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bill-splitting-linebot/proto"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	writer := &lumberjack.Logger{
		Filename:  "/var/log/app.log",
		MaxSize:   100, // megabytes
		LocalTime: true,
	}
	logger := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)
}

func main() {
	bot, err := linebot.New(os.Getenv("LINEBOT_SECRET"), os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		slog.Error("linebot init error", slog.Any("error", err))
	}

	client, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		slog.Error("NewMessagingApiAPI err:", slog.Any("error", err))
	}

	grpcServerHost := os.Getenv("GRPC_SERVER_HOST")
	grpcServerPort := os.Getenv("GRPC_SERVER_PORT")
	grpcServerAddr := fmt.Sprintf("%s:%s", grpcServerHost, grpcServerPort)
	conn, err := grpc.NewClient(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("grpc conn err:", slog.Any("error", err))
		return
	}
	defer conn.Close()
	grpcClient := proto.NewBillSplittingClient(conn)

	server := NewLineBotServer(bot, client, grpcClient)

	http.HandleFunc("/callback", server.callbackHandler)

	port := os.Getenv("LINEBOT_PORT")
	server.Start(fmt.Sprintf(":%s", port))
}
