package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSettlement(t *testing.T) Settlement {
	group := createRandomGroup(t)
	payer := createRandomMember(t, group.ID)
	payee := createRandomMember(t, group.ID)

	param := CreateSettlementParams{
		PayerID: payer.ID,
		PayeeID: payee.ID,
		Amount:  strconv.FormatInt(helper.RandomInt64(1, 1000), 10),
	}
	settlement, err := testStore.CreateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement)

	require.Equal(t, payer.ID, settlement.PayerID)
	require.Equal(t, payee.ID, settlement.PayeeID)
	require.Equal(t, param.Amount, settlement.Amount)

	return settlement
}

func TestCreateSettlement(t *testing.T) {
	createRandomSettlement(t)
}

func TestGetSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	settlement2, err := testStore.GetSettlement(context.Background(), GetSettlementParams{
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, settlement1.Amount, settlement2.Amount)
}

func TestUpdateSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	newAmount := strconv.FormatInt(helper.RandomInt64(1, 1000), 10)
	param := UpdateSettlementParams{
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
		Amount:  newAmount,
	}

	settlement2, err := testStore.UpdateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, newAmount, settlement2.Amount)
}

func TestDeleteSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	err := testStore.DeleteSettlement(context.Background(), DeleteSettlementParams{
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.NoError(t, err)

	settlement2, err := testStore.GetSettlement(context.Background(), GetSettlementParams{
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

	lastMember, err := testStore.GetMember(context.Background(), lastSettlement.PayerID)
	require.NoError(t, err)
	settlements, err := testStore.ListSettlements(context.Background(), lastMember.GroupID)

	require.NoError(t, err)
	require.NotEmpty(t, settlements)

	for _, settlement := range settlements {
		require.NotEmpty(t, settlement)

		member, err := testStore.GetMember(context.Background(), settlement.PayerID)
		require.NoError(t, err)
		require.Equal(t, member.GroupID, lastMember.GroupID)
	}
}
