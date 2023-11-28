package db_test

import (
	"bill-splitting/db/helper"
	db "bill-splitting/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomSettlement(t *testing.T) db.Settlement {
	group := createRandomGroup(t)
	payer := createRandomUser(t)
	payee := createRandomUser(t)
	fmt.Println(payer.ID, payee.ID)

	param := db.CreateSettlementParams{
		GroupID: group.ID,
		PayerID: payer.ID,
		PayeeID: payee.ID,
		Amount:  helper.RandomInt64(1, 1000),
		Date:    helper.RandomDate(),
	}
	settlement, err := testStore.Queries.CreateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement)

	require.Equal(t, group.ID, settlement.GroupID)
	require.Equal(t, payer.ID, settlement.PayerID)
	require.Equal(t, payee.ID, settlement.PayeeID)
	require.Equal(t, param.Amount, settlement.Amount)
	require.WithinDuration(t, param.Date, settlement.Date, time.Second)

	return settlement
}

func TestCreateSettlement(t *testing.T) {
	createRandomSettlement(t)
}

func TestGetSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	settlement2, err := testStore.Queries.GetSettlement(context.Background(), db.GetSettlementParams{
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
	require.WithinDuration(t, settlement1.Date, settlement2.Date, time.Second)
}

func TestUpdateSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	newAmount := helper.RandomInt64(1, 1000)
	newDate := helper.RandomDate()
	param := db.UpdateSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
		Amount:  newAmount,
		Date:    newDate,
	}

	settlement2, err := testStore.Queries.UpdateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.GroupID, settlement2.GroupID)
	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, newAmount, settlement2.Amount)
	require.WithinDuration(t, newDate, settlement2.Date, time.Second)
}

func TestDeleteSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	err := testStore.Queries.DeleteSettlement(context.Background(), db.DeleteSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.NoError(t, err)

	settlement2, err := testStore.Queries.GetSettlement(context.Background(), db.GetSettlementParams{
		GroupID: settlement1.GroupID,
		PayerID: settlement1.PayerID,
		PayeeID: settlement1.PayeeID,
	})

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, settlement2)
}
