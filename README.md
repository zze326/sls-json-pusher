看 help 就会用~

```bash
$ ./sls-json-pusher --help
将 JSON 数据推送到阿里云 SLS LogStore。
例 1 - 推送 json 对象：
        sls-json-pusher --access-key-id xxx --access-key-secret xxx --endpoint cn-hangzhou.log.aliyuncs.com --project test-project --log-store test-log-store '{"username": "张三", "age": 18}'
例 2 - 推送 json 数组：
        sls-json-pusher --access-key-id xxx --access-key-secret xxx --endpoint cn-hangzhou.log.aliyuncs.com --project test-project --log-store test-log-store '[{"username": "张三", "age": 18}, {"username":"李四", "age": 20}]'

Usage:
  sls-log-pusher [flags]

Flags:
      --access-key-id string       阿里云 Access Key Id
      --access-key-secret string   阿里云 Access Key Secret
      --base64                     JSON 是否使用 Base64 编码
      --endpoint string            阿里云日志服务端点，例如：cn-hangzhou.log.aliyuncs.com
  -h, --help                       help for sls-log-pusher
      --log-store string           日志库名称
      --project string             项目名称
      --topic string               主题，可用于归档分组日志 (default "sls-json-pusher")
```
