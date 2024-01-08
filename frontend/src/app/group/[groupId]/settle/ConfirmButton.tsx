'use client';

import { Button } from "@mui/material";
import { comfirmSettlement } from "./actions";

export default function ConfirmButton({ result }: { result: any }) {
    return (
        <Button onClick={() => comfirmSettlement(result.group_id, result.payer_id, result.payee_id)}>Comfirm</Button>
    )
}