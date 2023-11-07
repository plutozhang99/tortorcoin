package service

import (
	"tortorCoin/internal/utils"
	"tortorCoin/model"
)

// CreatePair creates a new pair.
func CreatePair(firstUserName string, secondUserName string) (*model.Pair, error) {
	firstUser, errFirst := model.GetPairByUserName(firstUserName)
	secondUser, errSecond := model.GetPairByUserName(secondUserName)
	if firstUser != nil || secondUser != nil {
		return nil, utils.ErrPairAlreadyExists
	}
	if errFirst != nil {
		return nil, errFirst
	}
	if errSecond != nil {
		return nil, errSecond
	}
	newPair, err := model.CreatePair(firstUserName, secondUserName)
	if err != nil {
		return nil, err
	}
	return newPair, nil
}
