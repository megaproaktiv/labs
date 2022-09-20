package labs
import (
	"context"
	"fmt"
	"log"
	"os"
    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)



func identity() *(sts.GetCallerIdentityOutput){
    var err error;
    cfg, err := config.LoadDefaultConfig(context.TODO())
     if err != nil {
        panic(fmt.Sprintf("failed loading config, %v", err))
     }
     client := sts.NewFromConfig(cfg)

	identity, err := client.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal("Error getting identity",err)
    }
    return identity
}

func cfg() ( aws.Config) {
    var err error;
    cfg, err := config.LoadDefaultConfig(context.TODO())
     if err != nil {
        panic(fmt.Sprintf("failed loading config, %v", err))
     }
     
    return cfg
}

// Getenv with fallback
func Getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

// GetRegion Current Region
func GetRegion()(string){
    cfg := cfg()
    region := cfg.Region
    return region;
}

// GetAccount Current Account
func GetAccount()(string){
    return *identity().Account
}