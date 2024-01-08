package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSettlement(t *testing.T) Settlement {
	group := createRandomGroup(t)
	payer := createRandomUser(t)
	payee := createRandomUser(t)

	param := CreateSettlementParams{
		GroupID: group.ID,
		PayerID: payer.ID,
		PayeeID: payee.ID,
		Amount:  helper.RandomInt64(1, 1000),
	}
	settlement, err := testStore.CreateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement)

	require.Equal(t, group.ID, settlement.GroupID)
	require.Equal(t, payer.ID, settlement.PayerID)
	require.Equal(t, payee.ID, settlement.PayeeID)
	require.Equal(t, param.Amount, settlement.Amount)
	require.False(t, settlement.IsConfirmed)

	return settlement
}

func TestCreateSettlement(t *testing.T) {
	createRandomSettlement(t)
}

func TestGetSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	settlement2, err := testStore.GetSettlement(context.Background(), GetSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.GroupID, settlement2.GroupID)
	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, settlement1.Amount, settlement2.Amount)
	require.Equal(t, settlement1.IsConfirmed, settlement2.IsConfirmed)
}

func TestUpdateSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	newAmount := helper.RandomInt64(1, 1000)
	newIsConfirmed := true
	param := UpdateSettlementParams{
		GroupID:     settlement1.GroupID,
		PayerID:     settlement1.PayerID,
		PayeeID:     settlement1.PayeeID,
		Amount:      newAmount,
		IsConfirmed: newIsConfirmed,
	}

	settlement2, err := testStore.UpdateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.GroupID, settlement2.GroupID)
	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, newAmount, settlement2.Amount)
	require.Equal(t, newIsConfirmed, settlement2.IsConfirmed)
}

func TestDeleteSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	err := testStore.DeleteSettlement(context.Background(), DeleteSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.NoError(t, err)

	settlement2, err := testStore.GetSettlement(context.Background(), GetSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, settlement2)
}

func TestListSettlements(t *testing.T) {
	var lastSettlement Settlement
	for i := 0; i < 10; i++ {
		lastSettlement = createRandomSettlement(t)
	}

	settlements, err := testStore.ListSettlements(context.Background(), lastSettlement.GroupID)

	require.NoError(t, err)
	require.NotEmpty(t, settlements)

	for _, settlement := range settlements {
		require.NotEmpty(t, settlement)
		require.Equal(t, lastSettlement.GroupID, settlement.GroupID)
	}
}
