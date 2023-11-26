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
	payer := createRandomUser(t)
	payee := createRandomUser(t)
	fmt.Println(payer.ID, payee.ID)

	param := db.CreateSettlementParams{
		PayerID: payer.ID,
		PayeeID: payee.ID,
		Amount:  helper.RandomInt64(1, 1000),
		Date:    helper.RandomDate(),
	}
	settlement, err := testQueries.CreateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement)

	require.NotZero(t, settlement.ID)
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

	settlement2, err := testQueries.GetSettlement(context.Background(), settlement1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement1.ID, settlement2.ID)
	require.Equal(t, settlement1.PayerID, settlement2.PayerID)
	require.Equal(t, settlement1.PayeeID, settlement2.PayeeID)
	require.Equal(t, settlement1.Amount, settlement2.Amount)
	require.WithinDuration(t, settlement1.Date, settlement2.Date, time.Second)
}

func TestUpdateSettlement(t *testing.T) {
	settlement := createRandomSettlement(t)

	newPayerID := settlement.PayeeID
	newPayeeID := settlement.PayerID
	newAmount := helper.RandomInt64(1, 1000)
	newDate := helper.RandomDate()
	param := db.UpdateSettlementParams{
		ID:      settlement.ID,
		PayerID: newPayerID,
		PayeeID: newPayeeID,
		Amount:  newAmount,
		Date:    newDate,
	}

	settlement2, err := testQueries.UpdateSettlement(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, settlement2)

	require.Equal(t, settlement.ID, settlement2.ID)
	require.Equal(t, newPayerID, settlement2.PayerID)
	require.Equal(t, newPayeeID, settlement2.PayeeID)
	require.Equal(t, newAmount, settlement2.Amount)
	require.WithinDuration(t, newDate, settlement2.Date, time.Second)
}

func TestDeleteSettlement(t *testing.T) {
	settlement1 := createRandomSettlement(t)

	err := testQueries.DeleteSettlement(context.Background(), settlement1.ID)

	require.NoError(t, err)

	settlement2, err := testQueries.GetSettlement(context.Background(), settlement1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, settlement2)
}
