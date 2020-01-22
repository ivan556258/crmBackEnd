package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyClient struct {
	mc *mongo.Client
	db *mongo.Database
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func NewMyClient(url, db string) (mc *MyClient, err error) {
	mc = &MyClient{}
	if mc.mc, err = mongo.NewClient(options.Client().ApplyURI(url)); err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mc.mc.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mc.db = mc.mc.Database(db)
	return
}

func main() {
	mc, err := NewMyClient("mongodb://localhost:27017", "crmTaxi")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/insertContractData", mc.insertContractData)
	http.HandleFunc("/updateContractData", mc.updateContractData)
	http.HandleFunc("/deleteContractData", mc.deleteContractData)
	http.HandleFunc("/selectContractData", mc.selectContractData)
	http.HandleFunc("/selectContractDataOne", mc.selectContractDataOne)

	http.HandleFunc("/insertDriverData", mc.insertDriverData)
	http.HandleFunc("/updateDriverData", mc.updateDriverData)
	http.HandleFunc("/deleteDriverData", mc.deleteDriverData)
	http.HandleFunc("/selectDriverData", mc.selectDriverData)
	http.HandleFunc("/selectDriverDataOne", mc.selectDriverDataOne)

	http.HandleFunc("/insertAutomobileData", mc.insertAutomobileData)
	http.HandleFunc("/updateAutomobileData", mc.updateAutomobileData)
	http.HandleFunc("/deleteAutomobileData", mc.deleteAutomobileData)
	http.HandleFunc("/selecAutomobileData", mc.selectAutomobileData)
	http.HandleFunc("/selectAutomobileDataOne", mc.selectAutomobileDataOne)

	http.HandleFunc("/insertTechnicalServiceData", mc.insertTechnicalServiceData)
	http.HandleFunc("/updateTechnicalServiceData", mc.updateTechnicalServiceData)
	http.HandleFunc("/deleteTechnicalServiceData", mc.deleteTechnicalServiceData)
	http.HandleFunc("/selectTechnicalServiceData", mc.selectTechnicalServiceData)
	http.HandleFunc("/selectTechnicalServiceDataOne", mc.selectTechnicalServiceDataOne)

	http.HandleFunc("/insertUserProfileData", mc.insertUserProfileData)
	http.HandleFunc("/updateUserProfileData", mc.updateUserProfileData)
	http.HandleFunc("/deleteUserProfileData", mc.deleteUserProfileData)
	http.HandleFunc("/selectUserProfileData", mc.selectUserProfileData)

	http.HandleFunc("/insertUserCommentData", mc.insertUserCommentData)
	http.HandleFunc("/updateUserCommentData", mc.updateUserCommentData)
	http.HandleFunc("/deleteUserCommentData", mc.deleteUserCommentData)
	http.HandleFunc("/selectUserCommentData", mc.selectUserCommentData)

	http.HandleFunc("/insertUserSmsData", mc.insertUserSmsData)
	http.HandleFunc("/updateUserSmsData", mc.updateUserSmsData)
	http.HandleFunc("/deleteUserSmsData", mc.deleteUserSmsData)
	http.HandleFunc("/selectUserSmsData", mc.selectUserSmsData)

	http.HandleFunc("/updateAccountData", mc.updateAccountData)
	http.HandleFunc("/insertAccountData", mc.insertAccountData)
	http.HandleFunc("/selectAccountData", mc.selectAccountData)
	http.HandleFunc("/deleteAccountData", mc.deleteAccountData)

	http.HandleFunc("/updateAccountBillData", mc.updateAccountBillData)
	http.HandleFunc("/insertAccountBillData", mc.insertAccountBillData)
	http.HandleFunc("/selectAccountBillData", mc.selectAccountBillData)
	http.HandleFunc("/selectAccountBillDataOne", mc.selectAccountBillDataOne)
	http.HandleFunc("/deleteAccountBillData", mc.deleteAccountBillData)

	http.HandleFunc("/updateAccountBillItemData", mc.updateAccountBillItemData)
	http.HandleFunc("/insertAccountBillItemData", mc.insertAccountBillItemData)
	http.HandleFunc("/selectAccountBillItemData", mc.selectAccountBillItemData)
	http.HandleFunc("/selectAccountBillItemDataOne", mc.selectAccountBillItemDataOne)
	http.HandleFunc("/deleteAccountBillItemData", mc.deleteAccountBillItemData)

	http.HandleFunc("/updateTariffData", mc.updateTariffData)
	http.HandleFunc("/insertTariffData", mc.insertTariffData)
	http.HandleFunc("/selectTariffData", mc.selectTariffData)
	http.HandleFunc("/selectTariffDataOne", mc.selectTariffDataOne)
	http.HandleFunc("/deleteTariffData", mc.deleteTariffData)

	http.HandleFunc("/updateTransactionData", mc.updateTransactionData)
	http.HandleFunc("/insertTransactionData", mc.insertTransactionData)
	http.HandleFunc("/selectTransactionData", mc.selectTransactionData)
	http.HandleFunc("/selectTransactionDataOne", mc.selectTransactionDataOne)
	http.HandleFunc("/deleteTransactionData", mc.deleteTransactionData)

	http.HandleFunc("/updatePenaltyData", mc.updatePenaltyData)
	http.HandleFunc("/insertPenaltyData", mc.insertPenaltyData)
	http.HandleFunc("/selectPenaltyData", mc.selectPenaltyData)
	http.HandleFunc("/selectPenaltyDataOne", mc.selectPenaltyDataOne)
	http.HandleFunc("/deletePenaltyData", mc.deletePenaltyData)

	http.HandleFunc("/updateNomenclatureData", mc.updateNomenclatureData)
	http.HandleFunc("/insertNomenclatureData", mc.insertNomenclatureData)
	http.HandleFunc("/selectNomenclatureData", mc.selectNomenclatureData)
	http.HandleFunc("/selectNomenclatureDataOne", mc.selectNomenclatureDataOne)
	http.HandleFunc("/deleteNomenclatureData", mc.deleteNomenclatureData)

	http.HandleFunc("/updateCounteragentData", mc.updateCounteragentData)
	http.HandleFunc("/insertCounteragentData", mc.insertCounteragentData)
	http.HandleFunc("/selectCounteragentData", mc.selectCounteragentData)
	http.HandleFunc("/selectCounteragentDataOne", mc.selectCounteragentDataOne)
	http.HandleFunc("/deleteCounteragentData", mc.deleteCounteragentData)

	http.HandleFunc("/updateReceptionNotesData", mc.updateReceptionNoteData)
	http.HandleFunc("/insertReceptionNotesData", mc.insertReceptionNoteData)
	http.HandleFunc("/selectReceptionNotesData", mc.selectReceptionNoteData)
	http.HandleFunc("/selectReceptionNotesDataOne", mc.selectReceptionNoteDataOne)
	http.HandleFunc("/deleteReceptionNotesData", mc.deleteReceptionNoteData)

	http.HandleFunc("/updateOverheadConsumablesData", mc.updateOverheadConsumableData)
	http.HandleFunc("/insertOverheadConsumablesData", mc.insertOverheadConsumableData)
	http.HandleFunc("/selectOverheadConsumablesData", mc.selectOverheadConsumableData)
	http.HandleFunc("/selectOverheadConsumablesDataOne", mc.selectOverheadConsumableDataOne)
	http.HandleFunc("/deleteOverheadConsumablesData", mc.deleteOverheadConsumableData)

	http.HandleFunc("/insertUserCompanyData", mc.insertUserCompanyData)
	http.HandleFunc("/saveUserCompanyData", mc.saveUserCompanyData)
	http.HandleFunc("/selectUserCompanyData", mc.selectUserCompanyDataOne)

	http.HandleFunc("/insertOwnerData", mc.insertOwnerData)
	http.HandleFunc("/updateOwnerData", mc.updateOwnerData)
	http.HandleFunc("/selectOwnerData", mc.selectOwnerData)
	http.HandleFunc("/selectAccountDataId", mc.selectAccountDataId)
	http.HandleFunc("/deleteOwnerData", mc.deleteOwnerData)

	http.HandleFunc("/insertAuthData", mc.insertAuthData)
	http.HandleFunc("/updateAuthData", mc.updateAuthData)
	http.HandleFunc("/checkAuthData", mc.checkAuthData)
	http.HandleFunc("/resetAuthData", mc.resetAuthData)
	http.HandleFunc("/deleteAuthData", mc.deleteAuthData)
	panic(http.ListenAndServe(":8081", nil))
}
