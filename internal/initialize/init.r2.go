package initialize

import (
	"schedule_gateway/global"

	"github.com/thanvuc/go-core-lib/storage"
)

func InitR2Client() {
	r2Config := global.Config.R2
	R2Clients, err := storage.NewClient(storage.Config{
		AccountID: r2Config.AccountID,
		Endpoint:  r2Config.Endpoint,
		AccessKey: r2Config.AccessKeyID,
		SecretKey: r2Config.SecrecAccessKey,
		Bucket:    r2Config.BucketName,
		UseSSL:    r2Config.UseSSL,
		PublicURL: r2Config.PublicURL,
	})

	if err != nil {
		panic("Failed to initialize R2 client: " + err.Error())
	}
	global.R2Client = R2Clients
}
