// Code generated by go-merge-types. DO NOT EDIT.

package contract_forta_staking

import (
	import_fmt "fmt"
	import_sync "sync"


	fortastaking011 "github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_1"

	fortastaking012 "github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_2"



	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	"github.com/ethereum/go-ethereum/common"

)

// FortaStakingCaller is a new type which can multiplex calls to different implementation types.
type FortaStakingCaller struct {

	typ0 *fortastaking011.FortaStakingCaller

	typ1 *fortastaking012.FortaStakingCaller

	currTag string
	mu import_sync.RWMutex
	unsafe bool // default: false
}

// NewFortaStakingCaller creates a new merged type.
func NewFortaStakingCaller(address common.Address, caller bind.ContractCaller) (*FortaStakingCaller, error) {
	var (
		mergedType FortaStakingCaller
		err error
	)
	mergedType.currTag = "0.1.1"


	mergedType.typ0, err = fortastaking011.NewFortaStakingCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize fortastaking011.FortaStakingCaller: %v", err)
	}

	mergedType.typ1, err = fortastaking012.NewFortaStakingCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize fortastaking012.FortaStakingCaller: %v", err)
	}


	return &mergedType, nil
}

// IsKnownTagForFortaStakingCaller tells if given tag is a known tag.
func IsKnownTagForFortaStakingCaller(tag string) bool {

	if tag == "0.1.1" {
		return true
	}

	if tag == "0.1.2" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *FortaStakingCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForFortaStakingCaller(tag) {
		tag = "0.1.1"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *FortaStakingCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *FortaStakingCaller) Safe() {
	merged.unsafe = false
}




// ActiveSharesToStake multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) ActiveSharesToStake(opts *bind.CallOpts, activeSharesId *big.Int, amount *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.ActiveSharesToStake(opts, activeSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.ActiveSharesToStake(opts, activeSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.ActiveSharesToStake not implemented (tag=%s)", merged.currTag)
	return
}



// ActiveStakeFor multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) ActiveStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.ActiveStakeFor(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.ActiveStakeFor(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.ActiveStakeFor not implemented (tag=%s)", merged.currTag)
	return
}



// AvailableReward multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) AvailableReward(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.AvailableReward(opts, subjectType, subject, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.AvailableReward not implemented (tag=%s)", merged.currTag)
	return
}



// BalanceOf multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.BalanceOf(opts, account, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.BalanceOf(opts, account, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.BalanceOf not implemented (tag=%s)", merged.currTag)
	return
}



// BalanceOfBatch multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) (retVal []*big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.BalanceOfBatch(opts, accounts, ids)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.BalanceOfBatch(opts, accounts, ids)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.BalanceOfBatch not implemented (tag=%s)", merged.currTag)
	return
}



// Exists multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) Exists(opts *bind.CallOpts, id *big.Int) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.Exists(opts, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.Exists(opts, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.Exists not implemented (tag=%s)", merged.currTag)
	return
}



// InactiveSharesOf multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) InactiveSharesOf(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.InactiveSharesOf(opts, subjectType, subject, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.InactiveSharesOf(opts, subjectType, subject, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.InactiveSharesOf not implemented (tag=%s)", merged.currTag)
	return
}



// InactiveSharesToStake multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) InactiveSharesToStake(opts *bind.CallOpts, inactiveSharesId *big.Int, amount *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.InactiveSharesToStake(opts, inactiveSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.InactiveSharesToStake(opts, inactiveSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.InactiveSharesToStake not implemented (tag=%s)", merged.currTag)
	return
}



// InactiveStakeFor multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) InactiveStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.InactiveStakeFor(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.InactiveStakeFor(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.InactiveStakeFor not implemented (tag=%s)", merged.currTag)
	return
}



// IsApprovedForAll multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.IsApprovedForAll(opts, account, operator)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.IsApprovedForAll(opts, account, operator)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.IsApprovedForAll not implemented (tag=%s)", merged.currTag)
	return
}



// IsFrozen multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) IsFrozen(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.IsFrozen(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.IsFrozen(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.IsFrozen not implemented (tag=%s)", merged.currTag)
	return
}



// IsTrustedForwarder multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.IsTrustedForwarder(opts, forwarder)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.IsTrustedForwarder(opts, forwarder)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.IsTrustedForwarder not implemented (tag=%s)", merged.currTag)
	return
}



// ProxiableUUID multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}



// SharesOf multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) SharesOf(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.SharesOf(opts, subjectType, subject, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.SharesOf(opts, subjectType, subject, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.SharesOf not implemented (tag=%s)", merged.currTag)
	return
}



// StakeToActiveShares multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) StakeToActiveShares(opts *bind.CallOpts, activeSharesId *big.Int, amount *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.StakeToActiveShares(opts, activeSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.StakeToActiveShares(opts, activeSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.StakeToActiveShares not implemented (tag=%s)", merged.currTag)
	return
}



// StakeToInactiveShares multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) StakeToInactiveShares(opts *bind.CallOpts, inactiveSharesId *big.Int, amount *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.StakeToInactiveShares(opts, inactiveSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.StakeToInactiveShares(opts, inactiveSharesId, amount)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.StakeToInactiveShares not implemented (tag=%s)", merged.currTag)
	return
}



// StakedToken multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) StakedToken(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.StakedToken(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.StakedToken(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.StakedToken not implemented (tag=%s)", merged.currTag)
	return
}



// SupportsInterface multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.SupportsInterface(opts, interfaceId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.SupportsInterface(opts, interfaceId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.SupportsInterface not implemented (tag=%s)", merged.currTag)
	return
}



// TotalActiveStake multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) TotalActiveStake(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.TotalActiveStake(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.TotalActiveStake(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.TotalActiveStake not implemented (tag=%s)", merged.currTag)
	return
}



// TotalInactiveShares multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) TotalInactiveShares(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.TotalInactiveShares(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.TotalInactiveShares(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.TotalInactiveShares not implemented (tag=%s)", merged.currTag)
	return
}



// TotalInactiveStake multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) TotalInactiveStake(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.TotalInactiveStake(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.TotalInactiveStake(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.TotalInactiveStake not implemented (tag=%s)", merged.currTag)
	return
}



// TotalShares multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) TotalShares(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.TotalShares(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.TotalShares(opts, subjectType, subject)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.TotalShares not implemented (tag=%s)", merged.currTag)
	return
}



// TotalSupply multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.TotalSupply(opts, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.TotalSupply(opts, id)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.TotalSupply not implemented (tag=%s)", merged.currTag)
	return
}



// Treasury multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) Treasury(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.Treasury(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.Treasury(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.Treasury not implemented (tag=%s)", merged.currTag)
	return
}



// Uri multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.Uri(opts, arg0)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.Uri(opts, arg0)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.Uri not implemented (tag=%s)", merged.currTag)
	return
}



// Version multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ0.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}



// MAXSLASHABLEPERCENT multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) MAXSLASHABLEPERCENT(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.MAXSLASHABLEPERCENT(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.MAXSLASHABLEPERCENT not implemented (tag=%s)", merged.currTag)
	return
}



// MAXWITHDRAWALDELAY multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) MAXWITHDRAWALDELAY(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.MAXWITHDRAWALDELAY(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.MAXWITHDRAWALDELAY not implemented (tag=%s)", merged.currTag)
	return
}



// MINWITHDRAWALDELAY multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) MINWITHDRAWALDELAY(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.MINWITHDRAWALDELAY(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.MINWITHDRAWALDELAY not implemented (tag=%s)", merged.currTag)
	return
}



// Allocator multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) Allocator(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.Allocator(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.Allocator not implemented (tag=%s)", merged.currTag)
	return
}



// GetDelegatedSubjectType multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) GetDelegatedSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.GetDelegatedSubjectType(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.GetDelegatedSubjectType not implemented (tag=%s)", merged.currTag)
	return
}



// GetDelegatorSubjectType multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) GetDelegatorSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.GetDelegatorSubjectType(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.GetDelegatorSubjectType not implemented (tag=%s)", merged.currTag)
	return
}



// GetSubjectTypeAgency multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) GetSubjectTypeAgency(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.GetSubjectTypeAgency(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.GetSubjectTypeAgency not implemented (tag=%s)", merged.currTag)
	return
}



// OpenProposals multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) OpenProposals(opts *bind.CallOpts, arg0 *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.OpenProposals(opts, arg0)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.OpenProposals not implemented (tag=%s)", merged.currTag)
	return
}



// SlashDelegatorsPercent multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) SlashDelegatorsPercent(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.SlashDelegatorsPercent(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.SlashDelegatorsPercent not implemented (tag=%s)", merged.currTag)
	return
}



// SubjectGateway multiplexes to different implementations of the method.
func (merged *FortaStakingCaller) SubjectGateway(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.2" {
		val, methodErr := merged.typ1.SubjectGateway(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("FortaStakingCaller.SubjectGateway not implemented (tag=%s)", merged.currTag)
	return
}