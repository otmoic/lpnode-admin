package adminapiservice

import (
	tokenmanager "admin-panel/gen/token_manager"
	database "admin-panel/mongo_database"
	redisbus "admin-panel/redis_bus"
	"admin-panel/service"
	"admin-panel/types"
	"admin-panel/utils"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/smithy-go/ptr"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// tokenManager service example implementation.
// The example methods log the requests and return zero values.
type tokenManagersrvc struct {
	logger *log.Logger
}

// NewTokenManager returns the tokenManager service implementation.
func NewTokenManager(logger *log.Logger) tokenmanager.Service {
	return &tokenManagersrvc{logger}
}

// TokenList implements tokenList.
func (s *tokenManagersrvc) TokenList(ctx context.Context) (res *tokenmanager.TokenListResult, err error) {
	res = &tokenmanager.TokenListResult{
		Code:    ptr.Int64(0),
		Message: ptr.String(""),
	}
	tmls := service.NewTokenManagerLogicService()
	tokenList, err := tmls.ListAll(bson.M{})
	for _, tokenItem := range tokenList {
		res.Result = append(res.Result, &tokenmanager.TokenItem{
			ID:         ptr.String(tokenItem.ID.Hex()),
			ChainID:    tokenItem.ChainId,
			Address:    tokenItem.Address,
			TokenName:  ptr.String(tokenItem.TokenName),
			MarketName: tokenItem.MarketName,
			Precision:  tokenItem.Precision,
			CoinType:   tokenItem.CoinType,
			ChainType:  ptr.String(tokenItem.ChainType),
		})
	}
	s.logger.Print("tokenManager.tokenList")
	return
}

// TokenCreate implements tokenCreate.
func (s *tokenManagersrvc) TokenCreate(ctx context.Context, p *tokenmanager.TokenItem) (res *tokenmanager.TokenCreateResult, err error) {
	res = &tokenmanager.TokenCreateResult{}
	chainRow := struct {
		ChainType string `bson:"chainType"`
	}{}
	err = database.FindOne("main", "chainList", bson.M{"chainId": p.ChainID}, &chainRow)
	if err != nil {
		return
	}
	if chainRow.ChainType == "" {
		err = errors.WithMessage(utils.GetNoEmptyError(err), "chainType is empty")
		return
	}
	uniqAddress, err := utils.GetUniqAddress(p.Address, chainRow.ChainType)
	if err != nil {
		return
	}
	filter := bson.M{
		"addressLower": uniqAddress,
		"chainId":      p.ChainID,
	}
	dataSet := bson.M{
		"$set": bson.M{
			"tokenId":      p.TokenID, // Near only
			"chainId":      p.ChainID,
			"addressIndex": strings.ToLower(p.Address),
			"address":      p.Address,
			"tokenName":    p.TokenName,
			"marketName":   p.MarketName,
			"precision":    p.Precision,
			"coinType":     p.CoinType,
			"chainType":    chainRow.ChainType,
		},
	}
	updateResult, err := database.FindOneAndUpdate("main", "tokens", filter, dataSet)
	if err != nil {
		return
	}
	setCount := updateResult.MatchedCount
	log.Println(setCount)
	res.Code = ptr.Int64(0)
	res.Message = ptr.String("")
	res.Result = ptr.Int64(setCount)
	redisbus.GetRedisBus().PublishEvent("LP_SYSTEM_Notice", `{"type":"tokenCreate","payload":"{}"}`)
	s.logger.Print("tokenManager.tokenCreate")
	return
}

// TokenDelete implements tokenDelete.
func (s *tokenManagersrvc) TokenDelete(ctx context.Context, p *tokenmanager.DeleteTokenFilter) (res *tokenmanager.TokenDeleteResult, err error) {
	res = &tokenmanager.TokenDeleteResult{
		Code:    ptr.Int64(0),
		Message: ptr.String(""),
	}

	objectId, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		err = errors.WithMessage(err, "generate objectID failed")
		return
	}
	v := struct {
		Id         primitive.ObjectID `bson:"_id"`
		BridgeName string             `bson:"bridgeName"`
		AmmName    string             `bson:"ammName"`
	}{}
	err = database.FindOne("main", "bridges", bson.M{
		"$or": bson.A{
			bson.M{"srcToken_id": objectId},
			bson.M{"dstToken_id": objectId}},
	}, &v)
	if err != nil {
		err = errors.WithMessage(err, "query bridges failed")
		return
	}
	if v.Id.Hex() != types.MongoEmptyIdHex {
		err = errors.WithMessage(utils.GetNoEmptyError(err), fmt.Sprintf("bridge already using this token, amm: %s, bridge: %s", v.AmmName, v.BridgeName))
		return
	}

	delCount, err := database.DeleteOne("main", "tokens", bson.M{
		"_id": objectId,
	})
	if delCount <= 0 {
		err = errors.New("cannot find record to delete, no operation")
		return
	}
	redisbus.GetRedisBus().PublishEvent("LP_SYSTEM_Notice", `{"type":"tokenDelete","payload":"{}"}`)
	res.Result = ptr.Int64(delCount)
	log.Println(objectId)
	s.logger.Print("tokenManager.tokenDelete")
	return
}
