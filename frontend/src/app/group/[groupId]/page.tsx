import {getGroupExpenses, getGroupMembers, getGroupName} from "./action";
import getToken from "@/app/actions";
import { redirect } from "next/navigation";
import { Button, List, ListItem, ListSubheader, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material";

const GroupPage = async ({ params }: { params: { groupId: string } }) => {
    const token = await getToken();
    if (!token) {
        redirect("/login");
    }

    const groupName = await getGroupName(params.groupId)

    const groupMembers: any[] = await getGroupMembers(params.groupId)

    const groupExpenses: any[] = await getGroupExpenses(params.groupId)

    return (
        <div className="grid grid-col-1 gap-4 m-10">
            <Typography variant="h5">Group: {groupName}</Typography>
            <div>
                <Typography variant="h6">Expenses</Typography>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>
                                Description
                            </TableCell>
                            <TableCell>
                                Amount
                            </TableCell>
                            <TableCell>
                                Payer
                            </TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody className="divide-y divide-x">
                        {groupExpenses && groupExpenses.map((expense) => {
                            return (
                                <TableRow key={expense.id} className={expense.is_settled ? "bg-green-200" : "bg-red-200"}>
                                    <TableCell>
                                        {expense.description}
                                    </TableCell>
                                    <TableCell>
                                        {expense.amount}
                                    </TableCell>
                                    <TableCell>
                                        {groupMembers.find((member) => member.id === expense.payer_id).username}
                                    </TableCell>
                                </TableRow>
                            )
                        })}
                    </TableBody>
                </Table>
            </div>
            <div>
                <List>
                    <ListSubheader>Members</ListSubheader>
                    {groupMembers.map((member) => <ListItem key={member.id}>{member.username}</ListItem>)}
                </List>
            </div>
            <Button href={`/group/${params.groupId}/invite`}>
                Invite Member
            </Button>
            <Button href={`/group/${params.groupId}/expense`}>
                Create Expense
            </Button>
            <Button href={`/group/${params.groupId}/settle`}>
                Settle Up
            </Button>
            <Button href="/group">
                Back
            </Button>
        </div>
    )
}

export default GroupPage;