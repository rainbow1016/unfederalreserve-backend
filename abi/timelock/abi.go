// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abitimelock

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbitimelockABI is the input ABI used to generate the binding from.
const AbitimelockABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"CancelTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"NewDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAXIMUM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"queueTransaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queuedTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingAdmin_\",\"type\":\"address\"}],\"name\":\"setPendingAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbitimelockBin is the compiled bytecode used for deploying new contracts.
var AbitimelockBin = "0x608060405234801561001057600080fd5b506040516118a33803806118a38339818101604052604081101561003357600080fd5b508051602090910151610e1081101561007d5760405162461bcd60e51b81526004018080602001828103825260378152602001806118346037913960400191505060405180910390fd5b62278d008111156100bf5760405162461bcd60e51b815260040180806020018281038252603881526020018061186b6038913960400191505060405180910390fd5b600080546001600160a01b039093166001600160a01b031990931692909217909155600255611741806100f36000396000f3fe6080604052600436106100c25760003560e01c80636a42b8f81161007f578063c1a287e211610059578063c1a287e2146105dd578063e177246e146105f2578063f2b065371461061c578063f851a4401461065a576100c2565b80636a42b8f81461059e5780637d645fab146105b3578063b1b43ae5146105c8576100c2565b80630825f38f146100c45780630e18b68114610279578063267822471461028e5780633a66f901146102bf5780634dd18bf51461041e578063591fcdfe14610451575b005b610204600480360360a08110156100da57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561010957600080fd5b82018360208201111561011b57600080fd5b803590602001918460018302840111600160201b8311171561013c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111600160201b831117156101c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061066f915050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561023e578181015183820152602001610226565b50505050905090810190601f16801561026b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561028557600080fd5b506100c2610b88565b34801561029a57600080fd5b506102a3610c24565b604080516001600160a01b039092168252519081900360200190f35b3480156102cb57600080fd5b5061040c600480360360a08110156102e257600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561031157600080fd5b82018360208201111561032357600080fd5b803590602001918460018302840111600160201b8311171561034457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561039657600080fd5b8201836020820111156103a857600080fd5b803590602001918460018302840111600160201b831117156103c957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610c33915050565b60408051918252519081900360200190f35b34801561042a57600080fd5b506100c26004803603602081101561044157600080fd5b50356001600160a01b0316610f44565b34801561045d57600080fd5b506100c2600480360360a081101561047457600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156104a357600080fd5b8201836020820111156104b557600080fd5b803590602001918460018302840111600160201b831117156104d657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561052857600080fd5b82018360208201111561053a57600080fd5b803590602001918460018302840111600160201b8311171561055b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610fd2915050565b3480156105aa57600080fd5b5061040c611288565b3480156105bf57600080fd5b5061040c61128e565b3480156105d457600080fd5b5061040c611295565b3480156105e957600080fd5b5061040c61129b565b3480156105fe57600080fd5b506100c26004803603602081101561061557600080fd5b50356112a2565b34801561062857600080fd5b506106466004803603602081101561063f57600080fd5b5035611396565b604080519115158252519081900360200190f35b34801561066657600080fd5b506102a36113ab565b6000546060906001600160a01b031633146106bb5760405162461bcd60e51b81526004018080602001828103825260388152602001806114206038913960400191505060405180910390fd5b6000868686868660405160200180866001600160a01b03166001600160a01b031681526020018581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561072a578181015183820152602001610712565b50505050905090810190601f1680156107575780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561078a578181015183820152602001610772565b50505050905090810190601f1680156107b75780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301206000818152600390935291205490995060ff16975061082896505050505050505760405162461bcd60e51b815260040180806020018281038252603d815260200180611573603d913960400191505060405180910390fd5b826108316113ba565b101561086e5760405162461bcd60e51b81526004018080602001828103825260458152602001806114c26045913960600191505060405180910390fd5b610881836212750063ffffffff6113be16565b6108896113ba565b11156108c65760405162461bcd60e51b815260040180806020018281038252603381526020018061148f6033913960400191505060405180910390fd5b6000818152600360205260409020805460ff1916905584516060906108ec575083610979565b85805190602001208560405160200180836001600160e01b0319166001600160e01b031916815260040182805190602001908083835b602083106109415780518252601f199092019160209182019101610922565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290505b60006060896001600160a01b031689846040518082805190602001908083835b602083106109b85780518252601f199092019160209182019101610999565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610a1a576040519150601f19603f3d011682016040523d82523d6000602084013e610a1f565b606091505b509150915081610a605760405162461bcd60e51b815260040180806020018281038252603d815260200180611656603d913960400191505060405180910390fd5b896001600160a01b0316847fa560e3198060a2f10670c1ec5b403077ea6ae93ca8de1c32b451dc1a943cd6e78b8b8b8b604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610add578181015183820152602001610ac5565b50505050905090810190601f168015610b0a5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610b3d578181015183820152602001610b25565b50505050905090810190601f168015610b6a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a39998505050505050505050565b6001546001600160a01b03163314610bd15760405162461bcd60e51b81526004018080602001828103825260388152602001806115b06038913960400191505060405180910390fd5b60008054336001600160a01b031991821617808355600180549092169091556040516001600160a01b03909116917f71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c91a2565b6001546001600160a01b031681565b600080546001600160a01b03163314610c7d5760405162461bcd60e51b81526004018080602001828103825260368152602001806116206036913960400191505060405180910390fd5b610c97600254610c8b6113ba565b9063ffffffff6113be16565b821015610cd55760405162461bcd60e51b81526004018080602001828103825260498152602001806116936049913960600191505060405180910390fd5b6000868686868660405160200180866001600160a01b03166001600160a01b031681526020018581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610d44578181015183820152602001610d2c565b50505050905090810190601f168015610d715780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610da4578181015183820152602001610d8c565b50505050905090810190601f168015610dd15780820380516001836020036101000a031916815260200191505b5097505050505050505060405160208183030381529060405280519060200120905060016003600083815260200190815260200160002060006101000a81548160ff021916908315150217905550866001600160a01b0316817f76e2796dc3a81d57b0e8504b647febcbeeb5f4af818e164f11eef8131a6a763f88888888604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610e9c578181015183820152602001610e84565b50505050905090810190601f168015610ec95780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610efc578181015183820152602001610ee4565b50505050905090810190601f168015610f295780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a39695505050505050565b333014610f825760405162461bcd60e51b81526004018080602001828103825260388152602001806115e86038913960400191505060405180910390fd5b600180546001600160a01b0319166001600160a01b0383811691909117918290556040519116907f69d78e38a01985fbb1462961809b4b2d65531bc93b2b94037f3334b82ca4a75690600090a250565b6000546001600160a01b0316331461101b5760405162461bcd60e51b81526004018080602001828103825260378152602001806114586037913960400191505060405180910390fd5b6000858585858560405160200180866001600160a01b03166001600160a01b031681526020018581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561108a578181015183820152602001611072565b50505050905090810190601f1680156110b75780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156110ea5781810151838201526020016110d2565b50505050905090810190601f1680156111175780820380516001836020036101000a031916815260200191505b5097505050505050505060405160208183030381529060405280519060200120905060006003600083815260200190815260200160002060006101000a81548160ff021916908315150217905550856001600160a01b0316817f2fffc091a501fd91bfbff27141450d3acb40fb8e6d8382b243ec7a812a3aaf8787878787604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156111e25781810151838201526020016111ca565b50505050905090810190601f16801561120f5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561124257818101518382015260200161122a565b50505050905090810190601f16801561126f5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a3505050505050565b60025481565b62278d0081565b610e1081565b6212750081565b3330146112e05760405162461bcd60e51b81526004018080602001828103825260318152602001806116dc6031913960400191505060405180910390fd5b610e108110156113215760405162461bcd60e51b81526004018080602001828103825260348152602001806115076034913960400191505060405180910390fd5b62278d008111156113635760405162461bcd60e51b815260040180806020018281038252603881526020018061153b6038913960400191505060405180910390fd5b600281905560405181907f948b1f6a42ee138b7e34058ba85a37f716d55ff25ff05a763f15bed6a04c8d2c90600090a250565b60036020526000908152604090205460ff1681565b6000546001600160a01b031681565b4290565b600082820183811015611418576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe54696d656c6f636b3a3a657865637574655472616e73616374696f6e3a2043616c6c206d75737420636f6d652066726f6d2061646d696e2e54696d656c6f636b3a3a63616e63656c5472616e73616374696f6e3a2043616c6c206d75737420636f6d652066726f6d2061646d696e2e54696d656c6f636b3a3a657865637574655472616e73616374696f6e3a205472616e73616374696f6e206973207374616c652e54696d656c6f636b3a3a657865637574655472616e73616374696f6e3a205472616e73616374696f6e206861736e2774207375727061737365642074696d65206c6f636b2e54696d656c6f636b3a3a73657444656c61793a2044656c6179206d75737420657863656564206d696e696d756d2064656c61792e54696d656c6f636b3a3a73657444656c61793a2044656c6179206d757374206e6f7420657863656564206d6178696d756d2064656c61792e54696d656c6f636b3a3a657865637574655472616e73616374696f6e3a205472616e73616374696f6e206861736e2774206265656e207175657565642e54696d656c6f636b3a3a61636365707441646d696e3a2043616c6c206d75737420636f6d652066726f6d2070656e64696e6741646d696e2e54696d656c6f636b3a3a73657450656e64696e6741646d696e3a2043616c6c206d75737420636f6d652066726f6d2054696d656c6f636b2e54696d656c6f636b3a3a71756575655472616e73616374696f6e3a2043616c6c206d75737420636f6d652066726f6d2061646d696e2e54696d656c6f636b3a3a657865637574655472616e73616374696f6e3a205472616e73616374696f6e20657865637574696f6e2072657665727465642e54696d656c6f636b3a3a71756575655472616e73616374696f6e3a20457374696d6174656420657865637574696f6e20626c6f636b206d75737420736174697366792064656c61792e54696d656c6f636b3a3a73657444656c61793a2043616c6c206d75737420636f6d652066726f6d2054696d656c6f636b2ea265627a7a72315820c91671fb8f6ceb84a782f41e4c13db179394d3b72afef8731e5b8191b60c6ebc64736f6c6343000510003254696d656c6f636b3a3a636f6e7374727563746f723a2044656c6179206d75737420657863656564206d696e696d756d2064656c61792e54696d656c6f636b3a3a73657444656c61793a2044656c6179206d757374206e6f7420657863656564206d6178696d756d2064656c61792e"

// DeployAbitimelock deploys a new Ethereum contract, binding an instance of Abitimelock to it.
func DeployAbitimelock(auth *bind.TransactOpts, backend bind.ContractBackend, admin_ common.Address, delay_ *big.Int) (common.Address, *types.Transaction, *Abitimelock, error) {
	parsed, err := abi.JSON(strings.NewReader(AbitimelockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbitimelockBin), backend, admin_, delay_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abitimelock{AbitimelockCaller: AbitimelockCaller{contract: contract}, AbitimelockTransactor: AbitimelockTransactor{contract: contract}, AbitimelockFilterer: AbitimelockFilterer{contract: contract}}, nil
}

// Abitimelock is an auto generated Go binding around an Ethereum contract.
type Abitimelock struct {
	AbitimelockCaller     // Read-only binding to the contract
	AbitimelockTransactor // Write-only binding to the contract
	AbitimelockFilterer   // Log filterer for contract events
}

// AbitimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbitimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbitimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbitimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbitimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbitimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbitimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbitimelockSession struct {
	Contract     *Abitimelock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbitimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbitimelockCallerSession struct {
	Contract *AbitimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AbitimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbitimelockTransactorSession struct {
	Contract     *AbitimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AbitimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbitimelockRaw struct {
	Contract *Abitimelock // Generic contract binding to access the raw methods on
}

// AbitimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbitimelockCallerRaw struct {
	Contract *AbitimelockCaller // Generic read-only contract binding to access the raw methods on
}

// AbitimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbitimelockTransactorRaw struct {
	Contract *AbitimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbitimelock creates a new instance of Abitimelock, bound to a specific deployed contract.
func NewAbitimelock(address common.Address, backend bind.ContractBackend) (*Abitimelock, error) {
	contract, err := bindAbitimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abitimelock{AbitimelockCaller: AbitimelockCaller{contract: contract}, AbitimelockTransactor: AbitimelockTransactor{contract: contract}, AbitimelockFilterer: AbitimelockFilterer{contract: contract}}, nil
}

// NewAbitimelockCaller creates a new read-only instance of Abitimelock, bound to a specific deployed contract.
func NewAbitimelockCaller(address common.Address, caller bind.ContractCaller) (*AbitimelockCaller, error) {
	contract, err := bindAbitimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbitimelockCaller{contract: contract}, nil
}

// NewAbitimelockTransactor creates a new write-only instance of Abitimelock, bound to a specific deployed contract.
func NewAbitimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*AbitimelockTransactor, error) {
	contract, err := bindAbitimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbitimelockTransactor{contract: contract}, nil
}

// NewAbitimelockFilterer creates a new log filterer instance of Abitimelock, bound to a specific deployed contract.
func NewAbitimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*AbitimelockFilterer, error) {
	contract, err := bindAbitimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbitimelockFilterer{contract: contract}, nil
}

// bindAbitimelock binds a generic wrapper to an already deployed contract.
func bindAbitimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbitimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abitimelock *AbitimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abitimelock.Contract.AbitimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abitimelock *AbitimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abitimelock.Contract.AbitimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abitimelock *AbitimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abitimelock.Contract.AbitimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abitimelock *AbitimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abitimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abitimelock *AbitimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abitimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abitimelock *AbitimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abitimelock.Contract.contract.Transact(opts, method, params...)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Abitimelock *AbitimelockCaller) GRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Abitimelock *AbitimelockSession) GRACEPERIOD() (*big.Int, error) {
	return _Abitimelock.Contract.GRACEPERIOD(&_Abitimelock.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Abitimelock *AbitimelockCallerSession) GRACEPERIOD() (*big.Int, error) {
	return _Abitimelock.Contract.GRACEPERIOD(&_Abitimelock.CallOpts)
}

// MAXIMUMDELAY is a free data retrieval call binding the contract method 0x7d645fab.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockCaller) MAXIMUMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "MAXIMUM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMDELAY is a free data retrieval call binding the contract method 0x7d645fab.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockSession) MAXIMUMDELAY() (*big.Int, error) {
	return _Abitimelock.Contract.MAXIMUMDELAY(&_Abitimelock.CallOpts)
}

// MAXIMUMDELAY is a free data retrieval call binding the contract method 0x7d645fab.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockCallerSession) MAXIMUMDELAY() (*big.Int, error) {
	return _Abitimelock.Contract.MAXIMUMDELAY(&_Abitimelock.CallOpts)
}

// MINIMUMDELAY is a free data retrieval call binding the contract method 0xb1b43ae5.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockCaller) MINIMUMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "MINIMUM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMDELAY is a free data retrieval call binding the contract method 0xb1b43ae5.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockSession) MINIMUMDELAY() (*big.Int, error) {
	return _Abitimelock.Contract.MINIMUMDELAY(&_Abitimelock.CallOpts)
}

// MINIMUMDELAY is a free data retrieval call binding the contract method 0xb1b43ae5.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (_Abitimelock *AbitimelockCallerSession) MINIMUMDELAY() (*big.Int, error) {
	return _Abitimelock.Contract.MINIMUMDELAY(&_Abitimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abitimelock *AbitimelockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abitimelock *AbitimelockSession) Admin() (common.Address, error) {
	return _Abitimelock.Contract.Admin(&_Abitimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abitimelock *AbitimelockCallerSession) Admin() (common.Address, error) {
	return _Abitimelock.Contract.Admin(&_Abitimelock.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Abitimelock *AbitimelockCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Abitimelock *AbitimelockSession) Delay() (*big.Int, error) {
	return _Abitimelock.Contract.Delay(&_Abitimelock.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Abitimelock *AbitimelockCallerSession) Delay() (*big.Int, error) {
	return _Abitimelock.Contract.Delay(&_Abitimelock.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abitimelock *AbitimelockCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abitimelock *AbitimelockSession) PendingAdmin() (common.Address, error) {
	return _Abitimelock.Contract.PendingAdmin(&_Abitimelock.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abitimelock *AbitimelockCallerSession) PendingAdmin() (common.Address, error) {
	return _Abitimelock.Contract.PendingAdmin(&_Abitimelock.CallOpts)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Abitimelock *AbitimelockCaller) QueuedTransactions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Abitimelock.contract.Call(opts, &out, "queuedTransactions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Abitimelock *AbitimelockSession) QueuedTransactions(arg0 [32]byte) (bool, error) {
	return _Abitimelock.Contract.QueuedTransactions(&_Abitimelock.CallOpts, arg0)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Abitimelock *AbitimelockCallerSession) QueuedTransactions(arg0 [32]byte) (bool, error) {
	return _Abitimelock.Contract.QueuedTransactions(&_Abitimelock.CallOpts, arg0)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_Abitimelock *AbitimelockTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_Abitimelock *AbitimelockSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abitimelock.Contract.AcceptAdmin(&_Abitimelock.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_Abitimelock *AbitimelockTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abitimelock.Contract.AcceptAdmin(&_Abitimelock.TransactOpts)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x591fcdfe.
//
// Solidity: function cancelTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns()
func (_Abitimelock *AbitimelockTransactor) CancelTransaction(opts *bind.TransactOpts, target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "cancelTransaction", target, value, signature, data, eta)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x591fcdfe.
//
// Solidity: function cancelTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns()
func (_Abitimelock *AbitimelockSession) CancelTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.CancelTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x591fcdfe.
//
// Solidity: function cancelTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns()
func (_Abitimelock *AbitimelockTransactorSession) CancelTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.CancelTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x0825f38f.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (_Abitimelock *AbitimelockTransactor) ExecuteTransaction(opts *bind.TransactOpts, target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "executeTransaction", target, value, signature, data, eta)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x0825f38f.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (_Abitimelock *AbitimelockSession) ExecuteTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.ExecuteTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x0825f38f.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (_Abitimelock *AbitimelockTransactorSession) ExecuteTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.ExecuteTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0x3a66f901.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (_Abitimelock *AbitimelockTransactor) QueueTransaction(opts *bind.TransactOpts, target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "queueTransaction", target, value, signature, data, eta)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0x3a66f901.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (_Abitimelock *AbitimelockSession) QueueTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.QueueTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// QueueTransaction is a paid mutator transaction binding the contract method 0x3a66f901.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (_Abitimelock *AbitimelockTransactorSession) QueueTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.QueueTransaction(&_Abitimelock.TransactOpts, target, value, signature, data, eta)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_Abitimelock *AbitimelockTransactor) SetDelay(opts *bind.TransactOpts, delay_ *big.Int) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "setDelay", delay_)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_Abitimelock *AbitimelockSession) SetDelay(delay_ *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.SetDelay(&_Abitimelock.TransactOpts, delay_)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (_Abitimelock *AbitimelockTransactorSession) SetDelay(delay_ *big.Int) (*types.Transaction, error) {
	return _Abitimelock.Contract.SetDelay(&_Abitimelock.TransactOpts, delay_)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address pendingAdmin_) returns()
func (_Abitimelock *AbitimelockTransactor) SetPendingAdmin(opts *bind.TransactOpts, pendingAdmin_ common.Address) (*types.Transaction, error) {
	return _Abitimelock.contract.Transact(opts, "setPendingAdmin", pendingAdmin_)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address pendingAdmin_) returns()
func (_Abitimelock *AbitimelockSession) SetPendingAdmin(pendingAdmin_ common.Address) (*types.Transaction, error) {
	return _Abitimelock.Contract.SetPendingAdmin(&_Abitimelock.TransactOpts, pendingAdmin_)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address pendingAdmin_) returns()
func (_Abitimelock *AbitimelockTransactorSession) SetPendingAdmin(pendingAdmin_ common.Address) (*types.Transaction, error) {
	return _Abitimelock.Contract.SetPendingAdmin(&_Abitimelock.TransactOpts, pendingAdmin_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abitimelock *AbitimelockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Abitimelock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abitimelock *AbitimelockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abitimelock.Contract.Fallback(&_Abitimelock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abitimelock *AbitimelockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abitimelock.Contract.Fallback(&_Abitimelock.TransactOpts, calldata)
}

// AbitimelockCancelTransactionIterator is returned from FilterCancelTransaction and is used to iterate over the raw logs and unpacked data for CancelTransaction events raised by the Abitimelock contract.
type AbitimelockCancelTransactionIterator struct {
	Event *AbitimelockCancelTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockCancelTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockCancelTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockCancelTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockCancelTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockCancelTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockCancelTransaction represents a CancelTransaction event raised by the Abitimelock contract.
type AbitimelockCancelTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelTransaction is a free log retrieval operation binding the contract event 0x2fffc091a501fd91bfbff27141450d3acb40fb8e6d8382b243ec7a812a3aaf87.
//
// Solidity: event CancelTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) FilterCancelTransaction(opts *bind.FilterOpts, txHash [][32]byte, target []common.Address) (*AbitimelockCancelTransactionIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "CancelTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockCancelTransactionIterator{contract: _Abitimelock.contract, event: "CancelTransaction", logs: logs, sub: sub}, nil
}

// WatchCancelTransaction is a free log subscription operation binding the contract event 0x2fffc091a501fd91bfbff27141450d3acb40fb8e6d8382b243ec7a812a3aaf87.
//
// Solidity: event CancelTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) WatchCancelTransaction(opts *bind.WatchOpts, sink chan<- *AbitimelockCancelTransaction, txHash [][32]byte, target []common.Address) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "CancelTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockCancelTransaction)
				if err := _Abitimelock.contract.UnpackLog(event, "CancelTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelTransaction is a log parse operation binding the contract event 0x2fffc091a501fd91bfbff27141450d3acb40fb8e6d8382b243ec7a812a3aaf87.
//
// Solidity: event CancelTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) ParseCancelTransaction(log types.Log) (*AbitimelockCancelTransaction, error) {
	event := new(AbitimelockCancelTransaction)
	if err := _Abitimelock.contract.UnpackLog(event, "CancelTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbitimelockExecuteTransactionIterator is returned from FilterExecuteTransaction and is used to iterate over the raw logs and unpacked data for ExecuteTransaction events raised by the Abitimelock contract.
type AbitimelockExecuteTransactionIterator struct {
	Event *AbitimelockExecuteTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockExecuteTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockExecuteTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockExecuteTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockExecuteTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockExecuteTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockExecuteTransaction represents a ExecuteTransaction event raised by the Abitimelock contract.
type AbitimelockExecuteTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransaction is a free log retrieval operation binding the contract event 0xa560e3198060a2f10670c1ec5b403077ea6ae93ca8de1c32b451dc1a943cd6e7.
//
// Solidity: event ExecuteTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) FilterExecuteTransaction(opts *bind.FilterOpts, txHash [][32]byte, target []common.Address) (*AbitimelockExecuteTransactionIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "ExecuteTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockExecuteTransactionIterator{contract: _Abitimelock.contract, event: "ExecuteTransaction", logs: logs, sub: sub}, nil
}

// WatchExecuteTransaction is a free log subscription operation binding the contract event 0xa560e3198060a2f10670c1ec5b403077ea6ae93ca8de1c32b451dc1a943cd6e7.
//
// Solidity: event ExecuteTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) WatchExecuteTransaction(opts *bind.WatchOpts, sink chan<- *AbitimelockExecuteTransaction, txHash [][32]byte, target []common.Address) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "ExecuteTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockExecuteTransaction)
				if err := _Abitimelock.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteTransaction is a log parse operation binding the contract event 0xa560e3198060a2f10670c1ec5b403077ea6ae93ca8de1c32b451dc1a943cd6e7.
//
// Solidity: event ExecuteTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) ParseExecuteTransaction(log types.Log) (*AbitimelockExecuteTransaction, error) {
	event := new(AbitimelockExecuteTransaction)
	if err := _Abitimelock.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbitimelockNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Abitimelock contract.
type AbitimelockNewAdminIterator struct {
	Event *AbitimelockNewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockNewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockNewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockNewAdmin represents a NewAdmin event raised by the Abitimelock contract.
type AbitimelockNewAdmin struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed newAdmin)
func (_Abitimelock *AbitimelockFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdmin []common.Address) (*AbitimelockNewAdminIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "NewAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockNewAdminIterator{contract: _Abitimelock.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed newAdmin)
func (_Abitimelock *AbitimelockFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *AbitimelockNewAdmin, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "NewAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockNewAdmin)
				if err := _Abitimelock.contract.UnpackLog(event, "NewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed newAdmin)
func (_Abitimelock *AbitimelockFilterer) ParseNewAdmin(log types.Log) (*AbitimelockNewAdmin, error) {
	event := new(AbitimelockNewAdmin)
	if err := _Abitimelock.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbitimelockNewDelayIterator is returned from FilterNewDelay and is used to iterate over the raw logs and unpacked data for NewDelay events raised by the Abitimelock contract.
type AbitimelockNewDelayIterator struct {
	Event *AbitimelockNewDelay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockNewDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockNewDelay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockNewDelay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockNewDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockNewDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockNewDelay represents a NewDelay event raised by the Abitimelock contract.
type AbitimelockNewDelay struct {
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewDelay is a free log retrieval operation binding the contract event 0x948b1f6a42ee138b7e34058ba85a37f716d55ff25ff05a763f15bed6a04c8d2c.
//
// Solidity: event NewDelay(uint256 indexed newDelay)
func (_Abitimelock *AbitimelockFilterer) FilterNewDelay(opts *bind.FilterOpts, newDelay []*big.Int) (*AbitimelockNewDelayIterator, error) {

	var newDelayRule []interface{}
	for _, newDelayItem := range newDelay {
		newDelayRule = append(newDelayRule, newDelayItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "NewDelay", newDelayRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockNewDelayIterator{contract: _Abitimelock.contract, event: "NewDelay", logs: logs, sub: sub}, nil
}

// WatchNewDelay is a free log subscription operation binding the contract event 0x948b1f6a42ee138b7e34058ba85a37f716d55ff25ff05a763f15bed6a04c8d2c.
//
// Solidity: event NewDelay(uint256 indexed newDelay)
func (_Abitimelock *AbitimelockFilterer) WatchNewDelay(opts *bind.WatchOpts, sink chan<- *AbitimelockNewDelay, newDelay []*big.Int) (event.Subscription, error) {

	var newDelayRule []interface{}
	for _, newDelayItem := range newDelay {
		newDelayRule = append(newDelayRule, newDelayItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "NewDelay", newDelayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockNewDelay)
				if err := _Abitimelock.contract.UnpackLog(event, "NewDelay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDelay is a log parse operation binding the contract event 0x948b1f6a42ee138b7e34058ba85a37f716d55ff25ff05a763f15bed6a04c8d2c.
//
// Solidity: event NewDelay(uint256 indexed newDelay)
func (_Abitimelock *AbitimelockFilterer) ParseNewDelay(log types.Log) (*AbitimelockNewDelay, error) {
	event := new(AbitimelockNewDelay)
	if err := _Abitimelock.contract.UnpackLog(event, "NewDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbitimelockNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Abitimelock contract.
type AbitimelockNewPendingAdminIterator struct {
	Event *AbitimelockNewPendingAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockNewPendingAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockNewPendingAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockNewPendingAdmin represents a NewPendingAdmin event raised by the Abitimelock contract.
type AbitimelockNewPendingAdmin struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0x69d78e38a01985fbb1462961809b4b2d65531bc93b2b94037f3334b82ca4a756.
//
// Solidity: event NewPendingAdmin(address indexed newPendingAdmin)
func (_Abitimelock *AbitimelockFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts, newPendingAdmin []common.Address) (*AbitimelockNewPendingAdminIterator, error) {

	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "NewPendingAdmin", newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockNewPendingAdminIterator{contract: _Abitimelock.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0x69d78e38a01985fbb1462961809b4b2d65531bc93b2b94037f3334b82ca4a756.
//
// Solidity: event NewPendingAdmin(address indexed newPendingAdmin)
func (_Abitimelock *AbitimelockFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *AbitimelockNewPendingAdmin, newPendingAdmin []common.Address) (event.Subscription, error) {

	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "NewPendingAdmin", newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockNewPendingAdmin)
				if err := _Abitimelock.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPendingAdmin is a log parse operation binding the contract event 0x69d78e38a01985fbb1462961809b4b2d65531bc93b2b94037f3334b82ca4a756.
//
// Solidity: event NewPendingAdmin(address indexed newPendingAdmin)
func (_Abitimelock *AbitimelockFilterer) ParseNewPendingAdmin(log types.Log) (*AbitimelockNewPendingAdmin, error) {
	event := new(AbitimelockNewPendingAdmin)
	if err := _Abitimelock.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbitimelockQueueTransactionIterator is returned from FilterQueueTransaction and is used to iterate over the raw logs and unpacked data for QueueTransaction events raised by the Abitimelock contract.
type AbitimelockQueueTransactionIterator struct {
	Event *AbitimelockQueueTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbitimelockQueueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbitimelockQueueTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbitimelockQueueTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbitimelockQueueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbitimelockQueueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbitimelockQueueTransaction represents a QueueTransaction event raised by the Abitimelock contract.
type AbitimelockQueueTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQueueTransaction is a free log retrieval operation binding the contract event 0x76e2796dc3a81d57b0e8504b647febcbeeb5f4af818e164f11eef8131a6a763f.
//
// Solidity: event QueueTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) FilterQueueTransaction(opts *bind.FilterOpts, txHash [][32]byte, target []common.Address) (*AbitimelockQueueTransactionIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.FilterLogs(opts, "QueueTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AbitimelockQueueTransactionIterator{contract: _Abitimelock.contract, event: "QueueTransaction", logs: logs, sub: sub}, nil
}

// WatchQueueTransaction is a free log subscription operation binding the contract event 0x76e2796dc3a81d57b0e8504b647febcbeeb5f4af818e164f11eef8131a6a763f.
//
// Solidity: event QueueTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) WatchQueueTransaction(opts *bind.WatchOpts, sink chan<- *AbitimelockQueueTransaction, txHash [][32]byte, target []common.Address) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Abitimelock.contract.WatchLogs(opts, "QueueTransaction", txHashRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbitimelockQueueTransaction)
				if err := _Abitimelock.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQueueTransaction is a log parse operation binding the contract event 0x76e2796dc3a81d57b0e8504b647febcbeeb5f4af818e164f11eef8131a6a763f.
//
// Solidity: event QueueTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (_Abitimelock *AbitimelockFilterer) ParseQueueTransaction(log types.Log) (*AbitimelockQueueTransaction, error) {
	event := new(AbitimelockQueueTransaction)
	if err := _Abitimelock.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
