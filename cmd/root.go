package cmd

import (
	"encoding/base64"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"os"
	"time"
)

var (
	accessKeyId     string
	accessKeySecret string
	endpoint        string
	project         string
	logStore        string
	topic           string
	isBase64        bool
	rootCmd         = &cobra.Command{
		Use:   "sls-log-pusher",
		Short: "阿里云日志推送工具",
		Long: `将 JSON 数据推送到阿里云 SLS LogStore。
例 1 - 推送 json 对象：
	sls-json-pusher --access-key-id xxx --access-key-secret xxx --endpoint cn-hangzhou.log.aliyuncs.com --project test-project --log-store test-log-store '{"username": "张三", "age": 18}'
例 2 - 推送 json 数组：
	sls-json-pusher --access-key-id xxx --access-key-secret xxx --endpoint cn-hangzhou.log.aliyuncs.com --project test-project --log-store test-log-store '[{"username": "张三", "age": 18}, {"username":"李四", "age": 20}]'`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			provider := sls.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")
			client := sls.CreateNormalInterfaceV2(endpoint, provider)
			jsonData := args[0]
			if isBase64 {
				jsonBytes, err := base64.StdEncoding.DecodeString(jsonData)
				if err != nil {
					fmt.Println("解码 base64 字符串出错, err: ", err)
					os.Exit(1)
				}
				jsonData = string(jsonBytes)
			}
			if !gjson.Valid(jsonData) {
				fmt.Println("Json 格式校验不通过")
				os.Exit(1)
			}

			jsonResult := gjson.Parse(jsonData)
			var (
				logs      = make([]*sls.Log, 0)
				jsonToLog = func(m map[string]gjson.Result) *sls.Log {
					content := make([]*sls.LogContent, 0)
					for k, v := range m {
						content = append(content, &sls.LogContent{
							Key:   proto.String(k),
							Value: proto.String(fmt.Sprintf("%v", v.Value())),
						})
					}
					return &sls.Log{
						Time:     proto.Uint32(uint32(time.Now().Unix())),
						Contents: content,
					}
				}
			)
			if jsonResult.IsArray() {
				for _, jsonObj := range jsonResult.Array() {
					logs = append(logs, jsonToLog(jsonObj.Map()))
				}
			} else {
				logs = append(logs, jsonToLog(jsonResult.Map()))
			}

			if err := client.PutLogs(project, logStore, &sls.LogGroup{
				Topic: proto.String(topic),
				Logs:  logs,
			}); err != nil {
				(fmt.Println("推送出错：", err))
				os.Exit(1)
			}
			fmt.Println("推送成功")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&accessKeyId, "access-key-id", "", "阿里云 Access Key Id")
	rootCmd.PersistentFlags().StringVar(&accessKeySecret, "access-key-secret", "", "阿里云 Access Key Secret")
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "阿里云日志服务端点，例如：cn-hangzhou.log.aliyuncs.com")
	rootCmd.Flags().BoolVar(&isBase64, "base64", false, "JSON 是否使用 Base64 编码")
	rootCmd.Flags().StringVar(&project, "project", "", "项目名称")
	rootCmd.Flags().StringVar(&logStore, "log-store", "", "日志库名称")
	rootCmd.Flags().StringVar(&topic, "topic", "sls-json-pusher", "主题，可用于归档分组日志")
	rootCmd.MarkPersistentFlagRequired("access-key-id")
	rootCmd.MarkPersistentFlagRequired("access-key-secret")
	rootCmd.MarkPersistentFlagRequired("endpoint")
	rootCmd.MarkFlagRequired("project")
	rootCmd.MarkFlagRequired("log-store")
	rootCmd.MarkFlagRequired("data")
}
