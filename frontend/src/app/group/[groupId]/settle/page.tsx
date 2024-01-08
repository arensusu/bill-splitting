
import { getGroupMembers } from "../action";
import ConfirmButton from "./ConfirmButton";
import { settleUp } from "./actions";
import { Button, List, ListItem } from "@mui/material";

export default async function GroupSettlePage({ params }: { params: { groupId: string } }) {
    const settleResult: {settlements: any[]} = await settleUp(params.groupId);

    const groupMembers: any[] = await getGroupMembers(params.groupId);

    return (
        <div className="flex max-w-md flex-col gap-4 m-10">
            <h1>Settle up</h1>
            <List>
                {settleResult.settlements.map((result) => {
                    return (
                        <ListItem key={`${result.group_id}${result.payer_id}${result.payee_id}`} className="flex gap-4 item-center">
                                {groupMembers.find((member) => member.id === result.payer_id).username} to {groupMembers.find((member) => member.id === result.payee_id).username}: {result.amount}
                                <ConfirmButton result={result} />
                        </ListItem>
                    )
                })}
            </List>
            <Button href={`/group/${params.groupId}`}>Back</Button>
        </div>
    )
}