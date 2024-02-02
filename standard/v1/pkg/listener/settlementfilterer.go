package listener

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	sg "github.com/primevprotocol/contracts-abi/clients/SettlementGateway"
	"github.com/rs/zerolog/log"
)

type SettlementFilterer struct {
	*sg.SettlementgatewayFilterer
}

func NewSettlementFilterer(
	gatewayAddr common.Address,
	client *ethclient.Client,
) *SettlementFilterer {
	f, err := sg.NewSettlementgatewayFilterer(gatewayAddr, client)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create settlement gateway filterer")
	}
	return &SettlementFilterer{f}
}

func (f *SettlementFilterer) ObtainTransferInitiatedEvents(opts *bind.FilterOpts) []TransferInitiatedEvent {
	iter, err := f.FilterTransferInitiated(opts, nil, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to filter transfer initiated")
	}
	toReturn := make([]TransferInitiatedEvent, 0)
	for iter.Next() {
		toReturn = append(toReturn, TransferInitiatedEvent{
			Sender:      iter.Event.Sender.String(),
			Recipient:   iter.Event.Recipient.String(),
			Amount:      iter.Event.Amount.Uint64(),
			TransferIdx: iter.Event.TransferIdx.Uint64(),
			srcChain:    l1,
		})
	}
	return toReturn
}