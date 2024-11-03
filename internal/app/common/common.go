package common

import (
	"awesomeProject/config"
	"awesomeProject/internal/app/constant"
	"fmt"
	"time"
)

func CheckExpiredContracts(init *config.Initialize) {
	fmt.Println("Checking expired contracts")
	now := time.Now()
	contracts, err := init.ContractRepo.GetAll(nil)
	if err != nil {
		return
	}

	for _, contract := range contracts {
		expired := contract.StartDate.AddDate(0, contract.RentalDuration, 0)
		if contract.Status == constant.LIQUIDITY_COMPLETED && expired.Before(now) {
			fmt.Println("Contract ID", contract.ID, "is expired")
			contract.Status = constant.CONTRACT_FINISHED
			_, err := init.ContractRepo.Update(contract)
			if err != nil {
				return
			}
		} else if contract.Status == constant.CONTRACT_ACTIVE && expired.Before(now) {
			fmt.Println("Contract ID", contract.ID, "is expired")
			contract.Status = constant.WAITING_FOR_LIQUIDITY
			_, err := init.ContractRepo.Update(contract)
			if err != nil {
				return
			}
		}
	}
}
